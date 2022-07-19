// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"upload-big-file-to-elma/internal/usecase"
	"upload-big-file-to-elma/pkg/logger"
)

// NewRouter -.
// Swagger spec:
// @title       Go Clean Template API
// @description Using a translation service as an example
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(handler *gin.Engine, l logger.Interface, b usecase.Bigfile) {
	// Optins -.
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// UploadFile
	h := handler.Group("/upload")
	{
		newBigfileRoutes(h, l, b)
	}
}
