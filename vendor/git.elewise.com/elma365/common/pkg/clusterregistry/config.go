package clusterregistry

// Config конфигурация реестра арендаторов и их вычислительных кластеров.
type Config interface {
	// GetClusterRegistryConnectionString возвращает строку подключения к реестру.
	GetClusterRegistryConnectionString() string
	// IsClusterRegistryEnabled проверяет включен ли поддержка реестр арендаторов и их вычислительных кластеров.
	IsClusterRegistryEnabled() bool
}
