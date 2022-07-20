package mw

import (
	"net/http"

	"git.elewise.com/elma365/common/pkg/md"
)

// ClientUserAgentKey ключ для user agent клиента
const ClientUserAgentKey = "clientUserAgent"

// WithUserAgent добавляет user agent клиента в контекст
func WithUserAgent(next http.Handler) http.Handler {
	userAgentHeader := http.CanonicalHeaderKey("User-Agent")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ua := r.Header.Get(userAgentHeader)
		ctx = md.AddKV(ctx, ClientUserAgentKey, ua)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
