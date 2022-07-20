package types

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// NewTableType конструктор
func NewTableType() TypeImplement {
	return &tableType{}
}

type tableType struct {
	defaultType
	canReplaceCheckArrayAndSingle
	canReplaceViewData
}

func (t *tableType) Validate(value json.RawMessage) (interface{}, error) {
	res := TableValue{}
	err := json.Unmarshal(value, &res)
	return res, errors.WithStack(err)
}

func (t *tableType) UnmarshalData(data json.RawMessage) (interface{}, error) {
	return t.unmarshalData(data)
}

func (t *tableType) unmarshalData(data json.RawMessage) (TableFieldData, error) {
	res := TableFieldData{}
	if len(data) == 0 {
		return res, nil
	}
	err := json.Unmarshal(data, &res)
	if err != nil {
		return res, errors.WithStack(err)
	}
	return res, nil
}

func (t *tableType) UnmarshalViewData(rawViewData json.RawMessage) (interface{}, error) {
	viewData := TableViewData{}
	err := json.Unmarshal(rawViewData, &viewData)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return viewData, nil
}

func (t *tableType) CanReplaceData(source, candidate json.RawMessage) (bool, error) {
	sourceData, err := t.unmarshalData(source)
	if err != nil {
		return false, errors.WithStack(err)
	}
	candidateData, err := t.unmarshalData(candidate)
	if err != nil {
		return false, errors.WithStack(err)
	}
	return t.canReplaceData(sourceData, candidateData)
}

func (t *tableType) canReplaceData(sourceData, candidateData TableFieldData) (bool, error) {
	canReplaceFields, err := sourceData.Fields.CanReplaceTo(candidateData.Fields)
	if err != nil {
		return false, errors.WithStack(err)
	}
	if !canReplaceFields {
		return false, nil
	}

	for key := range sourceData.Result {
		if sourceData.Result[key].Kind != ResultKindFormula {
			continue
		}
		if _, ok := candidateData.Result[key]; !ok {
			return false, nil
		}
	}

	return true, nil
}

// TableValue данные таблицы
type TableValue struct {
	Rows   []TableValueRow `json:"rows"`
	Result TableValueRow   `json:"result"`
	View   string          `json:"view"`
}

// TableValueRow строка таблицы
type TableValueRow = map[string]json.RawMessage

// TableFieldData описания поля типа TABLE
type TableFieldData struct {
	Fields Fields                 `json:"fields" patch:"po"`
	Result map[string]TableResult `json:"result" patch:"po"`
}

// TableResultKind тип агрегации итога
type TableResultKind string

const (
	// ResultKindNone не показывать итог
	ResultKindNone TableResultKind = "none"
	// ResultKindLabel текстовая
	ResultKindLabel TableResultKind = "label"
	// ResultKindFormula итог по формуле
	ResultKindFormula TableResultKind = "formula"
)

// TableResultFormula агрегатные функции
type TableResultFormula string

const (
	// TableResultFormulaSum сумма
	TableResultFormulaSum TableResultFormula = "sum"
	// TableResultFormulaMax максимум
	TableResultFormulaMax TableResultFormula = "max"
	// TableResultFormulaMin минимум
	TableResultFormulaMin TableResultFormula = "min"
	// TableResultFormulaAverage среднее
	TableResultFormulaAverage TableResultFormula = "average"
)

// TableResult описание результата по колонке
type TableResult struct {
	Kind    TableResultKind    `json:"kind"`
	Label   string             `json:"label" patch:"po"`
	Formula TableResultFormula `json:"formula"`
}

// TableViewData описание представления таблицы
type TableViewData struct {
	HeaderHidden          bool            `json:"headerHidden"`
	FooterHidden          bool            `json:"footerHidden"`
	RelativeWidth         bool            `json:"relativeWidth"`
	CollapseNestedHeaders bool            `json:"collapseNestedHeaders"`
	ColumnsView           json.RawMessage `json:"columnsView"`
	ViewVariant           string          `json:"viewVariant"`
	ViewTemplate          string          `json:"viewTemplate"`
	ShowOrderNumbers      bool            `json:"showOrderNumbers"`
}
