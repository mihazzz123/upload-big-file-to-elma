package sqldb

import (
	"time"
)

// Config конфигурация подключения к БД.
type Config interface {
	// GetDbConnectionString получает строку подключения.
	GetDbConnectionString() string
	// GetDbMinOpenConnections получает минимальное количество одновременно открытых подключений.
	GetDbMinOpenConnections() int32
	// GetDbMaxOpenConnections получает максимальное количество одновременно открытых подключений.
	GetDbMaxOpenConnections() int32
	// GetDbConnectionTimeout получает максимальное время установки соединения с базой данных.
	GetDbConnectionTimeout() time.Duration
	// GetDbMaxConnectionLifetime получает максимальное время в течении которого неиспользуемое подключение может быть переиспользовано.
	GetDbMaxConnectionLifetime() time.Duration
}
