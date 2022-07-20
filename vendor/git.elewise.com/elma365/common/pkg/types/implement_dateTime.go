package types

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// FieldDataTimeType дополнительный тип представления поля DATETIME
type FieldDataTimeType string

const (
	// FieldDataTimeTypeDate только дата
	FieldDataTimeTypeDate FieldDataTimeType = "date"
	// FieldDataTimeTypeTime только время
	FieldDataTimeTypeTime FieldDataTimeType = "time"
	// FieldDataTimeTypeDateTime дата и время
	FieldDataTimeTypeDateTime FieldDataTimeType = "datetime"
)

// DefaultTimeType предустановленные значения времени
type DefaultTimeType string

const (
	// DefaultTimeTypeStartOfDay 00:00:00
	DefaultTimeTypeStartOfDay DefaultTimeType = "startOfDay"
	// DefaultTimeTypeEndOfDay 23:59:59
	DefaultTimeTypeEndOfDay DefaultTimeType = "endOfDay"
	// DefaultTimeTypeNone не устанавливать время по умолчанию
	DefaultTimeTypeNone DefaultTimeType = "none"
)

// DataTimeViewData описание представления поля типа DATETIME
type DataTimeViewData struct {
	AdditionalType     FieldDataTimeType `json:"additionalType"`
	SetCurrentDatetime bool              `json:"setCurrentDatetime"`
	TimeOptional       bool              `json:"timeOptional"`
	DefaultTimeType    DefaultTimeType   `json:"defaultTimeType"`
}

// DataTimeData дополнительные настройки поля типа DATETIME
type DataTimeData struct {
}

// NewDateTimeType конструктор
func NewDateTimeType() TypeImplement {
	return &dateTimeType{}
}

type dateTimeType struct {
	defaultType
	canReplaceCheckArrayAndSingle
	canReplaceData
	canReplaceViewData
}

func (d *dateTimeType) Validate(value json.RawMessage) (interface{}, error) {
	var res time.Time
	err := json.Unmarshal(value, &res)

	return res, err
}

func (d *dateTimeType) UnmarshalViewData(rawViewData json.RawMessage) (interface{}, error) {
	viewData := DataTimeViewData{}
	if len(rawViewData) == 0 {
		return viewData, nil
	}
	err := json.Unmarshal(rawViewData, &viewData)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return viewData, nil
}
