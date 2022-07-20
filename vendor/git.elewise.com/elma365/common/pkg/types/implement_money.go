package types

import (
	"encoding/json"

	"github.com/pkg/errors"

	"git.elewise.com/elma365/common/pkg/types/complextypes/money"
)

// NewMoneyType конструктор
func NewMoneyType() TypeImplement {
	return &moneyType{}
}

type moneyType struct {
	defaultType
	canReplaceCheckArrayAndSingle
	canReplaceData
	canReplaceViewData
}

func (m *moneyType) Validate(value json.RawMessage) (interface{}, error) {
	var res money.Money
	err := json.Unmarshal(value, &res)
	return res, err
}

func (m *moneyType) UnmarshalViewData(rawViewData json.RawMessage) (interface{}, error) {
	viewData := money.ViewData{}
	if len(rawViewData) == 0 {
		return viewData, nil
	}
	err := json.Unmarshal(rawViewData, &viewData)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return viewData, nil
}
