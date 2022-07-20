package clusterregistry

//go:generate ../../tooling/bin/easyjson $GOFILE

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

// ClusterRegistry реестр конфигураций вычислительных кластеров.
type ClusterRegistry interface {
	// GetClusterConfig возвращает конфигурацию вычислительного кластера по его идентификатору.
	//
	// Если конфигурация по идентификатору не найдена, то возвращается (nil, errs.NotFound).
	GetClusterConfig(ctx context.Context, id uuid.UUID) (*ClusterConfig, error)
	// ListAllClustersConfigs возвращает список всех зарегистрированных конфигураций вычислительных кластеров.
	ListAllClustersConfigs(ctx context.Context) ([]ClusterConfig, error)
	// SaveClusterConfig сохраняет конфигурацию вычислительного кластера в хранилище.
	//
	// Если кластер с точно таким же идентификатором уже существует, то происходит его обновление и
	// возвращается обновленная конфигурация.
	//
	// Если создается кластер с уже существующим именем, то возвращается (nil, errs.Collision).
	SaveClusterConfig(ctx context.Context, clusterConfig *ClusterConfig) (*ClusterConfig, error)
}

// ClusterConfig конфигурация вычислительного кластера.
//
// easyjson:json
type ClusterConfig struct {
	// ID уникальный идентификатор кластера.
	ID uuid.UUID `json:"__id"`
	// Name название вычислительного кластера.
	Name string `json:"__description"`
	// SQLDatabase конфигурация подключения к базе данных.
	SQLDatabase *SQLDatabasesConfig `json:"sql_database"`
	// CreatedAt дата создания.
	CreatedAt time.Time `json:"__createdAt"`
	// UpdatedAt дата обновления.
	UpdatedAt time.Time `json:"__updatedAt"`
}

// SQLDatabasesConfig конфигурация подключений к базе данных и ее репликам.
//
// easyjson:json
type SQLDatabasesConfig struct {
	// Main конфигурация подключения к основной базе данных.
	Main *SQLDBConnectionConfig `json:"main"`
	// Replicas конфигурация подключения к RO-репликам базы данных.
	Replicas *SQLDBConnectionConfig `json:"replicas"`
}

// SQLDBConnectionConfig конфигурация подключения к базе данных SQL.
//
// easyjson:json
type SQLDBConnectionConfig struct {
	// ConnectionString Строка подключения к базе данных.
	ConnectionString string `json:"connectionString"`
	// MinOpenConnections минимальное количество одновременно открытых подключений.
	MinOpenConnections int32 `json:"minOpenConnections"`
	// MaxOpenConnections Максимальное количество одновременно открытых подключений.
	MaxOpenConnections int32 `json:"maxOpenConnections"`
	// ConnectionTimeout максимальное время установки соединения с базой данных.
	ConnectionTimeout time.Duration `json:"connectionTimeout"`
	// MaxConnectionLifetime Максимальное время в течении которого неиспользуемое подключение может быть переиспользовано.
	MaxConnectionLifetime time.Duration `json:"maxConnectionLifetime"`
}

// Value implements sql.Valuer interface
func (dbCfg SQLDatabasesConfig) Value() (driver.Value, error) {
	value, err := json.Marshal(dbCfg)
	return string(value), errors.WithStack(err)
}

// Scan implements sql.Scanner interface.
func (dbCfg *SQLDatabasesConfig) Scan(src interface{}) error {
	switch src := src.(type) {
	case string:
		data := []byte(src)
		if err := json.Unmarshal(data, dbCfg); err != nil {
			return errors.WithStack(err)
		}
	case []byte:
		if err := json.Unmarshal(src, dbCfg); err != nil {
			return errors.WithStack(err)
		}
	default:
		return errors.Errorf("DBConfig.Scan: cannot scan type %T into DBConfig", src)
	}
	return nil
}

// GetDbConnectionString получает строку подключения.
func (dbCfg SQLDBConnectionConfig) GetDbConnectionString() string {
	return dbCfg.ConnectionString
}

// GetDbMinOpenConnections получает минимальное количество одновременно открытых подключений.
func (dbCfg SQLDBConnectionConfig) GetDbMinOpenConnections() int32 {
	if dbCfg.MinOpenConnections <= 0 {
		return 1
	}
	return dbCfg.MinOpenConnections
}

// GetDbMaxOpenConnections получает максимальное количество одновременно открытых подключений.
func (dbCfg SQLDBConnectionConfig) GetDbMaxOpenConnections() int32 {
	if dbCfg.MaxOpenConnections <= 0 {
		return 10
	}
	return dbCfg.MaxOpenConnections
}

// GetDbConnectionTimeout получает максимальное время установки соединения с базой данных.
func (dbCfg SQLDBConnectionConfig) GetDbConnectionTimeout() time.Duration {
	if dbCfg.ConnectionTimeout < time.Second {
		return 30 * time.Second
	}
	return dbCfg.ConnectionTimeout
}

// GetDbMaxConnectionLifetime получает максимальное время в течении которого неиспользуемое подключение может быть переиспользовано.
func (dbCfg SQLDBConnectionConfig) GetDbMaxConnectionLifetime() time.Duration {
	return dbCfg.MaxConnectionLifetime
}
