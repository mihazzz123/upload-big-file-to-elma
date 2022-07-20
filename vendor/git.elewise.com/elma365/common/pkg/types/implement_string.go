package types

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// FieldStringType дополнительный тип для представления
type FieldStringType string

const (
	// FieldStringTypeString одна строка
	FieldStringTypeString FieldStringType = "string"
	// FieldStringTypeText многострочный
	FieldStringTypeText FieldStringType = "text"
)

// StringViewData описание представления типа
type StringViewData struct {
	AdditionalType FieldStringType `json:"additionalType"`
}

// StringData дополнительные параметры поля типа String
type StringData struct {
	Mask StringMask `json:"mask,omitempty" patch:"po"`
	Key  bool       `json:"key"`
}

// StringMask валидация строки по маске
type StringMask struct {
	Pattern      string `json:"pattern"`
	ErrorMessage string `json:"errorMessage,omitempty" patch:"po"`
}

// NewStringType конструктор
func NewStringType() TypeImplement {
	return &stringType{}
}

type stringType struct {
	defaultType
	canReplaceCheckArrayAndSingle
	canReplaceViewData
}

func (s *stringType) Validate(value json.RawMessage) (interface{}, error) {
	var res string
	err := json.Unmarshal(value, &res)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return res, nil
}

func (s *stringType) UnmarshalViewData(rawViewData json.RawMessage) (interface{}, error) {
	viewData := StringViewData{}
	if len(rawViewData) == 0 {
		return viewData, nil
	}
	err := json.Unmarshal(rawViewData, &viewData)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return viewData, nil
}

func (s *stringType) UnmarshalData(data json.RawMessage) (interface{}, error) {
	return s.unmarshalData(data)
}

func (s *stringType) unmarshalData(data json.RawMessage) (StringData, error) {
	res := StringData{}
	if len(data) == 0 {
		return res, nil
	}
	err := json.Unmarshal(data, &res)
	if err != nil {
		return res, errors.WithStack(err)
	}
	return res, nil
}

func (s *stringType) CanReplaceData(source, candidate json.RawMessage) (bool, error) {
	_, err := s.unmarshalData(source)
	if err != nil {
		return false, errors.WithStack(err)
	}
	_, err = s.unmarshalData(candidate)
	if err != nil {
		return false, errors.WithStack(err)
	}
	return true, nil
}
