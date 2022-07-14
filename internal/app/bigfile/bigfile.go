package bigfile

import (
	"net/http"
	"upload-big-file-to-elma/internal/app/store/bigfilestore"
)

func Start(config *Config) error {
	store := bigfilestore.New()
	srv := NewServer(store)
	return http.ListenAndServe(config.BindAddress, srv)
}
