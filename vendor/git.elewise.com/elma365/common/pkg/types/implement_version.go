package types

import (
	"encoding/json"
)

// NewVersionType конструктор
func NewVersionType() TypeImplement {
	return &versionType{}
}

type versionType struct {
	defaultType
	canReplaceCheckArrayAndSingle
	canReplaceData
	canReplaceViewData
}

func (v *versionType) Validate(value json.RawMessage) (interface{}, error) {
	var res string
	err := json.Unmarshal(value, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
