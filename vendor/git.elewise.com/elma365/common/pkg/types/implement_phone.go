package types

import (
	"encoding/json"

	"git.elewise.com/elma365/common/pkg/types/complextypes/phone"

	"github.com/pkg/errors"
)

// NewPhoneType конструктор
func NewPhoneType() TypeImplement {
	return &phoneType{}
}

type phoneType struct {
	defaultType
	canReplace
	canReplaceData
	canReplaceViewDataCheckType
}

func (p *phoneType) UnmarshalViewData(rawData json.RawMessage) (interface{}, error) {
	res := phone.ViewData{}
	if len(rawData) == 0 {
		return res, nil
	}
	err := json.Unmarshal(rawData, &res)
	return res, errors.WithStack(err)
}

func (p *phoneType) Validate(value json.RawMessage) (interface{}, error) {
	var res phone.Phone
	err := json.Unmarshal(value, &res)
	return res, err
}
