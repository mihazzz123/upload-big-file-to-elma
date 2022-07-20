package util

import (
	"context"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"hash/crc64"
	"strings"

	"git.elewise.com/elma365/common/pkg/md"
	"git.elewise.com/elma365/common/pkg/namespace"
)

// TableName имя таблицы в postgres с возможностью префикса компанией
type TableName string

const namespaceCodeSeparator = ":"

// NewTableName — возвращает имя таблицы в postgres для заданной пары namespace и code
func NewTableName(ns namespace.Namespace, code string) TableName {
	if ns == namespace.System {
		return TableName(code)
	}

	return TableName(string(ns) + namespaceCodeSeparator + code)
}

// это разделитель схемы в postgres, поэтому нельзя его менять
const companySchemaSeparator = "."

// WithCompany добавляет префикс компании
func (tn TableName) WithCompany(ctx context.Context) string {
	company := md.CompanyFromContext(ctx)

	return quoteIdentifier(company) + companySchemaSeparator + quoteIdentifier(tn.PGIdentifier())
}

// Alias короткая запись для использования в FROM tn AS alias
func (tn TableName) Alias(ctx context.Context, alias string) string {
	return fmt.Sprintf("%s AS %s", tn.WithCompany(ctx), alias)
}

// String возвращает название таблицы как строку
func (tn TableName) String() string {
	return string(tn)
}

// PGIdentifier возвращает название таблицы обрезанное с помощью TruncatePGIdentifier
func (tn TableName) PGIdentifier() string {
	return TruncatePGIdentifier(string(tn))
}

// Format форматирует название таблицы в зависимости от ключа
func (tn TableName) Format(s fmt.State, verb rune) {
	switch verb {
	default:
		_, _ = s.Write([]byte(tn))
	case 'q':
		_, _ = s.Write([]byte(quoteIdentifier(tn.PGIdentifier())))
	}
}

const categorySeparator = "@"

// CategoryCode возвращает код коллекции для категории
//
// Если коллекция уже является коллекцией категорий, то вернётся код коллекции категорий родительской коллекции.
func CategoryCode(collectionCode, fieldCode string) string {
	index := strings.Index(collectionCode, categorySeparator)
	if index != -1 {
		collectionCode = collectionCode[:index]
	}

	return collectionCode + categorySeparator + fieldCode
}

// IndexName is a name of elastic index
type IndexName string

const companyIndexSeparator = "@"

// WithCompany prefix
func (in IndexName) WithCompany(ctx context.Context) string {
	company := md.CompanyFromContext(ctx)

	return company + companyIndexSeparator + string(in)
}

// TruncatePGIdentifier обрезает строку до 63 символов, добавляя CRC64 в конце, если строка длинее
func TruncatePGIdentifier(s string) string {
	if len(s) < 64 {
		return s
	}

	hash := crc64.New(crc64.MakeTable(crc64.ISO))
	_, _ = hash.Write([]byte(s))
	sum := make([]byte, 10)
	binary.PutUvarint(sum, hash.Sum64())
	suffix := base32.StdEncoding.EncodeToString(sum)
	suffix = strings.ToLower(strings.TrimRight(suffix, string(base32.StdPadding)))

	return s[:63-len(suffix)] + suffix
}

func quoteIdentifier(name string) string {
	return `"` + strings.ReplaceAll(name, `"`, `""`) + `"`
}
