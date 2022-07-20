package mw

import (
	"bytes"
	"context"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"git.elewise.com/elma365/common/pkg/edition"
	"git.elewise.com/elma365/common/pkg/i18n"
	"git.elewise.com/elma365/common/pkg/md"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// HTTPServiceOption опции настройки
type HTTPServiceOption interface {
	Apply(chi.Router) chi.Router
}

type httpServiceOptionFunc func(chi.Router) chi.Router

func (fn httpServiceOptionFunc) Apply(router chi.Router) chi.Router {
	return fn(router)
}

// WithReadyHandler подменить стандартную функцию готовности
func WithReadyHandler(ready http.HandlerFunc) HTTPServiceOption {
	return httpServiceOptionFunc(func(router chi.Router) chi.Router {
		// Deprecated: оставлено для обратной совместимости, необходимо использовать /readyz
		router = WithHandler("/ready", ready).Apply(router)
		router = WithHandler("/readyz", ready).Apply(router)
		return router
	})
}

// WithHealthHandler подменить стандартную функцию живости сервера
func WithHealthHandler(health http.HandlerFunc) HTTPServiceOption {
	return WithHandler("/healthz", health)
}

// WithMigrator включить интерфейс мигратора
func WithMigrator(mig http.Handler) HTTPServiceOption {
	return WithHandler(
		"/migration",
		mig,
		middleware.RealIP,
		HTTPMetricsMiddleware,
		HTTPLoggerMiddleware,
		HTTPTracingMiddleware,
		HTTPLangMiddleware,
		HTTPKVMiddleware,
	)
}

// WithEventsCollector добавить интерфейс получения событий системы
func WithEventsCollector(events http.Handler) HTTPServiceOption {
	return WithHandler(
		"/events",
		events,
		HTTPLangMiddleware,
	)
}

// WithHandler добавить обработчик к роутеру по переданному пути, применив переданные мидлвары
func WithHandler(pattern string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) HTTPServiceOption {
	return httpServiceOptionFunc(func(router chi.Router) chi.Router {
		router.With(middlewares...).Mount(pattern, handler)
		return router
	})
}

// Config конфигурация необходимая для настройки http-интерфейса
type Config interface {
	DebugConfig
	Name() string
	GetTimeout() time.Duration
}

// NewHTTPServer создаёт новый сервер поверх переданного хэндлера с роутами pprof и prometheus
func NewHTTPServer(cfg Config, handler http.Handler, opts ...HTTPServiceOption) *http.Server {
	defaultHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router := chi.NewRouter()

	for _, opt := range opts {
		opt.Apply(router)
	}
	if !router.Match(chi.NewRouteContext(), http.MethodGet, "/readyz") {
		router.Mount("/ready", defaultHandler)
		router.Mount("/readyz", defaultHandler)
	}
	if !router.Match(chi.NewRouteContext(), http.MethodGet, "/healthz") {
		router.Mount("/healthz", defaultHandler)
	}

	router.Mount("/debug", middleware.Profiler())
	router.Mount("/metrics", promhttp.Handler())
	router.Mount("/switch-debug", NewDebugSwitcher(cfg))

	if handler != nil {
		router.Group(func(router chi.Router) {
			router.Use(
				middleware.RealIP,
				HTTPMetricsMiddleware,
				HTTPLoggerMiddleware,
				HTTPTracingMiddleware,
				HTTPLangMiddleware,
				HTTPGatewayMiddleware(cfg),
				HTTPKVMiddleware,
				HTTPEditionMiddleware,
			)
			router.Group(func(router chi.Router) {
				if cfg.GetTimeout() > 0 {
					router.Use(middleware.Timeout(cfg.GetTimeout()))
				} else {
					zap.S().Warn("Default timeout has been disabled. Make sure you applied custom timeouts to your routes.")
				}
				router.Mount("/", handler)
			})
		})
	}

	errorLogger, err := zap.NewStdLogAt(zap.L(), zapcore.ErrorLevel)
	if err != nil {
		panic("fail to create logger: " + err.Error())
	}

	return &http.Server{
		Handler:           router,
		ErrorLog:          errorLogger,
		ReadHeaderTimeout: 10 * time.Second,
	}
}

//nolint: gochecknoglobals // метрики должны быть глобальными
var (
	inFlightGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "http_server",
		Name:      "in_flight_requests",
		Help:      "A gauge of requests currently being served by the wrapped handler.",
	})

	counter = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "http_server",
		Name:      "requests_total",
		Help:      "A counter for requests to the wrapped handler.",
	}, []string{"code"})

	requestDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "http_server",
		Name:      "request_duration_seconds",
		Help:      "A histogram of latencies for requests.",
		Buckets:   []float64{.25, .5, 1, 2.5, 5, 10},
	}, []string{})

	responseSize = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "http_server",
		Name:      "response_size_bytes",
		Help:      "A histogram of response sizes for requests.",
		Buckets:   []float64{200, 500, 900, 1500},
	}, []string{})
)

type httpHeaders http.Header

// Add implements md.OutgoingHeaders
func (headers httpHeaders) Add(key, value string) {
	http.Header(headers).Add(key, value)
}

// Values implements md.IncomingHeaders
func (headers httpHeaders) Values(key string) []string {
	return http.Header(headers).Values(key)
}

// Range implements md.IncomingHeaders
func (headers httpHeaders) Range(visitor func(key string, values []string) bool) bool {
	for key, values := range headers {
		if !visitor(key, values) {
			return false
		}
	}
	return true
}

// HTTPMetricsMiddleware добавляет метрики к HTTP запросу
func HTTPMetricsMiddleware(next http.Handler) http.Handler {
	return promhttp.InstrumentHandlerInFlight(inFlightGauge,
		promhttp.InstrumentHandlerCounter(counter,
			promhttp.InstrumentHandlerDuration(requestDuration,
				promhttp.InstrumentHandlerResponseSize(responseSize, next),
			),
		),
	)
}

// HTTPLoggerMiddleware добавляет логгер в контекст
func HTTPLoggerMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger := zap.L()
		ctx = ctxzap.ToContext(ctx, logger)

		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

// HTTPTracingMiddleware middleware add tracing span to context and it's id in response header
func HTTPTracingMiddleware(next http.Handler) http.Handler {
	// HTTPTracingMiddleware middleware add tracing span to context and it's id in response header
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var sp opentracing.Span
		spCtx, _ := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(r.Header),
		)
		// В общем случае, мидлварь может обрабатывать запросы не только между сервисами, но и от внешних клиентов,
		// которые знать не знают про заголовки трассировки. Задача по добавлению заголовков должна быть переложена
		// на клиента, который инициирует запрос. Для внутренних сервисов, заголовки будут автоматически добавлены
		// в рамках https://git.elewise.com/elma365/common/issues/15.
		// Таким образом, проверять наличие ошибки и писать сообщение в лог об отсутствии заголовков трассировки
		// становится избыточным.

		sp = opentracing.StartSpan(r.URL.Path, opentracing.ChildOf(spCtx))
		ext.HTTPMethod.Set(sp, r.Method)
		ext.HTTPUrl.Set(sp, r.URL.String())
		ctx = opentracing.ContextWithSpan(ctx, sp)

		traceID, _ := TraceIDFromContext(ctx)
		w.Header().Set(TraceIDHeader, traceID)

		logger := ctxzap.Extract(ctx)
		logger = LoggerWithTraceID(ctx, logger)
		logger.Debug(
			r.URL.Path,
			zap.String("method", r.Method),
			zap.String("clientIP", r.RemoteAddr),
		)
		ctx = ctxzap.ToContext(ctx, logger)

		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		buf := &bytes.Buffer{}
		ww.Tee(buf)

		defer func() {
			ext.HTTPStatusCode.Set(sp, uint16(ww.Status()))
			if ww.Status() >= 400 {
				ext.Error.Set(sp, true)
				sp.LogFields(log.Error(errors.New(buf.String())))
			}
			sp.Finish()
		}()

		next.ServeHTTP(ww, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

// HTTPGatewayMiddleware добавляет md.KV в контекст из заголовков
func HTTPGatewayMiddleware(cfg interface{ Name() string }) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = md.ExtractGateways(ctx, GatewayHTTPHeader, httpHeaders(r.Header))
			ctx = md.ContextWithGateway(ctx, cfg.Name())
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}

// HTTPKVMiddleware добавляет md.KV в контекст из заголовков
func HTTPKVMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = md.ExtractKV(ctx, KVHTTPHeaderPrefix, httpHeaders(r.Header))
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

var companyRE = regexp.MustCompile(`^[0-9a-zA-Z][\w\-]+$`)

// HTTPCompanyMiddleware извлекает информацию о компании из заголовка запроса
func HTTPCompanyMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger := ctxzap.Extract(ctx)
		sp := opentracing.SpanFromContext(ctx)

		company := r.Header.Get(CompanyHTTPHeader)
		if company == "" || !companyRE.MatchString(company) {
			logger.Sugar().Debugf("invalid company name %q", company)
			http.Error(w, "invalid company name", http.StatusPreconditionFailed)
			return
		}
		ctx = md.ContextWithCompany(ctx, company)
		sp.SetTag(CompanyLogEntry, company)
		logger = logger.With(zap.String(CompanyLogEntry, company))

		alias := r.Header.Get(CompanyAliasHTTPHeader)
		if len(alias) > 0 {
			ctx = md.ContextWithCompanyAlias(ctx, alias)
			sp.SetTag(CompanyAliasLogEntry, alias)
			logger = logger.With(zap.String(CompanyAliasLogEntry, alias))
		}

		lang := r.Header.Get(CompanyLangHTTPHeader)
		if len(lang) > 0 {
			ctx = i18n.ContextWithCompanyLang(ctx, lang)
			sp.SetTag(CompanyLangLogEntry, lang)
			logger = logger.With(zap.String(CompanyLangLogEntry, lang))
		}

		ctx = ctxzap.ToContext(ctx, logger)

		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

// HTTPLangMiddleware извлекает информацию о языке и устанавливает в текущий контекст
func HTTPLangMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// Текущий язык
		lang := r.Header.Get(LangHTTPHeader)
		if len(lang) > 0 {
			ctx = i18n.ContextWithLang(ctx, lang)
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

// HTTPTimestampNowMiddleware временная метка из time.Now()
func HTTPTimestampNowMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger := ctxzap.Extract(ctx)
		sp := opentracing.SpanFromContext(ctx)
		timestamp := time.Now().UTC()
		ctx = md.ContextWithTimestamp(ctx, timestamp)
		sp.SetTag(TimestampLogEntry, timestamp)
		logger = logger.With(zap.Time(TimestampLogEntry, timestamp))
		ctx = ctxzap.ToContext(ctx, logger)
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

// HTTPTimestampFromHeaderMiddleware временная метка из заголовка
func HTTPTimestampFromHeaderMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger := ctxzap.Extract(ctx)
		sp := opentracing.SpanFromContext(ctx)

		timestampSt := r.Header.Get(TimestampHTTPHeader)
		timestamp, err := time.Parse(time.RFC3339Nano, timestampSt)
		if err != nil {
			logger.Sugar().Debugf("invalid timestamp %s", timestampSt)
			http.Error(w, "invalid timestamp", http.StatusPreconditionFailed)
			return
		}
		ctx = md.ContextWithTimestamp(ctx, timestamp)
		sp.SetTag(TimestampLogEntry, timestamp)
		logger = logger.With(zap.Time(TimestampLogEntry, timestamp))
		ctx = ctxzap.ToContext(ctx, logger)
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

// HTTPUserIDMiddleware извлекает информацию о пользователе из заголовка запроса
func HTTPUserIDMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger := ctxzap.Extract(ctx)
		sp := opentracing.SpanFromContext(ctx)

		userIDRaw := r.Header.Get(UserIDHTTPHeader)
		userID, err := uuid.FromString(userIDRaw)
		if err != nil {
			logger.Sugar().Debugf("invalid user id %q: %s", userIDRaw, err.Error())
			http.Error(w, "invalid user id", http.StatusPreconditionFailed)
			return
		}
		ctx = md.ContextWithUserID(ctx, userID)
		sp.SetTag(UserIDLogEntry, userID.String())
		logger = logger.With(zap.Stringer(UserIDLogEntry, userID))

		isAdmin := r.Header.Get(IsAdminHTTPHeader) == "true"
		ctx = md.ContextWithIsAdmin(ctx, isAdmin)
		sp.SetTag(IsAdminLogEntry, isAdmin)
		logger = logger.With(zap.Bool(IsAdminLogEntry, isAdmin))

		ctx = ctxzap.ToContext(ctx, logger)

		isPortal := r.Header.Get(IsPortalUserHTTPHeader) == "true"
		ctx = md.ContextWithIsPortalUser(ctx, isPortal)

		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

// HTTPEditionMiddleware извлекает информацию о edition компании и устанавливает в текущий контекст
func HTTPEditionMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger := ctxzap.Extract(ctx)

		// Текущий edition
		editionStr := r.Header.Get(EditionHTTPHeader)
		if len(editionStr) > 0 {
			editionCompany, err := edition.EditionString(editionStr)
			if err != nil {
				logger.Sugar().Debugf("invalid edition:  %q", editionStr)
				http.Error(w, "invalid edition", http.StatusPreconditionFailed)
				return
			}

			ctx = md.ContextWithEdition(ctx, editionCompany)
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

// InjectMDToHTTPHeaders добавляет заголовки с метадатой запроса
func InjectMDToHTTPHeaders(ctx context.Context, headers http.Header) {
	company, ok := md.TryCompanyFromContext(ctx)
	if ok {
		headers.Set(CompanyHTTPHeader, company)
	}

	alias := md.CompanyAliasFromContext(ctx)
	if len(alias) > 0 {
		headers.Set(CompanyAliasHTTPHeader, alias)
	} else {
		headers.Del(CompanyAliasHTTPHeader)
	}

	companyLang := i18n.CompanyLangFromContext(ctx)
	headers.Set(CompanyLangHTTPHeader, companyLang)

	userID, ok := md.TryUserIDFromContext(ctx)
	if ok {
		headers.Set(UserIDHTTPHeader, userID.String())
	}

	isAdmin := md.IsAdminFromContext(ctx)
	headers.Set(IsAdminHTTPHeader, strconv.FormatBool(isAdmin))

	isPortalUser := md.IsPortalUserFromContext(ctx)
	headers.Set(IsPortalUserHTTPHeader, strconv.FormatBool(isPortalUser))

	lang := i18n.LangFromContext(ctx)
	headers.Set(LangHTTPHeader, lang)

	editionCompany, ok := md.TryEditionFromContext(ctx)
	if ok {
		headers.Set(EditionHTTPHeader, editionCompany.String())
	}

	if sp := opentracing.SpanFromContext(ctx); sp != nil {
		_ = opentracing.GlobalTracer().
			Inject(sp.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(headers))
	}
	md.InjectGateways(ctx, GatewayHTTPHeader, httpHeaders(headers))
	md.InjectKV(ctx, KVHTTPHeaderPrefix, httpHeaders(headers))
}

// AdminGuard проверяет наличие у человека привилегии администратора
func AdminGuard(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if !md.IsAdminFromContext(ctx) {
			ctxzap.Extract(ctx).Debug("insufficient privileges")
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
