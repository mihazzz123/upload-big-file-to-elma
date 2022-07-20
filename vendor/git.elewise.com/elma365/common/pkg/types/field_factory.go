package types

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"git.elewise.com/elma365/common/pkg/types/complextypes/boolean"
	"git.elewise.com/elma365/common/pkg/types/complextypes/enum"
)

// Translator сервис переводов
type Translator interface {
	TranslateString(context.Context, string, ...interface{}) string
}

// FieldFactory фабрика описания полей
type FieldFactory struct {
	tr     Translator
	format string
}

// NewFieldFactory создать новую фабрику описания полей
//
// nameFormat будет использоваться для перевода названий полей и должен иметь одну строковую
// подстановку для кода поля.
func NewFieldFactory(tr Translator, nameFormat string) *FieldFactory {
	return &FieldFactory{tr, nameFormat}
}

// AdditionalType возвращает объект для подстановки поля additionalType в view.data
func (*FieldFactory) AdditionalType(value string) json.RawMessage {
	return json.RawMessage(fmt.Sprintf(`{ "additionalType": %q }`, value))
}

// BoolView формирует описание отображения поля Да/Нет
func (factory *FieldFactory) BoolView(ctx context.Context, yesKey, noKey string) boolean.ViewData {
	return boolean.ViewData{
		YesValue: factory.tr.TranslateString(ctx, yesKey),
		NoValue:  factory.tr.TranslateString(ctx, noKey),
	}
}

// EnumData формирует описание типа Enum
//
// format будет использоваться для перевода названий вариантов и должен иметь одну строковую подстановку
// для кода варианта.
//
// varinats должен быть срезом строк или объектов, удовлетворяющих fmt.Stringer.
func (factory *FieldFactory) EnumData(ctx context.Context, format string, variants interface{}) enum.EnumData {
	variantsValue := reflect.ValueOf(variants)
	length := variantsValue.Len()
	items := make([]enum.EnumItem, length)
	for i := 0; i < length; i++ {
		var code string
		variantValue := variantsValue.Index(i)
		switch variant := variantValue.Interface().(type) {
		case fmt.Stringer:
			code = variant.String()
		case string:
			code = variant
		}
		items[i] = enum.EnumItem{
			Code: code,
			Name: factory.tr.TranslateString(ctx, fmt.Sprintf(format, code)),
		}
	}
	return enum.EnumData{EnumItems: items}
}

// FieldBuilder конструктор поля
type FieldBuilder struct {
	field Field
}

// New новый конструктор поля
func (factory *FieldFactory) New(ctx context.Context, code string, t Type) *FieldBuilder {
	return &FieldBuilder{
		field: Field{
			Code: code,
			Type: t,
			View: FieldView{
				Name: factory.tr.TranslateString(ctx, fmt.Sprintf(factory.format, code)),
			},
		},
	}
}

// Required включить флаг обязательности
func (builder *FieldBuilder) Required() *FieldBuilder {
	builder.field.Required = true
	return builder
}

// ColumnName указать имя колонки в БД
func (builder *FieldBuilder) ColumnName(name string) *FieldBuilder {
	builder.field.ColumnName = name
	return builder
}

// Array включить флаг хранения данных списком
func (builder *FieldBuilder) Array(single bool) *FieldBuilder {
	builder.field.Array = true
	builder.field.Single = single
	return builder
}

// Searchable включить индексы по полю
func (builder *FieldBuilder) Searchable(indexed bool) *FieldBuilder {
	builder.field.Searchable = true
	builder.field.Indexed = indexed
	return builder
}

// Data добавить описание поля
func (builder *FieldBuilder) Data(data interface{}) *FieldBuilder {
	builder.field.Data, _ = json.Marshal(data)
	return builder
}

// View добавить описание отображения поля
func (builder *FieldBuilder) View(data interface{}) *FieldBuilder {
	builder.field.View.Data, _ = json.Marshal(data)
	return builder
}

// Field вернуть построенное поле
func (builder *FieldBuilder) Field() Field {
	return builder.field
}
