package bigfile

import "net/http"

func Start(config *Config) error {
	store := store.
	//srv := NewServer(store)
	return http.ListenAndServe(config.BindAddress)
}