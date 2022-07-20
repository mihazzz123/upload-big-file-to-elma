package http

import (
	"net/http"

	"git.elewise.com/elma365/upload-big-file-elma365/internal/action"

	"github.com/go-chi/chi"
)

// Service новый http сервис
type Service struct {
	di action.DIContainer
}

// NewService новый http
func NewService(di action.DIContainer) http.Handler {
	hs := Service{di: di}
	return hs.newRouter()
}

func (hs Service) newRouter() chi.Router {
	r := chi.NewRouter()

	return r
}
