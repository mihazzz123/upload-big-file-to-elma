package util

import (
	"context"
	"fmt"

	"git.elewise.com/elma365/common/pkg/md"
)

// CollectionName определяет тип для название коллекции mongoDB
type CollectionName string

// WithCompany возвращает название коллекции mongoDB с префиксом компании
func (cn CollectionName) WithCompany(ctx context.Context) string {
	company := md.CompanyFromContext(ctx)

	return fmt.Sprintf("%s.%s", company, cn)
}

// String implements fmt.Stringer
func (cn CollectionName) String() string {
	return string(cn)
}
