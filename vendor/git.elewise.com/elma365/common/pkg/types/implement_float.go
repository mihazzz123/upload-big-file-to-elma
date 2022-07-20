package types

import "encoding/json"

// FieldFloatType разновидности типов
type FieldFloatType string

const (
	// FieldFloatTypeFloat float
	FieldFloatTypeFloat FieldFloatType = "float"
	// FieldFloatTypeInteger integer
	FieldFloatTypeInteger FieldFloatType = "integer"
)

// FloatViewData описание представления поля типа FLOAT
type FloatViewData struct {
	AdditionalType   FieldFloatType `json:"additionalType"`
	ShowRowSeparator bool           `json:"showRowSeparator"`
}

// NewFloatType конструктор
func NewFloatType() TypeImplement {
	return &floatType{}
}

type floatType struct {
	defaultType
	canReplaceCheckArrayAndSingle
	canReplaceData
	canReplaceViewData
}

func (f *floatType) Validate(value json.RawMessage) (interface{}, error) {
	var res float64
	err := json.Unmarshal(value, &res)
	return res, err
}
