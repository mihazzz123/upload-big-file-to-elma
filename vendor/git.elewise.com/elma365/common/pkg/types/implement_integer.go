package types

import "encoding/json"

// NewIntegerType конструктор
func NewIntegerType() TypeImplement {
	return &integerType{}
}

type integerType struct {
	defaultType
	canReplaceCheckArrayAndSingle
	canReplaceData
	canReplaceViewData
}

func (i *integerType) Validate(value json.RawMessage) (interface{}, error) {
	var res int32
	err := json.Unmarshal(value, &res)
	return res, err
}
