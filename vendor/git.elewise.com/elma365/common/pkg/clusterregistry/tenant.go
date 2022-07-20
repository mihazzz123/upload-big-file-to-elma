package clusterregistry

import (
	"context"
)

// TenantRegistry реестр конфигураций арендаторов.
type TenantRegistry interface {
	// GetTenantClusterConfig возвращает конфигурацию кластера арендатора.
	// Если конфигурация кластера для указанного арендатора отсутствует, то будет возвращена конфигурацию по-умолчанию.
	GetTenantClusterConfig(ctx context.Context, tenant string) (*ClusterConfig, error)
	// SaveTenantClusterConfig сохраняет указанную конфигурацию кластера для арендатора.
	// Если кластера ранее не существовало, то он будет создан и сохранен в хранилище.
	// Если для указанного арендатора уже существовала привязка к кластеру, то она будет обновлена.
	SaveTenantClusterConfig(ctx context.Context, tenant string, clusterConfig *ClusterConfig) error
}
