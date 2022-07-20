package sqldb

// Expr Выражение.
type Expr map[string]interface{}

// Eq Выражение эквивалентности параметров.
type Eq interface {
	QueryBuilder
}

// NotEq Выражение не эквивалентности параметров.
type NotEq interface {
	QueryBuilder
}

// Like Выражение эквивалентности строк.
type Like interface {
	QueryBuilder
}

// NotLike Выражение не эквивалентности строк.
type NotLike interface {
	QueryBuilder
}

// Lt Выражение меньше.
type Lt interface {
	QueryBuilder
}

// LtOrEq Выражение меньше или равно.
type LtOrEq interface {
	QueryBuilder
}

// Gt Выражение больше.
type Gt interface {
	QueryBuilder
}

// GtOrEq Выражение больше или равно.
type GtOrEq interface {
	QueryBuilder
}

// And Конъюнкция выражений.
type And interface {
	QueryBuilder
}

// Or Дизъюнкция выражений.
type Or interface {
	QueryBuilder
}
