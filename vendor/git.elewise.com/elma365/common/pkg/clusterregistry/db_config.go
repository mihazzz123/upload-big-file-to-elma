package clusterregistry

import (
	"time"
)

// PublicDBConfig конфигурация подключения к базе данных содержащей схему "public".
type PublicDBConfig interface {
	// GetPublicDbConnectionString получает строку подключения.
	GetPublicDbConnectionString() string
	// GetPublicDbMinOpenConnections получает минимальное количество одновременно открытых подключений.
	GetPublicDbMinOpenConnections() int32
	// GetPublicDbMaxOpenConnections получает максимальное количество одновременно открытых подключений.
	GetPublicDbMaxOpenConnections() int32
	// GetPublicDbConnectionTimeout получает максимальное время установки соединения с базой данных.
	GetPublicDbConnectionTimeout() time.Duration
	// GetPublicDbMaxConnectionLifetime получает максимальное время в течении которого неиспользуемое подключение может быть переиспользовано.
	GetPublicDbMaxConnectionLifetime() time.Duration
}
