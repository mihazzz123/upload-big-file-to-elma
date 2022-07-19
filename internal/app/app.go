// Package app configures and runs application.
package app

import (
	"github.com/gin-gonic/gin"
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

}
