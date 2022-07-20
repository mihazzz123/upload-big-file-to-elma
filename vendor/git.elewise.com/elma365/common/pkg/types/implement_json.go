package types

import (
	"encoding/json"
)

// NewJSONType конструктор
func NewJSONType() TypeImplement {
	return &jsonType{}
}

type jsonType struct {
	defaultType
	canReplaceCheckArrayAndSingle
	canReplaceData
	canReplaceViewData
}

func (d *jsonType) Validate(value json.RawMessage) (interface{}, error) {
	return value, nil
}
