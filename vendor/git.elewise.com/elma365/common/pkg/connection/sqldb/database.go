package sqldb

import (
	"context"
)

// Type тип базы данных.
type Type int8

const (
	// PostgreSQL Postgre SQL Server.
	PostgreSQL Type = iota
)

// Database база данных.
type Database interface {
	Querier

	// Acquire создает подключение к БД или получает его из пула.
	Acquire(ctx context.Context) (Connection, error)

	// WithAcquire выполняет указанный делегат в рамках отдельного подключения к БД.
	WithAcquire(ctx context.Context, delegate func(context.Context, Connection) error) error

	// WithTx выполняет указанный делегат в рамках транзакции.
	WithTx(ctx context.Context, delegate func(context.Context, ReduceTransaction) error) error
	// WithTxOptions выполняет указанный делегат в рамках транзакции с дополнительными опциями.
	WithTxOptions(ctx context.Context, opt TransactionOptions, delegate func(context.Context, ReduceTransaction) error) error

	// Ping проверяет подключение к БД.
	Ping() error

	// Close закрывает подключение к БД.
	Close()
}
