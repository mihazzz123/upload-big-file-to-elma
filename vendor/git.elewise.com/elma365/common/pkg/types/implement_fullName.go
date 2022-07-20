package types

import (
	"encoding/json"

	"git.elewise.com/elma365/common/pkg/types/complextypes/fullname"
)

// NewFullNameType конструктор
func NewFullNameType() TypeImplement {
	return &fullNameType{}
}

type fullNameType struct {
	defaultType
	canReplaceCheckArrayAndSingle
	canReplaceData
	canReplaceViewData
}

func (f *fullNameType) Validate(value json.RawMessage) (interface{}, error) {
	var res fullname.FullName
	err := json.Unmarshal(value, &res)

	return res, err
}
