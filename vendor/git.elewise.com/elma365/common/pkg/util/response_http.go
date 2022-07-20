package util

import (
	"context"
	"encoding/json"
	"net/http"

	"git.elewise.com/elma365/common/pkg/errs"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
)

type responseConfig struct {
	logger *zap.Logger
	status int
	err    error
}

// ResponseHTTPOption опция ответа
type ResponseHTTPOption interface {
	Apply(responseConfig) responseConfig
}

type responseHTTPOptionFunc func(responseConfig) responseConfig

func (fn responseHTTPOptionFunc) Apply(cfg responseConfig) responseConfig {
	return fn(cfg)
}

// ResponseWithStatus ответить с определённым статусом (по умолчанию 200)
func ResponseWithStatus(status int) ResponseHTTPOption {
	return responseHTTPOptionFunc(func(cfg responseConfig) responseConfig {
		cfg.status = status

		return cfg
	})
}

// ResponseWithError вернуть ошибку, если она не пустая
func ResponseWithError(ctx context.Context, err error) ResponseHTTPOption {
	return responseHTTPOptionFunc(func(cfg responseConfig) responseConfig {
		cfg.logger = ctxzap.Extract(ctx)
		cfg.err = err

		return cfg
	})
}

func newResponseConfig(opts []ResponseHTTPOption) responseConfig {
	cfg := responseConfig{
		status: http.StatusOK,
	}

	for _, opt := range opts {
		cfg = opt.Apply(cfg)
	}

	return cfg
}

// ResponseHTTP ответить json на http-запрос
func ResponseHTTP(w http.ResponseWriter, data interface{}, opts ...ResponseHTTPOption) {
	cfg := newResponseConfig(opts)
	if cfg.err != nil {
		errs.Handle(cfg.logger, w, cfg.err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(cfg.status)
	je := json.NewEncoder(w)
	_ = je.Encode(data)
}
