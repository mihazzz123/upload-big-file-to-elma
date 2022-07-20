package clusterregistry

import (
	"context"

	"git.elewise.com/elma365/common/pkg/connection/sqldb"
)

// Type тип хранилища реестра арендаторов.
type Type int8

const (
	// PostgreSQL Postgre SQL Server.
	PostgreSQL Type = iota
)

// Registry реестр арендаторов и их вычислительных кластеров.
// Объединяет функционал ClusterRegistry и TenantRegistry.
type Registry interface {
	ClusterRegistry
	TenantRegistry

	// GetPublicDBConfig возвращает конфигурацию подключения к БД содержащей схему "public".
	GetPublicDBConfig(ctx context.Context) (sqldb.Config, error)
	// Close закрывает реестр.
	Close()
	// Ping проверяет доступность реестра.
	Ping() error
}

type clusterConfigCtxKey struct{}

// ContextWithClusterConfig сохраняет конфигурацию виртуального вычислительного кластера арендатора в контекст.
func ContextWithClusterConfig(ctx context.Context, cfg *ClusterConfig) context.Context {
	return context.WithValue(ctx, clusterConfigCtxKey{}, cfg)
}

// TryClusterConfigFromContext извлекается конфигурация виртуального вычислительного кластера арендатора из контекста.
func TryClusterConfigFromContext(ctx context.Context) (*ClusterConfig, bool) {
	clusterConfig, ok := ctx.Value(clusterConfigCtxKey{}).(*ClusterConfig)
	return clusterConfig, ok
}
