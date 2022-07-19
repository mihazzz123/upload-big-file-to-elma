// Package app configures and runs application.
package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
	"upload-big-file-to-elma/config"
	v1 "upload-big-file-to-elma/internal/controllers/http/v1"
	"upload-big-file-to-elma/internal/usecase"
	"upload-big-file-to-elma/pkg/httpserver"
	"upload-big-file-to-elma/pkg/logger"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Use case
	bigfileUseCase := usecase.New()

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, l, bigfileUseCase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run -- httpServer.Notify: %w", err))
	}

	// Shutdown
	if err := httpServer.Shutdown(); err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
