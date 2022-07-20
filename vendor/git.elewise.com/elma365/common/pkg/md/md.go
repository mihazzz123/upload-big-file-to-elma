package md

import (
	"context"
	"time"

	pkgEdition "git.elewise.com/elma365/common/pkg/edition"

	uuid "github.com/satori/go.uuid"
)

type companyCtxKey struct{}

// ContextWithCompany сохраняет имя компании в контексте
func ContextWithCompany(ctx context.Context, company string) context.Context {
	return context.WithValue(ctx, companyCtxKey{}, company)
}

// CompanyFromContext извлекает имя компании из контекста
func CompanyFromContext(ctx context.Context) string {
	return ctx.Value(companyCtxKey{}).(string)
}

// TryCompanyFromContext извлекает имя компании из контекста, если она есть
func TryCompanyFromContext(ctx context.Context) (string, bool) {
	company, ok := ctx.Value(companyCtxKey{}).(string)

	return company, ok
}

type companyAliasCtxKey struct{}

// ContextWithCompanyAlias сохранить алиса компании (человеческое название) в контексте
func ContextWithCompanyAlias(ctx context.Context, alias string) context.Context {
	return context.WithValue(ctx, companyAliasCtxKey{}, alias)
}

// CompanyAliasFromContext извлечь алиас компании из контекста
func CompanyAliasFromContext(ctx context.Context) string {
	val := ctx.Value(companyAliasCtxKey{})
	if val == nil {
		return ""
	}

	return val.(string)
}

// CompanyNameFromContext извлечь финальное имя компании из контекста
// Если у компании есть Alias, то извлекает его, если нет - возвращает сгенерированное имя компании (CompanyFromContext)
// которое используется в PG схемах
func CompanyNameFromContext(ctx context.Context) string {
	cval := ctx.Value(companyAliasCtxKey{})
	if cval == nil {
		return CompanyFromContext(ctx)
	}
	val, ok := cval.(string)
	if !ok {
		return CompanyFromContext(ctx)
	}

	if val == "" {
		return CompanyFromContext(ctx)
	}

	return val
}

type userIDCtxKey struct{}

// ContextWithUserID сохраняет id пользователя в контексте
func ContextWithUserID(ctx context.Context, userID uuid.UUID) context.Context {
	return context.WithValue(ctx, userIDCtxKey{}, userID)
}

// UserIDFromContext извлекает id пользователя из контекста
func UserIDFromContext(ctx context.Context) uuid.UUID {
	return ctx.Value(userIDCtxKey{}).(uuid.UUID)
}

// TryUserIDFromContext извлекает id пользователя из контекста если оно там есть
func TryUserIDFromContext(ctx context.Context) (uuid.UUID, bool) {
	userID, ok := ctx.Value(userIDCtxKey{}).(uuid.UUID)

	return userID, ok
}

type isAdminCtxKey struct{}

// ContextWithIsAdmin сохраняет признак того, что пользователь администратор
func ContextWithIsAdmin(ctx context.Context, isAdmin bool) context.Context {
	return context.WithValue(ctx, isAdminCtxKey{}, isAdmin)
}

// IsAdminFromContext извлекает из контекста признак того, что пользователь администратор
func IsAdminFromContext(ctx context.Context) bool {
	res, ok := ctx.Value(isAdminCtxKey{}).(bool)

	return ok && res
}

type isPortalUserCtxKey struct{}

// ContextWithIsPortalUser сохраняет признак того, что пользователь портальный
func ContextWithIsPortalUser(ctx context.Context, isPortal bool) context.Context {
	return context.WithValue(ctx, isPortalUserCtxKey{}, isPortal)
}

// IsPortalUserFromContext извлекает из контекста признак того, что пользователь портальный
func IsPortalUserFromContext(ctx context.Context) bool {
	res, ok := ctx.Value(isPortalUserCtxKey{}).(bool)

	return ok && res
}

type timestamp struct{}

// ContextWithTimestamp создает новый контекст с временной меткой
func ContextWithTimestamp(ctx context.Context, ts time.Time) context.Context {
	return context.WithValue(ctx, timestamp{}, ts)
}

// TimestampFromContext извлекает временную метку из контекста
func TimestampFromContext(ctx context.Context) time.Time {
	return ctx.Value(timestamp{}).(time.Time)
}

type userCompanyIDCtxKey struct{}

// ContextWithUserCompanyID сохраняет id компании с которой связан текущий пользователь в контексте
func ContextWithUserCompanyID(ctx context.Context, userCompanyID uuid.UUID) context.Context {
	return context.WithValue(ctx, userCompanyIDCtxKey{}, userCompanyID)
}

// UserCompanyIDFromContext извлекате из контекста ид компании с которой связан текущий пользователь
func UserCompanyIDFromContext(ctx context.Context) uuid.UUID {
	return ctx.Value(userCompanyIDCtxKey{}).(uuid.UUID)
}

type editionCtxKey struct{}

// ContextWithEdition сохраняет edition компании в контексте
func ContextWithEdition(ctx context.Context, edition pkgEdition.Edition) context.Context {
	return context.WithValue(ctx, editionCtxKey{}, edition)
}

// EditionFromContext извлекает edition компании из контекста
func EditionFromContext(ctx context.Context) pkgEdition.Edition {
	return ctx.Value(editionCtxKey{}).(pkgEdition.Edition)
}

// TryEditionFromContext извлекает edition компании из контекста, если он есть
func TryEditionFromContext(ctx context.Context) (pkgEdition.Edition, bool) {
	edition, ok := ctx.Value(editionCtxKey{}).(pkgEdition.Edition)
	return edition, ok
}

type gatewayCtxKey struct{}

type gatewayRecord struct {
	name string
	next *gatewayRecord
}

// ContextWithGateway добавляет запись о точке входа/проксирования в контекст
func ContextWithGateway(ctx context.Context, gateway string) context.Context {
	head, _ := ctx.Value(gatewayCtxKey{}).(*gatewayRecord)
	head = &gatewayRecord{gateway, head}
	return context.WithValue(ctx, gatewayCtxKey{}, head)
}

// GatewaysFromContext извлекает последовательность шлюзов и прокси запроса из контекста
func GatewaysFromContext(ctx context.Context) []string {
	var res []string
	for head, _ := ctx.Value(gatewayCtxKey{}).(*gatewayRecord); head != nil; head = head.next {
		res = append(res, head.name)
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

// InjectGateways записать шлюзы в заголовки для последующего извлечения методом ExtractGateways
func InjectGateways(ctx context.Context, key string, headers OutgoingHeaders) {
	for _, gateway := range GatewaysFromContext(ctx) {
		headers.Add(key, gateway)
	}
}

// ExtractGateways извлечь шлюзы из заголовков, записанные методом InjectGateways
func ExtractGateways(ctx context.Context, key string, headers IncomingHeaders) context.Context {
	head, _ := ctx.Value(gatewayCtxKey{}).(*gatewayRecord)
	for _, gateway := range headers.Values(key) {
		head = &gatewayRecord{gateway, head}
	}
	return context.WithValue(ctx, gatewayCtxKey{}, head)
}
