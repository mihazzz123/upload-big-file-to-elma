package action

//go:generate mkdir -p dimock
//go:generate mockgen -source=adaptor_interfaces.go -destination=dimock/adaptor_interfaces_mock.go -package=dimock -imports uuid=github.com/satori/go.uuid

import (
	"git.elewise.com/elma365/upload-big-file-elma365/internal/config"
)

// DIContainer предоставляет контейнер с методами получения необходимых адаптеров
type DIContainer interface {
	GetConfig() *config.Config
}
