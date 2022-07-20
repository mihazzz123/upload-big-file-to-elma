package collection

import "fmt"

//go:generate ../../tooling/bin/enumer -json -sql -transform=snake -type=Alias -trimprefix=Alias -output=alias_string.go alias.go

// Alias типы ограничения доступа к коллекции
type Alias int8

const (
	// AliasMain Алиас основной таблицы
	AliasMain Alias = iota
	// AliasDir Алиас для таблицы директорий
	AliasDir
)

// WithAlias поле с алиасом
func WithAlias(alias Alias, column string) string {
	return fmt.Sprintf("%s.%s", alias, column)
}

// WithMainAlias поле с дефолтным алиасом
func WithMainAlias(column string) string {
	return WithAlias(AliasMain, column)
}
