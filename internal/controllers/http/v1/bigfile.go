package v1

import (
	"github.com/gin-gonic/gin"
	"upload-big-file-to-elma/internal/usecase"
	"upload-big-file-to-elma/pkg/logger"
)

type bigfileRoutes struct {
	t usecase.Bigfile
	l logger.Interface
}

func newBigfileRoutes(handler *gin.RouterGroup, l logger.Interface, b usecase.Bigfile) {

}
