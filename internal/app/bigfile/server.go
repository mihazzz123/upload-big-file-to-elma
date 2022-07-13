package bigfile

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mihazzz123/upload-big-file-to-elma/internal/app/store"
	"github.com/sirupsen/logrus"
	"net/http"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

func NewServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}
	s.ConfigureRouter()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) ConfigureRouter() {
	s.router.HandleFunc("/test", s.handlerTest()).Methods("POST", "GET")
}

func (s *server) handlerTest() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			s.respond(w, r, http.StatusOK, "Hello world!!!")
			return
		}
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
