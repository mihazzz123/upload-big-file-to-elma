package refitem

import (
	"fmt"

	uuid "github.com/satori/go.uuid"

	"git.elewise.com/elma365/common/pkg/namespace"
)

//go:generate ../../../../tooling/bin/easyjson refitem.go

// RefItem ссылка на элемент произвольного приложения
//
// easyjson:json
type RefItem struct {
	Namespace namespace.Namespace `json:"namespace" validate:"required"` // Раздел справочника элемента
	Code      string              `json:"code"      validate:"required"` // Код справочника элемента
	ID        uuid.UUID           `json:"id"        validate:"required"` // ИД элемента
}

func (ri RefItem) String() string {
	return fmt.Sprintf("%s:%s:%s", ri.Namespace, ri.Code, ri.ID.String())
}
