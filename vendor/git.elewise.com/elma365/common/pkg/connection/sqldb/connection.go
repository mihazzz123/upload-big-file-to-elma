package sqldb

import (
	"context"
)

// Connection подключение к базе данных.
type Connection interface {
	QuerierEx

	// BeginTx открывает транзакцию в БД.
	BeginTx(ctx context.Context) (Transaction, error)
	// BeginTxOptions открывает транзакцию в БД с дополнительными опциями.
	BeginTxOptions(ctx context.Context, opt TransactionOptions) (Transaction, error)

	// WithTx выполняет указанный делегат в рамках транзакции.
	WithTx(ctx context.Context, delegate func(context.Context, ReduceTransaction) error) error
	// WithTxOptions выполняет указанный делегат в рамках транзакции с дополнительными опциями.
	WithTxOptions(ctx context.Context, opt TransactionOptions, delegate func(context.Context, ReduceTransaction) error) error

	// Release освобождает подключение к БД.
	Release()
}
