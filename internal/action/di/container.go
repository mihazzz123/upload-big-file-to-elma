package di

import "github.com/mihazzz123/upload-big-file-to-elma/internal/model"

//go:generate mkdir -p dimock
//go:generate mockgen -source=adaptor_interfaces.go -destination=dimock/adaptor_interfaces_mock.go -package=dimock -imports uuid=github.com/satori/go.uuid

import (
	"github.com/mihazzz123/upload-big-file-to-elma/internal/config"
)

// Container предоставляет контейнер с методами получения необходимых адаптеров
type Container interface {
	GetConfig() *config.Config
	GetBigfileRepository() BigfileRepository
}

// BigfileRepository ...
type BigfileRepository interface {
	DownloadByLink(bigfile *model.Bigfile) error
	SendFileToElma(bigfile *model.Bigfile) error
}
