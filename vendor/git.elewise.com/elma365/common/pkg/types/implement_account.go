package types

import (
	"encoding/json"

	"git.elewise.com/elma365/common/pkg/types/complextypes/account"
)

// NewAccountType конструктор
func NewAccountType() TypeImplement {
	return &accountType{}
}

type accountType struct {
	defaultType
	canReplace
	canReplaceData
	canReplaceViewDataCheckType
}

func (a *accountType) UnmarshalViewData(rawData json.RawMessage) (interface{}, error) {
	return a.unmarshalViewData(account.ViewData{})(rawData)
}

func (a *accountType) Validate(value json.RawMessage) (interface{}, error) {
	var res account.Account
	err := json.Unmarshal(value, &res)
	return res, err
}
