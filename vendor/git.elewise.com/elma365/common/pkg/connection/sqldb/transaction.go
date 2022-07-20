package sqldb

import (
	"context"
)

// TransactionIsolationLevel уровень изоляции транзакции.
type TransactionIsolationLevel int8

const (
	// ReadUncommitted каждая транзакция видит незафиксированные изменения другой транзакции (феномен грязного чтения).
	ReadUncommitted TransactionIsolationLevel = iota
	// ReadCommitted параллельно исполняющиеся транзакции видят только зафиксированные изменения из других транзакций.
	ReadCommitted
	// RepeatableRead параллельно исполняющиеся транзакции не видят измененные и удаленные записи другой транзакции, но видят вставленные записи.
	RepeatableRead
	// Serializable параллельные транзакции не оказывают влияния друг на друга.
	Serializable
)

// TransactionAccessMode режим доступа транзакции.
type TransactionAccessMode int8

const (
	// ReadWrite транзакция для чтения и записи.
	ReadWrite TransactionAccessMode = iota
	// ReadOnly транзакция только для чтения.
	ReadOnly
)

// TransactionOptions настройки транзакции.
type TransactionOptions struct {
	IsolationLevel TransactionIsolationLevel
	AccessMode     TransactionAccessMode
}

// ReduceTransaction усеченная транзакция базы данных.
//
// Если в сигнатуре метода передается экземпляр этого интерфейса, то это означает, что транзакция была создана выше и
// мы можем лишь использовать транзакцию без возможности ее фиксации или откаа в БД.
type ReduceTransaction interface {
	QuerierEx

	// AfterCommit добавляет действие, которое будет выполнено после фиксации изменения в транзакции.
	AfterCommit(ctx context.Context, delegate func(ctx context.Context)) error
	// SavePoint создает точку восстановления в рамках текущей транзакции.
	SavePoint(ctx context.Context) (SavePoint, error)
}

// Transaction транзакция базы данных.
type Transaction interface {
	ReduceTransaction

	// Commit фиксирует изменений в транзакции.
	Commit(ctx context.Context) error
	// Rollback откатывает изменения в транзакции.
	Rollback(ctx context.Context) error
}
