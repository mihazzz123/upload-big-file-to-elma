package types

//go:generate ../../tooling/bin/easyjson field.go
//go:generate ../../tooling/bin/easylocalizer field.go

import (
	"encoding/json"
	"regexp"
	"strings"

	"git.elewise.com/elma365/common/pkg/types/complextypes/enum"

	"github.com/pkg/errors"
)

const (
	IdFieldCode        = "__id" // nolint: golint
	KeyFieldCode       = "__name"
	IndexFieldCode     = "__index"
	StatusFieldCode    = "__status"
	DirectoryFieldCode = "__directory" // UUID директории (appviews_dirs), в которой лежит элемент приложения
	CreatedAtFieldCode = "__createdAt"
	CreatedByFieldCode = "__createdBy"
	UpdatedAtFieldCode = "__updatedAt"
	UpdatedByFieldCode = "__updatedBy"
	DeletedAtFieldCode = "__deletedAt"
)

// Field is a definition of application field
//
// easyjson:json
// localizer:field
// nolint: maligned // поля сгруппированы по смыслу, поэтому этот линтер отключен
type Field struct {
	Code       string `json:"code"         validate:"fieldCode,required" localizer:"id"`
	Type       Type   `json:"type"         validate:"required"`
	Searchable bool   `json:"searchable"` // Возможность поиска по этому полю
	Indexed    bool   `json:"indexed"`    // Полнотекстовый поиск
	Deleted    bool   `json:"deleted"`
	Array      bool   `json:"array"` // Для валидации данных
	Required   bool   `json:"required"`
	ColumnName string `json:"-"` // название поля в БД, если оно отличается от Code
	/**
	 * Поле c признаком `Array` может хранить как множество, так и одно значение
	 * (например для типов Phone, Email, Application) для удобства быстрого переключения
	 * отображения значений с конструкторе формы и исключения преобразования формата
	 * храненияф данных. Признак `Single` как раз и определяет сколько элементов хранится в
	 * поле типа массив.
	 *
	 * Refactoring: поле стоит переименовать в ArrayWithSingleItem
	 */
	Single        bool            `json:"single"`
	Default       json.RawMessage `json:"defaultValue"`
	CalcByFormula bool            `json:"calcByFormula"`
	Formula       string          `json:"formula" patch:"po"`
	Data          json.RawMessage `json:"data"` // Данные, специфичные для конкретного типа поля
	View          FieldView       `json:"view" `
}

// FieldView - данные для представления значения поля или контролов для поля
//
// easyjson:json
// localizer:field_view
type FieldView struct {
	Name          string          `json:"name,omitempty" patch:"po" localizer:"value"`
	Sort          int             `json:"sort,omitempty"`
	Tooltip       string          `json:"tooltip,omitempty" patch:"po" localizer:"value"`
	TooltipAsHTML bool            `json:"tooltipAsHtml,omitempty"` // Обрабатывать поле tooltip как html
	System        bool            `json:"system,omitempty"`        // Признак системного поля
	Hidden        bool            `json:"hidden,omitempty"`        // Видимость поля на клиенте (пример - обратные ссылки типа "Приложение")
	Data          json.RawMessage `json:"data,omitempty"`          // Параметры отображения, специфичные для конкретного типа поля
	Disabled      bool            `json:"disabled,omitempty"`      // Признак блокировки для редактирования поля
	Input         *bool           `json:"input,omitempty"`         // Признак входного поля
	Output        *bool           `json:"output,omitempty"`        // Признак выходного поля
}

// IsKey проверяет признак ключевого поля
func (f Field) IsKey() bool {
	return f.Code == KeyFieldCode
}

// GetData возвращает данные типа поля
func (f Field) GetData() (interface{}, error) {
	return f.Type.UnmarshalData(f.Data)
}

// GetViewData возвращает данные для представления, спецефичные для каждого типа
func (f Field) GetViewData() (interface{}, error) {
	return f.Type.UnmarshalViewData(f.View.Data)
}

// GetDefault получить значение по умолчанию
func (f *Field) GetDefault(timeForDateTime json.RawMessage) (json.RawMessage, error) {
	switch f.Type {
	default:
		return f.Default, nil
	case DateTime:
		var viewData DataTimeViewData
		if len(f.View.Data) == 0 {
			return f.Default, nil
		}
		if err := json.Unmarshal(f.View.Data, &viewData); err != nil {
			return nil, err
		}
		if !viewData.SetCurrentDatetime {
			return f.Default, nil
		}
		return timeForDateTime, nil
	case Enum:
		var fData enum.EnumData
		if err := json.Unmarshal(f.Data, &fData); err != nil {
			return nil, err
		}
		var enums []enum.Enum
		for _, v := range fData.EnumItems {
			if v.Checked {
				enums = append(enums, enum.Enum{Code: v.Code, Name: v.Name})
			}
		}
		return json.Marshal(enums)
	}
}

// Validate value for this field
func (f Field) Validate(value json.RawMessage) (interface{}, error) {
	// nolint:goconst
	if value == nil || string(value) == "null" {
		return nil, nil
	}
	if f.Type == String && string(value) == `""` {
		return "", nil
	}

	var res interface{}
	var err error
	if f.Array {
		res, err = f.validateArray(value)
	} else {
		res, err = f.validateSingle(value)
	}

	return res, errors.WithStack(err)
}

// ValidateDefault value of the field
func (f Field) ValidateDefault() error {
	// nolint:goconst
	if len(f.Default) == 0 || string(f.Default) == "null" || string(f.Default) == `""` {
		return nil
	}

	var err error
	if f.Array {
		_, err = f.validateArray(f.Default)
	} else {
		_, err = f.validateSingle(f.Default)
	}

	return errors.WithStack(err)
}

func (f Field) validateArray(value json.RawMessage) (interface{}, error) {
	valarr := make([]json.RawMessage, 0)
	if err := json.Unmarshal(value, &valarr); err != nil {
		return nil, err
	}

	var res interface{}
	var ares []interface{}
	for _, val := range valarr {
		// nolint:goconst
		if string(val) == "null" {
			return nil, errors.WithStack(errors.New("null - is invalid value of array"))
		}
		valitdatedFieldValue, err := f.validateSingle(val)
		if err != nil {
			return nil, err
		}
		ares = append(ares, valitdatedFieldValue)
	}
	res = ares
	return res, nil
}

func (f Field) validateSingle(value json.RawMessage) (interface{}, error) {
	res, err := f.Type.Validate(value)
	err = f.errorParse(err)
	return res, err
}

//nolint: gochecknoglobals // регулярки должны быть глобальными
var (
	jsonErrorRE = regexp.MustCompile(`^json: cannot unmarshal (.*) into Go value of type (.*)$`)
	uuidErrorRE = regexp.MustCompile(`^uuid: (.*): invalid$`)
	timeErrorRE = regexp.MustCompile(`^parsing time "(.*)" as ""2006-01-02T15:04:05Z07:00"": cannot parse "(.*)" as "(.*)"$`)
	timeParts   = map[string]string{
		"2006":   "year",
		"01":     "month",
		"02":     "day",
		"15":     "hour",
		"04":     "minutes",
		"05":     "seconds",
		"Z07:00": "time zone",
	}
)

func (f Field) errorParse(err error) error {
	if err == nil {
		return nil
	}

	if parts := jsonErrorRE.FindStringSubmatch(err.Error()); len(parts) == 3 {
		err = errors.Errorf("cannot unmarshal %s into field of type %s", parts[1], f.Type)
	} else if parts := uuidErrorRE.FindStringSubmatch(err.Error()); len(parts) == 2 {
		err = errors.New("invalid uuid")
	} else if parts := timeErrorRE.FindStringSubmatch(err.Error()); len(parts) == 4 {
		timePart, ok := timeParts[parts[3]]
		if !ok {
			timePart = parts[3]
		}
		err = errors.Errorf("cannot parse %q as %s", strings.Trim(parts[2], `"`), timePart)
	}

	return errors.Wrap(err, f.Code)
}

// методы для протобафа

// Marshal marshaler interfacer
func (f Field) Marshal() ([]byte, error) {
	return json.Marshal(f)
}

// MarshalTo protobuf marshaler
func (f *Field) MarshalTo(data []byte) (n int, err error) {
	d, err := json.Marshal(f)
	if err != nil {
		return 0, err
	}
	return copy(data, d), nil
}

// Unmarshal unmarshaller interface
func (f *Field) Unmarshal(data []byte) error {
	return json.Unmarshal(data, f)
}

// Size resturn size for protobuf
func (f *Field) Size() int {
	if f == nil {
		return 0
	}

	d, _ := json.Marshal(f)

	return len(d)
}

// Compare сравнить значения на равенство
func (f *Field) Compare(value1, value2 json.RawMessage) (bool, error) {
	if value1 == nil {
		return value2 == nil, nil
	}
	return f.Type.Compare(value1, value2)
}

// CanReplaceTo можно ли обновить описание поля
func (f *Field) CanReplaceTo(candidate *Field) (bool, error) {
	if ok, err := f.Type.CanReplace(*f, *candidate); !ok || err != nil {
		return false, errors.WithStack(err)
	}
	if ok, err := f.Type.CanReplaceData(f.Data, candidate.Data); !ok || err != nil {
		return false, errors.WithStack(err)
	}
	if ok, err := f.Type.CanReplaceViewData(f.View.Data, candidate.View.Data); !ok || err != nil {
		return false, errors.WithStack(err)
	}
	return true, nil
}

// GetFieldName возвращается название поля в БД
// Если явно не указано название поля (ColumnName), то возвращается его код (Code)
func (f Field) GetFieldName() string {
	if len(f.ColumnName) > 0 {
		return f.ColumnName
	}
	return f.Code
}
