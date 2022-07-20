# types
`import "git.elewise.com/elma365/common/pkg/types"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [type CategoryData](#CategoryData)
* [type Content](#Content)
  * [func NewContent(fields []Field) *Content](#NewContent)
  * [func (c *Content) Fields() []Field](#Content.Fields)
  * [func (c *Content) Validate(inputData map[string]json.RawMessage, checkOriginalRequired bool) error](#Content.Validate)
* [type DataTimeData](#DataTimeData)
* [type DataTimeViewData](#DataTimeViewData)
* [type DefaultTimeType](#DefaultTimeType)
* [type Field](#Field)
  * [func (f *Field) CanReplaceTo(candidate *Field) (bool, error)](#Field.CanReplaceTo)
  * [func (f *Field) Compare(value1, value2 json.RawMessage) (bool, error)](#Field.Compare)
  * [func (f Field) GetData() (interface{}, error)](#Field.GetData)
  * [func (f *Field) GetDefault(timeForDateTime json.RawMessage) (json.RawMessage, error)](#Field.GetDefault)
  * [func (f Field) GetExtracts(target string) (patch.Extracts, error)](#Field.GetExtracts)
  * [func (f Field) GetFieldName() string](#Field.GetFieldName)
  * [func (f Field) GetViewData() (interface{}, error)](#Field.GetViewData)
  * [func (f Field) IsKey() bool](#Field.IsKey)
  * [func (f Field) Marshal() ([]byte, error)](#Field.Marshal)
  * [func (v Field) MarshalEasyJSON(w *jwriter.Writer)](#Field.MarshalEasyJSON)
  * [func (v Field) MarshalJSON() ([]byte, error)](#Field.MarshalJSON)
  * [func (f *Field) MarshalTo(data []byte) (n int, err error)](#Field.MarshalTo)
  * [func (f *Field) Size() int](#Field.Size)
  * [func (f *Field) Unmarshal(data []byte) error](#Field.Unmarshal)
  * [func (v *Field) UnmarshalEasyJSON(l *jlexer.Lexer)](#Field.UnmarshalEasyJSON)
  * [func (v *Field) UnmarshalJSON(data []byte) error](#Field.UnmarshalJSON)
  * [func (f Field) Validate(value json.RawMessage) (interface{}, error)](#Field.Validate)
  * [func (f Field) ValidateDefault() error](#Field.ValidateDefault)
* [type FieldBuilder](#FieldBuilder)
  * [func (builder *FieldBuilder) Array(single bool) *FieldBuilder](#FieldBuilder.Array)
  * [func (builder *FieldBuilder) ColumnName(name string) *FieldBuilder](#FieldBuilder.ColumnName)
  * [func (builder *FieldBuilder) Data(data interface{}) *FieldBuilder](#FieldBuilder.Data)
  * [func (builder *FieldBuilder) Field() Field](#FieldBuilder.Field)
  * [func (builder *FieldBuilder) Required() *FieldBuilder](#FieldBuilder.Required)
  * [func (builder *FieldBuilder) Searchable(indexed bool) *FieldBuilder](#FieldBuilder.Searchable)
  * [func (builder *FieldBuilder) View(data interface{}) *FieldBuilder](#FieldBuilder.View)
* [type FieldDataTimeType](#FieldDataTimeType)
* [type FieldFactory](#FieldFactory)
  * [func NewFieldFactory(tr Translator, nameFormat string) *FieldFactory](#NewFieldFactory)
  * [func (*FieldFactory) AdditionalType(value string) json.RawMessage](#FieldFactory.AdditionalType)
  * [func (factory *FieldFactory) BoolView(ctx context.Context, yesKey, noKey string) boolean.ViewData](#FieldFactory.BoolView)
  * [func (factory *FieldFactory) EnumData(ctx context.Context, format string, variants interface{}) enum.EnumData](#FieldFactory.EnumData)
  * [func (factory *FieldFactory) New(ctx context.Context, code string, t Type) *FieldBuilder](#FieldFactory.New)
* [type FieldFloatType](#FieldFloatType)
* [type FieldListWithValue](#FieldListWithValue)
  * [func NewFieldListWithValue(fields []Field, values map[string]json.RawMessage) FieldListWithValue](#NewFieldListWithValue)
  * [func (fl FieldListWithValue) Get(code string) (FieldWithValue, error)](#FieldListWithValue.Get)
  * [func (fl FieldListWithValue) GetField(code string) (Field, error)](#FieldListWithValue.GetField)
  * [func (fl FieldListWithValue) GetFieldList() []Field](#FieldListWithValue.GetFieldList)
  * [func (fl FieldListWithValue) GetValue(code string) (json.RawMessage, error)](#FieldListWithValue.GetValue)
  * [func (fl FieldListWithValue) GetValuesMap() map[string]json.RawMessage](#FieldListWithValue.GetValuesMap)
* [type FieldStringType](#FieldStringType)
* [type FieldView](#FieldView)
  * [func (v FieldView) MarshalEasyJSON(w *jwriter.Writer)](#FieldView.MarshalEasyJSON)
  * [func (v FieldView) MarshalJSON() ([]byte, error)](#FieldView.MarshalJSON)
  * [func (v *FieldView) UnmarshalEasyJSON(l *jlexer.Lexer)](#FieldView.UnmarshalEasyJSON)
  * [func (v *FieldView) UnmarshalJSON(data []byte) error](#FieldView.UnmarshalJSON)
* [type FieldWithValue](#FieldWithValue)
* [type Fields](#Fields)
  * [func (o *Fields) CanReplaceTo(candidate Fields) (bool, error)](#Fields.CanReplaceTo)
  * [func (o *Fields) FindByCode(code string) *Field](#Fields.FindByCode)
  * [func (o *Fields) GetIndexByCode(code string) int](#Fields.GetIndexByCode)
  * [func (o Fields) Has(code string) bool](#Fields.Has)
  * [func (o Fields) Marshal() ([]byte, error)](#Fields.Marshal)
  * [func (v Fields) MarshalEasyJSON(w *jwriter.Writer)](#Fields.MarshalEasyJSON)
  * [func (v Fields) MarshalJSON() ([]byte, error)](#Fields.MarshalJSON)
  * [func (o *Fields) MarshalTo(data []byte) (n int, err error)](#Fields.MarshalTo)
  * [func (o Fields) Merge(fields Fields) (Fields, error)](#Fields.Merge)
  * [func (o *Fields) Scan(pSrc interface{}) error](#Fields.Scan)
  * [func (o *Fields) Size() int](#Fields.Size)
  * [func (o *Fields) Unmarshal(data []byte) error](#Fields.Unmarshal)
  * [func (v *Fields) UnmarshalEasyJSON(l *jlexer.Lexer)](#Fields.UnmarshalEasyJSON)
  * [func (v *Fields) UnmarshalJSON(data []byte) error](#Fields.UnmarshalJSON)
  * [func (o Fields) Value() (value driver.Value, err error)](#Fields.Value)
* [type FloatViewData](#FloatViewData)
* [type StringViewData](#StringViewData)
* [type TableFieldData](#TableFieldData)
* [type TableResult](#TableResult)
* [type TableResultFormula](#TableResultFormula)
* [type TableResultKind](#TableResultKind)
* [type TableValue](#TableValue)
* [type TableValueRow](#TableValueRow)
* [type TableViewData](#TableViewData)
* [type Translator](#Translator)
* [type Type](#Type)
  * [func FromString(s string) (Type, error)](#FromString)
  * [func (t Type) CanReplace(field, candidate Field) (bool, error)](#Type.CanReplace)
  * [func (t Type) CanReplaceData(sourceData, candidateData json.RawMessage) (bool, error)](#Type.CanReplaceData)
  * [func (t Type) CanReplaceViewData(sourceData, candidateData json.RawMessage) (bool, error)](#Type.CanReplaceViewData)
  * [func (t Type) Compare(value1, value2 json.RawMessage) (bool, error)](#Type.Compare)
  * [func (Type) Generate(r *rand.Rand, _ int) reflect.Value](#Type.Generate)
  * [func (t Type) GetImplement() (TypeImplement, error)](#Type.GetImplement)
  * [func (t Type) MarshalJSON() ([]byte, error)](#Type.MarshalJSON)
  * [func (t Type) String() string](#Type.String)
  * [func (t Type) UnmarshalData(rawData json.RawMessage) (interface{}, error)](#Type.UnmarshalData)
  * [func (t *Type) UnmarshalJSON(data []byte) error](#Type.UnmarshalJSON)
  * [func (t Type) UnmarshalViewData(rawViewData json.RawMessage) (interface{}, error)](#Type.UnmarshalViewData)
  * [func (t Type) Validate(message json.RawMessage) (interface{}, error)](#Type.Validate)
* [type TypeImplement](#TypeImplement)
  * [func NewAccountType() TypeImplement](#NewAccountType)
  * [func NewBooleanType() TypeImplement](#NewBooleanType)
  * [func NewCategoryType() TypeImplement](#NewCategoryType)
  * [func NewDateTimeType() TypeImplement](#NewDateTimeType)
  * [func NewDurationType() TypeImplement](#NewDurationType)
  * [func NewEmailType() TypeImplement](#NewEmailType)
  * [func NewEnumType() TypeImplement](#NewEnumType)
  * [func NewFileType() TypeImplement](#NewFileType)
  * [func NewFloatType() TypeImplement](#NewFloatType)
  * [func NewFullNameType() TypeImplement](#NewFullNameType)
  * [func NewIntegerType() TypeImplement](#NewIntegerType)
  * [func NewJSONType() TypeImplement](#NewJSONType)
  * [func NewLinkType() TypeImplement](#NewLinkType)
  * [func NewMoneyType() TypeImplement](#NewMoneyType)
  * [func NewPhoneType() TypeImplement](#NewPhoneType)
  * [func NewRefItemType() TypeImplement](#NewRefItemType)
  * [func NewStatusType() TypeImplement](#NewStatusType)
  * [func NewStringType() TypeImplement](#NewStringType)
  * [func NewSysCollectionType() TypeImplement](#NewSysCollectionType)
  * [func NewSysOSNodeType() TypeImplement](#NewSysOSNodeType)
  * [func NewSysUserType() TypeImplement](#NewSysUserType)
  * [func NewTableType() TypeImplement](#NewTableType)
  * [func NewTagType() TypeImplement](#NewTagType)
  * [func NewVersionType() TypeImplement](#NewVersionType)

#### <a name="pkg-examples">Examples</a>
* [FieldFactory](#example-fieldfactory)

#### <a name="pkg-files">Package files</a>
[content.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/content.go) [field.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field.go) [field_easyjson.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field_easyjson.go) [field_extractor.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field_extractor.go) [field_factory.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field_factory.go) [field_with_value.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field_with_value.go) [fields.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/fields.go) [fields_easyjson.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/fields_easyjson.go) [implement_account.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_account.go) [implement_boolean.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_boolean.go) [implement_category.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_category.go) [implement_dateTime.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_dateTime.go) [implement_duration.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_duration.go) [implement_email.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_email.go) [implement_enum.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_enum.go) [implement_file.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_file.go) [implement_float.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_float.go) [implement_fullName.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_fullName.go) [implement_integer.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_integer.go) [implement_json.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_json.go) [implement_link.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_link.go) [implement_money.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_money.go) [implement_phone.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_phone.go) [implement_refItem.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_refItem.go) [implement_status.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_status.go) [implement_string.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_string.go) [implement_sysCollection.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_sysCollection.go) [implement_sysOSNode.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_sysOSNode.go) [implement_sysUser.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_sysUser.go) [implement_table.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_table.go) [implement_tag.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_tag.go) [implement_version.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_version.go) [implements.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implements.go) [type.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/type.go)


## <a name="pkg-constants">Constants</a>
``` go
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
```




## <a name="CategoryData">type</a> [CategoryData](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_category.go?s=304:374#L17)
``` go
type CategoryData struct {
    Fields Fields `json:"fields" patch:"po"`
}

```
CategoryData дополнительные параметры поля типа CATEGORY










## <a name="Content">type</a> [Content](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/content.go?s=92:121#L10)
``` go
type Content map[string]Field
```
Content is a map of fields







### <a name="NewContent">func</a> [NewContent](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/content.go?s=154:194#L13)
``` go
func NewContent(fields []Field) *Content
```
NewContent is a constructor





### <a name="Content.Fields">func</a> (\*Content) [Fields](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/content.go?s=338:372#L24)
``` go
func (c *Content) Fields() []Field
```
Fields return list of fields




### <a name="Content.Validate">func</a> (\*Content) [Validate](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/content.go?s=495:593#L34)
``` go
func (c *Content) Validate(inputData map[string]json.RawMessage, checkOriginalRequired bool) error
```
Validate data




## <a name="DataTimeData">type</a> [DataTimeData](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_dateTime.go?s=1449:1477#L43)
``` go
type DataTimeData struct {
}

```
DataTimeData дополнительные настройки поля типа DATETIME










## <a name="DataTimeViewData">type</a> [DataTimeViewData](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_dateTime.go?s=1073:1356#L35)
``` go
type DataTimeViewData struct {
    AdditionalType     FieldDataTimeType `json:"additionalType"`
    SetCurrentDatetime bool              `json:"setCurrentDatetime"`
    TimeOptional       bool              `json:"timeOptional"`
    DefaultTimeType    DefaultTimeType   `json:"defaultTimeType"`
}

```
DataTimeViewData описание представления поля типа DATETIME










## <a name="DefaultTimeType">type</a> [DefaultTimeType](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_dateTime.go?s=617:644#L23)
``` go
type DefaultTimeType string
```
DefaultTimeType предустановленные значения времени


``` go
const (
    // DefaultTimeTypeStartOfDay 00:00:00
    DefaultTimeTypeStartOfDay DefaultTimeType = "startOfDay"
    // DefaultTimeTypeEndOfDay 23:59:59
    DefaultTimeTypeEndOfDay DefaultTimeType = "endOfDay"
    // DefaultTimeTypeNone не устанавливать время по умолчанию
    DefaultTimeTypeNone DefaultTimeType = "none"
)
```









## <a name="Field">type</a> [Field](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field.go?s=884:2572#L32)
``` go
type Field struct {
    Code       string `json:"code"         validate:"fieldCode,required"`
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
    Formula       string          `json:"formula"`
    Data          json.RawMessage `json:"data"` // Данные, специфичные для конкретного типа поля
    View          FieldView       `json:"view" `
}

```
Field is a definition of application field

easyjson:json
nolint: maligned // поля сгруппированы по смыслу, поэтому этот линтер отключен










### <a name="Field.CanReplaceTo">func</a> (\*Field) [CanReplaceTo](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field.go?s=8648:8708#L263)
``` go
func (f *Field) CanReplaceTo(candidate *Field) (bool, error)
```
CanReplaceTo можно ли обновить описание поля




### <a name="Field.Compare">func</a> (\*Field) [Compare](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field.go?s=8408:8477#L255)
``` go
func (f *Field) Compare(value1, value2 json.RawMessage) (bool, error)
```
Compare сравнить значения на равенство




### <a name="Field.GetData">func</a> (Field) [GetData](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field.go?s=3895:3940#L78)
``` go
func (f Field) GetData() (interface{}, error)
```
GetData возвращает данные типа поля




### <a name="Field.GetDefault">func</a> (\*Field) [GetDefault](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field.go?s=4294:4378#L88)
``` go
func (f *Field) GetDefault(timeForDateTime json.RawMessage) (json.RawMessage, error)
```
GetDefault получить значение по умолчанию




### <a name="Field.GetExtracts">func</a> (Field) [GetExtracts](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field_extractor.go?s=233:298#L13)
``` go
func (f Field) GetExtracts(target string) (patch.Extracts, error)
```
GetExtracts - реализация интерфейса patch.Extractor




### <a name="Field.GetFieldName">func</a> (Field) [GetFieldName](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field.go?s=9310:9346#L278)
``` go
func (f Field) GetFieldName() string
```
GetFieldName возвращается название поля в БД
Если явно не указано название поля (ColumnName), то возвращается его код (Code)




### <a name="Field.GetViewData">func</a> (Field) [GetViewData](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field.go?s=4121:4170#L83)
``` go
func (f Field) GetViewData() (interface{}, error)
```
GetViewData возвращает данные для представления, спецефичные для каждого типа




### <a name="Field.IsKey">func</a> (Field) [IsKey](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field.go?s=3768:3795#L73)
``` go
func (f Field) IsKey() bool
```
IsKey проверяет признак ключевого поля




### <a name="Field.Marshal">func</a> (Field) [Marshal](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field.go?s=7829:7869#L225)
``` go
func (f Field) Marshal() ([]byte, error)
```
Marshal marshaler interfacer




### <a name="Field.MarshalEasyJSON">func</a> (Field) [MarshalEasyJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field_easyjson.go?s=6889:6938#L314)
``` go
func (v Field) MarshalEasyJSON(w *jwriter.Writer)
```
MarshalEasyJSON supports easyjson.Marshaler interface




### <a name="Field.MarshalJSON">func</a> (Field) [MarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field_easyjson.go?s=6654:6698#L307)
``` go
func (v Field) MarshalJSON() ([]byte, error)
```
MarshalJSON supports json.Marshaler interface




### <a name="Field.MarshalTo">func</a> (\*Field) [MarshalTo](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field.go?s=7931:7988#L230)
``` go
func (f *Field) MarshalTo(data []byte) (n int, err error)
```
MarshalTo protobuf marshaler




### <a name="Field.Size">func</a> (\*Field) [Size](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field.go?s=8236:8262#L244)
``` go
func (f *Field) Size() int
```
Size resturn size for protobuf




### <a name="Field.Unmarshal">func</a> (\*Field) [Unmarshal](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field.go?s=8120:8164#L239)
``` go
func (f *Field) Unmarshal(data []byte) error
```
Unmarshal unmarshaller interface




### <a name="Field.UnmarshalEasyJSON">func</a> (\*Field) [UnmarshalEasyJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field_easyjson.go?s=7292:7342#L326)
``` go
func (v *Field) UnmarshalEasyJSON(l *jlexer.Lexer)
```
UnmarshalEasyJSON supports easyjson.Unmarshaler interface




### <a name="Field.UnmarshalJSON">func</a> (\*Field) [UnmarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field_easyjson.go?s=7062:7110#L319)
``` go
func (v *Field) UnmarshalJSON(data []byte) error
```
UnmarshalJSON supports json.Unmarshaler interface




### <a name="Field.Validate">func</a> (Field) [Validate](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field.go?s=5061:5128#L120)
``` go
func (f Field) Validate(value json.RawMessage) (interface{}, error)
```
Validate value for this field




### <a name="Field.ValidateDefault">func</a> (Field) [ValidateDefault](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field.go?s=5499:5537#L141)
``` go
func (f Field) ValidateDefault() error
```
ValidateDefault value of the field




## <a name="FieldBuilder">type</a> [FieldBuilder](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field_factory.go?s=2560:2601#L73)
``` go
type FieldBuilder struct {
    // contains filtered or unexported fields
}

```
FieldBuilder конструктор поля










### <a name="FieldBuilder.Array">func</a> (\*FieldBuilder) [Array](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field_factory.go?s=3358:3419#L103)
``` go
func (builder *FieldBuilder) Array(single bool) *FieldBuilder
```
Array включить флаг хранения данных списком




### <a name="FieldBuilder.ColumnName">func</a> (\*FieldBuilder) [ColumnName](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field_factory.go?s=3157:3223#L97)
``` go
func (builder *FieldBuilder) ColumnName(name string) *FieldBuilder
```
ColumnName указать имя колонки в БД




### <a name="FieldBuilder.Data">func</a> (\*FieldBuilder) [Data](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field_factory.go?s=3766:3831#L117)
``` go
func (builder *FieldBuilder) Data(data interface{}) *FieldBuilder
```
Data добавить описание поля




### <a name="FieldBuilder.Field">func</a> (\*FieldBuilder) [Field](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field_factory.go?s=4163:4205#L129)
``` go
func (builder *FieldBuilder) Field() Field
```
Field вернуть построенное поле




### <a name="FieldBuilder.Required">func</a> (\*FieldBuilder) [Required](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field_factory.go?s=2992:3045#L91)
``` go
func (builder *FieldBuilder) Required() *FieldBuilder
```
Required включить флаг обязательности




### <a name="FieldBuilder.Searchable">func</a> (\*FieldBuilder) [Searchable](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field_factory.go?s=3560:3627#L110)
``` go
func (builder *FieldBuilder) Searchable(indexed bool) *FieldBuilder
```
Searchable включить индексы по полю




### <a name="FieldBuilder.View">func</a> (\*FieldBuilder) [View](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field_factory.go?s=3971:4036#L123)
``` go
func (builder *FieldBuilder) View(data interface{}) *FieldBuilder
```
View добавить описание отображения поля




## <a name="FieldDataTimeType">type</a> [FieldDataTimeType](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_dateTime.go?s=180:209#L11)
``` go
type FieldDataTimeType string
```
FieldDataTimeType дополнительный тип представления поля DATETIME


``` go
const (
    // FieldDataTimeTypeDate только дата
    FieldDataTimeTypeDate FieldDataTimeType = "date"
    // FieldDataTimeTypeTime только время
    FieldDataTimeTypeTime FieldDataTimeType = "time"
    // FieldDataTimeTypeDateTime дата и время
    FieldDataTimeTypeDateTime FieldDataTimeType = "datetime"
)
```









## <a name="FieldFactory">type</a> [FieldFactory](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field_factory.go?s=402:464#L19)
``` go
type FieldFactory struct {
    // contains filtered or unexported fields
}

```
FieldFactory фабрика описания полей



##### Example FieldFactory:
``` go
ctx := context.Background()
ff := types.NewFieldFactory(FakeTranslator{}, "my-collection.fields@%s")

stringField := ff.New(ctx, "foo", types.String).Required().Searchable(false).View(ff.AdditionalType("text")).Field()
blob, _ := stringField.Marshal()
fmt.Print(string(blob))
// Output: {"code":"foo","type":"STRING","searchable":true,"indexed":false,"deleted":false,"array":false,"required":true,"single":false,"defaultValue":null,"calcByFormula":false,"formula":"","data":null,"view":{"name":"my-collection.fields@foo","data":{"additionalType":"text"}}}
```

Output:

```
{"code":"foo","type":"STRING","searchable":true,"indexed":false,"deleted":false,"array":false,"required":true,"single":false,"defaultValue":null,"calcByFormula":false,"formula":"","data":null,"view":{"name":"my-collection.fields@foo","data":{"additionalType":"text"}}}
```





### <a name="NewFieldFactory">func</a> [NewFieldFactory](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field_factory.go?s=770:838#L28)
``` go
func NewFieldFactory(tr Translator, nameFormat string) *FieldFactory
```
NewFieldFactory создать новую фабрику описания полей

nameFormat будет использоваться для перевода названий полей и должен иметь одну строковую
подстановку для кода поля.





### <a name="FieldFactory.AdditionalType">func</a> (\*FieldFactory) [AdditionalType](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field_factory.go?s=1001:1066#L33)
``` go
func (*FieldFactory) AdditionalType(value string) json.RawMessage
```
AdditionalType возвращает объект для подстановки поля additionalType в view.data




### <a name="FieldFactory.BoolView">func</a> (\*FieldFactory) [BoolView](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field_factory.go?s=1236:1333#L38)
``` go
func (factory *FieldFactory) BoolView(ctx context.Context, yesKey, noKey string) boolean.ViewData
```
BoolView формирует описание отображения поля Да/Нет




### <a name="FieldFactory.EnumData">func</a> (\*FieldFactory) [EnumData](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field_factory.go?s=1894:2003#L51)
``` go
func (factory *FieldFactory) EnumData(ctx context.Context, format string, variants interface{}) enum.EnumData
```
EnumData формирует описание типа Enum

format будет использоваться для перевода названий вариантов и должен иметь одну строковую подстановку
для кода варианта.

varinats должен быть срезом строк или объектов, удовлетворяющих fmt.Stringer.




### <a name="FieldFactory.New">func</a> (\*FieldFactory) [New](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field_factory.go?s=2653:2741#L78)
``` go
func (factory *FieldFactory) New(ctx context.Context, code string, t Type) *FieldBuilder
```
New новый конструктор поля




## <a name="FieldFloatType">type</a> [FieldFloatType](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_float.go?s=95:121#L6)
``` go
type FieldFloatType string
```
FieldFloatType разновидности типов


``` go
const (
    // FieldFloatTypeFloat float
    FieldFloatTypeFloat FieldFloatType = "float"
    // FieldFloatTypeInteger integer
    FieldFloatTypeInteger FieldFloatType = "integer"
)
```









## <a name="FieldListWithValue">type</a> [FieldListWithValue](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field_with_value.go?s=290:330#L15)
``` go
type FieldListWithValue []FieldWithValue
```
FieldListWithValue список пар из описания поля и значения







### <a name="NewFieldListWithValue">func</a> [NewFieldListWithValue](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field_with_value.go?s=408:504#L18)
``` go
func NewFieldListWithValue(fields []Field, values map[string]json.RawMessage) FieldListWithValue
```
NewFieldListWithValue возвращает новый FieldListWithValue





### <a name="FieldListWithValue.Get">func</a> (FieldListWithValue) [Get](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field_with_value.go?s=731:800#L28)
``` go
func (fl FieldListWithValue) Get(code string) (FieldWithValue, error)
```
Get получить поле со значением по коду




### <a name="FieldListWithValue.GetField">func</a> (FieldListWithValue) [GetField](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field_with_value.go?s=1339:1404#L48)
``` go
func (fl FieldListWithValue) GetField(code string) (Field, error)
```
GetField получить поле по коду




### <a name="FieldListWithValue.GetFieldList">func</a> (FieldListWithValue) [GetFieldList](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field_with_value.go?s=1637:1688#L58)
``` go
func (fl FieldListWithValue) GetFieldList() []Field
```
GetFieldList возвращает список полей




### <a name="FieldListWithValue.GetValue">func</a> (FieldListWithValue) [GetValue](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field_with_value.go?s=1044:1119#L38)
``` go
func (fl FieldListWithValue) GetValue(code string) (json.RawMessage, error)
```
GetValue получить значение по коду поля




### <a name="FieldListWithValue.GetValuesMap">func</a> (FieldListWithValue) [GetValuesMap](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field_with_value.go?s=1868:1938#L67)
``` go
func (fl FieldListWithValue) GetValuesMap() map[string]json.RawMessage
```
GetValuesMap получить мапу значений




## <a name="FieldStringType">type</a> [FieldStringType](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_string.go?s=159:186#L10)
``` go
type FieldStringType string
```
FieldStringType дополнительный тип для представления


``` go
const (
    // FieldStringTypeString одна строка
    FieldStringTypeString FieldStringType = "string"
    // FieldStringTypeText многострочный
    FieldStringTypeText FieldStringType = "text"
)
```









## <a name="FieldView">type</a> [FieldView](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field.go?s=2724:3695#L61)
``` go
type FieldView struct {
    Name          string          `json:"name,omitempty" patch:"po"`
    Sort          int             `json:"sort,omitempty"`
    Tooltip       string          `json:"tooltip,omitempty" patch:"po"`
    TooltipAsHTML bool            `json:"tooltipAsHtml,omitempty"` // Обрабатывать поле tooltip как html
    System        bool            `json:"system,omitempty"`        // Признак системного поля
    Hidden        bool            `json:"hidden,omitempty"`        // Видимость поля на клиенте (пример - обратные ссылки типа "Приложение")
    Data          json.RawMessage `json:"data,omitempty"`          // Параметры отображения, специфичные для конкретного типа поля
    Disabled      bool            `json:"disabled,omitempty"`      // Признак блокировки для редактирования поля
}

```
FieldView - данные для представления значения поля или контролов для поля

easyjson:json










### <a name="FieldView.MarshalEasyJSON">func</a> (FieldView) [MarshalEasyJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field_easyjson.go?s=3255:3308#L158)
``` go
func (v FieldView) MarshalEasyJSON(w *jwriter.Writer)
```
MarshalEasyJSON supports easyjson.Marshaler interface




### <a name="FieldView.MarshalJSON">func</a> (FieldView) [MarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field_easyjson.go?s=3017:3065#L151)
``` go
func (v FieldView) MarshalJSON() ([]byte, error)
```
MarshalJSON supports json.Marshaler interface




### <a name="FieldView.UnmarshalEasyJSON">func</a> (\*FieldView) [UnmarshalEasyJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field_easyjson.go?s=3664:3718#L170)
``` go
func (v *FieldView) UnmarshalEasyJSON(l *jlexer.Lexer)
```
UnmarshalEasyJSON supports easyjson.Unmarshaler interface




### <a name="FieldView.UnmarshalJSON">func</a> (\*FieldView) [UnmarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field_easyjson.go?s=3431:3483#L163)
``` go
func (v *FieldView) UnmarshalJSON(data []byte) error
```
UnmarshalJSON supports json.Unmarshaler interface




## <a name="FieldWithValue">type</a> [FieldWithValue](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field_with_value.go?s=129:195#L9)
``` go
type FieldWithValue struct {
    Field Field
    Value json.RawMessage
}

```
FieldWithValue пара из описания поля и значения










## <a name="Fields">type</a> [Fields](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/fields.go?s=194:213#L15)
``` go
type Fields []Field
```
Fields is slice of fields

easyjson:json










### <a name="Fields.CanReplaceTo">func</a> (\*Fields) [CanReplaceTo](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/fields.go?s=1780:1841#L96)
``` go
func (o *Fields) CanReplaceTo(candidate Fields) (bool, error)
```
CanReplaceTo можно ли обновить описание полей




### <a name="Fields.FindByCode">func</a> (\*Fields) [FindByCode](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/fields.go?s=282:329#L18)
``` go
func (o *Fields) FindByCode(code string) *Field
```
FindByCode поиск поля по заданному коду




### <a name="Fields.GetIndexByCode">func</a> (\*Fields) [GetIndexByCode](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/fields.go?s=497:545#L29)
``` go
func (o *Fields) GetIndexByCode(code string) int
```
GetIndexByCode получить индекс поля по коду




### <a name="Fields.Has">func</a> (Fields) [Has](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/fields.go?s=2877:2914#L139)
``` go
func (o Fields) Has(code string) bool
```
Has проверка на существование нужного филда по коду




### <a name="Fields.Marshal">func</a> (Fields) [Marshal](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/fields.go?s=1189:1230#L66)
``` go
func (o Fields) Marshal() ([]byte, error)
```
Marshal marshaler interfacer




### <a name="Fields.MarshalEasyJSON">func</a> (Fields) [MarshalEasyJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/fields_easyjson.go?s=1505:1555#L71)
``` go
func (v Fields) MarshalEasyJSON(w *jwriter.Writer)
```
MarshalEasyJSON supports easyjson.Marshaler interface




### <a name="Fields.MarshalJSON">func</a> (Fields) [MarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/fields_easyjson.go?s=1271:1316#L64)
``` go
func (v Fields) MarshalJSON() ([]byte, error)
```
MarshalJSON supports json.Marshaler interface




### <a name="Fields.MarshalTo">func</a> (\*Fields) [MarshalTo](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/fields.go?s=1292:1350#L71)
``` go
func (o *Fields) MarshalTo(data []byte) (n int, err error)
```
MarshalTo protobuf marshaler




### <a name="Fields.Merge">func</a> (Fields) [Merge](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/fields.go?s=2258:2310#L115)
``` go
func (o Fields) Merge(fields Fields) (Fields, error)
```
Merge объединить наборы полей




### <a name="Fields.Scan">func</a> (\*Fields) [Scan](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/fields.go?s=677:722#L40)
``` go
func (o *Fields) Scan(pSrc interface{}) error
```
Scan implements sql.Scanner interface




### <a name="Fields.Size">func</a> (\*Fields) [Size](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/fields.go?s=1599:1626#L85)
``` go
func (o *Fields) Size() int
```
Size resturn size for protobuf




### <a name="Fields.Unmarshal">func</a> (\*Fields) [Unmarshal](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/fields.go?s=1482:1527#L80)
``` go
func (o *Fields) Unmarshal(data []byte) error
```
Unmarshal unmarshaller interface




### <a name="Fields.UnmarshalEasyJSON">func</a> (\*Fields) [UnmarshalEasyJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/fields_easyjson.go?s=1906:1957#L83)
``` go
func (v *Fields) UnmarshalEasyJSON(l *jlexer.Lexer)
```
UnmarshalEasyJSON supports easyjson.Unmarshaler interface




### <a name="Fields.UnmarshalJSON">func</a> (\*Fields) [UnmarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/fields_easyjson.go?s=1677:1726#L76)
``` go
func (v *Fields) UnmarshalJSON(data []byte) error
```
UnmarshalJSON supports json.Unmarshaler interface




### <a name="Fields.Value">func</a> (Fields) [Value](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/fields.go?s=1029:1084#L59)
``` go
func (o Fields) Value() (value driver.Value, err error)
```
Value implements sql.Valuer interface




## <a name="FloatViewData">type</a> [FloatViewData](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_float.go?s=379:524#L16)
``` go
type FloatViewData struct {
    AdditionalType   FieldFloatType `json:"additionalType"`
    ShowRowSeparator bool           `json:"showRowSeparator"`
}

```
FloatViewData описание представления поля типа FLOAT










## <a name="StringViewData">type</a> [StringViewData](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_string.go?s=465:551#L20)
``` go
type StringViewData struct {
    AdditionalType FieldStringType `json:"additionalType"`
}

```
StringViewData описание представления типа










## <a name="TableFieldData">type</a> [TableFieldData](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_table.go?s=2224:2370#L95)
``` go
type TableFieldData struct {
    Fields Fields                 `json:"fields" patch:"po"`
    Result map[string]TableResult `json:"result" patch:"po"`
}

```
TableFieldData описания поля типа TABLE










## <a name="TableResult">type</a> [TableResult](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_table.go?s=3291:3459#L127)
``` go
type TableResult struct {
    Kind    TableResultKind    `json:"kind"`
    Label   string             `json:"label" patch:"po"`
    Formula TableResultFormula `json:"formula"`
}

```
TableResult описание результата по колонке










## <a name="TableResultFormula">type</a> [TableResultFormula](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_table.go?s=2801:2831#L113)
``` go
type TableResultFormula string
```
TableResultFormula агрегатные функции


``` go
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
```









## <a name="TableResultKind">type</a> [TableResultKind](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_table.go?s=2428:2455#L101)
``` go
type TableResultKind string
```
TableResultKind тип агрегации итога


``` go
const (
    // ResultKindNone не показывать итог
    ResultKindNone TableResultKind = "none"
    // ResultKindLabel текстовая
    ResultKindLabel TableResultKind = "label"
    // ResultKindFormula итог по формуле
    ResultKindFormula TableResultKind = "formula"
)
```









## <a name="TableValue">type</a> [TableValue](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_table.go?s=1927:2069#L85)
``` go
type TableValue struct {
    Rows   []TableValueRow `json:"rows"`
    Result TableValueRow   `json:"result"`
    View   string          `json:"view"`
}

```
TableValue данные таблицы










## <a name="TableValueRow">type</a> [TableValueRow](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_table.go?s=2116:2163#L92)
``` go
type TableValueRow = map[string]json.RawMessage
```
TableValueRow строка таблицы










## <a name="TableViewData">type</a> [TableViewData](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_table.go?s=3537:4066#L134)
``` go
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

```
TableViewData описание представления таблицы










## <a name="Translator">type</a> [Translator](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/field_factory.go?s=247:341#L14)
``` go
type Translator interface {
    TranslateString(context.Context, string, ...interface{}) string
}
```
Translator сервис переводов










## <a name="Type">type</a> [Type](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/type.go?s=150:163#L14)
``` go
type Type int
```
Type is an enum of available types


``` go
const (

    // String is a plain string
    String Type

    // Float used for numbers by default
    Float

    // Integer used only for system collections
    Integer

    // Boolean is a boolean
    Boolean

    // DateTime is a RFC3339 timestamp
    DateTime

    // Duration is a number in seconds
    Duration

    // Category is a hierarchical tree linked with collection
    Category

    // Tag is an enum
    //
    // Tag can be static or enhancable (non-privileged user may add tags on item editing)
    Tag

    // Money is a pair of currency and cents value (int)
    Money

    // File is a hash code of file in storage
    File

    // Phone record
    Phone

    // Email record
    Email

    // Image is a file with preview cache?
    //
    // Warning: not implemented yet
    Image

    // Status is an static enum. It must be no more than one per collection.
    Status

    // Version semver compatible?
    //
    // Warning: not implemented yet
    Version

    // JSON object
    JSON

    // SysUser reference to user
    SysUser

    // FullName record
    FullName

    // Link record
    Link

    // SysOSNode reference to orgstruct node
    SysOSNode

    // SysCollection reference to collection (not element of collection)
    SysCollection

    // RefItem references to element of any collection
    RefItem

    // Enum type
    Enum

    // Table type
    Table

    // Account type
    Account
)
```






### <a name="FromString">func</a> [FromString](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/type.go?s=3142:3181#L160)
``` go
func FromString(s string) (Type, error)
```
FromString сформировать Type из строки





### <a name="Type.CanReplace">func</a> (Type) [CanReplace](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implements.go?s=3682:3744#L107)
``` go
func (t Type) CanReplace(field, candidate Field) (bool, error)
```
CanReplace проверить возможность обновления описания




### <a name="Type.CanReplaceData">func</a> (Type) [CanReplaceData](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implements.go?s=3343:3428#L98)
``` go
func (t Type) CanReplaceData(sourceData, candidateData json.RawMessage) (bool, error)
```
CanReplaceData проверить возможность обновления описания поля




### <a name="Type.CanReplaceViewData">func</a> (Type) [CanReplaceViewData](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implements.go?s=2983:3072#L89)
``` go
func (t Type) CanReplaceViewData(sourceData, candidateData json.RawMessage) (bool, error)
```
CanReplaceViewData проверить возможность обновления описания представления поля




### <a name="Type.Compare">func</a> (Type) [Compare](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implements.go?s=2636:2703#L80)
``` go
func (t Type) Compare(value1, value2 json.RawMessage) (bool, error)
```
Compare сравнить пару значений на равенство




### <a name="Type.Generate">func</a> (Type) [Generate](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/type.go?s=4178:4233#L211)
``` go
func (Type) Generate(r *rand.Rand, _ int) reflect.Value
```
Generate implements testing/quick.Generator

Deprecated: only for tests




### <a name="Type.GetImplement">func</a> (Type) [GetImplement](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implements.go?s=4105:4156#L122)
``` go
func (t Type) GetImplement() (TypeImplement, error)
```
GetImplement получить конкретную имплементацию




### <a name="Type.MarshalJSON">func</a> (Type) [MarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/type.go?s=3515:3558#L179)
``` go
func (t Type) MarshalJSON() ([]byte, error)
```
MarshalJSON implements json.Marshaler interface




### <a name="Type.String">func</a> (Type) [String](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/type.go?s=3340:3369#L170)
``` go
func (t Type) String() string
```
String implements sys.Stringer interface




### <a name="Type.UnmarshalData">func</a> (Type) [UnmarshalData](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implements.go?s=2038:2111#L62)
``` go
func (t Type) UnmarshalData(rawData json.RawMessage) (interface{}, error)
```
UnmarshalData извлечь описание поля




### <a name="Type.UnmarshalJSON">func</a> (\*Type) [UnmarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/type.go?s=3813:3860#L193)
``` go
func (t *Type) UnmarshalJSON(data []byte) error
```
UnmarshalJSON implements json.Unmarshaler interface




### <a name="Type.UnmarshalViewData">func</a> (Type) [UnmarshalViewData](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implements.go?s=2330:2411#L71)
``` go
func (t Type) UnmarshalViewData(rawViewData json.RawMessage) (interface{}, error)
```
UnmarshalViewData извлечь описание представления




### <a name="Type.Validate">func</a> (Type) [Validate](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implements.go?s=1778:1846#L53)
``` go
func (t Type) Validate(message json.RawMessage) (interface{}, error)
```
Validate провалидировать значение для поля




## <a name="TypeImplement">type</a> [TypeImplement](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implements.go?s=214:715#L13)
``` go
type TypeImplement interface {
    Validate(message json.RawMessage) (interface{}, error)
    UnmarshalData(data json.RawMessage) (interface{}, error)
    UnmarshalViewData(viewData json.RawMessage) (interface{}, error)
    Compare(value1 json.RawMessage, value2 json.RawMessage) (bool, error)
    CanReplace(oldField, newField Field) (bool, error)
    CanReplaceData(source json.RawMessage, candidate json.RawMessage) (bool, error)
    CanReplaceViewData(source json.RawMessage, candidate json.RawMessage) (bool, error)
}
```
TypeImplement имплементация поля определённого типа







### <a name="NewAccountType">func</a> [NewAccountType](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_account.go?s=151:186#L10)
``` go
func NewAccountType() TypeImplement
```
NewAccountType конструктор


### <a name="NewBooleanType">func</a> [NewBooleanType](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_boolean.go?s=177:212#L12)
``` go
func NewBooleanType() TypeImplement
```
NewBooleanType конструктор


### <a name="NewCategoryType">func</a> [NewCategoryType](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_category.go?s=147:183#L12)
``` go
func NewCategoryType() TypeImplement
```
NewCategoryType конструктор


### <a name="NewDateTimeType">func</a> [NewDateTimeType](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_dateTime.go?s=1521:1557#L47)
``` go
func NewDateTimeType() TypeImplement
```
NewDateTimeType конструктор


### <a name="NewDurationType">func</a> [NewDurationType](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_duration.go?s=81:117#L6)
``` go
func NewDurationType() TypeImplement
```
NewDurationType конструктор


### <a name="NewEmailType">func</a> [NewEmailType](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_email.go?s=173:206#L12)
``` go
func NewEmailType() TypeImplement
```
NewEmailType конструктор


### <a name="NewEnumType">func</a> [NewEnumType](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_enum.go?s=215:247#L14)
``` go
func NewEnumType() TypeImplement
```
NewEnumType конструктор


### <a name="NewFileType">func</a> [NewFileType](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_file.go?s=145:177#L10)
``` go
func NewFileType() TypeImplement
```
NewFileType конструктор


### <a name="NewFloatType">func</a> [NewFloatType](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_float.go?s=565:598#L22)
``` go
func NewFloatType() TypeImplement
```
NewFloatType конструктор


### <a name="NewFullNameType">func</a> [NewFullNameType](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_fullName.go?s=153:189#L10)
``` go
func NewFullNameType() TypeImplement
```
NewFullNameType конструктор


### <a name="NewIntegerType">func</a> [NewIntegerType](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_integer.go?s=80:115#L6)
``` go
func NewIntegerType() TypeImplement
```
NewIntegerType конструктор


### <a name="NewJSONType">func</a> [NewJSONType](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_json.go?s=82:114#L8)
``` go
func NewJSONType() TypeImplement
```
NewJSONType конструктор


### <a name="NewLinkType">func</a> [NewLinkType](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_link.go?s=145:177#L10)
``` go
func NewLinkType() TypeImplement
```
NewLinkType конструктор


### <a name="NewMoneyType">func</a> [NewMoneyType](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_money.go?s=173:206#L12)
``` go
func NewMoneyType() TypeImplement
```
NewMoneyType конструктор


### <a name="NewPhoneType">func</a> [NewPhoneType](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_phone.go?s=173:206#L12)
``` go
func NewPhoneType() TypeImplement
```
NewPhoneType конструктор


### <a name="NewRefItemType">func</a> [NewRefItemType](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_refItem.go?s=149:184#L10)
``` go
func NewRefItemType() TypeImplement
```
NewRefItemType конструтор


### <a name="NewStatusType">func</a> [NewStatusType](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_status.go?s=149:183#L10)
``` go
func NewStatusType() TypeImplement
```
NewStatusType конструктор


### <a name="NewStringType">func</a> [NewStringType](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_string.go?s=593:627#L25)
``` go
func NewStringType() TypeImplement
```
NewStringType конструктор


### <a name="NewSysCollectionType">func</a> [NewSysCollectionType](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_sysCollection.go?s=189:230#L12)
``` go
func NewSysCollectionType() TypeImplement
```
NewSysCollectionType конструктор


### <a name="NewSysOSNodeType">func</a> [NewSysOSNodeType](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_sysOSNode.go?s=152:189#L10)
``` go
func NewSysOSNodeType() TypeImplement
```
NewSysOSNodeType конструктор


### <a name="NewSysUserType">func</a> [NewSysUserType](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_sysUser.go?s=150:185#L10)
``` go
func NewSysUserType() TypeImplement
```
NewSysUserType конструктор


### <a name="NewTableType">func</a> [NewTableType](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_table.go?s=109:142#L10)
``` go
func NewTableType() TypeImplement
```
NewTableType конструктор


### <a name="NewTagType">func</a> [NewTagType](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_tag.go?s=76:107#L6)
``` go
func NewTagType() TypeImplement
```
NewTagType конструктор


### <a name="NewVersionType">func</a> [NewVersionType](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/implement_version.go?s=85:120#L8)
``` go
func NewVersionType() TypeImplement
```
NewVersionType конструктор










## <a name="pkg-subdirectories">Subdirectories</a>

| Name | Synopsis |
| ---- | -------- |
| [..](..) | |
| [complextypes](complextypes/) |  |
| [complextypes/account](complextypes/account/) |  |
| [complextypes/boolean](complextypes/boolean/) |  |
| [complextypes/category](complextypes/category/) |  |
| [complextypes/email](complextypes/email/) |  |
| [complextypes/enum](complextypes/enum/) |  |
| [complextypes/file](complextypes/file/) |  |
| [complextypes/fullname](complextypes/fullname/) |  |
| [complextypes/link](complextypes/link/) |  |
| [complextypes/money](complextypes/money/) |  |
| [complextypes/phone](complextypes/phone/) |  |
| [complextypes/ref](complextypes/ref/) |  |
| [complextypes/refitem](complextypes/refitem/) |  |
| [complextypes/status](complextypes/status/) |  |
| [complextypes/syscollection](complextypes/syscollection/) |  |
| [complextypes/sysref](complextypes/sysref/) |  |
| [test_data](test_data/) |  |
| [test_data/field_factory](test_data/field_factory/) |  |
| [test_data/golden](test_data/golden/) |  |
| [typestest](typestest/) |  |
| [uuids](uuids/) |  |


- - -
Generated by [godoc2md](https://github.com/Exa-Networks/godoc2md)
