# namespace
`import "git.elewise.com/elma365/common/pkg/namespace"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [type ConvertMap](#ConvertMap)
  * [func ConvertMapFromNamespaceLegacy(namespace Namespace) ConvertMap](#ConvertMapFromNamespaceLegacy)
  * [func (m ConvertMap) ExcludeSystemNamespaces() ConvertMap](#ConvertMap.ExcludeSystemNamespaces)
  * [func (m ConvertMap) Merge(second ConvertMap) ConvertMap](#ConvertMap.Merge)
  * [func (m ConvertMap) Rebase(namespace Namespace) Namespace](#ConvertMap.Rebase)
* [type Namespace](#Namespace)
  * [func FromParts(parts []string) Namespace](#FromParts)
  * [func FromPartsArg(parts ...string) Namespace](#FromPartsArg)
  * [func (ns Namespace) AsAbsolute(context Namespace) Namespace](#Namespace.AsAbsolute)
  * [func (ns Namespace) AsRelative(context Namespace) Namespace](#Namespace.AsRelative)
  * [func (ns Namespace) CutFrom(ns2 Namespace) []string](#Namespace.CutFrom)
  * [func (ns Namespace) GetDeepestSection() string](#Namespace.GetDeepestSection)
  * [func (ns Namespace) GetLevel() int](#Namespace.GetLevel)
  * [func (ns Namespace) GetParent() Namespace](#Namespace.GetParent)
  * [func (ns Namespace) GetRoot() Namespace](#Namespace.GetRoot)
  * [func (ns Namespace) GetSections() []string](#Namespace.GetSections)
  * [func (ns Namespace) GetTopLevel() Namespace](#Namespace.GetTopLevel)
  * [func (ns Namespace) IsExtension() bool](#Namespace.IsExtension)
  * [func (ns Namespace) IsSubNamespaceOf(context Namespace) bool](#Namespace.IsSubNamespaceOf)
  * [func (ns Namespace) IsSuperNamespaceFor(ns2 Namespace) bool](#Namespace.IsSuperNamespaceFor)
  * [func (ns Namespace) IsSystem() bool](#Namespace.IsSystem)
  * [func (ns Namespace) IsSystemCode(code string) bool](#Namespace.IsSystemCode)
  * [func (ns Namespace) MakeCollectionID(code string) uuid.UUID](#Namespace.MakeCollectionID)
  * [func (ns Namespace) MakeSubNamespace(subSections ...string) Namespace](#Namespace.MakeSubNamespace)
  * [func (ns Namespace) String() string](#Namespace.String)
* [type Namespaces](#Namespaces)
  * [func (n Namespaces) GetAllSuperFor(ns Namespace) Namespaces](#Namespaces.GetAllSuperFor)
  * [func (n Namespaces) GetHighSuperFor(ns Namespace) (Namespace, error)](#Namespaces.GetHighSuperFor)
  * [func (n Namespaces) HasSuperFor(ns Namespace) bool](#Namespaces.HasSuperFor)
  * [func (n Namespaces) SortByLevel() Namespaces](#Namespaces.SortByLevel)
* [type SortByLevel](#SortByLevel)
  * [func (n SortByLevel) Len() int](#SortByLevel.Len)
  * [func (n SortByLevel) Less(i, j int) bool](#SortByLevel.Less)
  * [func (n SortByLevel) Swap(i, j int)](#SortByLevel.Swap)


#### <a name="pkg-files">Package files</a>
[const.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/namespace/const.go) [convert_map.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/namespace/convert_map.go) [namespace.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/namespace/namespace.go) [namespaces.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/namespace/namespaces.go)


## <a name="pkg-constants">Constants</a>
``` go
const (
    // SuperNamespaceNotFoundError - отсутствие надNamespace
    SuperNamespaceNotFoundError = "super namespace not found"
)
```
``` go
const SystemCodePrefix = "_"
```
SystemCodePrefix - префикс для системных кодов, которые иначе не получается пометить как системный





## <a name="ConvertMap">type</a> [ConvertMap](https://git.elewise.com/elma365/common/-/tree/develop/pkg/namespace/convert_map.go?s=85:124#L4)
``` go
type ConvertMap map[Namespace]Namespace
```
ConvertMap - карта преобразований namespace







### <a name="ConvertMapFromNamespaceLegacy">func</a> [ConvertMapFromNamespaceLegacy](https://git.elewise.com/elma365/common/-/tree/develop/pkg/namespace/convert_map.go?s=1862:1928#L49)
``` go
func ConvertMapFromNamespaceLegacy(namespace Namespace) ConvertMap
```
ConvertMapFromNamespaceLegacy - преобразовывает namespace в ConvertMap нужно для устаревшего механизма
Deprecated: только для совместимости. Как всезде будет переведено на новый механизм можно удалять





### <a name="ConvertMap.ExcludeSystemNamespaces">func</a> (ConvertMap) [ExcludeSystemNamespaces](https://git.elewise.com/elma365/common/-/tree/develop/pkg/namespace/convert_map.go?s=740:796#L18)
``` go
func (m ConvertMap) ExcludeSystemNamespaces() ConvertMap
```
ExcludeSystemNamespaces фильтрует из карты системные области видимости, используемые в качестве ключей




### <a name="ConvertMap.Merge">func</a> (ConvertMap) [Merge](https://git.elewise.com/elma365/common/-/tree/develop/pkg/namespace/convert_map.go?s=1329:1384#L33)
``` go
func (m ConvertMap) Merge(second ConvertMap) ConvertMap
```
Merge - слияние карт. В новой карте будут ключи и значения из основной карты,
дополененной ключами и значениями из второй. При совпадении ключей приоритет у основной карты.
Исходные карты остаются без изменения




### <a name="ConvertMap.Rebase">func</a> (ConvertMap) [Rebase](https://git.elewise.com/elma365/common/-/tree/develop/pkg/namespace/convert_map.go?s=353:410#L8)
``` go
func (m ConvertMap) Rebase(namespace Namespace) Namespace
```
Rebase - получение нового namespace в соответствии с картой преобразования
вернёт тот же namespace если в карте нет подходящего пути




## <a name="Namespace">type</a> [Namespace](https://git.elewise.com/elma365/common/-/tree/develop/pkg/namespace/namespace.go?s=157:178#L12)
``` go
type Namespace string
```
Namespace - тип для пространства имен


``` go
const (
    // Separator - разделитель пространств имен
    Separator = "."
    // System - системное пространство имен
    System Namespace = "system"
    // Global - глобальное пространство имен
    Global Namespace = "global"
    // Virtual - виртуальное пространство имен
    Virtual Namespace = "virtual"
    // Current - ссылочное пространство имен. При вычислении абсолютного заменяется на текущее пространство имен
    // Deprecated: отказались от относительных namespace
    Current Namespace = "$current"
    // Parent - ссылочное пространство имен. При вычислении абсолютного заменяется на родительское пространство имен
    // Deprecated: отказались от относительных namespace
    Parent Namespace = "$parent"
)
```






### <a name="FromParts">func</a> [FromParts](https://git.elewise.com/elma365/common/-/tree/develop/pkg/namespace/namespace.go?s=613:653#L20)
``` go
func FromParts(parts []string) Namespace
```
FromParts - получение пространства имён из массива составляющих


### <a name="FromPartsArg">func</a> [FromPartsArg](https://git.elewise.com/elma365/common/-/tree/develop/pkg/namespace/namespace.go?s=847:891#L28)
``` go
func FromPartsArg(parts ...string) Namespace
```
FromPartsArg - получение пространства имён из аргументов





### <a name="Namespace.AsAbsolute">func</a> (Namespace) [AsAbsolute](https://git.elewise.com/elma365/common/-/tree/develop/pkg/namespace/namespace.go?s=3479:3538#L103)
``` go
func (ns Namespace) AsAbsolute(context Namespace) Namespace
```
AsAbsolute - получение абсолютного пространства имен по относительному
Deprecated: отказываемся от относительных путей




### <a name="Namespace.AsRelative">func</a> (Namespace) [AsRelative](https://git.elewise.com/elma365/common/-/tree/develop/pkg/namespace/namespace.go?s=2482:2541#L75)
``` go
func (ns Namespace) AsRelative(context Namespace) Namespace
```
AsRelative - построение относительно пути в контексте пространства имён context
Deprecated: отказываемся от относительных путей




### <a name="Namespace.CutFrom">func</a> (Namespace) [CutFrom](https://git.elewise.com/elma365/common/-/tree/develop/pkg/namespace/namespace.go?s=5430:5481#L154)
``` go
func (ns Namespace) CutFrom(ns2 Namespace) []string
```
CutFrom - вырезает ns из начала ns2 и возвращает оставшиеся части списком




### <a name="Namespace.GetDeepestSection">func</a> (Namespace) [GetDeepestSection](https://git.elewise.com/elma365/common/-/tree/develop/pkg/namespace/namespace.go?s=4838:4884#L131)
``` go
func (ns Namespace) GetDeepestSection() string
```
GetDeepestSection - возвращает наиболее глубокую часть. Например достать код приложения из полного неймспейса




### <a name="Namespace.GetLevel">func</a> (Namespace) [GetLevel](https://git.elewise.com/elma365/common/-/tree/develop/pkg/namespace/namespace.go?s=1953:1987#L61)
``` go
func (ns Namespace) GetLevel() int
```
GetLevel - уроваень вложенности пространства имён




### <a name="Namespace.GetParent">func</a> (Namespace) [GetParent](https://git.elewise.com/elma365/common/-/tree/develop/pkg/namespace/namespace.go?s=3061:3102#L90)
``` go
func (ns Namespace) GetParent() Namespace
```
GetParent - возвращает родительское пространство имён




### <a name="Namespace.GetRoot">func</a> (Namespace) [GetRoot](https://git.elewise.com/elma365/common/-/tree/develop/pkg/namespace/namespace.go?s=1460:1499#L47)
``` go
func (ns Namespace) GetRoot() Namespace
```
GetRoot - возвращает корневое пространство имён (global или system)




### <a name="Namespace.GetSections">func</a> (Namespace) [GetSections](https://git.elewise.com/elma365/common/-/tree/develop/pkg/namespace/namespace.go?s=2170:2212#L69)
``` go
func (ns Namespace) GetSections() []string
```
GetSections - составляющие пространства имён в виде массива строк




### <a name="Namespace.GetTopLevel">func</a> (Namespace) [GetTopLevel](https://git.elewise.com/elma365/common/-/tree/develop/pkg/namespace/namespace.go?s=6073:6116#L174)
``` go
func (ns Namespace) GetTopLevel() Namespace
```
GetTopLevel - получить namespace верхнего уровня (код раздела)




### <a name="Namespace.IsExtension">func</a> (Namespace) [IsExtension](https://git.elewise.com/elma365/common/-/tree/develop/pkg/namespace/namespace.go?s=6297:6335#L180)
``` go
func (ns Namespace) IsExtension() bool
```
IsExtension проверяет, является ли пространством имён расширения




### <a name="Namespace.IsSubNamespaceOf">func</a> (Namespace) [IsSubNamespaceOf](https://git.elewise.com/elma365/common/-/tree/develop/pkg/namespace/namespace.go?s=1083:1143#L33)
``` go
func (ns Namespace) IsSubNamespaceOf(context Namespace) bool
```
IsSubNamespaceOf - проверяет является ли текущее пространство подпространством контекстного




### <a name="Namespace.IsSuperNamespaceFor">func</a> (Namespace) [IsSuperNamespaceFor](https://git.elewise.com/elma365/common/-/tree/develop/pkg/namespace/namespace.go?s=1767:1826#L56)
``` go
func (ns Namespace) IsSuperNamespaceFor(ns2 Namespace) bool
```
IsSuperNamespaceFor - проверяет является ли текущее пространство имён надпространством для ns2




### <a name="Namespace.IsSystem">func</a> (Namespace) [IsSystem](https://git.elewise.com/elma365/common/-/tree/develop/pkg/namespace/namespace.go?s=5059:5094#L140)
``` go
func (ns Namespace) IsSystem() bool
```
IsSystem - признак системности




### <a name="Namespace.IsSystemCode">func</a> (Namespace) [IsSystemCode](https://git.elewise.com/elma365/common/-/tree/develop/pkg/namespace/namespace.go?s=4541:4591#L126)
``` go
func (ns Namespace) IsSystemCode(code string) bool
```
IsSystemCode - является ли приложение с неймспейсом ns и кодом code системным.
текущий признак системного приложения - код начинается с символа "_"




### <a name="Namespace.MakeCollectionID">func</a> (Namespace) [MakeCollectionID](https://git.elewise.com/elma365/common/-/tree/develop/pkg/namespace/namespace.go?s=3989:4048#L115)
``` go
func (ns Namespace) MakeCollectionID(code string) uuid.UUID
```
MakeCollectionID формирует id коллекции для данного раздела и кода




### <a name="Namespace.MakeSubNamespace">func</a> (Namespace) [MakeSubNamespace](https://git.elewise.com/elma365/common/-/tree/develop/pkg/namespace/namespace.go?s=5697:5766#L162)
``` go
func (ns Namespace) MakeSubNamespace(subSections ...string) Namespace
```
MakeSubNamespace - создать подNamespace добавив части из аргумента




### <a name="Namespace.String">func</a> (Namespace) [String](https://git.elewise.com/elma365/common/-/tree/develop/pkg/namespace/namespace.go?s=4225:4260#L120)
``` go
func (ns Namespace) String() string
```
String имплементация fmt.Stringer - строковое представление ns




## <a name="Namespaces">type</a> [Namespaces](https://git.elewise.com/elma365/common/-/tree/develop/pkg/namespace/namespaces.go?s=88:115#L9)
``` go
type Namespaces []Namespace
```
Namespaces - список namespace










### <a name="Namespaces.GetAllSuperFor">func</a> (Namespaces) [GetAllSuperFor](https://git.elewise.com/elma365/common/-/tree/develop/pkg/namespace/namespaces.go?s=1798:1857#L63)
``` go
func (n Namespaces) GetAllSuperFor(ns Namespace) Namespaces
```
GetAllSuperFor - возвращает все namespace из списка которые являются надNamespace для данного




### <a name="Namespaces.GetHighSuperFor">func</a> (Namespaces) [GetHighSuperFor](https://git.elewise.com/elma365/common/-/tree/develop/pkg/namespace/namespaces.go?s=726:794#L28)
``` go
func (n Namespaces) GetHighSuperFor(ns Namespace) (Namespace, error)
```
GetHighSuperFor - вернуть самый верхний (с самым коротким путём) namespace из списка,
который является надNamespace для данного




### <a name="Namespaces.HasSuperFor">func</a> (Namespaces) [HasSuperFor](https://git.elewise.com/elma365/common/-/tree/develop/pkg/namespace/namespaces.go?s=370:420#L17)
``` go
func (n Namespaces) HasSuperFor(ns Namespace) bool
```
HasSuperFor - признак того, что в списке есть надNamespace для данного




### <a name="Namespaces.SortByLevel">func</a> (Namespaces) [SortByLevel](https://git.elewise.com/elma365/common/-/tree/develop/pkg/namespace/namespaces.go?s=1072:1116#L38)
``` go
func (n Namespaces) SortByLevel() Namespaces
```
SortByLevel - отсортировать по уровню




## <a name="SortByLevel">type</a> [SortByLevel](https://git.elewise.com/elma365/common/-/tree/develop/pkg/namespace/namespaces.go?s=1281:1308#L45)
``` go
type SortByLevel Namespaces
```
SortByLevel - релизация сортировки по уровню Namespace










### <a name="SortByLevel.Len">func</a> (SortByLevel) [Len](https://git.elewise.com/elma365/common/-/tree/develop/pkg/namespace/namespaces.go?s=1355:1385#L48)
``` go
func (n SortByLevel) Len() int
```
Len - реализация sort.Interface




### <a name="SortByLevel.Less">func</a> (SortByLevel) [Less](https://git.elewise.com/elma365/common/-/tree/develop/pkg/namespace/namespaces.go?s=1452:1492#L53)
``` go
func (n SortByLevel) Less(i, j int) bool
```
Less - реализация sort.Interface




### <a name="SortByLevel.Swap">func</a> (SortByLevel) [Swap](https://git.elewise.com/elma365/common/-/tree/develop/pkg/namespace/namespaces.go?s=1586:1621#L58)
``` go
func (n SortByLevel) Swap(i, j int)
```
Swap - реализация sort.Interface







- - -
Generated by [godoc2md](https://github.com/Exa-Networks/godoc2md)
