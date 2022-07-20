package types

//go:generate ../../tooling/bin/easyjson fields.go

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/pkg/errors"
)

// Fields is slice of fields
//
// easyjson:json
type Fields []Field

// FindByCode поиск поля по заданному коду
func (o *Fields) FindByCode(code string) *Field {
	for _, field := range *o {
		if field.Code == code {
			return &field
		}
	}

	return nil
}

// GetIndexByCode получить индекс поля по коду
func (o *Fields) GetIndexByCode(code string) int {
	for i, field := range *o {
		if field.Code == code {
			return i
		}
	}

	return -1
}

// Scan implements sql.Scanner interface
func (o *Fields) Scan(pSrc interface{}) error {
	var arrfs []Field

	switch src := pSrc.(type) {
	case string:
		data := []byte(src)
		if err := json.Unmarshal(data, &arrfs); err != nil {
			return err
		}
	case []byte:
		if err := json.Unmarshal(src, &arrfs); err != nil {
			return err
		}
	default:
		return errors.Errorf("Fields.Scan: cannot scan type %T into Fields", pSrc)
	}
	*o = Fields(arrfs)

	return nil
}

// Value implements sql.Valuer interface
func (o Fields) Value() (value driver.Value, err error) {
	val, err := json.Marshal(o)
	return string(val), errors.WithStack(err)
}

// методы для протобафа

// Marshal marshaler interfacer
func (o Fields) Marshal() ([]byte, error) {
	return json.Marshal(o)
}

// MarshalTo protobuf marshaler
func (o *Fields) MarshalTo(data []byte) (n int, err error) {
	d, err := json.Marshal(o)
	if err != nil {
		return 0, err
	}
	return copy(data, d), nil
}

// Unmarshal unmarshaller interface
func (o *Fields) Unmarshal(data []byte) error {
	return json.Unmarshal(data, o)
}

// Size resturn size for protobuf
func (o *Fields) Size() int {
	if o == nil {
		return 0
	}

	d, _ := json.Marshal(o)

	return len(d)
}

// CanReplaceTo можно ли обновить описание полей
func (o *Fields) CanReplaceTo(candidate Fields) (bool, error) {
	for i := range *o {
		sourceField := (*o)[i]
		candidateField := candidate.FindByCode(sourceField.Code)
		if candidateField == nil {
			return false, nil
		}
		canReplaceField, err := sourceField.CanReplaceTo(candidateField)
		if err != nil {
			return false, errors.WithStack(err)
		}
		if !canReplaceField {
			return false, nil
		}
	}
	return true, nil
}

// Merge объединить наборы полей
func (o Fields) Merge(fields Fields) (Fields, error) {
	res := o
	for i := range fields {
		newField := fields[i]

		oldFieldIndex := res.GetIndexByCode(newField.Code)
		if oldFieldIndex == -1 {
			res = append(res, newField)
		} else {
			oldField := res[oldFieldIndex]
			canReplace, err := oldField.CanReplaceTo(&newField)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			if !canReplace {
				return nil, errors.New("incompatible fields")
			}
			res[oldFieldIndex] = newField
		}
	}
	return res, nil
}

// Has проверка на существование нужного филда по коду
func (o Fields) Has(code string) bool {
	for _, field := range o {
		if field.Code == code {
			return true
		}
	}
	return false
}
