package action

//go:generate mkdir -p dimock
//go:generate mockgen -source=adaptor_interfaces.go -destination=dimock/adaptor_interfaces_mock.go -package=dimock -imports uuid=github.com/satori/go.uuid

import (
	"github.com/mihazzz123/upload-big-file-to-elma/internal/config"
)

// DIContainer предоставляет контейнер с методами получения необходимых адаптеров
type DIContainer interface {
	GetConfig() *config.Config
}
