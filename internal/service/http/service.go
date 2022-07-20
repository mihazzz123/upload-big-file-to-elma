package http

import (
	"encoding/json"
	"net/http"

	"github.com/mihazzz123/upload-big-file-to-elma/internal/action"

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

	r.Get("/test", hs.testHandler())
	r.Post("/upload", hs.uploadFilelink())

	return r
}

func (hs Service) testHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hs.respond(w, r, http.StatusOK, "test")
	})
}

func (hs Service) uploadFilelink() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hs.respond(w, r, http.StatusOK, "upload file link")

	})
}

func (hs Service) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	hs.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (hs Service) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

// модель -/model
//service/http(роуты/хендлеры) -> actions(загрузить файл) -> adapter(отправить файл)
