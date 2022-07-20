package types

import (
	"encoding/json"

	"github.com/pkg/errors"

	"git.elewise.com/elma365/common/pkg/errs"

	"git.elewise.com/elma365/common/pkg/types/complextypes/enum"
)

// NewEnumType конструктор
func NewEnumType() TypeImplement {
	return &enumType{}
}

type enumType struct {
	defaultType
	canReplaceCheckArrayAndSingle
	canReplaceViewData
}

func (e *enumType) Validate(value json.RawMessage) (interface{}, error) {
	var res enum.Enum
	err := json.Unmarshal(value, &res)

	return res, err
}

func (e *enumType) UnmarshalData(rawData json.RawMessage) (interface{}, error) {
	return e.unmarshalData(rawData)
}

func (e *enumType) unmarshalData(rawData json.RawMessage) (enum.EnumData, error) {
	data := enum.EnumData{}
	if len(rawData) == 0 {
		return data, nil
	}
	err := json.Unmarshal(rawData, &data)
	return data, err
}

func (e *enumType) CanReplaceData(source, candidate json.RawMessage) (bool, error) {
	sourceData, err := e.unmarshalData(source)
	if err != nil {
		return false, errors.WithStack(err)
	}
	candidateData, err := e.unmarshalData(candidate)
	if err != nil {
		return false, errors.WithStack(err)
	}

	for i := range sourceData.EnumItems {
		item1 := sourceData.EnumItems[i]
		_, err := candidateData.GetItemByCode(item1.Code)
		if err != nil {
			if errors.Cause(err) == errs.NotFound {
				return false, nil
			}
			return false, errors.WithStack(err)
		}
	}
	return true, nil
}
