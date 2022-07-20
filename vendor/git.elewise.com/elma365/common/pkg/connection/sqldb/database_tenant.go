package sqldb

import (
	"context"
)

// TenantDatabase "кластеризуемая" база данных.
//
// В рамках "кластеризуемой" базы данных схема "public" может располагаться на выделенном для нее сервере и быть
// не доступной в рамках прямого подключения.
type TenantDatabase interface {
	Database

	// Public создает подключение к БД содержащей схему "public".
	Public(ctx context.Context) (Database, error)
}
