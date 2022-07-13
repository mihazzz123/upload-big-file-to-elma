package bigfile

import (
	"net/http"
)

func Start(config *Config) error {
	srv := NewServer()
	return http.ListenAndServe(config.BindAddress, srv)
}
