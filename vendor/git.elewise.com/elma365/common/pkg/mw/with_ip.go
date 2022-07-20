package mw

import (
	"net"
	"net/http"
	"strings"

	"git.elewise.com/elma365/common/pkg/md"
)

// ClientIPKey ключ ip-адреса клиента
const ClientIPKey = "clientIP"

// WithIP добавляет ip-адрес клиента в контекст
func WithIP(next http.Handler) http.Handler {
	xForwardedFor := http.CanonicalHeaderKey("X-Forwarded-For")
	xRealIP := http.CanonicalHeaderKey("X-Real-IP")
	realIP := func(r *http.Request) string {
		var ip string

		if xrip := r.Header.Get(xRealIP); xrip != "" {
			ip = xrip
		} else if xff := r.Header.Get(xForwardedFor); xff != "" {
			i := strings.Index(xff, ", ")
			if i == -1 {
				i = len(xff)
			}
			ip = xff[:i]
		}
		if ip != "" {
			return ip
		}

		return r.RemoteAddr
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ip := realIP(r)
		if host, _, err := net.SplitHostPort(ip); err == nil {
			ip = host
		}
		ctx = md.AddKV(ctx, ClientIPKey, ip)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
