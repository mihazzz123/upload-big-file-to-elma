package sqldb

import (
	"context"
)

// Statement подготовленный запрос в БД.
type Statement interface {
	// Get получает одну запись из БД.
	Get(ctx context.Context, dst interface{}, args ...interface{}) error
	// Select получает несколько записей из БД.
	Select(ctx context.Context, dst interface{}, args ...interface{}) error
	// Exec выполняет запрос в БД (insert/update/delete).
	Exec(ctx context.Context, args ...interface{}) (int64, error)

	// Close закрывает подготовленный запрос.
	Close(ctx context.Context) error
}
