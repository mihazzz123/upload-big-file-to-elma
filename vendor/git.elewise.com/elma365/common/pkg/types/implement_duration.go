package types

import "encoding/json"

// NewDurationType конструктор
func NewDurationType() TypeImplement {
	return &durationType{}
}

type durationType struct {
	defaultType
	canReplaceCheckArrayAndSingle
	canReplaceData
	canReplaceViewData
}

func (d *durationType) Validate(value json.RawMessage) (interface{}, error) {
	var res int32
	err := json.Unmarshal(value, &res)
	return res, err
}
