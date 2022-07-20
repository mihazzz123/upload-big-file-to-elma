package mw

//go:generate ../../tooling/bin/minimock -g -i ../company.RedisStorage

import (
	"net/http"

	"git.elewise.com/elma365/common/pkg/company"
	"git.elewise.com/elma365/common/pkg/md"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

// WithCompanyAlias middleware которая заменяет алиса компании на её имя, если находится алиас
// требует наличие компании (CompanyFromContext()) в контексте
func WithCompanyAlias(red company.RedisStorage, namespacePrefix string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			l := ctxzap.Extract(ctx)
			sp := opentracing.SpanFromContext(ctx)

			requestCompany := md.CompanyFromContext(ctx)
			companyName, err := company.FromAlias(red, namespacePrefix, requestCompany)
			if err != nil {
				sp.SetTag("company", requestCompany)
				// так хотябы продолжат прямые ссылки работать, без алиасов
				next.ServeHTTP(w, r)
				return
			}

			// если нашелся alias, значит надо заменить начальный хэдер
			if len(companyName) > 0 {
				l.Debug("found alias", zap.String("alias", requestCompany), zap.String("company", companyName))
				ctx = md.ContextWithCompany(ctx, companyName)
				ctx = md.ContextWithCompanyAlias(ctx, requestCompany)
				sp.SetTag("company_alias", requestCompany)
				sp.SetTag("company", companyName)
				r = r.WithContext(ctx)
			}

			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
