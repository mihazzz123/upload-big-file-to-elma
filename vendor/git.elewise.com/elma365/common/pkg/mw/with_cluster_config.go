package mw

import (
	"net/http"

	"git.elewise.com/elma365/common/pkg/clusterregistry"
	"git.elewise.com/elma365/common/pkg/errs"
	"git.elewise.com/elma365/common/pkg/md"

	"github.com/pkg/errors"
)

// WithClusterConfig добавляет в контекст запросов конфигурацию виртуального вычислительного кластера арендатора.
func WithClusterConfig(registry clusterregistry.TenantRegistry) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			company, ok := md.TryCompanyFromContext(ctx)
			if !ok {
				http.Error(
					w,
					errors.Wrap(errs.Precondition, "company name not defined").Error(),
					http.StatusInternalServerError,
				)
			}
			clusterConfig, err := registry.GetTenantClusterConfig(ctx, company)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			ctx = clusterregistry.ContextWithClusterConfig(ctx, clusterConfig)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
