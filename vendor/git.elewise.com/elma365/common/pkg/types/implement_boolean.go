package types

import (
	"encoding/json"

	"git.elewise.com/elma365/common/pkg/types/complextypes/boolean"

	"github.com/pkg/errors"
)

// NewBooleanType конструктор
func NewBooleanType() TypeImplement {
	return &booleanType{}
}

type booleanType struct {
	defaultType
	canReplaceCheckArrayAndSingle
	canReplaceData
	canReplaceViewData
}

func (b *booleanType) Validate(value json.RawMessage) (interface{}, error) {
	var res bool
	err := json.Unmarshal(value, &res)
	return res, err
}

func (b *booleanType) UnmarshalViewData(rawViewData json.RawMessage) (interface{}, error) {
	viewData := boolean.ViewData{}
	if len(rawViewData) == 0 {
		return viewData, nil
	}
	err := json.Unmarshal(rawViewData, &viewData)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return viewData, nil
}
