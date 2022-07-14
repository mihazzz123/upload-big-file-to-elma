package bigfile

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"time"
	"upload-big-file-to-elma/internal/app/model"
	"upload-big-file-to-elma/internal/app/store"
)

const (
	ctxKeyRequestID ctxKey = iota
)

var (
	errRequiredHeaderMissing = errors.New("the Content-Type header was expected to be set to multipart/form-data or application/json")
)

type ctxKey int8

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
	s.router.Use(s.setRequestID)
	s.router.Use(s.logRequest)
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))

	s.router.HandleFunc("/test", s.handlerTest()).Methods("GET")
	s.router.HandleFunc("/upload", s.handlerUploadFile()).Methods("POST")
}

func (s *server) handlerTest() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			s.respond(w, r, http.StatusOK, "Hello world!!!")
			return
		}

		if r.Method == http.MethodPost {
			s.respond(w, r, http.StatusOK, "POST Hello world!!!")
			return
		}
	}
}

func (s *server) handlerUploadFile() http.HandlerFunc {
	type request struct {
		Name string `json:"name"`
		Link string `json:"link"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		ctype := r.Header.Get("Content-Type")

		if ctype == "multipart/form-data; boundary=-------------573cf973d5228" {
			err := r.ParseMultipartForm(32 << 20)
			if err != nil {
				s.error(w, r, http.StatusBadRequest, err)
			}

			file, handler, err := r.FormFile("file")
			if err != nil {
				s.error(w, r, http.StatusBadRequest, err)
				return
			}
			defer file.Close()

			buf := bytes.NewBuffer(nil)
			if _, err := io.Copy(buf, file); err != nil {
				s.error(w, r, http.StatusBadRequest, err)
				return
			}
			bigFile := &model.Bigfile{
				Name:      time.Now().Format(time.RFC3339Nano) + handler.Filename,
				Size:      int(handler.Size),
				FileBytes: buf.Bytes(),
			}

			s.store.BigFile().SaveLocal(bigFile)
			dst, err := os.Create("tempfile/" + time.Now().Format(time.RFC3339Nano) + handler.Filename)
			if err != nil {
				s.error(w, r, http.StatusBadRequest, err)
				return
			}

			defer dst.Close()
			// SavaLocal ...
			if _, err := io.Copy(dst, file); err != nil {
				s.error(w, r, http.StatusBadRequest, err)
				return
			}

			s.respond(w, r, http.StatusOK, nil)
			return
		}

		if ctype == "application/json" {
			req := &request{}
			if err := json.NewDecoder(r.Body).Decode(req); err != nil {
				s.error(w, r, http.StatusBadRequest, err)
				return
			}

			s.respond(w, r, http.StatusOK, nil)
			return
		}

		s.error(w, r, http.StatusBadRequest, errRequiredHeaderMissing)
		return
	}
}

func (s *server) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyRequestID, id)))
	})
}

func (s *server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
			"remote_id":   r.Context().Value(ctxKeyRequestID),
		})
		logger.Infof("started % %", r.Method, r.RequestURI)

		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		logger.Infof(
			"complited with %d %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Now().Sub(start))
	})
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
