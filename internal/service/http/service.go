package http

import (
	"encoding/json"
	"github.com/mihazzz123/upload-big-file-to-elma/internal/action/di"
	"github.com/mihazzz123/upload-big-file-to-elma/internal/model"
	"net/http"

	"github.com/mihazzz123/upload-big-file-to-elma/internal/action"

	"github.com/go-chi/chi"
)

// Service новый http сервис
type Service struct {
	di di.Container
}

// NewService новый http
func NewService(di di.Container) http.Handler {
	hs := Service{di: di}
	return hs.newRouter()
}

func (hs Service) newRouter() chi.Router {
	r := chi.NewRouter()

	r.Get("/test", hs.testHandler())
	r.Post("/upload", hs.uploadFile())

	return r
}

func (hs Service) testHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hs.respond(w, r, http.StatusOK, "test")
	})
}

func (hs Service) uploadFile() http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bf := &model.Bigfile{}

		if err := json.NewDecoder(r.Body).Decode(bf); err != nil {
			hs.error(w, r, http.StatusBadRequest, err)
			return
		}

		if err := action.CreateBigfile(bf, hs.di); err != nil {
			hs.error(w, r, http.StatusInternalServerError, err)
			return
		}

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
