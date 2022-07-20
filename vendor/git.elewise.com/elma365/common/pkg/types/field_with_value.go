package types

import (
	"encoding/json"
	"fmt"
)

// FieldWithValue пара из описания поля и значения
type FieldWithValue struct {
	Field Field
	Value json.RawMessage
}

// FieldListWithValue список пар из описания поля и значения
type FieldListWithValue []FieldWithValue

// NewFieldListWithValue возвращает новый FieldListWithValue
func NewFieldListWithValue(fields []Field, values map[string]json.RawMessage) FieldListWithValue {
	res := make([]FieldWithValue, len(fields))
	for i, field := range fields {
		res[i].Field = field
		res[i].Value = values[field.Code]
	}
	return res
}

// Get получить поле со значением по коду
func (fl FieldListWithValue) Get(code string) (FieldWithValue, error) {
	for _, fieldWithValue := range fl {
		if fieldWithValue.Field.Code == code {
			return fieldWithValue, nil
		}
	}
	return FieldWithValue{}, fl.errorFieldNotFound(code)
}

// GetValue получить значение по коду поля
func (fl FieldListWithValue) GetValue(code string) (json.RawMessage, error) {
	for _, fieldWithValue := range fl {
		if fieldWithValue.Field.Code == code {
			return fieldWithValue.Value, nil
		}
	}
	return nil, fl.errorFieldNotFound(code)
}

// GetField получить поле по коду
func (fl FieldListWithValue) GetField(code string) (Field, error) {
	for _, fieldWithValue := range fl {
		if fieldWithValue.Field.Code == code {
			return fieldWithValue.Field, nil
		}
	}
	return Field{}, fl.errorFieldNotFound(code)
}

// GetFieldList возвращает список полей
func (fl FieldListWithValue) GetFieldList() []Field {
	res := make([]Field, len(fl))
	for i, fieldWithValue := range fl {
		res[i] = fieldWithValue.Field
	}
	return res
}

// GetValuesMap получить мапу значений
func (fl FieldListWithValue) GetValuesMap() map[string]json.RawMessage {
	values := make(map[string]json.RawMessage, len(fl))
	for _, fieldWithValue := range fl {
		values[fieldWithValue.Field.Code] = fieldWithValue.Value
	}
	return values
}

func (fl FieldListWithValue) errorFieldNotFound(code string) error {
	return fmt.Errorf("field code = \"%s\" not found", code)
}
