package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mihazzz123/upload-big-file-to-elma/internal/action"
	"github.com/mihazzz123/upload-big-file-to-elma/internal/action/di"
	"github.com/mihazzz123/upload-big-file-to-elma/internal/model"
	"go.uber.org/zap"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

// Service новый http сервис
type Service struct {
	di di.Container

	baseCtx context.Context
	workers map[string]context.CancelFunc
}

const (
	ctxKeyPartsUpload ctxKey = iota
)

type ctxKey int8

var (
	errQueryParametrID = errors.New("query parameter uuid is required")
)

// NewService новый http
func NewService(di di.Container) http.Handler {
	hs := Service{
		di:      di,
		baseCtx: context.Background(),
		workers: make(map[string]context.CancelFunc),
	}
	return hs.newRouter()
}

func (hs Service) newRouter() chi.Router {
	r := chi.NewRouter()

	r.Get("/test", hs.testHandler())
	r.Post("/upload", hs.uploadFile())
	r.Get("/cancel", hs.cancelUploadFile())

	return r
}

func (hs Service) testHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hs.respond(w, r, http.StatusOK, "test")
	})
}

func (hs Service) uploadFile() http.HandlerFunc {
	type response struct {
		Uuid string `json:"uuid"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bf := &model.Bigfile{}

		if err := json.NewDecoder(r.Body).Decode(bf); err != nil {
			hs.error(w, r, http.StatusBadRequest, err)
			return
		}

		bf, err := action.CreateBigfile(bf, hs.di)
		if err != nil {
			hs.error(w, r, http.StatusInternalServerError, err)
			return
		}

		bfUuid := bf.Uuid.String()
		ctx, cancelFunc := context.WithCancel(hs.baseCtx)
		hs.workers[bfUuid] = cancelFunc

		go func(ctx context.Context, id string) {
			//for {
			select {
			case <-ctx.Done():
				zap.L().Info("worker stopped", zap.String("uuid", bfUuid))
				return
			default:
				bf, err = action.DownloadByLink(bf, hs.di)
				if err != nil {
					zap.L().Error("downloadByLink error", zap.Error(err))
					return
				}
				if err = action.SendFileToElma(r.Context(), bf, hs.di); err != nil {
					zap.L().Error("SendFileToElma error", zap.Error(err))
					return
				}
				time.Sleep(time.Second)
				fmt.Printf("worker '%s' is working\n", id)
			}
			//}
		}(ctx, bfUuid)

		resp := &response{
			Uuid: bfUuid,
		}
		hs.respond(w, r, http.StatusOK, resp)
	})
}

func (hs Service) cancelUploadFile() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		bfUuid := r.URL.Query().Get("uuid")

		if bfUuid != "" {
			cancelFunc, found := hs.workers[bfUuid]
			if found {
				cancelFunc()

				hs.respond(w, r, http.StatusOK, nil)
				return
			} else {
				hs.error(w, r, http.StatusNotFound, nil)
				return
			}
		}
		hs.error(w, r, http.StatusBadRequest, errQueryParametrID)
	}
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
