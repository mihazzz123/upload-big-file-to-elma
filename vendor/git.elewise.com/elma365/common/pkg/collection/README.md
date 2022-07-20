# collection
`import "git.elewise.com/elma365/common/pkg/collection"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [func WithAlias(alias Alias, column string) string](#WithAlias)
* [func WithMainAlias(column string) string](#WithMainAlias)
* [type AccessType](#AccessType)
  * [func AccessTypeString(s string) (AccessType, error)](#AccessTypeString)
  * [func AccessTypeValues() []AccessType](#AccessTypeValues)
  * [func (i AccessType) IsAAccessType() bool](#AccessType.IsAAccessType)
  * [func (i AccessType) MarshalJSON() ([]byte, error)](#AccessType.MarshalJSON)
  * [func (i *AccessType) Scan(value interface{}) error](#AccessType.Scan)
  * [func (i AccessType) String() string](#AccessType.String)
  * [func (i *AccessType) UnmarshalJSON(data []byte) error](#AccessType.UnmarshalJSON)
  * [func (i AccessType) Value() (driver.Value, error)](#AccessType.Value)
* [type Alias](#Alias)
  * [func AliasString(s string) (Alias, error)](#AliasString)
  * [func AliasValues() []Alias](#AliasValues)
  * [func (i Alias) IsAAlias() bool](#Alias.IsAAlias)
  * [func (i Alias) MarshalJSON() ([]byte, error)](#Alias.MarshalJSON)
  * [func (i *Alias) Scan(value interface{}) error](#Alias.Scan)
  * [func (i Alias) String() string](#Alias.String)
  * [func (i *Alias) UnmarshalJSON(data []byte) error](#Alias.UnmarshalJSON)
  * [func (i Alias) Value() (driver.Value, error)](#Alias.Value)
* [type BodyItem](#BodyItem)
  * [func NewBodyItem(itemJSON json.RawMessage) (*BodyItem, error)](#NewBodyItem)
  * [func (item BodyItem) GetID() uuid.UUID](#BodyItem.GetID)
  * [func (item BodyItem) MarshalJSON() ([]byte, error)](#BodyItem.MarshalJSON)
  * [func (item *BodyItem) Scan(pSrc interface{}) error](#BodyItem.Scan)
  * [func (item *BodyItem) UnmarshalJSON(data []byte) error](#BodyItem.UnmarshalJSON)
* [type Collection](#Collection)
  * [func (c *Collection) GetCode() string](#Collection.GetCode)
  * [func (c *Collection) GetFields() types.Fields](#Collection.GetFields)
  * [func (c *Collection) GetNamespace() namespace.Namespace](#Collection.GetNamespace)
  * [func (v Collection) MarshalEasyJSON(w *jwriter.Writer)](#Collection.MarshalEasyJSON)
  * [func (v Collection) MarshalJSON() ([]byte, error)](#Collection.MarshalJSON)
  * [func (v *Collection) UnmarshalEasyJSON(l *jlexer.Lexer)](#Collection.UnmarshalEasyJSON)
  * [func (v *Collection) UnmarshalJSON(data []byte) error](#Collection.UnmarshalJSON)
* [type Item](#Item)
  * [func (item Item) GetID() uuid.UUID](#Item.GetID)
  * [func (v Item) MarshalEasyJSON(w *jwriter.Writer)](#Item.MarshalEasyJSON)
  * [func (v Item) MarshalJSON() ([]byte, error)](#Item.MarshalJSON)
  * [func (item *Item) Scan(pSrc interface{}) error](#Item.Scan)
  * [func (v *Item) UnmarshalEasyJSON(l *jlexer.Lexer)](#Item.UnmarshalEasyJSON)
  * [func (v *Item) UnmarshalJSON(data []byte) error](#Item.UnmarshalJSON)
* [type ItemIdentifier](#ItemIdentifier)
* [type Metadata](#Metadata)
* [type SimpleMetadata](#SimpleMetadata)
  * [func (md SimpleMetadata) GetCode() string](#SimpleMetadata.GetCode)
  * [func (md SimpleMetadata) GetFields() types.Fields](#SimpleMetadata.GetFields)
  * [func (md SimpleMetadata) GetNamespace() namespace.Namespace](#SimpleMetadata.GetNamespace)
* [type Type](#Type)
  * [func TypeString(s string) (Type, error)](#TypeString)
  * [func TypeValues() []Type](#TypeValues)
  * [func (i Type) IsAType() bool](#Type.IsAType)
  * [func (i Type) MarshalJSON() ([]byte, error)](#Type.MarshalJSON)
  * [func (i *Type) Scan(value interface{}) error](#Type.Scan)
  * [func (i Type) String() string](#Type.String)
  * [func (i *Type) UnmarshalJSON(data []byte) error](#Type.UnmarshalJSON)
  * [func (i Type) Value() (driver.Value, error)](#Type.Value)


#### <a name="pkg-files">Package files</a>
[access_type.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/access_type.go) [access_type_string.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/access_type_string.go) [alias.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/alias.go) [alias_string.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/alias_string.go) [collection.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/collection.go) [collection_easyjson.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/collection_easyjson.go) [const.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/const.go) [item.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/item.go) [item_easyjson.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/item_easyjson.go) [item_type.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/item_type.go) [item_type_string.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/item_type_string.go)


## <a name="pkg-constants">Constants</a>
``` go
const (
    // DiskCode системная коллекция файлов
    DiskCode string = "disk_files"
    // GroupsCode системная коллекция групп
    GroupsCode string = "groups"
    // TasksCode системная коллекция задач
    TasksCode string = "tasks"
    // UsersCode системная коллекция пользователей
    UsersCode string = "users"

    // ProcessInstanceCodePrefix префикс кода коллеуции экземпляра процесса
    ProcessInstanceCodePrefix = "_process_"

    // CollectionBodyCol — название колонки, в которой хранится тело элемента
    CollectionBodyCol = "body"
    // CollectionPermissionsCol — название колонки, в которой хранятся разрешения объекта
    CollectionPermissionsCol = "permissions"
    // CollectionReadCol — название колонки, в которой хранится развёрнутый список групп, имеющих право на чтение
    CollectionReadCol = "read"
    // CollectionInheritCol — название колонки, в которой хранится признак наследования разрешений объекта
    CollectionInheritCol = "inherit"

    // CollectionType - поле типа коллекции
    CollectionType = "type"
)
```



## <a name="WithAlias">func</a> [WithAlias](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/alias.go?s=469:518#L18)
``` go
func WithAlias(alias Alias, column string) string
```
WithAlias поле с алиасом



## <a name="WithMainAlias">func</a> [WithMainAlias](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/alias.go?s=631:671#L23)
``` go
func WithMainAlias(column string) string
```
WithMainAlias поле с дефолтным алиасом




## <a name="AccessType">type</a> [AccessType](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/access_type.go?s=256:276#L6)
``` go
type AccessType int8
```
AccessType типы ограничения доступа к коллекции


``` go
const (
    // AccessTypeNone не ограничивать доступ к коллекции
    AccessTypeNone AccessType = iota
    // AccessTypeCollection права назначаются на всю коллекцию целиком
    AccessTypeCollection
    // AccessTypeRow права назначаются на каждую строчку коллекции
    AccessTypeRow
    // AccessTypeDirectory права назначаются папке коллекции
    AccessTypeDirectory
)
```






### <a name="AccessTypeString">func</a> [AccessTypeString](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/access_type_string.go?s=914:965#L34)
``` go
func AccessTypeString(s string) (AccessType, error)
```
AccessTypeString retrieves an enum value from the enum constants string name.
Throws an error if the param is not part of the enum.


### <a name="AccessTypeValues">func</a> [AccessTypeValues](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/access_type_string.go?s=1161:1197#L42)
``` go
func AccessTypeValues() []AccessType
```
AccessTypeValues returns all values of the enum





### <a name="AccessType.IsAAccessType">func</a> (AccessType) [IsAAccessType](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/access_type_string.go?s=1326:1366#L47)
``` go
func (i AccessType) IsAAccessType() bool
```
IsAAccessType returns "true" if the value is listed in the enum definition. "false" otherwise




### <a name="AccessType.MarshalJSON">func</a> (AccessType) [MarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/access_type_string.go?s=1531:1580#L57)
``` go
func (i AccessType) MarshalJSON() ([]byte, error)
```
MarshalJSON implements the json.Marshaler interface for AccessType




### <a name="AccessType.Scan">func</a> (\*AccessType) [Scan](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/access_type_string.go?s=2024:2074#L77)
``` go
func (i *AccessType) Scan(value interface{}) error
```



### <a name="AccessType.String">func</a> (AccessType) [String](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/access_type_string.go?s=347:382#L16)
``` go
func (i AccessType) String() string
```



### <a name="AccessType.UnmarshalJSON">func</a> (\*AccessType) [UnmarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/access_type_string.go?s=1693:1746#L62)
``` go
func (i *AccessType) UnmarshalJSON(data []byte) error
```
UnmarshalJSON implements the json.Unmarshaler interface for AccessType




### <a name="AccessType.Value">func</a> (AccessType) [Value](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/access_type_string.go?s=1945:1994#L73)
``` go
func (i AccessType) Value() (driver.Value, error)
```



## <a name="Alias">type</a> [Alias](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/alias.go?s=243:258#L8)
``` go
type Alias int8
```
Alias типы ограничения доступа к коллекции


``` go
const (
    // AliasMain Алиас основной таблицы
    AliasMain Alias = iota
    // AliasDir Алиас для таблицы директорий
    AliasDir
)
```






### <a name="AliasString">func</a> [AliasString](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/alias_string.go?s=718:759#L32)
``` go
func AliasString(s string) (Alias, error)
```
AliasString retrieves an enum value from the enum constants string name.
Throws an error if the param is not part of the enum.


### <a name="AliasValues">func</a> [AliasValues](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/alias_string.go?s=940:966#L40)
``` go
func AliasValues() []Alias
```
AliasValues returns all values of the enum





### <a name="Alias.IsAAlias">func</a> (Alias) [IsAAlias](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/alias_string.go?s=1085:1115#L45)
``` go
func (i Alias) IsAAlias() bool
```
IsAAlias returns "true" if the value is listed in the enum definition. "false" otherwise




### <a name="Alias.MarshalJSON">func</a> (Alias) [MarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/alias_string.go?s=1270:1314#L55)
``` go
func (i Alias) MarshalJSON() ([]byte, error)
```
MarshalJSON implements the json.Marshaler interface for Alias




### <a name="Alias.Scan">func</a> (\*Alias) [Scan](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/alias_string.go?s=1733:1778#L75)
``` go
func (i *Alias) Scan(value interface{}) error
```



### <a name="Alias.String">func</a> (Alias) [String](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/alias_string.go?s=287:317#L16)
``` go
func (i Alias) String() string
```



### <a name="Alias.UnmarshalJSON">func</a> (\*Alias) [UnmarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/alias_string.go?s=1422:1470#L60)
``` go
func (i *Alias) UnmarshalJSON(data []byte) error
```
UnmarshalJSON implements the json.Unmarshaler interface for Alias




### <a name="Alias.Value">func</a> (Alias) [Value](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/alias_string.go?s=1659:1703#L71)
``` go
func (i Alias) Value() (driver.Value, error)
```



## <a name="BodyItem">type</a> [BodyItem](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/item.go?s=1757:1833#L57)
``` go
type BodyItem struct {
    Item
    Body json.RawMessage `db:"body" json:"body"`
}

```
BodyItem описывает основные системные поля элемента коллекции, и хранит остальные его поля в "сыром" виде.
При сериализации в json, возвращает только "сырые значения".
Реализует интерфейс collection.ItemIdentifier.







### <a name="NewBodyItem">func</a> [NewBodyItem](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/item.go?s=1893:1954#L63)
``` go
func NewBodyItem(itemJSON json.RawMessage) (*BodyItem, error)
```
NewBodyItem создает экземпляр BodyItem





### <a name="BodyItem.GetID">func</a> (BodyItem) [GetID](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/item.go?s=2312:2350#L80)
``` go
func (item BodyItem) GetID() uuid.UUID
```
GetID реализует интерфейс collection.ItemIdentifier




### <a name="BodyItem.MarshalJSON">func</a> (BodyItem) [MarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/item.go?s=2440:2490#L85)
``` go
func (item BodyItem) MarshalJSON() ([]byte, error)
```
MarshalJSON реализует интерфейс json.Marshaler




### <a name="BodyItem.Scan">func</a> (\*BodyItem) [Scan](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/item.go?s=2920:2970#L107)
``` go
func (item *BodyItem) Scan(pSrc interface{}) error
```
Scan implements sql.Scanner interface




### <a name="BodyItem.UnmarshalJSON">func</a> (\*BodyItem) [UnmarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/item.go?s=2701:2755#L95)
``` go
func (item *BodyItem) UnmarshalJSON(data []byte) error
```
UnmarshalJSON implements the json.Unmarshaler interface




## <a name="Collection">type</a> [Collection](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/collection.go?s=356:1781#L18)
``` go
type Collection struct {
    ID          uuid.UUID               `db:"id"             json:"__id"`
    Namespace   namespace.Namespace     `db:"namespace"      json:"namespace"`
    Name        string                  `db:"name"           json:"name"`
    Code        string                  `db:"code"           json:"code"`
    Type        Type                    `db:"type"           json:"type"`
    Fields      types.Fields            `db:"fields"         json:"fields"`
    AccessType  AccessType              `db:"access_type"    json:"accessType"`
    Subordinate bool                    `db:"subordinate"    json:"subordinate"` // Разрешать доступ к элементам по иерархии оргструктуры
    Permissions permissions.Permissions `db:"permissions"    json:"permissions"`
    Indices     *bool                   `db:"indices"        json:"indices"` // Флаг влюченных индексов коллекции
    CreatedAt   time.Time               `db:"created_at"     json:"__createdAt"`
    CreatedBy   uuid.UUID               `db:"created_by"     json:"__createdBy"`
    UpdatedAt   time.Time               `db:"updated_at"     json:"__updatedAt"`
    UpdatedBy   uuid.UUID               `db:"updated_by"     json:"__updatedBy"`
    DeletedAt   *time.Time              `db:"deleted_at"     json:"__deletedAt"`
    Columned    bool                    `db:"-"              json:"-"               default:"false"`
}

```
Collection описание коллекции

easyjson:json










### <a name="Collection.GetCode">func</a> (\*Collection) [GetCode](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/collection.go?s=2073:2110#L43)
``` go
func (c *Collection) GetCode() string
```
GetCode имплементирует метод интерфейса connection.CollectionMetadata




### <a name="Collection.GetFields">func</a> (\*Collection) [GetFields](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/collection.go?s=2235:2280#L48)
``` go
func (c *Collection) GetFields() types.Fields
```
GetFields имплементирует метод интерфейса connection.CollectionMetadata




### <a name="Collection.GetNamespace">func</a> (\*Collection) [GetNamespace](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/collection.go?s=1890:1945#L38)
``` go
func (c *Collection) GetNamespace() namespace.Namespace
```
GetNamespace имплементирует метод интерфейса connection.CollectionMetadata




### <a name="Collection.MarshalEasyJSON">func</a> (Collection) [MarshalEasyJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/collection_easyjson.go?s=4769:4823#L211)
``` go
func (v Collection) MarshalEasyJSON(w *jwriter.Writer)
```
MarshalEasyJSON supports easyjson.Marshaler interface




### <a name="Collection.MarshalJSON">func</a> (Collection) [MarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/collection_easyjson.go?s=4525:4574#L204)
``` go
func (v Collection) MarshalJSON() ([]byte, error)
```
MarshalJSON supports json.Marshaler interface




### <a name="Collection.UnmarshalEasyJSON">func</a> (\*Collection) [UnmarshalEasyJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/collection_easyjson.go?s=5190:5245#L223)
``` go
func (v *Collection) UnmarshalEasyJSON(l *jlexer.Lexer)
```
UnmarshalEasyJSON supports easyjson.Unmarshaler interface




### <a name="Collection.UnmarshalJSON">func</a> (\*Collection) [UnmarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/collection_easyjson.go?s=4951:5004#L216)
``` go
func (v *Collection) UnmarshalJSON(data []byte) error
```
UnmarshalJSON supports json.Unmarshaler interface




## <a name="Item">type</a> [Item](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/item.go?s=507:937#L22)
``` go
type Item struct {
    ID        uuid.UUID  `db:"id"   json:"__id"`
    CreatedAt time.Time  `db:"-"    json:"__createdAt"`
    CreatedBy uuid.UUID  `db:"-"    json:"__createdBy"`
    UpdatedAt time.Time  `db:"-"    json:"__updatedAt"`
    UpdatedBy uuid.UUID  `db:"-"    json:"__updatedBy"`
    DeletedAt *time.Time `db:"-"    json:"__deletedAt"`
    Version   uint32     `db:"-"    json:"__version"`
    Name      string     `db:"-"    json:"__name"`
}

```
Item описывает основные системные поля элемента коллекции

easyjson:json










### <a name="Item.GetID">func</a> (Item) [GetID](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/item.go?s=1012:1046#L34)
``` go
func (item Item) GetID() uuid.UUID
```
GetID реализует интерфейс collection.ItemIdentifier




### <a name="Item.MarshalEasyJSON">func</a> (Item) [MarshalEasyJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/item_easyjson.go?s=3243:3291#L145)
``` go
func (v Item) MarshalEasyJSON(w *jwriter.Writer)
```
MarshalEasyJSON supports easyjson.Marshaler interface




### <a name="Item.MarshalJSON">func</a> (Item) [MarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/item_easyjson.go?s=3005:3048#L138)
``` go
func (v Item) MarshalJSON() ([]byte, error)
```
MarshalJSON supports json.Marshaler interface




### <a name="Item.Scan">func</a> (\*Item) [Scan](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/item.go?s=1109:1155#L39)
``` go
func (item *Item) Scan(pSrc interface{}) error
```
Scan implements sql.Scanner interface




### <a name="Item.UnmarshalEasyJSON">func</a> (\*Item) [UnmarshalEasyJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/item_easyjson.go?s=3652:3701#L157)
``` go
func (v *Item) UnmarshalEasyJSON(l *jlexer.Lexer)
```
UnmarshalEasyJSON supports easyjson.Unmarshaler interface




### <a name="Item.UnmarshalJSON">func</a> (\*Item) [UnmarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/item_easyjson.go?s=3419:3466#L150)
``` go
func (v *Item) UnmarshalJSON(data []byte) error
```
UnmarshalJSON supports json.Unmarshaler interface




## <a name="ItemIdentifier">type</a> [ItemIdentifier](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/item.go?s=325:377#L15)
``` go
type ItemIdentifier interface {
    GetID() uuid.UUID
}
```
ItemIdentifier описывает интерфейс c методом получения ИД элемента коллекции










## <a name="Metadata">type</a> [Metadata](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/collection.go?s=2390:2497#L53)
``` go
type Metadata interface {
    GetNamespace() namespace.Namespace
    GetCode() string
    GetFields() types.Fields
}
```
Metadata - метаданные коллекции (раздел, код, поля)










## <a name="SimpleMetadata">type</a> [SimpleMetadata](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/collection.go?s=2653:2756#L60)
``` go
type SimpleMetadata struct {
    Namespace namespace.Namespace
    Code      string
    Fields    types.Fields
}

```
SimpleMetadata - дефолтная реализация интерфейса метаданные коллекции (раздел, код, поля)










### <a name="SimpleMetadata.GetCode">func</a> (SimpleMetadata) [GetCode](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/collection.go?s=2912:2953#L72)
``` go
func (md SimpleMetadata) GetCode() string
```
GetCode - Collection code




### <a name="SimpleMetadata.GetFields">func</a> (SimpleMetadata) [GetFields](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/collection.go?s=3008:3057#L77)
``` go
func (md SimpleMetadata) GetFields() types.Fields
```
GetFields - Collection fields




### <a name="SimpleMetadata.GetNamespace">func</a> (SimpleMetadata) [GetNamespace](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/collection.go?s=2797:2856#L67)
``` go
func (md SimpleMetadata) GetNamespace() namespace.Namespace
```
GetNamespace - Collection namespace




## <a name="Type">type</a> [Type](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/item_type.go?s=212:226#L6)
``` go
type Type int8
```
Type - тип элементов коллекции


``` go
const (
    // TypeApplication приложение, обрабатывается в main
    TypeApplication Type = iota
    // TypeContract - контракт, обрабатывается сервисом contractor
    TypeContract
)
```






### <a name="TypeString">func</a> [TypeString](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/item_type_string.go?s=726:765#L32)
``` go
func TypeString(s string) (Type, error)
```
TypeString retrieves an enum value from the enum constants string name.
Throws an error if the param is not part of the enum.


### <a name="TypeValues">func</a> [TypeValues](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/item_type_string.go?s=943:967#L40)
``` go
func TypeValues() []Type
```
TypeValues returns all values of the enum





### <a name="Type.IsAType">func</a> (Type) [IsAType](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/item_type_string.go?s=1084:1112#L45)
``` go
func (i Type) IsAType() bool
```
IsAType returns "true" if the value is listed in the enum definition. "false" otherwise




### <a name="Type.MarshalJSON">func</a> (Type) [MarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/item_type_string.go?s=1265:1308#L55)
``` go
func (i Type) MarshalJSON() ([]byte, error)
```
MarshalJSON implements the json.Marshaler interface for Type




### <a name="Type.Scan">func</a> (\*Type) [Scan](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/item_type_string.go?s=1722:1766#L75)
``` go
func (i *Type) Scan(value interface{}) error
```



### <a name="Type.String">func</a> (Type) [String](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/item_type_string.go?s=305:334#L16)
``` go
func (i Type) String() string
```



### <a name="Type.UnmarshalJSON">func</a> (\*Type) [UnmarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/item_type_string.go?s=1415:1462#L60)
``` go
func (i *Type) UnmarshalJSON(data []byte) error
```
UnmarshalJSON implements the json.Unmarshaler interface for Type




### <a name="Type.Value">func</a> (Type) [Value](https://git.elewise.com/elma365/common/-/tree/develop/pkg/collection/item_type_string.go?s=1649:1692#L71)
``` go
func (i Type) Value() (driver.Value, error)
```








## <a name="pkg-subdirectories">Subdirectories</a>

| Name | Synopsis |
| ---- | -------- |
| [..](..) | |
| [pb](pb/) |  |
| [saver](saver/) |  |
| [saver/mock](saver/mock/) |  |


- - -
Generated by [godoc2md](https://github.com/Exa-Networks/godoc2md)
