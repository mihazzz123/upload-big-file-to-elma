package mw

import (
	"context"
	"strconv"

	"git.elewise.com/elma365/common/pkg/edition"
	"git.elewise.com/elma365/common/pkg/errs"
	"git.elewise.com/elma365/common/pkg/errs/validation"
	"git.elewise.com/elma365/common/pkg/i18n"
	"git.elewise.com/elma365/common/pkg/md"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcHeaders metautils.NiceMD

// Add implements md.OutgoingHeaders
func (headers grpcHeaders) Add(key, value string) {
	metautils.NiceMD(headers).Add(key, value)
}

// Values implements md.IncomingHeaders
func (headers grpcHeaders) Values(key string) []string {
	return metautils.NiceMD(headers)[key]
}

// Range implements md.IncomingHeaders
func (headers grpcHeaders) Range(visitor func(key string, values []string) bool) bool {
	for key, values := range metautils.NiceMD(headers) {
		if !visitor(key, values) {
			return false
		}
	}
	return true
}

// MDClientInterceptor добавляет имя компании и id пользователя в контекст запроса
func MDClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		var err error
		ctx, err = generateClientInterceptorContext(ctx)
		if err != nil {
			return errors.WithStack(err)
		}

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

// StreamMDClientInterceptor добавляет имя компании, алиас(если есть) и id пользователя в контекст запроса стримингового grpc клиента
func StreamMDClientInterceptor() grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		var err error
		ctx, err = generateClientInterceptorContext(ctx)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		return streamer(ctx, desc, cc, method, opts...)
	}
}

func generateClientInterceptorContext(ctx context.Context) (context.Context, error) {
	data := metautils.ExtractOutgoing(ctx).Clone()

	company, ok := md.TryCompanyFromContext(ctx)
	if !ok {
		return nil, errors.Errorf("company required in context")
	}
	data.Set(CompanyGRPCHeader, company)
	data.Set(CompanyLangGRPCHeader, i18n.CompanyLangFromContext(ctx))

	userID, ok := md.TryUserIDFromContext(ctx)
	if !ok {
		return nil, errors.Errorf("userID required in context")
	}
	data.Set(UserIDGRPCHeader, userID.String())

	data.Set(IsAdminGRPCHeader, strconv.FormatBool(md.IsAdminFromContext(ctx)))
	data.Set(IsPortalUserGRPCHeader, strconv.FormatBool(md.IsPortalUserFromContext(ctx)))
	if len(md.CompanyAliasFromContext(ctx)) > 0 {
		data.Set(CompanyAliasGRPCHeader, md.CompanyAliasFromContext(ctx))
	} else {
		data.Del(CompanyAliasGRPCHeader)
	}
	data.Set(LangGRPCHeader, i18n.LangFromContext(ctx))
	editionCtx, ok := md.TryEditionFromContext(ctx)
	if ok {
		data.Set(EditionGRPCHeader, editionCtx.String())
	}
	ctx = data.ToOutgoing(ctx)

	md.InjectGateways(ctx, GatewayGRPCHeader, grpcHeaders(data))
	md.InjectKV(ctx, KVGRPCHeaderPrefix, grpcHeaders(data))

	return ctx, nil
}

// NewGRPCServer возвращает gRPC сервер с настроенными middleware
func NewGRPCServer() *grpc.Server {
	return grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_zap.UnaryServerInterceptor(zap.L(), grpc_zap.WithLevels(codeToLevelServerFunc)),
			UnaryMetadataInterceptor,
			UnaryErrorInterceptor,
			UnaryPayloadInterceptor,
			UnaryValidateInterceptor,
		)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_opentracing.StreamServerInterceptor(),
			grpc_prometheus.StreamServerInterceptor,
			grpc_zap.StreamServerInterceptor(zap.L(), grpc_zap.WithLevels(codeToLevelServerFunc)),
			StreamMetadataInterceptor,
			StreamErrorInterceptor,
		)),
	)
}

//nolint:gocyclo // Нечего упрощать
func codeToLevelServerFunc(code codes.Code) zapcore.Level {
	switch code {
	case codes.OK:
		return zap.DebugLevel
	case codes.Canceled:
		return zap.DebugLevel
	case codes.Unknown:
		return zap.ErrorLevel
	case codes.InvalidArgument:
		return zap.DebugLevel
	case codes.DeadlineExceeded:
		return zap.ErrorLevel
	case codes.NotFound:
		return zap.DebugLevel
	case codes.AlreadyExists:
		return zap.DebugLevel
	case codes.PermissionDenied:
		return zap.DebugLevel
	case codes.Unauthenticated:
		return zap.DebugLevel
	case codes.ResourceExhausted:
		return zap.WarnLevel
	case codes.FailedPrecondition:
		return zap.WarnLevel
	case codes.Aborted:
		return zap.WarnLevel
	case codes.OutOfRange:
		return zap.WarnLevel
	case codes.Unimplemented:
		return zap.ErrorLevel
	case codes.Internal:
		return zap.ErrorLevel
	case codes.Unavailable:
		return zap.WarnLevel
	case codes.DataLoss:
		return zap.ErrorLevel
	default:
		return zap.ErrorLevel
	}
}

// UnaryMetadataInterceptor извлекает имя компании и id пользователя из контекста запроса
func UnaryMetadataInterceptor(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	var err error

	ctx, err = buildContextWithMetadata(ctx)
	if err != nil {
		return nil, err
	}

	return handler(ctx, req)
}

// StreamMetadataInterceptor extract company, user, privileges from metadata and put it in context
func StreamMetadataInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	ctx, err := buildContextWithMetadata(ss.Context())
	if err != nil {
		return err
	}
	wrappedStream := grpc_middleware.WrapServerStream(ss)
	wrappedStream.WrappedContext = ctx

	return handler(srv, wrappedStream)
}

//nolint: gocyclo, funlen
func buildContextWithMetadata(ctx context.Context) (context.Context, error) {
	data := metautils.ExtractIncoming(ctx)
	tags := grpc_ctxtags.Extract(ctx)

	// extract company
	company, ok := data[CompanyGRPCHeader]
	if !ok || len(company) == 0 {
		return nil, errs.InvalidArgument.New("company required")
	}
	tags.Set(CompanyLogEntry, company[0])
	ctx = md.ContextWithCompany(ctx, company[0])

	companyAlias, ok := data[CompanyAliasGRPCHeader]
	if ok && len(companyAlias) > 0 {
		tags.Set(CompanyAliasLogEntry, companyAlias[0])
		ctx = md.ContextWithCompanyAlias(ctx, companyAlias[0])
	}

	companyLang, ok := data[CompanyLangGRPCHeader]
	if ok && len(companyLang) > 0 {
		tags.Set(CompanyLangLogEntry, companyLang[0])
		ctx = i18n.ContextWithCompanyLang(ctx, companyLang[0])
	}

	// extract user id
	userIDStr, ok := data[UserIDGRPCHeader]
	if !ok || len(userIDStr) == 0 {
		return nil, errs.InvalidArgument.New("user_id required")
	}
	userID, err := uuid.FromString(userIDStr[0])
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "user_id: %s", err.Error())
	}
	tags.Set(UserIDLogEntry, userIDStr[0])
	ctx = md.ContextWithUserID(ctx, userID)

	// extract is admin
	isAdminStr, ok := data[IsAdminGRPCHeader]
	// nolint:goconst // требует "true" сделать константой
	isAdmin := ok && len(isAdminStr) == 1 && isAdminStr[0] == "true"
	tags.Set(IsAdminLogEntry, isAdmin)
	ctx = md.ContextWithIsAdmin(ctx, isAdmin)

	// extract is portal user
	isPortalStr, ok := data[IsPortalUserGRPCHeader]
	isPortalUser := ok && len(isPortalStr) == 1 && isPortalStr[0] == "true"
	ctx = md.ContextWithIsPortalUser(ctx, isPortalUser)

	// Текущий язык пользователя
	userLang, ok := data[LangGRPCHeader]
	if ok && len(userLang) > 0 {
		ctx = i18n.ContextWithLang(ctx, userLang[0])
	}
	ctx = md.ExtractGateways(ctx, GatewayGRPCHeader, grpcHeaders(data))
	ctx = md.ExtractKV(ctx, KVGRPCHeaderPrefix, grpcHeaders(data))

	// Текущий edition
	editionStr, ok := data[EditionGRPCHeader]
	if ok && len(editionStr) > 0 {
		editionCompany, err := edition.EditionString(editionStr[0])
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid edition: %s", err.Error())
		}
		ctx = md.ContextWithEdition(ctx, editionCompany)
	}

	return ctx, nil
}

// UnaryErrorInterceptor извлекает код ответа из ошибки и явно передаёт его дальше
func UnaryErrorInterceptor(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	res, err := handler(ctx, req)

	if err != nil {
		code := errs.GRPCCodeFromError(err)
		err = status.Error(code, err.Error())
	}

	return res, err
}

// StreamErrorInterceptor извлекает код ответа из ошибки и явно передаёт его дальше
func StreamErrorInterceptor(srv interface{}, ss grpc.ServerStream, _ *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	err := handler(srv, ss)
	if err != nil {
		code := errs.GRPCCodeFromError(err)
		err = status.Error(code, err.Error())
	}

	return err
}

const errorText = "error occurred"

// UnaryPayloadInterceptor логирует запросы и ответы
func UnaryPayloadInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	l := ctxzap.Extract(ctx)

	l.Debug(
		"request payload",
		zap.String("method", info.FullMethod),
		zap.Any("request", req),
	)

	res, err := handler(ctx, req)

	if err != nil {
		logLevel := errs.ErrorLevel(err)
		if ce := l.Check(logLevel, errorText); ce != nil {
			ce.Write(zap.Error(err))
		}
	} else {
		l.Debug(
			"response payload",
			zap.String("method", info.FullMethod),
			zap.Any("response", res),
		)
	}

	return res, err
}

// UnaryValidateInterceptor валидирует входные данные и если они невалидны, тут же отвечает соответствующей ошибкой
func UnaryValidateInterceptor(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	verrs := validation.ValidateStruct(ctx, req)
	if !verrs.IsEmpty() {
		return nil, verrs
	}

	return handler(ctx, req)
}
