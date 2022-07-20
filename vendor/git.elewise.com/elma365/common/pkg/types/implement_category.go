package types

import (
	"encoding/json"

	"github.com/pkg/errors"

	uuid "github.com/satori/go.uuid"
)

// NewCategoryType конструктор
func NewCategoryType() TypeImplement {
	return &categoryType{}
}

// CategoryData дополнительные параметры поля типа CATEGORY
type CategoryData struct {
	Fields Fields `json:"fields" patch:"po"`
}

type categoryType struct {
	defaultType
	canReplaceCheckArrayAndSingle
	canReplaceViewData
}

func (t *categoryType) Validate(value json.RawMessage) (interface{}, error) {
	var res uuid.UUID
	err := json.Unmarshal(value, &res)
	return res, err
}

func (t *categoryType) UnmarshalData(data json.RawMessage) (interface{}, error) {
	return t.unmarshalData(data)
}

func (t *categoryType) unmarshalData(data json.RawMessage) (CategoryData, error) {
	res := CategoryData{}
	if len(data) == 0 {
		return res, nil
	}
	err := json.Unmarshal(data, &res)
	if err != nil {
		return res, errors.WithStack(err)
	}
	return res, nil
}

func (t *categoryType) CanReplaceData(source, candidate json.RawMessage) (bool, error) {
	sourceData, err := t.unmarshalData(source)
	if err != nil {
		return false, errors.WithStack(err)
	}
	candidateData, err := t.unmarshalData(candidate)
	if err != nil {
		return false, errors.WithStack(err)
	}
	return sourceData.Fields.CanReplaceTo(candidateData.Fields)
}
