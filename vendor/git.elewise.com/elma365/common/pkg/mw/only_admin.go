package mw

import (
	"git.elewise.com/elma365/common/pkg/md"

	"net/http"
)

// OnlyAdmin запрещает доступ к роуту запросам без признака администратора в контексте
func OnlyAdmin(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if !md.IsAdminFromContext(r.Context()) {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
