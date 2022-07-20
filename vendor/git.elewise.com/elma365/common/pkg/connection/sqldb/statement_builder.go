package sqldb

import "fmt"

// QueryStatementBuilder интерфейс fluent-построителя SQL запросов.
type QueryStatementBuilder interface {
	// Select возвращает построитель SELECT запроса.
	Select(columns ...string) QuerySelectStatementBuilder
	// Insert возвращает построитель INSERT запроса.
	Insert(into string) QueryInsertStatementBuilder
	// Update возвращает построитель UPDATE запроса.
	Update(table string) QueryUpdateStatementBuilder
	// Delete возвращает построитель DELETE запроса.
	Delete(from string) QueryDeleteStatementBuilder
	// CreateIndex возвращает построитель CREATE INDEX запроса.
	CreateIndex(name string) QueryCreateIndexStatementBuilder

	// Eq возвращает новый экземпляр Eq.
	Eq(expr Expr) Eq
	// NotEq возвращает новый экземпляр NotEq.
	NotEq(expr Expr) NotEq
	// Like возвращает новый экземпляр Like.
	Like(expr Expr) Like
	// NotLike возвращает новый экземпляр NotLike.
	NotLike(expr Expr) NotLike
	// Lt возвращает новый экземпляр Lt.
	Lt(expr Expr) Lt
	// LtOrEq возвращает новый экземпляр LtOrEq.
	LtOrEq(expr Expr) LtOrEq
	// Gt возвращает новый экземпляр Gt.
	Gt(expr Expr) Gt
	// GtOrEq возвращает новый экземпляр GtOrEq.
	GtOrEq(expr Expr) GtOrEq
	// And возвращает новый экземпляр And.
	And(exprs ...QueryBuilder) And
	// Or возвращает новый экземпляр Or.
	Or(exprs ...QueryBuilder) Or
	// Expr строит выражение по тексту запроса и аргументам.
	Expr(sql string, args ...interface{}) QueryBuilder
}

// PlaceholderFormat префикс переменной запроса.
type PlaceholderFormat int8

const (
	// Question по умолчанию используется знак ?.
	Question PlaceholderFormat = iota
	// Dollar указывает что необходимо заменить ? на $ при формировании SQL-запроса.
	Dollar
	// Colon указывает что необходимо заменить ? на : при формировании SQL-запроса.
	Colon
	// AtP указывает что необходимо заменить ? на @ при формировании SQL-запроса.
	AtP
)

// QueryBuilder интерфейс построителя запроса.
type QueryBuilder interface {
	// ToSql возвращает сконструированный запрос и его параметры.
	ToSql() (string, []interface{}, error)
}

// QuerySelectStatementBuilder интерфейс fluent-построителя SELECT запроса.
type QuerySelectStatementBuilder interface {
	QueryBuilder
	fmt.Stringer

	// Columns добавляет результирующие колонки в запрос.
	Columns(columns ...string) QuerySelectStatementBuilder
	// Column добавляет результирующую колонку в запрос.
	Column(column interface{}, args ...interface{}) QuerySelectStatementBuilder
	// Distinct добавляет ограничение DISTINCT в запрос.
	Distinct() QuerySelectStatementBuilder
	// From добавляет ограничение FROM в запрос.
	From(from string) QuerySelectStatementBuilder
	// FromSelect добавляет подзапрос в ограничение FROM запроса.
	FromSelect(from QuerySelectStatementBuilder, alias string) QuerySelectStatementBuilder
	// Join добавляет ограничение JOIN в запрос.
	Join(join string, rest ...interface{}) QuerySelectStatementBuilder
	// LeftJoin добавляет ограничение LEFT JOIN в запрос.
	LeftJoin(join string, rest ...interface{}) QuerySelectStatementBuilder
	// RightJoin добавляет ограничение RIGHT JOIN в запрос.
	RightJoin(join string, rest ...interface{}) QuerySelectStatementBuilder
	// InnerJoin добавляет ограничение INNER JOIN в запрос.
	InnerJoin(join string, rest ...interface{}) QuerySelectStatementBuilder
	// CrossJoin добавляет ограничение CROSS JOIN в запрос.
	CrossJoin(join string, rest ...interface{}) QuerySelectStatementBuilder
	// Where добавляет ограничение WHERE в запрос.
	Where(pred interface{}, args ...interface{}) QuerySelectStatementBuilder
	// GroupBy добавляет ограничение GROUP BY в запрос.
	GroupBy(groupBys ...string) QuerySelectStatementBuilder
	// Having добавляет ограничение HAVING в запрос.
	Having(pred interface{}, rest ...interface{}) QuerySelectStatementBuilder
	// OrderBy добавляет ограничение ORDER BY в запрос.
	OrderBy(orderBys ...string) QuerySelectStatementBuilder
	// Limit ограничивает максимальное количество записей в результате запроса.
	Limit(limit uint64) QuerySelectStatementBuilder
	// Prefix добавляет префикс в запрос.
	Prefix(sql string, args ...interface{}) QuerySelectStatementBuilder
	// Suffix добавляет суффикс в запрос
	Suffix(sql string, args ...interface{}) QuerySelectStatementBuilder
	// Offset добавляет смещение относительно начала результирующего набора данных.
	Offset(offset uint64) QuerySelectStatementBuilder
	// Union добавляет ограничение UNION в запрос.
	Union(with QuerySelectStatementBuilder) QuerySelectStatementBuilder
}

// QueryInsertStatementBuilder интерфейс fluent-построителя INSERT запроса.
type QueryInsertStatementBuilder interface {
	QueryBuilder
	fmt.Stringer

	// Columns добавляет вставляемые колонки в запрос.
	Columns(columns ...string) QueryInsertStatementBuilder
	// Values добавляет вставляемые значения в запрос
	Values(values ...interface{}) QueryInsertStatementBuilder
	// SetMap устанавливает значения столбцов и значений согласно карте вставки.
	SetMap(clauses map[string]interface{}) QueryInsertStatementBuilder
	// Select устанавливает вставляемые значения из результатов подзапроса.
	// Если Values и Select используются одновременно, то Select имеет приоритет.
	Select(sb QuerySelectStatementBuilder) QueryInsertStatementBuilder
	// Prefix добавляет префикс в запрос.
	Prefix(sql string, args ...interface{}) QueryInsertStatementBuilder
	// Suffix добавляет суффикс в запрос
	Suffix(sql string, args ...interface{}) QueryInsertStatementBuilder
}

// QueryUpdateStatementBuilder интерфейс fluent-построителя UPDATE запроса.
type QueryUpdateStatementBuilder interface {
	QueryBuilder
	fmt.Stringer

	// Set добавляет ограничение SET в запрос.
	Set(column string, value interface{}) QueryUpdateStatementBuilder
	// SetMap устанавливает значения столбцов и значений согласно карте вставки.
	SetMap(clauses map[string]interface{}) QueryUpdateStatementBuilder
	// Where добавляет ограничение WHERE в запрос.
	Where(pred interface{}, args ...interface{}) QueryUpdateStatementBuilder
	// OrderBy добавляет ограничение ORDER BY в запрос.
	OrderBy(orderBys ...string) QueryUpdateStatementBuilder
	// Limit ограничивает максимальное количество записей в результате запроса.
	Limit(limit uint64) QueryUpdateStatementBuilder
	// Prefix добавляет префикс в запрос.
	Prefix(sql string, args ...interface{}) QueryUpdateStatementBuilder
	// Suffix добавляет суффикс в запрос
	Suffix(sql string, args ...interface{}) QueryUpdateStatementBuilder
}

// QueryDeleteStatementBuilder интерфейс fluent-построителя DELETE запроса.
type QueryDeleteStatementBuilder interface {
	QueryBuilder
	fmt.Stringer

	// Where добавляет ограничение WHERE в запрос.
	Where(pred interface{}, args ...interface{}) QueryDeleteStatementBuilder
	// OrderBy добавляет ограничение ORDER BY в запрос.
	OrderBy(orderBys ...string) QueryDeleteStatementBuilder
	// Limit ограничивает максимальное количество записей в результате запроса.
	Limit(limit uint64) QueryDeleteStatementBuilder
	// Prefix добавляет префикс в запрос.
	Prefix(sql string, args ...interface{}) QueryDeleteStatementBuilder
	// Suffix добавляет суффикс в запрос
	Suffix(sql string, args ...interface{}) QueryDeleteStatementBuilder
}

// QueryCreateIndexStatementBuilder интерфейс fluent-построителя запроса CREATE INDEX.
type QueryCreateIndexStatementBuilder interface {
	QueryBuilder
	fmt.Stringer

	// Table указывает название таблицы для которой создается индекс.
	Table(table string) QueryCreateIndexStatementBuilder
	// Unique указывает что индекс является уникальным.
	Unique(unique bool) QueryCreateIndexStatementBuilder
	// Fields задает названия колонок таблицы по которым будет создан индекс.
	Fields(columns []string) QueryCreateIndexStatementBuilder
	// IfNotExists указывает что индекс должен быть создан только в том случае, если он не существует.
	IfNotExists(notExists bool) QueryCreateIndexStatementBuilder
	// Where добавляет ограничение WHERE в запрос.
	Where(pred interface{}, args ...interface{}) QueryCreateIndexStatementBuilder
}
