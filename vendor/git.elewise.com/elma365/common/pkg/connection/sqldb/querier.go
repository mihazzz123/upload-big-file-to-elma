package sqldb

import (
	"context"
)

// Querier базовый абстрактный интерфейс позволяющий выполнять запросы в БД.
// Логически объединяет sqldb.Transaction, sqldb.Connection, sqldb.Database.
type Querier interface {
	// GetSTB получает fluent-построитель SQL запросов.
	GetSTB() QueryStatementBuilder
	// QuoteIdentifier экранирует идентификатор.
	QuoteIdentifier(name string) string

	// Get получает одну запись из БД.
	Get(ctx context.Context, dst interface{}, query string, args ...interface{}) error
	// Select получает несколько записей из БД.
	Select(ctx context.Context, dst interface{}, query string, args ...interface{}) error
	// Exec выполняет запрос в БД (insert/update/delete).
	Exec(ctx context.Context, query string, args ...interface{}) (int64, error)

	// Upsert сохраняет сущность в БД.
	//
	// Передавать надо объект со структурными тэгами `db`. В качестве индекса чаще всего передаётся строка `id`.
	Upsert(ctx context.Context, table string, v interface{}, index string, skipNils bool) error
}

// QuerierEx абстрактный интерфейс позволяющий выполнять запросы в БД.
// Логически объединяет sqldb.Transaction, sqldb.Connection.
type QuerierEx interface {
	Querier

	// Prepare подготавливает запрос для выполнения.
	Prepare(ctx context.Context, name, query string) (Statement, error)
	// SelectRows выполняет запрос в БД и возвращает "сырые" данные.
	SelectRows(ctx context.Context, query string, args ...interface{}) (Rows, error)
}
