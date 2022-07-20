package md

import (
	"context"

	"git.elewise.com/elma365/common/pkg/i18n"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/opentracing/opentracing-go"
)

// Copy возвращает новый контекст с информацией из текущего
func Copy(ctx context.Context) context.Context {
	res := context.Background()

	res = ctxzap.ToContext(res, ctxzap.Extract(ctx))
	res = opentracing.ContextWithSpan(res, opentracing.SpanFromContext(ctx))

	if id, ok := TryUserIDFromContext(ctx); ok {
		res = ContextWithUserID(res, id)
	}
	if company, ok := TryCompanyFromContext(ctx); ok {
		res = ContextWithCompany(res, company)
	}
	res = ContextWithIsAdmin(res, IsAdminFromContext(ctx))
	res = ContextWithIsPortalUser(res, IsPortalUserFromContext(ctx))
	res = i18n.ContextWithLang(res, i18n.LangFromContext(ctx))

	if edition, ok := TryEditionFromContext(ctx); ok {
		res = ContextWithEdition(res, edition)
	}

	alias := CompanyAliasFromContext(ctx)
	if len(alias) > 0 {
		res = ContextWithCompanyAlias(res, alias)
	}
	if kv, ok := ctx.Value(kvContextKey{}).(*kvRecord); ok {
		res = context.WithValue(res, kvContextKey{}, kv)
	}

	return res
}
