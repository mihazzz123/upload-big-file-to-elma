package di

import (
	"context"
	"github.com/mihazzz123/upload-big-file-to-elma/internal/config"
	"github.com/mihazzz123/upload-big-file-to-elma/internal/model"
)

//go:generate mkdir -p dimock
//go:generate mockgen -source=adaptor_interfaces.go -destination=dimock/adaptor_interfaces_mock.go -package=dimock -imports uuid=github.com/satori/go.uuid

// Container предоставляет контейнер с методами получения необходимых адаптеров
type Container interface {
	GetConfig() *config.Config
	GetBigfileRepository() BigfileRepository
}

// BigfileRepository ...
type BigfileRepository interface {
	DownloadByLink(bigfile *model.Bigfile) error
	SendFileToElma(ctx context.Context, bigfile *model.Bigfile, cfgElma *config.Config) error
}
