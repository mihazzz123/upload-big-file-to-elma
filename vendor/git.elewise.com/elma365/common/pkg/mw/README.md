# mw
`import "git.elewise.com/elma365/common/pkg/mw"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [func AdminGuard(next http.Handler) http.Handler](#AdminGuard)
* [func HTTPCompanyMiddleware(next http.Handler) http.Handler](#HTTPCompanyMiddleware)
* [func HTTPEditionMiddleware(next http.Handler) http.Handler](#HTTPEditionMiddleware)
* [func HTTPGatewayMiddleware(cfg interface{Name() string}) func(next http.Handler) http.Handler](#HTTPGatewayMiddleware)
* [func HTTPKVMiddleware(next http.Handler) http.Handler](#HTTPKVMiddleware)
* [func HTTPLangMiddleware(next http.Handler) http.Handler](#HTTPLangMiddleware)
* [func HTTPLoggerMiddleware(next http.Handler) http.Handler](#HTTPLoggerMiddleware)
* [func HTTPMetricsMiddleware(next http.Handler) http.Handler](#HTTPMetricsMiddleware)
* [func HTTPTimestampFromHeaderMiddleware(next http.Handler) http.Handler](#HTTPTimestampFromHeaderMiddleware)
* [func HTTPTimestampNowMiddleware(next http.Handler) http.Handler](#HTTPTimestampNowMiddleware)
* [func HTTPTracingMiddleware(next http.Handler) http.Handler](#HTTPTracingMiddleware)
* [func HTTPUserIDMiddleware(next http.Handler) http.Handler](#HTTPUserIDMiddleware)
* [func InjectMDToHTTPHeaders(ctx context.Context, headers http.Header)](#InjectMDToHTTPHeaders)
* [func LoggerWithTraceID(ctx context.Context, logger *zap.Logger) *zap.Logger](#LoggerWithTraceID)
* [func MDClientInterceptor() grpc.UnaryClientInterceptor](#MDClientInterceptor)
* [func NewGRPCServer() *grpc.Server](#NewGRPCServer)
* [func NewHTTPServer(cfg Config, handler http.Handler, opts ...HTTPServiceOption) *http.Server](#NewHTTPServer)
* [func ParseBody(bodyInstance interface{}) func(next http.Handler) http.Handler](#ParseBody)
* [func ParsedBodyFromContext(ctx context.Context) interface{}](#ParsedBodyFromContext)
* [func StreamMDClientInterceptor() grpc.StreamClientInterceptor](#StreamMDClientInterceptor)
* [func StreamMetadataInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error](#StreamMetadataInterceptor)
* [func TraceIDFromContext(ctx context.Context) (traceID, spanID string)](#TraceIDFromContext)
* [func UnaryErrorInterceptor(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error)](#UnaryErrorInterceptor)
* [func UnaryMetadataInterceptor(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error)](#UnaryMetadataInterceptor)
* [func UnaryPayloadInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error)](#UnaryPayloadInterceptor)
* [func UnaryValidateInterceptor(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error)](#UnaryValidateInterceptor)
* [func WithCompanyAlias(red company.RedisStorage, namespacePrefix string) func(next http.Handler) http.Handler](#WithCompanyAlias)
* [func WithIP(next http.Handler) http.Handler](#WithIP)
* [func WithUserAgent(next http.Handler) http.Handler](#WithUserAgent)
* [type Config](#Config)
* [type DebugConfig](#DebugConfig)
* [type DebugSwitcher](#DebugSwitcher)
  * [func NewDebugSwitcher(cfg DebugConfig) *DebugSwitcher](#NewDebugSwitcher)
  * [func (de *DebugSwitcher) ServeHTTP(w http.ResponseWriter, r *http.Request)](#DebugSwitcher.ServeHTTP)
* [type HTTPServiceOption](#HTTPServiceOption)
  * [func WithEventsCollector(events http.Handler) HTTPServiceOption](#WithEventsCollector)
  * [func WithHandler(pattern string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) HTTPServiceOption](#WithHandler)
  * [func WithHealthHandler(health http.HandlerFunc) HTTPServiceOption](#WithHealthHandler)
  * [func WithMigrator(mig http.Handler) HTTPServiceOption](#WithMigrator)
  * [func WithReadyHandler(ready http.HandlerFunc) HTTPServiceOption](#WithReadyHandler)


#### <a name="pkg-files">Package files</a>
[company_alias.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/company_alias.go) [const.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/const.go) [debug_switcher.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/debug_switcher.go) [grpc.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/grpc.go) [http.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/http.go) [parse_body.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/parse_body.go) [trace_id.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/trace_id.go) [with_ip.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/with_ip.go) [with_user_agent.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/with_user_agent.go)


## <a name="pkg-constants">Constants</a>
``` go
const (
    // CompanyHTTPHeader заголовок HTTP-запроса, из которого (и в который) будет извлекаться (записываться) имя компании
    CompanyHTTPHeader = "X-Company"
    // CompanyGRPCHeader заголовок gRPC-запроса, из которого (и в который) будет извлекаться (записываться) имя компании
    CompanyGRPCHeader = "company"
    // CompanyAMQPHeader заголовок AMQP-сообщения, из которого (и в который) будет извлекаться (записываться) имя компании
    CompanyAMQPHeader = "company"
    // CompanyLogEntry название поля в логах и трейсах
    CompanyLogEntry = "company"

    // CompanyAliasHTTPHeader заголовок HTTP-запроса, из которого (и в который) будет извлекаться (записываться) алиас компании
    CompanyAliasHTTPHeader = "X-Company-Alias"
    // CompanyAliasGRPCHeader заголовок gRPC-запроса, из которого (и в который) будет извлекаться (записываться) алиас компании
    CompanyAliasGRPCHeader = "company_alias"
    // CompanyAliasAMQPHeader заголовок AMQP-сообщения, из которого (и в который) будет извлекаться (записываться) алиас компании
    CompanyAliasAMQPHeader = "company_alias"
    // CompanyAliasLogEntry название поля в логах и трейсах
    CompanyAliasLogEntry = "company_alias"

    // TimestampHTTPHeader заголовок HTTP-запрос, из которого (и в который) будет извлекаться (записываться) временная метка
    TimestampHTTPHeader = "X-Timestamp"
    // TimestampLogEntry название поля в логах и трейсах
    TimestampLogEntry = "timestamp"

    // IsPortalUserHTTPHeader заголовок HTTP-запроса, из которого (и в который) будет извлекаться (записываться) что пользователь портальный
    IsPortalUserHTTPHeader = "X-Is-Portal-User"
    // IsPortalUserGRPCHeader заголовок gRPC-запроса, из которого (и в который) будет извлекаться (записываться) что пользователь портальный
    IsPortalUserGRPCHeader = "is_portal_user"
    // IsPortalUserAMQPHeader заголовок AMQP-сообщения, из которого (и в который) будет извлекаться (записываться) что пользователь портальный
    IsPortalUserAMQPHeader = "is_portal_user"

    // UserIDHTTPHeader заголовок HTTP-запроса, из которого (и в который) будет извлекаться (записываться) id пользователя
    UserIDHTTPHeader = "X-User-ID"
    // UserIDGRPCHeader заголовок gRPC-запроса, из которого (и в который) будет извлекаться (записываться) id пользователя
    UserIDGRPCHeader = "user_id"
    // UserIDAMQPHeader заголовок AMQP-сообщения, из которого (и в который) будет извлекаться (записываться) id пользователя
    UserIDAMQPHeader = "user_id"
    // UserIDLogEntry название поля в логах и трейсах
    UserIDLogEntry = "user_id"

    // IsAdminHTTPHeader заголовок HTTP-запроса, из которого (и в который) будет извлекаться (записываться) признак того, что пользователь администратор
    IsAdminHTTPHeader = "X-Is-Admin"
    // IsAdminGRPCHeader заголовок gRPC-запроса, из которого (и в который) будет извлекаться (записываться) признак того, что пользователь администратор
    IsAdminGRPCHeader = "is_admin"
    // IsAdminAMQPHeader заголовок AMQP-сообщения, из которого (и в который) будет извлекаться (записываться) признак того, что пользователь администратор
    IsAdminAMQPHeader = "is_admin"
    // IsAdminLogEntry название поля в логах и трейсах
    IsAdminLogEntry = "is_admin"

    // TraceIDHeader название заголовка ответа, в которое будет положен id трейса
    TraceIDHeader = "X-Trace-ID"

    // LangHTTPHeader заголовок HTTP-запроса, из которого (и в который) будет извлекаться (записываться) текущий язык
    LangHTTPHeader = "X-Language"
    // LangGRPCHeader заголовок gRPC-запроса, из которого (и в который) будет извлекаться (записываться) текущий язык
    LangGRPCHeader = "language"
    // LangAMQPHeader заголовок AMQP-сообщения, из которого (и в который) будет извелекаться (записываться) значение языка контекста вызова
    LangAMQPHeader = "language"

    // GatewayHTTPHeader заголовок для списка шлюзов, через которые прошёл запрос
    GatewayHTTPHeader = "X-Gateway"
    // GatewayGRPCHeader заголовок для списка шлюзов, через которые прошёл запрос
    GatewayGRPCHeader = "gateway"
    // GatewayAMQPHeader заголовок для списка шлюзов, через которые прошёл запрос
    GatewayAMQPHeader = "gateway"

    // KVHTTPHeaderPrefix префикс для заголовков дополнительной информации запроса
    KVHTTPHeaderPrefix = "X-KV-"
    // KVGRPCHeaderPrefix префикс для заголовков дополнительной информации запроса
    KVGRPCHeaderPrefix = "kv-"
    // KVAMQPHeaderPrefix префикс для заголовков дополнительной информации запроса
    KVAMQPHeaderPrefix = "kv_"

    // EditionHTTPHeader заголовок HTTP-запроса, из которого (и в который) будет извлекаться (записываться) edition компании
    EditionHTTPHeader = "X-Edition"
    // EditionGRPCHeader заголовок gRPC-запроса, из которого (и в который) будет извлекаться (записываться) edition компании
    EditionGRPCHeader = "edition"
)
```
``` go
const ClientIPKey = "clientIP"
```
ClientIPKey ключ ip-адреса клиента

``` go
const ClientUserAgentKey = "clientUserAgent"
```
ClientUserAgentKey ключ для user agent клиента




## <a name="AdminGuard">func</a> [AdminGuard](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/http.go?s=16485:16532#L498)
``` go
func AdminGuard(next http.Handler) http.Handler
```
AdminGuard проверяет наличие у человека привилегии администратора



## <a name="HTTPCompanyMiddleware">func</a> [HTTPCompanyMiddleware](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/http.go?s=10257:10315#L308)
``` go
func HTTPCompanyMiddleware(next http.Handler) http.Handler
```
HTTPCompanyMiddleware извлекает информацию о компании из заголовка запроса



## <a name="HTTPEditionMiddleware">func</a> [HTTPEditionMiddleware](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/http.go?s=14470:14528#L432)
``` go
func HTTPEditionMiddleware(next http.Handler) http.Handler
```
HTTPEditionMiddleware извлекает информацию о edition компании и устанавливает в текущий контекст



## <a name="HTTPGatewayMiddleware">func</a> [HTTPGatewayMiddleware](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/http.go?s=9280:9375#L282)
``` go
func HTTPGatewayMiddleware(cfg interface{ Name() string }) func(next http.Handler) http.Handler
```
HTTPGatewayMiddleware добавляет md.KV в контекст из заголовков



## <a name="HTTPKVMiddleware">func</a> [HTTPKVMiddleware](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/http.go?s=9797:9850#L295)
``` go
func HTTPKVMiddleware(next http.Handler) http.Handler
```
HTTPKVMiddleware добавляет md.KV в контекст из заголовков



## <a name="HTTPLangMiddleware">func</a> [HTTPLangMiddleware](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/http.go?s=11345:11400#L340)
``` go
func HTTPLangMiddleware(next http.Handler) http.Handler
```
HTTPLangMiddleware извлекает информацию о языке и устанавливает в текущий контекст



## <a name="HTTPLoggerMiddleware">func</a> [HTTPLoggerMiddleware](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/http.go?s=6652:6709#L215)
``` go
func HTTPLoggerMiddleware(next http.Handler) http.Handler
```
HTTPLoggerMiddleware добавляет логгер в контекст



## <a name="HTTPMetricsMiddleware">func</a> [HTTPMetricsMiddleware](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/http.go?s=6276:6334#L204)
``` go
func HTTPMetricsMiddleware(next http.Handler) http.Handler
```
HTTPMetricsMiddleware добавляет метрики к HTTP запросу



## <a name="HTTPTimestampFromHeaderMiddleware">func</a> [HTTPTimestampFromHeaderMiddleware](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/http.go?s=12370:12440#L374)
``` go
func HTTPTimestampFromHeaderMiddleware(next http.Handler) http.Handler
```
HTTPTimestampFromHeaderMiddleware временная метка из заголовка



## <a name="HTTPTimestampNowMiddleware">func</a> [HTTPTimestampNowMiddleware](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/http.go?s=11766:11829#L357)
``` go
func HTTPTimestampNowMiddleware(next http.Handler) http.Handler
```
HTTPTimestampNowMiddleware временная метка из time.Now()



## <a name="HTTPTracingMiddleware">func</a> [HTTPTracingMiddleware](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/http.go?s=7017:7075#L228)
``` go
func HTTPTracingMiddleware(next http.Handler) http.Handler
```
HTTPTracingMiddleware middleware add tracing span to context and it's id in response header



## <a name="HTTPUserIDMiddleware">func</a> [HTTPUserIDMiddleware](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/http.go?s=13266:13323#L398)
``` go
func HTTPUserIDMiddleware(next http.Handler) http.Handler
```
HTTPUserIDMiddleware извлекает информацию о пользователе из заголовка запроса



## <a name="InjectMDToHTTPHeaders">func</a> [InjectMDToHTTPHeaders](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/http.go?s=15201:15269#L457)
``` go
func InjectMDToHTTPHeaders(ctx context.Context, headers http.Header)
```
InjectMDToHTTPHeaders добавляет заголовки с метадатой запроса



## <a name="LoggerWithTraceID">func</a> [LoggerWithTraceID](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/trace_id.go?s=402:477#L17)
``` go
func LoggerWithTraceID(ctx context.Context, logger *zap.Logger) *zap.Logger
```
LoggerWithTraceID добавляет поля TraceID и SpanID к полям логгера



## <a name="MDClientInterceptor">func</a> [MDClientInterceptor](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/grpc.go?s=1612:1666#L51)
``` go
func MDClientInterceptor() grpc.UnaryClientInterceptor
```
MDClientInterceptor добавляет имя компании и id пользователя в контекст запроса



## <a name="NewGRPCServer">func</a> [NewGRPCServer](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/grpc.go?s=3849:3882#L112)
``` go
func NewGRPCServer() *grpc.Server
```
NewGRPCServer возвращает gRPC сервер с настроенными middleware



## <a name="NewHTTPServer">func</a> [NewHTTPServer](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/http.go?s=3071:3163#L95)
``` go
func NewHTTPServer(cfg Config, handler http.Handler, opts ...HTTPServiceOption) *http.Server
```
NewHTTPServer создаёт новый сервер поверх переданного хэндлера с роутами pprof и prometheus



## <a name="ParseBody">func</a> [ParseBody](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/parse_body.go?s=604:681#L22)
``` go
func ParseBody(bodyInstance interface{}) func(next http.Handler) http.Handler
```
ParseBody получает из тела запроса данные и валидирует их согласно тэгам валидации структуры bodyInstance

В качестве валидатора используется gopkg.in/go-playground/validator.v9



## <a name="ParsedBodyFromContext">func</a> [ParsedBodyFromContext](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/parse_body.go?s=2101:2160#L61)
``` go
func ParsedBodyFromContext(ctx context.Context) interface{}
```
ParsedBodyFromContext получает контекст с разобранным и отвалидированным телом запроса

Т.к. метод общий, то возвращает он пустой интерфейс, хотя реально в нем лежит ссылка на объект нужного типа.
Пример использования: req := commonmw.ParsedBodyFromContext(ctx).(createMessageRequest)



## <a name="StreamMDClientInterceptor">func</a> [StreamMDClientInterceptor](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/grpc.go?s=2222:2283#L64)
``` go
func StreamMDClientInterceptor() grpc.StreamClientInterceptor
```
StreamMDClientInterceptor добавляет имя компании, алиас(если есть) и id пользователя в контекст запроса стримингового grpc клиента



## <a name="StreamMetadataInterceptor">func</a> [StreamMetadataInterceptor](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/grpc.go?s=5052:5184#L147)
``` go
func StreamMetadataInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error
```
StreamMetadataInterceptor extract company, user, privileges from metadata and put it in context



## <a name="TraceIDFromContext">func</a> [TraceIDFromContext](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/trace_id.go?s=192:261#L12)
``` go
func TraceIDFromContext(ctx context.Context) (traceID, spanID string)
```
TraceIDFromContext extract trace and span ids if span started or empty string otherwise



## <a name="UnaryErrorInterceptor">func</a> [UnaryErrorInterceptor](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/grpc.go?s=7605:7742#L223)
``` go
func UnaryErrorInterceptor(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error)
```
UnaryErrorInterceptor извлекает код ответа из ошибки и явно передаёт его дальше



## <a name="UnaryMetadataInterceptor">func</a> [UnaryMetadataInterceptor](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/grpc.go?s=4684:4824#L135)
``` go
func UnaryMetadataInterceptor(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error)
```
UnaryMetadataInterceptor извлекает имя компании и id пользователя из контекста запроса



## <a name="UnaryPayloadInterceptor">func</a> [UnaryPayloadInterceptor](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/grpc.go?s=8007:8149#L237)
``` go
func UnaryPayloadInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error)
```
UnaryPayloadInterceptor логирует запросы и ответы



## <a name="UnaryValidateInterceptor">func</a> [UnaryValidateInterceptor](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/grpc.go?s=8785:8925#L265)
``` go
func UnaryValidateInterceptor(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error)
```
UnaryValidateInterceptor валидирует входные данные и если они невалидны, тут же отвечает соответствующей ошибкой



## <a name="WithCompanyAlias">func</a> [WithCompanyAlias](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/company_alias.go?s=567:675#L18)
``` go
func WithCompanyAlias(red company.RedisStorage, namespacePrefix string) func(next http.Handler) http.Handler
```
WithCompanyAlias middleware которая заменяет алиса компании на её имя, если находится алиас
требует наличие компании (CompanyFromContext()) в контексте



## <a name="WithIP">func</a> [WithIP](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/with_ip.go?s=261:304#L15)
``` go
func WithIP(next http.Handler) http.Handler
```
WithIP добавляет ip-адрес клиента в контекст



## <a name="WithUserAgent">func</a> [WithUserAgent](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/with_user_agent.go?s=270:320#L13)
``` go
func WithUserAgent(next http.Handler) http.Handler
```
WithUserAgent добавляет user agent клиента в контекст




## <a name="Config">type</a> [Config](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/http.go?s=2841:2922#L88)
``` go
type Config interface {
    DebugConfig
    Name() string
    GetTimeout() time.Duration
}
```
Config конфигурация необходимая для настройки http-интерфейса










## <a name="DebugConfig">type</a> [DebugConfig](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/debug_switcher.go?s=178:255#L11)
``` go
type DebugConfig interface {
    IsDebug() bool
    EnableDebug()
    DisableDebug()
}
```
DebugConfig конфигурация необходимая для переключения дебаг режима










## <a name="DebugSwitcher">type</a> [DebugSwitcher](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/debug_switcher.go?s=343:442#L18)
``` go
type DebugSwitcher struct {
    sync.Mutex
    // contains filtered or unexported fields
}

```
DebugSwitcher сервис для переключения дебаг режима







### <a name="NewDebugSwitcher">func</a> [NewDebugSwitcher](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/debug_switcher.go?s=487:540#L26)
``` go
func NewDebugSwitcher(cfg DebugConfig) *DebugSwitcher
```
NewDebugSwitcher конструктор





### <a name="DebugSwitcher.ServeHTTP">func</a> (\*DebugSwitcher) [ServeHTTP](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/debug_switcher.go?s=764:838#L40)
``` go
func (de *DebugSwitcher) ServeHTTP(w http.ResponseWriter, r *http.Request)
```
ServeHTTP implements http.Handler




## <a name="HTTPServiceOption">type</a> [HTTPServiceOption](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/http.go?s=800:866#L31)
``` go
type HTTPServiceOption interface {
    Apply(chi.Router) chi.Router
}
```
HTTPServiceOption опции настройки







### <a name="WithEventsCollector">func</a> [WithEventsCollector](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/http.go?s=2162:2225#L71)
``` go
func WithEventsCollector(events http.Handler) HTTPServiceOption
```
WithEventsCollector добавить интерфейс получения событий системы


### <a name="WithHandler">func</a> [WithHandler](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/http.go?s=2464:2584#L80)
``` go
func WithHandler(pattern string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) HTTPServiceOption
```
WithHandler добавить обработчик к роутеру по переданному пути, применив переданные мидлвары


### <a name="WithHealthHandler">func</a> [WithHealthHandler](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/http.go?s=1629:1694#L52)
``` go
func WithHealthHandler(health http.HandlerFunc) HTTPServiceOption
```
WithHealthHandler подменить стандартную функцию живости сервера


### <a name="WithMigrator">func</a> [WithMigrator](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/http.go?s=1811:1864#L57)
``` go
func WithMigrator(mig http.Handler) HTTPServiceOption
```
WithMigrator включить интерфейс мигратора


### <a name="WithReadyHandler">func</a> [WithReadyHandler](https://git.elewise.com/elma365/common/-/tree/develop/pkg/mw/http.go?s=1114:1177#L42)
``` go
func WithReadyHandler(ready http.HandlerFunc) HTTPServiceOption
```
WithReadyHandler подменить стандартную функцию готовности








- - -
Generated by [godoc2md](https://github.com/Exa-Networks/godoc2md)
