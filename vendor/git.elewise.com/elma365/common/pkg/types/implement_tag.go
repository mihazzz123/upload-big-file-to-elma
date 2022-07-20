package types

import "encoding/json"

// NewTagType конструктор
func NewTagType() TypeImplement {
	return &tagType{}
}

type tagType struct {
	defaultType
	canReplaceCheckArrayAndSingle
	canReplaceData
	canReplaceViewData
}

func (t *tagType) Validate(value json.RawMessage) (interface{}, error) {
	var res string
	err := json.Unmarshal(value, &res)
	return res, err
}
