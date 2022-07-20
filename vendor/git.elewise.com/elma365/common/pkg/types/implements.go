package types

import (
	"encoding/json"
	"reflect"

	"git.elewise.com/elma365/common/pkg/errs"

	"github.com/pkg/errors"
)

// TypeImplement имплементация поля определённого типа
type TypeImplement interface {
	Validate(message json.RawMessage) (interface{}, error)
	UnmarshalData(data json.RawMessage) (interface{}, error)
	UnmarshalViewData(viewData json.RawMessage) (interface{}, error)
	Compare(value1 json.RawMessage, value2 json.RawMessage) (bool, error)
	CanReplace(oldField, newField Field) (bool, error)
	CanReplaceData(source json.RawMessage, candidate json.RawMessage) (bool, error)
	CanReplaceViewData(source json.RawMessage, candidate json.RawMessage) (bool, error)
}

//nolint: gochecknoglobals // это должна быть константа, но Go так не умеет
var implements = map[Type]TypeImplement{
	String:        NewStringType(),
	Float:         NewFloatType(),
	Integer:       NewIntegerType(),
	Boolean:       NewBooleanType(),
	DateTime:      NewDateTimeType(),
	Duration:      NewDurationType(),
	Category:      NewCategoryType(),
	Tag:           NewTagType(),
	Money:         NewMoneyType(),
	File:          NewFileType(),
	Image:         NewFileType(),
	Link:          NewLinkType(),
	Phone:         NewPhoneType(),
	Email:         NewEmailType(),
	Version:       NewVersionType(),
	JSON:          NewJSONType(),
	SysUser:       NewSysUserType(),
	FullName:      NewFullNameType(),
	SysOSNode:     NewSysOSNodeType(),
	SysCollection: NewSysCollectionType(),
	Status:        NewStatusType(),
	Enum:          NewEnumType(),
	RefItem:       NewRefItemType(),
	Table:         NewTableType(),
	Account:       NewAccountType(),
	Role:          NewRoleType(),
}

// Validate провалидировать значение для поля
func (t Type) Validate(message json.RawMessage) (interface{}, error) {
	implement, err := t.GetImplement()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return implement.Validate(message)
}

// UnmarshalData извлечь описание поля
func (t Type) UnmarshalData(rawData json.RawMessage) (interface{}, error) {
	implement, err := t.GetImplement()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return implement.UnmarshalData(rawData)
}

// UnmarshalViewData извлечь описание представления
func (t Type) UnmarshalViewData(rawViewData json.RawMessage) (interface{}, error) {
	implement, err := t.GetImplement()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return implement.UnmarshalViewData(rawViewData)
}

// Compare сравнить пару значений на равенство
func (t Type) Compare(value1, value2 json.RawMessage) (bool, error) {
	implement, err := t.GetImplement()
	if err != nil {
		return false, errors.WithStack(err)
	}
	return implement.Compare(value1, value2)
}

// CanReplaceViewData проверить возможность обновления описания представления поля
func (t Type) CanReplaceViewData(sourceData, candidateData json.RawMessage) (bool, error) {
	implement, err := t.GetImplement()
	if err != nil {
		return false, errors.WithStack(err)
	}
	return implement.CanReplaceViewData(sourceData, candidateData)
}

// CanReplaceData проверить возможность обновления описания поля
func (t Type) CanReplaceData(sourceData, candidateData json.RawMessage) (bool, error) {
	implement, err := t.GetImplement()
	if err != nil {
		return false, errors.WithStack(err)
	}
	return implement.CanReplaceData(sourceData, candidateData)
}

// CanReplace проверить возможность обновления описания
func (t Type) CanReplace(field, candidate Field) (bool, error) {
	if field.Type != t {
		return false, errors.New("wrong type of field")
	}
	if field.Type != candidate.Type {
		return false, nil
	}
	implement, err := t.GetImplement()
	if err != nil {
		return false, errors.WithStack(err)
	}
	return implement.CanReplace(field, candidate)
}

// GetImplement получить конкретную имплементацию
func (t Type) GetImplement() (TypeImplement, error) {
	implement, ok := implements[t]
	if !ok {
		return nil, errors.New("type not implemented")
	}
	return implement, nil
}

type defaultType struct {
}

func (d defaultType) Validate(value json.RawMessage) (interface{}, error) {
	var res interface{}
	err := json.Unmarshal(value, &res)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return res, nil
}

func (d defaultType) UnmarshalData(data json.RawMessage) (interface{}, error) {
	return nil, errors.WithStack(errs.NotImplemented)
}

func (d defaultType) UnmarshalViewData(viewData json.RawMessage) (interface{}, error) {
	return nil, errors.WithStack(errs.NotImplemented)
}

func (defaultType) unmarshalViewData(res interface{}) func(json.RawMessage) (interface{}, error) {
	return func(raw json.RawMessage) (interface{}, error) {
		if len(raw) == 0 {
			return res, nil
		}
		err := json.Unmarshal(raw, &res)
		return res, errors.WithStack(err)
	}
}

func (d defaultType) Compare(rawValue1, rawValue2 json.RawMessage) (bool, error) {
	value1, err := d.Validate(rawValue1)
	if err != nil {
		return false, errors.WithStack(err)
	}
	value2, err := d.Validate(rawValue2)
	if err != nil {
		return false, errors.WithStack(err)
	}
	return reflect.DeepEqual(value1, value2), nil
}

// canReplace тип, игнорирующий признаки Array и Single
type canReplace struct{}

func (canReplace) CanReplace(_, _ Field) (bool, error) {
	return true, nil
}

// canReplaceCheckArrayAndSingle тип разрешающий подмену только при возможности обновить признаки Array и Single
type canReplaceCheckArrayAndSingle struct{}

func (canReplaceCheckArrayAndSingle) CanReplace(field, candidate Field) (bool, error) {
	if field.Array != candidate.Array {
		return false, nil
	}
	return !field.Array || field.Single || !candidate.Single, nil
}

// canReplaceData всегда разрешать обновлять настройки типа
type canReplaceData struct{}

func (canReplaceData) CanReplaceData(_, _ json.RawMessage) (bool, error) {
	return true, nil
}

// canReplaceViewData всегда разрешать обновлять настройки представления
type canReplaceViewData struct{}

func (canReplaceViewData) CanReplaceViewData(_, _ json.RawMessage) (bool, error) {
	return true, nil
}

// canReplaceViewDataCheckType проверять поле type в описании (применимо для телефона, почты и аккаунта)
type canReplaceViewDataCheckType struct{}

func (canReplaceViewDataCheckType) CanReplaceViewData(sourceRaw, candidateRaw json.RawMessage) (bool, error) {
	var source, candidate struct {
		Type string `json:"type"`
	}
	if len(sourceRaw) > 0 {
		if err := json.Unmarshal(sourceRaw, &source); err != nil {
			return false, errors.Wrap(err, "cannot unmarshal source view data")
		}
	}
	if len(candidateRaw) > 0 {
		if err := json.Unmarshal(candidateRaw, &candidate); err != nil {
			return false, errors.Wrap(err, "cannot unmarshal candidate view data")
		}
	}
	return source.Type != "" || candidate.Type == "", nil
}
