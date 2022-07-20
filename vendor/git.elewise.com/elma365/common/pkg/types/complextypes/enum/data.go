package enum

import (
	"github.com/pkg/errors"

	"git.elewise.com/elma365/common/pkg/errs"
)

//go:generate ../../../../tooling/bin/easyjson data.go
//go:generate ../../../../tooling/bin/easylocalizer data.go

// EnumData - структура поля Data
//
// easyjson:json
// localizer:enum_data
// nolint:golint
type EnumData struct {
	EnumItems []EnumItem `json:"enumItems" patch:"po"`
}

// GetItemByCode получить вариант по коду
func (e EnumData) GetItemByCode(code string) (EnumItem, error) {
	for i := range e.EnumItems {
		if e.EnumItems[i].Code == code {
			return e.EnumItems[i], nil
		}
	}
	return EnumItem{}, errors.WithStack(errs.NotFound)
}

// EnumItem - вариант значения поля
//
// easyjson:json
// localizer:enum_item
// nolint:golint
type EnumItem struct {
	Code    string `json:"code" localizer:"id"`
	Name    string `json:"name" patch:"po" localizer:"value"`
	Checked bool   `json:"checked"`
}
