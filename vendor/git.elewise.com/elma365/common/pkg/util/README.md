# util
`import "git.elewise.com/elma365/common/pkg/util"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [func CategoryCode(collectionCode, fieldCode string) string](#CategoryCode)
* [func ExtractTagsFromStruct(tagName string, i interface{}) []string](#ExtractTagsFromStruct)
* [func ResponseHTTP(w http.ResponseWriter, data interface{}, opts ...ResponseHTTPOption)](#ResponseHTTP)
* [func StructToMapByTags(tagName string, i interface{}, skipNils bool) map[string]interface{}](#StructToMapByTags)
* [func TruncatePGIdentifier(s string) string](#TruncatePGIdentifier)
* [type BodyBuilder](#BodyBuilder)
  * [func (bb *BodyBuilder) Add(code, from string, codes ...string) *BodyBuilder](#BodyBuilder.Add)
  * [func (bb *BodyBuilder) AddRaw(code, query string) *BodyBuilder](#BodyBuilder.AddRaw)
  * [func (bb *BodyBuilder) AddSubquery(code, aggregate string, stmt squirrel.Sqlizer) *BodyBuilder](#BodyBuilder.AddSubquery)
  * [func (bb *BodyBuilder) Build() string](#BodyBuilder.Build)
  * [func (bb *BodyBuilder) BuildAs(aliase string) string](#BodyBuilder.BuildAs)
  * [func (bb *BodyBuilder) BuildExpr() squirrel.Sqlizer](#BodyBuilder.BuildExpr)
  * [func (bb *BodyBuilder) RemapJSONQuery(from string, prependFields, appendFields []string) string](#BodyBuilder.RemapJSONQuery)
* [type BodyJSON](#BodyJSON)
  * [func Body(codes ...string) BodyJSON](#Body)
  * [func Column(column string, codes ...string) BodyJSON](#Column)
  * [func RawBody(codes ...string) BodyJSON](#RawBody)
  * [func RawColumn(column string, codes ...string) BodyJSON](#RawColumn)
  * [func (bj BodyJSON) As(alias string) string](#BodyJSON.As)
  * [func (bj BodyJSON) Cast(t string) BodyJSON](#BodyJSON.Cast)
  * [func (bj BodyJSON) CastTS(ctx context.Context) BodyJSON](#BodyJSON.CastTS)
  * [func (bj BodyJSON) Eq(to interface{}) squirrel.Sqlizer](#BodyJSON.Eq)
  * [func (bj BodyJSON) Gt(to interface{}) squirrel.Sqlizer](#BodyJSON.Gt)
  * [func (bj BodyJSON) Gte(to interface{}) squirrel.Sqlizer](#BodyJSON.Gte)
  * [func (bj BodyJSON) IsNotNull() squirrel.Sqlizer](#BodyJSON.IsNotNull)
  * [func (bj BodyJSON) IsNull() squirrel.Sqlizer](#BodyJSON.IsNull)
  * [func (bj BodyJSON) Lt(to interface{}) squirrel.Sqlizer](#BodyJSON.Lt)
  * [func (bj BodyJSON) Lte(to interface{}) squirrel.Sqlizer](#BodyJSON.Lte)
  * [func (bj BodyJSON) Neq(to interface{}) squirrel.Sqlizer](#BodyJSON.Neq)
  * [func (bj BodyJSON) String() string](#BodyJSON.String)
  * [func (bj BodyJSON) Table(t string) BodyJSON](#BodyJSON.Table)
  * [func (bj BodyJSON) ToSql() (query string, args []interface{}, err error)](#BodyJSON.ToSql)
* [type Caller](#Caller)
  * [func GetCaller(exclude string) Caller](#GetCaller)
* [type CollectionName](#CollectionName)
  * [func (cn CollectionName) String() string](#CollectionName.String)
  * [func (cn CollectionName) WithCompany(ctx context.Context) string](#CollectionName.WithCompany)
* [type DebugTransport](#DebugTransport)
  * [func (t DebugTransport) RoundTrip(req *http.Request) (*http.Response, error)](#DebugTransport.RoundTrip)
* [type ExtractOption](#ExtractOption)
  * [func WithoutEmbedded() ExtractOption](#WithoutEmbedded)
  * [func WithoutEmpty() ExtractOption](#WithoutEmpty)
  * [func WithoutMinus() ExtractOption](#WithoutMinus)
* [type IndexName](#IndexName)
  * [func (in IndexName) WithCompany(ctx context.Context) string](#IndexName.WithCompany)
* [type Reflector](#Reflector)
  * [func FromValue(v reflect.Value) Reflector](#FromValue)
  * [func NewReflector(i interface{}) Reflector](#NewReflector)
  * [func (r Reflector) Apply(m map[string]string) error](#Reflector.Apply)
  * [func (r Reflector) ExtractTags(tagName string, opts ...ExtractOption) map[string]string](#Reflector.ExtractTags)
  * [func (r Reflector) ExtractValues(tagName string, skipNils bool, opts ...ExtractOption) map[string]interface{}](#Reflector.ExtractValues)
  * [func (r Reflector) Value() interface{}](#Reflector.Value)
* [type ResponseHTTPOption](#ResponseHTTPOption)
  * [func ResponseWithError(ctx context.Context, err error) ResponseHTTPOption](#ResponseWithError)
  * [func ResponseWithStatus(status int) ResponseHTTPOption](#ResponseWithStatus)
* [type TableName](#TableName)
  * [func NewTableName(ns namespace.Namespace, code string) TableName](#NewTableName)
  * [func (tn TableName) Alias(ctx context.Context, alias string) string](#TableName.Alias)
  * [func (tn TableName) Format(s fmt.State, verb rune)](#TableName.Format)
  * [func (tn TableName) PGIdentifier() string](#TableName.PGIdentifier)
  * [func (tn TableName) String() string](#TableName.String)
  * [func (tn TableName) WithCompany(ctx context.Context) string](#TableName.WithCompany)

#### <a name="pkg-examples">Examples</a>
* [BodyBuilder](#example-bodybuilder)

#### <a name="pkg-files">Package files</a>
[body.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/body.go) [collection.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/collection.go) [debug_transport.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/debug_transport.go) [get_caller.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/get_caller.go) [reflect.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/reflect.go) [reflector.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/reflector.go) [response_http.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/response_http.go) [table.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/table.go)





## <a name="CategoryCode">func</a> [CategoryCode](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/table.go?s=2330:2388#L71)
``` go
func CategoryCode(collectionCode, fieldCode string) string
```
CategoryCode возвращает код коллекции для категории

Если коллекция уже является коллекцией категорий, то вернётся код коллекции категорий родительской коллекции.



## <a name="ExtractTagsFromStruct">func</a> [ExtractTagsFromStruct](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/reflect.go?s=227:293#L6)
``` go
func ExtractTagsFromStruct(tagName string, i interface{}) []string
```
ExtractTagsFromStruct извлекает список тэгов по имени из структуры

Подробнее смотри <a href="https://godoc.org/github.com/jmoiron/sqlx/reflectx#Mapper.FieldMap">https://godoc.org/github.com/jmoiron/sqlx/reflectx#Mapper.FieldMap</a>



## <a name="ResponseHTTP">func</a> [ResponseHTTP](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/response_http.go?s=1401:1487#L63)
``` go
func ResponseHTTP(w http.ResponseWriter, data interface{}, opts ...ResponseHTTPOption)
```
ResponseHTTP ответить json на http-запрос



## <a name="StructToMapByTags">func</a> [StructToMapByTags](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/reflect.go?s=708:799#L20)
``` go
func StructToMapByTags(tagName string, i interface{}, skipNils bool) map[string]interface{}
```
StructToMapByTags преобразует структуру в хэш-таблицу по значениям тэга

Подробнее смотри <a href="https://godoc.org/github.com/jmoiron/sqlx/reflectx#Mapper.FieldMap">https://godoc.org/github.com/jmoiron/sqlx/reflectx#Mapper.FieldMap</a>



## <a name="TruncatePGIdentifier">func</a> [TruncatePGIdentifier](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/table.go?s=3003:3045#L93)
``` go
func TruncatePGIdentifier(s string) string
```
TruncatePGIdentifier обрезает строку до 63 символов, добавляя CRC64 в конце, если строка длинее




## <a name="BodyBuilder">type</a> [BodyBuilder](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/body.go?s=4385:4411#L130)
``` go
type BodyBuilder []bbField
```
BodyBuilder конструктор объекта body для сложных запросов



##### Example BodyBuilder:
``` go
bb := util.BodyBuilder{}
stmt := squirrel.Select(bb.
    Add("namespace", "foo", "namespace").
    Add("code", "bar", "code").
    Build()).
    From("foo").
    Join("bar ON foo.id = bar.id")
fmt.Println(squirrel.DebugSqlizer(stmt))
// Output: SELECT jsonb_build_object('namespace',(foo.body#>'{namespace}'),'code',(bar.body#>'{code}')) FROM foo JOIN bar ON foo.id = bar.id
```

Output:

```
SELECT jsonb_build_object('namespace',(foo.body#>'{namespace}'),'code',(bar.body#>'{code}')) FROM foo JOIN bar ON foo.id = bar.id
```








### <a name="BodyBuilder.Add">func</a> (\*BodyBuilder) [Add](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/body.go?s=4561:4636#L133)
``` go
func (bb *BodyBuilder) Add(code, from string, codes ...string) *BodyBuilder
```
Add добавляет к строящемуся объекту поле с указанным кодом из указанной таблицы




### <a name="BodyBuilder.AddRaw">func</a> (\*BodyBuilder) [AddRaw](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/body.go?s=4821:4883#L142)
``` go
func (bb *BodyBuilder) AddRaw(code, query string) *BodyBuilder
```
AddRaw - добавляем query к текущему запросу




### <a name="BodyBuilder.AddSubquery">func</a> (\*BodyBuilder) [AddSubquery](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/body.go?s=5332:5426#L152)
``` go
func (bb *BodyBuilder) AddSubquery(code, aggregate string, stmt squirrel.Sqlizer) *BodyBuilder
```
AddSubquery добавляет к строящемуся объекту поле с указанным кодом как подзапрос

**Внимание!** Подзапрос должен быть SELECT без аргументов.
По сути он вычисляется с откидыванием аргументов, оставляя лишь query.




### <a name="BodyBuilder.Build">func</a> (\*BodyBuilder) [Build](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/body.go?s=6086:6123#L173)
``` go
func (bb *BodyBuilder) Build() string
```
Build построить запрос




### <a name="BodyBuilder.BuildAs">func</a> (\*BodyBuilder) [BuildAs](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/body.go?s=5935:5987#L168)
``` go
func (bb *BodyBuilder) BuildAs(aliase string) string
```
BuildAs построить запрос с aliase




### <a name="BodyBuilder.BuildExpr">func</a> (\*BodyBuilder) [BuildExpr](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/body.go?s=5770:5821#L163)
``` go
func (bb *BodyBuilder) BuildExpr() squirrel.Sqlizer
```
BuildExpr - возврат запроса, как выражение для squirrel
используется для подстановки выражений как подзапросов




### <a name="BodyBuilder.RemapJSONQuery">func</a> (\*BodyBuilder) [RemapJSONQuery](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/body.go?s=6902:6997#L190)
``` go
func (bb *BodyBuilder) RemapJSONQuery(from string, prependFields, appendFields []string) string
```
RemapJSONQuery - построение запроса, для селекта из таблицы с определенным мэпингом json
используется для запросов insert from select, update from select
Пример:
мэпим namespace, __id приложения к namespace, id другого приложения

bb := util.BodyBuilder{}
bb.Add("namespace", "some_namespace:some_app", "namespace").
Add("id", "some_namespace:some_app", "__id")
query := bb.RemapJSONQuery("someTable", []string{}, []string{})




## <a name="BodyJSON">type</a> [BodyJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/body.go?s=287:307#L15)
``` go
type BodyJSON string
```
BodyJSON шорткат для работы с подполями элемента коллекции







### <a name="Body">func</a> [Body](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/body.go?s=873:908#L28)
``` go
func Body(codes ...string) BodyJSON
```
Body шорткат для обращения к подполям элемента коллекции


### <a name="Column">func</a> [Column](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/body.go?s=430:482#L18)
``` go
func Column(column string, codes ...string) BodyJSON
```
Column - шорткат обращения к подполяем необходимой колонки таблицы


### <a name="RawBody">func</a> [RawBody](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/body.go?s=1109:1147#L33)
``` go
func RawBody(codes ...string) BodyJSON
```
RawBody шорткат для обращения к подполям элемента коллекции в виде jsonb


### <a name="RawColumn">func</a> [RawColumn](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/body.go?s=661:716#L23)
``` go
func RawColumn(column string, codes ...string) BodyJSON
```
RawColumn - шорткат для обращения к полям необходимой колонки в виде jsonb





### <a name="BodyJSON.As">func</a> (BodyJSON) [As](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/body.go?s=2543:2585#L81)
``` go
func (bj BodyJSON) As(alias string) string
```
As формирует выражения для алиаса к полю




### <a name="BodyJSON.Cast">func</a> (BodyJSON) [Cast](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/body.go?s=1616:1658#L53)
``` go
func (bj BodyJSON) Cast(t string) BodyJSON
```
Cast привести значение поля к определённому типу




### <a name="BodyJSON.CastTS">func</a> (BodyJSON) [CastTS](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/body.go?s=2026:2081#L63)
``` go
func (bj BodyJSON) CastTS(ctx context.Context) BodyJSON
```
CastTS привести значение поля к формату временной метки и добавить компанию из контекста




### <a name="BodyJSON.Eq">func</a> (BodyJSON) [Eq](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/body.go?s=2736:2790#L86)
``` go
func (bj BodyJSON) Eq(to interface{}) squirrel.Sqlizer
```
Eq возвращает выражение для выражения WHERE на равенство




### <a name="BodyJSON.Gt">func</a> (BodyJSON) [Gt](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/body.go?s=3139:3193#L96)
``` go
func (bj BodyJSON) Gt(to interface{}) squirrel.Sqlizer
```
Gt возвращает выражение для выражения WHERE на равенство




### <a name="BodyJSON.Gte">func</a> (BodyJSON) [Gte](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/body.go?s=3339:3394#L101)
``` go
func (bj BodyJSON) Gte(to interface{}) squirrel.Sqlizer
```
Gte возвращает выражение для выражения WHERE на равенство




### <a name="BodyJSON.IsNotNull">func</a> (BodyJSON) [IsNotNull](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/body.go?s=4148:4195#L121)
``` go
func (bj BodyJSON) IsNotNull() squirrel.Sqlizer
```
IsNotNull возвращает выражение для выражения WHERE на равенство




### <a name="BodyJSON.IsNull">func</a> (BodyJSON) [IsNull](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/body.go?s=3951:3995#L116)
``` go
func (bj BodyJSON) IsNull() squirrel.Sqlizer
```
IsNull возвращает выражение для выражения WHERE на равенство




### <a name="BodyJSON.Lt">func</a> (BodyJSON) [Lt](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/body.go?s=3543:3597#L106)
``` go
func (bj BodyJSON) Lt(to interface{}) squirrel.Sqlizer
```
Lt возвращает выражение для выражения WHERE на равенство




### <a name="BodyJSON.Lte">func</a> (BodyJSON) [Lte](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/body.go?s=3743:3798#L111)
``` go
func (bj BodyJSON) Lte(to interface{}) squirrel.Sqlizer
```
Lte возвращает выражение для выражения WHERE на равенство




### <a name="BodyJSON.Neq">func</a> (BodyJSON) [Neq](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/body.go?s=2936:2991#L91)
``` go
func (bj BodyJSON) Neq(to interface{}) squirrel.Sqlizer
```
Neq возвращает выражение для выражения WHERE на равенство




### <a name="BodyJSON.String">func</a> (BodyJSON) [String](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/body.go?s=2224:2258#L69)
``` go
func (bj BodyJSON) String() string
```
String implements fmt.Stringer




### <a name="BodyJSON.Table">func</a> (BodyJSON) [Table](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/body.go?s=1761:1804#L58)
``` go
func (bj BodyJSON) Table(t string) BodyJSON
```
Table добавляет имя таблицы




### <a name="BodyJSON.ToSql">func</a> (BodyJSON) [ToSql](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/body.go?s=2360:2432#L76)
``` go
func (bj BodyJSON) ToSql() (query string, args []interface{}, err error)
```
ToSql implements squirrel.Sqlizer




## <a name="Caller">type</a> [Caller](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/get_caller.go?s=183:228#L12)
``` go
type Caller struct {
    File, Function string
}

```
Caller описание вызывающего

Название файла и функции или тип.метод







### <a name="GetCaller">func</a> [GetCaller](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/get_caller.go?s=763:800#L20)
``` go
func GetCaller(exclude string) Caller
```
GetCaller возвращает первого вызывающего за пределами указанного файла

По сути берётся стектрейс вызова и в нём берётся первая запись, которая не содержит указанный файл (файл указывается
без расширения .go), из этой записи возвращается имя функции (метода). Более подробно смотри в тестах.





## <a name="CollectionName">type</a> [CollectionName](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/collection.go?s=183:209#L11)
``` go
type CollectionName string
```
CollectionName определяет тип для название коллекции mongoDB










### <a name="CollectionName.String">func</a> (CollectionName) [String](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/collection.go?s=516:556#L21)
``` go
func (cn CollectionName) String() string
```
String implements fmt.Stringer




### <a name="CollectionName.WithCompany">func</a> (CollectionName) [WithCompany](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/collection.go?s=330:394#L14)
``` go
func (cn CollectionName) WithCompany(ctx context.Context) string
```
WithCompany возвращает название коллекции mongoDB с префиксом компании




## <a name="DebugTransport">type</a> [DebugTransport](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/debug_transport.go?s=190:272#L12)
``` go
type DebugTransport struct {
    Transport http.RoundTripper
    Logger    *zap.Logger
}

```
DebugTransport http transport с логированием запроса и ответа










### <a name="DebugTransport.RoundTrip">func</a> (DebugTransport) [RoundTrip](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/debug_transport.go?s=324:400#L18)
``` go
func (t DebugTransport) RoundTrip(req *http.Request) (*http.Response, error)
```
RoundTrip Реализует http.RoundTripper




## <a name="ExtractOption">type</a> [ExtractOption](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/reflector.go?s=1146:1214#L62)
``` go
type ExtractOption interface {
    Apply(extractConfig) extractConfig
}
```
ExtractOption опция обработчика







### <a name="WithoutEmbedded">func</a> [WithoutEmbedded](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/reflector.go?s=1444:1480#L73)
``` go
func WithoutEmbedded() ExtractOption
```
WithoutEmbedded игнорировать встроенные структуры


### <a name="WithoutEmpty">func</a> [WithoutEmpty](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/reflector.go?s=1662:1695#L82)
``` go
func WithoutEmpty() ExtractOption
```
WithoutEmpty игнорировать поля без тэга


### <a name="WithoutMinus">func</a> [WithoutMinus](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/reflector.go?s=1889:1922#L91)
``` go
func WithoutMinus() ExtractOption
```
WithoutMinus игнорировать поля с тэгом равным "-"





## <a name="IndexName">type</a> [IndexName](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/table.go?s=2612:2633#L81)
``` go
type IndexName string
```
IndexName is a name of elastic index










### <a name="IndexName.WithCompany">func</a> (IndexName) [WithCompany](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/table.go?s=2692:2751#L86)
``` go
func (in IndexName) WithCompany(ctx context.Context) string
```
WithCompany prefix




## <a name="Reflector">type</a> [Reflector](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/reflector.go?s=207:273#L16)
``` go
type Reflector struct {
    // contains filtered or unexported fields
}

```
Reflector рефлексия вокруг структурных тэгов







### <a name="FromValue">func</a> [FromValue](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/reflector.go?s=643:684#L36)
``` go
func FromValue(v reflect.Value) Reflector
```
FromValue конструирует рефлектор по рефлексии объекта


### <a name="NewReflector">func</a> [NewReflector](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/reflector.go?s=318:360#L22)
``` go
func NewReflector(i interface{}) Reflector
```
NewReflector — конструктор





### <a name="Reflector.Apply">func</a> (Reflector) [Apply](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/reflector.go?s=4214:4265#L164)
``` go
func (r Reflector) Apply(m map[string]string) error
```
Apply применяет данные из переданной хеш-таблицы на структуру, вокруг которой построен рефлектор




### <a name="Reflector.ExtractTags">func</a> (Reflector) [ExtractTags](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/reflector.go?s=2777:2864#L116)
``` go
func (r Reflector) ExtractTags(tagName string, opts ...ExtractOption) map[string]string
```
ExtractTags возвращает хеш-таблицу, где имени поля сопоставлено значение тэга




### <a name="Reflector.ExtractValues">func</a> (Reflector) [ExtractValues](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/reflector.go?s=2180:2289#L100)
``` go
func (r Reflector) ExtractValues(tagName string, skipNils bool, opts ...ExtractOption) map[string]interface{}
```
ExtractValues возвращает хеш-таблицу, где значению тэга сопоставлено значение поля




### <a name="Reflector.Value">func</a> (Reflector) [Value](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/reflector.go?s=907:945#L50)
``` go
func (r Reflector) Value() interface{}
```
Value возвращает объект рефлексии




## <a name="ResponseHTTPOption">type</a> [ResponseHTTPOption](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/response_http.go?s=321:396#L21)
``` go
type ResponseHTTPOption interface {
    Apply(responseConfig) responseConfig
}
```
ResponseHTTPOption опция ответа







### <a name="ResponseWithError">func</a> [ResponseWithError](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/response_http.go?s=929:1002#L41)
``` go
func ResponseWithError(ctx context.Context, err error) ResponseHTTPOption
```
ResponseWithError вернуть ошибку, если она не пустая


### <a name="ResponseWithStatus">func</a> [ResponseWithStatus](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/response_http.go?s=672:726#L32)
``` go
func ResponseWithStatus(status int) ResponseHTTPOption
```
ResponseWithStatus ответить с определённым статусом (по умолчанию 200)





## <a name="TableName">type</a> [TableName](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/table.go?s=330:351#L18)
``` go
type TableName string
```
TableName имя таблицы в postgres с возможностью префикса компанией







### <a name="NewTableName">func</a> [NewTableName](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/table.go?s=515:579#L23)
``` go
func NewTableName(ns namespace.Namespace, code string) TableName
```
NewTableName — возвращает имя таблицы в postgres для заданной пары namespace и code





### <a name="TableName.Alias">func</a> (TableName) [Alias](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/table.go?s=1211:1278#L42)
``` go
func (tn TableName) Alias(ctx context.Context, alias string) string
```
Alias короткая запись для использования в FROM tn AS alias




### <a name="TableName.Format">func</a> (TableName) [Format](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/table.go?s=1811:1861#L57)
``` go
func (tn TableName) Format(s fmt.State, verb rune)
```
Format форматирует название таблицы в зависимости от ключа




### <a name="TableName.PGIdentifier">func</a> (TableName) [PGIdentifier](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/table.go?s=1616:1657#L52)
``` go
func (tn TableName) PGIdentifier() string
```
PGIdentifier вовзращает название таблицы обрезанное с помощью TruncatePGIdentifier




### <a name="TableName.String">func</a> (TableName) [String](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/table.go?s=1427:1462#L47)
``` go
func (tn TableName) String() string
```
String возвращает название таблицы как строку




### <a name="TableName.WithCompany">func</a> (TableName) [WithCompany](https://git.elewise.com/elma365/common/-/tree/develop/pkg/util/table.go?s=912:971#L35)
``` go
func (tn TableName) WithCompany(ctx context.Context) string
```
WithCompany добавляет префикс компании







- - -
Generated by [godoc2md](https://github.com/Exa-Networks/godoc2md)
