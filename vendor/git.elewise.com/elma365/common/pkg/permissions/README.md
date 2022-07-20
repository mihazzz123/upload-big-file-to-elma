# permissions
`import "git.elewise.com/elma365/common/pkg/permissions"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [type Access](#Access)
  * [func (a Access) Has(t AccessAtom) bool](#Access.Has)
  * [func (a Access) MarshalJSON() ([]byte, error)](#Access.MarshalJSON)
  * [func (a *Access) UnmarshalJSON(data []byte) error](#Access.UnmarshalJSON)
* [type AccessAtom](#AccessAtom)
  * [func AccessAtomString(s string) (AccessAtom, error)](#AccessAtomString)
  * [func AccessAtomValues() []AccessAtom](#AccessAtomValues)
  * [func (i AccessAtom) IsAAccessAtom() bool](#AccessAtom.IsAAccessAtom)
  * [func (i AccessAtom) MarshalJSON() ([]byte, error)](#AccessAtom.MarshalJSON)
  * [func (i AccessAtom) String() string](#AccessAtom.String)
  * [func (i *AccessAtom) UnmarshalJSON(data []byte) error](#AccessAtom.UnmarshalJSON)
* [type Group](#Group)
  * [func (v Group) MarshalEasyJSON(w *jwriter.Writer)](#Group.MarshalEasyJSON)
  * [func (v Group) MarshalJSON() ([]byte, error)](#Group.MarshalJSON)
  * [func (v *Group) UnmarshalEasyJSON(l *jlexer.Lexer)](#Group.UnmarshalEasyJSON)
  * [func (v *Group) UnmarshalJSON(data []byte) error](#Group.UnmarshalJSON)
* [type GroupType](#GroupType)
  * [func GroupTypeString(s string) (GroupType, error)](#GroupTypeString)
  * [func GroupTypeValues() []GroupType](#GroupTypeValues)
  * [func (i GroupType) IsAGroupType() bool](#GroupType.IsAGroupType)
  * [func (i GroupType) MarshalJSON() ([]byte, error)](#GroupType.MarshalJSON)
  * [func (i GroupType) String() string](#GroupType.String)
  * [func (i *GroupType) UnmarshalJSON(data []byte) error](#GroupType.UnmarshalJSON)
* [type Permission](#Permission)
  * [func (v Permission) MarshalEasyJSON(w *jwriter.Writer)](#Permission.MarshalEasyJSON)
  * [func (v Permission) MarshalJSON() ([]byte, error)](#Permission.MarshalJSON)
  * [func (pn *Permission) Scan(pSrc interface{}) error](#Permission.Scan)
  * [func (v *Permission) UnmarshalEasyJSON(l *jlexer.Lexer)](#Permission.UnmarshalEasyJSON)
  * [func (v *Permission) UnmarshalJSON(data []byte) error](#Permission.UnmarshalJSON)
  * [func (pn Permission) Value() (driver.Value, error)](#Permission.Value)
* [type Permissions](#Permissions)
  * [func (p Permissions) Add(plus []Permission) Permissions](#Permissions.Add)
  * [func (p Permissions) AsInherited() []Permission](#Permissions.AsInherited)
  * [func (p Permissions) Delete(minus []Permission) Permissions](#Permissions.Delete)
  * [func (p Permissions) GetIDsByAtom(atom AccessAtom) uuids.UUIDS](#Permissions.GetIDsByAtom)
  * [func (p Permissions) Inherited() []Permission](#Permissions.Inherited)
  * [func (p Permissions) Marshal() ([]byte, error)](#Permissions.Marshal)
  * [func (v Permissions) MarshalEasyJSON(w *jwriter.Writer)](#Permissions.MarshalEasyJSON)
  * [func (v Permissions) MarshalJSON() ([]byte, error)](#Permissions.MarshalJSON)
  * [func (p *Permissions) MarshalTo(data []byte) (n int, err error)](#Permissions.MarshalTo)
  * [func (p Permissions) NonInherited() []Permission](#Permissions.NonInherited)
  * [func (p *Permissions) Scan(pSrc interface{}) error](#Permissions.Scan)
  * [func (p *Permissions) Size() int](#Permissions.Size)
  * [func (p *Permissions) Unmarshal(data []byte) error](#Permissions.Unmarshal)
  * [func (v *Permissions) UnmarshalEasyJSON(l *jlexer.Lexer)](#Permissions.UnmarshalEasyJSON)
  * [func (v *Permissions) UnmarshalJSON(data []byte) error](#Permissions.UnmarshalJSON)
  * [func (p Permissions) Value() (driver.Value, error)](#Permissions.Value)


#### <a name="pkg-files">Package files</a>
[access.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/access.go) [access_string.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/access_string.go) [group_type.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/group_type.go) [group_type_string.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/group_type_string.go) [permissions.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/permissions.go) [permissions_easyjson.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/permissions_easyjson.go)


## <a name="pkg-constants">Constants</a>
``` go
const (
    // FullAccess Полный доступ
    FullAccess = Access(READ | CREATE | UPDATE | DELETE | ASSIGN | BPMANAGE | EXPORT | IMPORT)
    // FullFileAccess Полный доступ к файлу
    FullFileAccess = Access(READ | UPDATE | DELETE | ASSIGN | EXPORT | IMPORT)
    // FullDirAccess Полный доступ к директории
    FullDirAccess = Access(READ | CREATE | UPDATE | DELETE | ASSIGN | EXPORT | IMPORT)
    // EmptyAccess Отсутствие доступа
    EmptyAccess = Access(0)
)
```




## <a name="Access">type</a> [Access](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/access.go?s=838:860#L35)
``` go
type Access AccessAtom
```
Access набор разрешений










### <a name="Access.Has">func</a> (Access) [Has](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/access.go?s=1469:1507#L49)
``` go
func (a Access) Has(t AccessAtom) bool
```
Has проверяет наличие атомарного разрешения в наборе




### <a name="Access.MarshalJSON">func</a> (Access) [MarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/access.go?s=1578:1623#L54)
``` go
func (a Access) MarshalJSON() ([]byte, error)
```
MarshalJSON implements json.Marshaler




### <a name="Access.UnmarshalJSON">func</a> (\*Access) [UnmarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/access.go?s=2139:2188#L74)
``` go
func (a *Access) UnmarshalJSON(data []byte) error
```
UnmarshalJSON implements json.Unmarshaler




## <a name="AccessAtom">type</a> [AccessAtom](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/access.go?s=268:287#L13)
``` go
type AccessAtom int
```
AccessAtom атомарное разрешение


``` go
const (
    // READ чтение элемента
    READ AccessAtom = 1 << iota
    // CREATE создание элемента (в папке или справочнике)
    CREATE
    // UPDATE изменение элемента
    UPDATE
    // DELETE удаление элемента
    DELETE
    // ASSIGN изменение прав
    ASSIGN
    // BPMANAGE управление процессами элемента
    BPMANAGE
    // EXPORT экспорт данных
    EXPORT
    // IMPORT импорт данных
    IMPORT
)
```






### <a name="AccessAtomString">func</a> [AccessAtomString](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/access_string.go?s=1680:1731#L68)
``` go
func AccessAtomString(s string) (AccessAtom, error)
```
AccessAtomString retrieves an enum value from the enum constants string name.
Throws an error if the param is not part of the enum.


### <a name="AccessAtomValues">func</a> [AccessAtomValues](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/access_string.go?s=1927:1963#L76)
``` go
func AccessAtomValues() []AccessAtom
```
AccessAtomValues returns all values of the enum





### <a name="AccessAtom.IsAAccessAtom">func</a> (AccessAtom) [IsAAccessAtom](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/access_string.go?s=2092:2132#L81)
``` go
func (i AccessAtom) IsAAccessAtom() bool
```
IsAAccessAtom returns "true" if the value is listed in the enum definition. "false" otherwise




### <a name="AccessAtom.MarshalJSON">func</a> (AccessAtom) [MarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/access_string.go?s=2297:2346#L91)
``` go
func (i AccessAtom) MarshalJSON() ([]byte, error)
```
MarshalJSON implements the json.Marshaler interface for AccessAtom




### <a name="AccessAtom.String">func</a> (AccessAtom) [String](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/access_string.go?s=716:751#L31)
``` go
func (i AccessAtom) String() string
```



### <a name="AccessAtom.UnmarshalJSON">func</a> (\*AccessAtom) [UnmarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/access_string.go?s=2459:2512#L96)
``` go
func (i *AccessAtom) UnmarshalJSON(data []byte) error
```
UnmarshalJSON implements the json.Unmarshaler interface for AccessAtom




## <a name="Group">type</a> [Group](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/permissions.go?s=357:478#L19)
``` go
type Group struct {
    ID   uuid.UUID `json:"id"   validate:"required"`
    Type GroupType `json:"type" validate:"required"`
}

```
Group группа владельца доступа

easyjson:json










### <a name="Group.MarshalEasyJSON">func</a> (Group) [MarshalEasyJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/permissions_easyjson.go?s=6868:6917#L292)
``` go
func (v Group) MarshalEasyJSON(w *jwriter.Writer)
```
MarshalEasyJSON supports easyjson.Marshaler interface




### <a name="Group.MarshalJSON">func</a> (Group) [MarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/permissions_easyjson.go?s=6627:6671#L285)
``` go
func (v Group) MarshalJSON() ([]byte, error)
```
MarshalJSON supports json.Marshaler interface




### <a name="Group.UnmarshalEasyJSON">func</a> (\*Group) [UnmarshalEasyJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/permissions_easyjson.go?s=7283:7333#L304)
``` go
func (v *Group) UnmarshalEasyJSON(l *jlexer.Lexer)
```
UnmarshalEasyJSON supports easyjson.Unmarshaler interface




### <a name="Group.UnmarshalJSON">func</a> (\*Group) [UnmarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/permissions_easyjson.go?s=7047:7095#L297)
``` go
func (v *Group) UnmarshalJSON(data []byte) error
```
UnmarshalJSON supports json.Unmarshaler interface




## <a name="GroupType">type</a> [GroupType](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/group_type.go?s=232:251#L6)
``` go
type GroupType int8
```
GroupType тип группы владельца доступа


``` go
const (
    // GroupTypeUser — пользователь
    GroupTypeUser GroupType = iota
    // GroupTypeGroup — группа
    GroupTypeGroup
    // GroupTypeOrgstruct — элемент оргуструктуры
    GroupTypeOrgstruct
    // GroupTypeRole — роль (группа с единственным участником
    GroupTypeRole
)
```






### <a name="GroupTypeString">func</a> [GroupTypeString](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/group_type_string.go?s=860:909#L33)
``` go
func GroupTypeString(s string) (GroupType, error)
```
GroupTypeString retrieves an enum value from the enum constants string name.
Throws an error if the param is not part of the enum.


### <a name="GroupTypeValues">func</a> [GroupTypeValues](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/group_type_string.go?s=1102:1136#L41)
``` go
func GroupTypeValues() []GroupType
```
GroupTypeValues returns all values of the enum





### <a name="GroupType.IsAGroupType">func</a> (GroupType) [IsAGroupType](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/group_type_string.go?s=1263:1301#L46)
``` go
func (i GroupType) IsAGroupType() bool
```
IsAGroupType returns "true" if the value is listed in the enum definition. "false" otherwise




### <a name="GroupType.MarshalJSON">func</a> (GroupType) [MarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/group_type_string.go?s=1464:1512#L56)
``` go
func (i GroupType) MarshalJSON() ([]byte, error)
```
MarshalJSON implements the json.Marshaler interface for GroupType




### <a name="GroupType.String">func</a> (GroupType) [String](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/group_type_string.go?s=309:343#L15)
``` go
func (i GroupType) String() string
```



### <a name="GroupType.UnmarshalJSON">func</a> (\*GroupType) [UnmarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/group_type_string.go?s=1624:1676#L61)
``` go
func (i *GroupType) UnmarshalJSON(data []byte) error
```
UnmarshalJSON implements the json.Unmarshaler interface for GroupType




## <a name="Permission">type</a> [Permission](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/permissions.go?s=675:892#L27)
``` go
type Permission struct {
    Group     Group  `json:"group"`
    Types     Access `json:"types"`
    Inherited bool   `json:"inherited"` // флаг - это наследованная пермиссия, или родная
}

```
Permission отдельный доступ, связанный с определённой группой пользователей и набором разрешений

easyjson:json










### <a name="Permission.MarshalEasyJSON">func</a> (Permission) [MarshalEasyJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/permissions_easyjson.go?s=4947:5001#L215)
``` go
func (v Permission) MarshalEasyJSON(w *jwriter.Writer)
```
MarshalEasyJSON supports easyjson.Marshaler interface




### <a name="Permission.MarshalJSON">func</a> (Permission) [MarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/permissions_easyjson.go?s=4701:4750#L208)
``` go
func (v Permission) MarshalJSON() ([]byte, error)
```
MarshalJSON supports json.Marshaler interface




### <a name="Permission.Scan">func</a> (\*Permission) [Scan](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/permissions.go?s=935:985#L34)
``` go
func (pn *Permission) Scan(pSrc interface{}) error
```
Scan implements sql.Scanner interface




### <a name="Permission.UnmarshalEasyJSON">func</a> (\*Permission) [UnmarshalEasyJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/permissions_easyjson.go?s=5372:5427#L227)
``` go
func (v *Permission) UnmarshalEasyJSON(l *jlexer.Lexer)
```
UnmarshalEasyJSON supports easyjson.Unmarshaler interface




### <a name="Permission.UnmarshalJSON">func</a> (\*Permission) [UnmarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/permissions_easyjson.go?s=5131:5184#L220)
``` go
func (v *Permission) UnmarshalJSON(data []byte) error
```
UnmarshalJSON supports json.Unmarshaler interface




### <a name="Permission.Value">func</a> (Permission) [Value](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/permissions.go?s=1254:1304#L50)
``` go
func (pn Permission) Value() (driver.Value, error)
```
Value implements sql.Valuer interface




## <a name="Permissions">type</a> [Permissions](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/permissions.go?s=1416:1991#L57)
``` go
type Permissions struct {
    InheritParent bool             `json:"inheritParent"`     // глобальный флаг, наследовать или нет родительские пермиссии
    Values        []Permission     `json:"values"`            // список пермиссий
    Timestamp     int64            `json:"timestamp"`         // дата изменения пермишшенов, в unixtime
    RefItem       *refitem.RefItem `json:"refItem,omitempty"` // ссылка на элемент приложения, учитываем его права
}

```
Permissions набор доступов к объекту

easyjson:json










### <a name="Permissions.Add">func</a> (Permissions) [Add](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/permissions.go?s=4050:4105#L140)
``` go
func (p Permissions) Add(plus []Permission) Permissions
```
Add добавляет ненаследованные права в список

Если права для группы уже были в списке и были не наследованы,
то новые права просто добавляются к имеющимся.
Иначе добавляется новая строка.

Добавляемые наследованные права откидываются.




### <a name="Permissions.AsInherited">func</a> (Permissions) [AsInherited](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/permissions.go?s=2634:2681#L89)
``` go
func (p Permissions) AsInherited() []Permission
```
AsInherited returns all permissions as inherited




### <a name="Permissions.Delete">func</a> (Permissions) [Delete](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/permissions.go?s=4634:4693#L165)
``` go
func (p Permissions) Delete(minus []Permission) Permissions
```
Delete удаляет ненаследованные права из списка

Наследованные права не затрагиваются.




### <a name="Permissions.GetIDsByAtom">func</a> (Permissions) [GetIDsByAtom](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/permissions.go?s=3356:3418#L121)
``` go
func (p Permissions) GetIDsByAtom(atom AccessAtom) uuids.UUIDS
```
GetIDsByAtom — возвращает идентификаторы групп, имеющих указанный атом в правах




### <a name="Permissions.Inherited">func</a> (Permissions) [Inherited](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/permissions.go?s=2090:2135#L65)
``` go
func (p Permissions) Inherited() []Permission
```
Inherited возвращает только отнаследованные пермиссии




### <a name="Permissions.Marshal">func</a> (Permissions) [Marshal](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/permissions.go?s=5532:5578#L209)
``` go
func (p Permissions) Marshal() ([]byte, error)
```
Marshal marshaler interface




### <a name="Permissions.MarshalEasyJSON">func</a> (Permissions) [MarshalEasyJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/permissions_easyjson.go?s=2902:2957#L133)
``` go
func (v Permissions) MarshalEasyJSON(w *jwriter.Writer)
```
MarshalEasyJSON supports easyjson.Marshaler interface




### <a name="Permissions.MarshalJSON">func</a> (Permissions) [MarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/permissions_easyjson.go?s=2656:2706#L126)
``` go
func (v Permissions) MarshalJSON() ([]byte, error)
```
MarshalJSON supports json.Marshaler interface




### <a name="Permissions.MarshalTo">func</a> (\*Permissions) [MarshalTo](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/permissions.go?s=5650:5713#L214)
``` go
func (p *Permissions) MarshalTo(data []byte) (n int, err error)
```
MarshalTo protobuf marshaler interface




### <a name="Permissions.NonInherited">func</a> (Permissions) [NonInherited](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/permissions.go?s=2392:2440#L77)
``` go
func (p Permissions) NonInherited() []Permission
```
NonInherited возвращает только родные пермиссии (не наследованные)




### <a name="Permissions.Scan">func</a> (\*Permissions) [Scan](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/permissions.go?s=5060:5110#L186)
``` go
func (p *Permissions) Scan(pSrc interface{}) error
```
Scan implements sql.Scanner interface




### <a name="Permissions.Size">func</a> (\*Permissions) [Size](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/permissions.go?s=5960:5992#L228)
``` go
func (p *Permissions) Size() int
```
Size - size for protobuf




### <a name="Permissions.Unmarshal">func</a> (\*Permissions) [Unmarshal](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/permissions.go?s=5844:5894#L223)
``` go
func (p *Permissions) Unmarshal(data []byte) error
```
Unmarshal unmarshaler interface




### <a name="Permissions.UnmarshalEasyJSON">func</a> (\*Permissions) [UnmarshalEasyJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/permissions_easyjson.go?s=3327:3383#L145)
``` go
func (v *Permissions) UnmarshalEasyJSON(l *jlexer.Lexer)
```
UnmarshalEasyJSON supports easyjson.Unmarshaler interface




### <a name="Permissions.UnmarshalJSON">func</a> (\*Permissions) [UnmarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/permissions_easyjson.go?s=3086:3140#L138)
``` go
func (v *Permissions) UnmarshalJSON(data []byte) error
```
UnmarshalJSON supports json.Unmarshaler interface




### <a name="Permissions.Value">func</a> (Permissions) [Value](https://git.elewise.com/elma365/common/-/tree/develop/pkg/permissions/permissions.go?s=5380:5430#L202)
``` go
func (p Permissions) Value() (driver.Value, error)
```
Value implements sql.Valuer interface







- - -
Generated by [godoc2md](https://github.com/Exa-Networks/godoc2md)
