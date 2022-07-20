package adaptor

import (
	"git.elewise.com/elma365/upload-big-file-elma365/internal/config"
	"github.com/pkg/errors"
)

// DIContainer dependency injection
type DIContainer struct {
	cfg         *config.Config
}

// NewDIContainer возвращает реализацию action.DIContainer поверх подключений
//
// Сам контейнер хранит в себе подключения и создаёт адаптеры по запросу. Контейнер должен создаваться один раз
// при старте приложения и передаваться в сервисы для инстанцирования действий.
func NewDIContainer(cfg config.Config) (*DIContainer, error) {

	return &DIContainer{
		&cfg,
		
		
		
		
		
	}, nil
}

// GetConfig возвращает конфиг
func (di *DIContainer) GetConfig() *config.Config {
	return di.cfg
}
