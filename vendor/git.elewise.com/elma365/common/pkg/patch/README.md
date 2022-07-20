# patch
`import "git.elewise.com/elma365/common/pkg/patch"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [func CheckExtracts(blob json.RawMessage, extracts Extracts) ([]string, error)](#CheckExtracts)
* [func ExtractPO(blob json.RawMessage, extracts Extracts) ([]byte, error)](#ExtractPO)
* [func ReplacePO(blob json.RawMessage, po []byte) (json.RawMessage, error)](#ReplacePO)
* [type Extract](#Extract)
  * [func ExtractFromString(s string) (Extract, error)](#ExtractFromString)
  * [func NewExtract(path ...interface{}) Extract](#NewExtract)
  * [func (extract Extract) Apply(container *gabs.Container) interface{}](#Extract.Apply)
  * [func (extract Extract) Prefix(path ...interface{}) Extract](#Extract.Prefix)
  * [func (extract Extract) String() string](#Extract.String)
  * [func (extract Extract) ToReplace(val interface{}) Replace](#Extract.ToReplace)
* [type Extracter](#Extracter)
* [type Extracts](#Extracts)
  * [func CreateExtracts(val interface{}, target string) (Extracts, error)](#CreateExtracts)
  * [func (extracts Extracts) Apply(blob json.RawMessage) (map[string]interface{}, error)](#Extracts.Apply)
  * [func (extracts Extracts) Prefix(path ...interface{}) Extracts](#Extracts.Prefix)
  * [func (extracts Extracts) ToStrings() []string](#Extracts.ToStrings)
* [type Path](#Path)
  * [func FromString(s string) (Path, error)](#FromString)
  * [func NewPath(parts ...interface{}) Path](#NewPath)
  * [func (path Path) Get(container *gabs.Container) interface{}](#Path.Get)
  * [func (path Path) Prefix(parts ...interface{}) Path](#Path.Prefix)
  * [func (path Path) Search(container *gabs.Container) *gabs.Container](#Path.Search)
  * [func (path Path) Set(container *gabs.Container, value interface{}) error](#Path.Set)
  * [func (path Path) String() string](#Path.String)
* [type PathPart](#PathPart)
  * [func NewPart(x interface{}) PathPart](#NewPart)
* [type Replace](#Replace)
  * [func NewReplace(value interface{}, path ...interface{}) Replace](#NewReplace)
  * [func (replace Replace) Apply(container *gabs.Container) error](#Replace.Apply)
  * [func (replace Replace) Prefix(parts ...interface{}) Replace](#Replace.Prefix)
* [type Replaces](#Replaces)
  * [func (replaces Replaces) Apply(blob json.RawMessage) (json.RawMessage, error)](#Replaces.Apply)
  * [func (replaces Replaces) ApplyToContainer(container *gabs.Container) error](#Replaces.ApplyToContainer)
  * [func (replaces Replaces) Prefix(path ...interface{}) Replaces](#Replaces.Prefix)


#### <a name="pkg-files">Package files</a>
[extract.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/patch/extract.go) [extracts_creater.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/patch/extracts_creater.go) [path.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/patch/path.go) [po.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/patch/po.go) [replace.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/patch/replace.go)





## <a name="CheckExtracts">func</a> [CheckExtracts](https://git.elewise.com/elma365/common/-/tree/develop/pkg/patch/po.go?s=662:739#L37)
``` go
func CheckExtracts(blob json.RawMessage, extracts Extracts) ([]string, error)
```
CheckExtracts - проверка полноты экстрактов



## <a name="ExtractPO">func</a> [ExtractPO](https://git.elewise.com/elma365/common/-/tree/develop/pkg/patch/po.go?s=1351:1422#L59)
``` go
func ExtractPO(blob json.RawMessage, extracts Extracts) ([]byte, error)
```
ExtractPO извлекает ресурсы в po-файл



## <a name="ReplacePO">func</a> [ReplacePO](https://git.elewise.com/elma365/common/-/tree/develop/pkg/patch/po.go?s=2055:2127#L87)
``` go
func ReplacePO(blob json.RawMessage, po []byte) (json.RawMessage, error)
```
ReplacePO заменяет ресурсы в json-документе




## <a name="Extract">type</a> [Extract](https://git.elewise.com/elma365/common/-/tree/develop/pkg/patch/extract.go?s=164:198#L11)
``` go
type Extract struct {
    // contains filtered or unexported fields
}

```
Extract вытаскивает значения по путям







### <a name="ExtractFromString">func</a> [ExtractFromString](https://git.elewise.com/elma365/common/-/tree/develop/pkg/patch/extract.go?s=1073:1122#L45)
``` go
func ExtractFromString(s string) (Extract, error)
```
ExtractFromString преобразует строку в Extract


### <a name="NewExtract">func</a> [NewExtract](https://git.elewise.com/elma365/common/-/tree/develop/pkg/patch/extract.go?s=241:285#L16)
``` go
func NewExtract(path ...interface{}) Extract
```
NewExtract — конструктор





### <a name="Extract.Apply">func</a> (Extract) [Apply](https://git.elewise.com/elma365/common/-/tree/develop/pkg/patch/extract.go?s=745:812#L35)
``` go
func (extract Extract) Apply(container *gabs.Container) interface{}
```
Apply применить изменение




### <a name="Extract.Prefix">func</a> (Extract) [Prefix](https://git.elewise.com/elma365/common/-/tree/develop/pkg/patch/extract.go?s=576:634#L28)
``` go
func (extract Extract) Prefix(path ...interface{}) Extract
```
Prefix префиксить путь изменения




### <a name="Extract.String">func</a> (Extract) [String](https://git.elewise.com/elma365/common/-/tree/develop/pkg/patch/extract.go?s=931:969#L40)
``` go
func (extract Extract) String() string
```
String возвращает строковое представление




### <a name="Extract.ToReplace">func</a> (Extract) [ToReplace](https://git.elewise.com/elma365/common/-/tree/develop/pkg/patch/extract.go?s=416:473#L23)
``` go
func (extract Extract) ToReplace(val interface{}) Replace
```
ToReplace - создает из извлекателя заменитель




## <a name="Extracter">type</a> [Extracter](https://git.elewise.com/elma365/common/-/tree/develop/pkg/patch/extracts_creater.go?s=243:317#L14)
``` go
type Extracter interface {
    GetExtracts(target string) (Extracts, error)
}
```
Extracter - интерфейс объекта, который сам создает экстракты










## <a name="Extracts">type</a> [Extracts](https://git.elewise.com/elma365/common/-/tree/develop/pkg/patch/extract.go?s=1311:1334#L54)
``` go
type Extracts []Extract
```
Extracts набор изменений JSON-документа







### <a name="CreateExtracts">func</a> [CreateExtracts](https://git.elewise.com/elma365/common/-/tree/develop/pkg/patch/extracts_creater.go?s=557:626#L21)
``` go
func CreateExtracts(val interface{}, target string) (Extracts, error)
```
CreateExtracts - пытается создать экстракты анализируя структуру

Если значение реализует интерфейс patch.Extracter, то он имеет приоритет





### <a name="Extracts.Apply">func</a> (Extracts) [Apply](https://git.elewise.com/elma365/common/-/tree/develop/pkg/patch/extract.go?s=1871:1955#L77)
``` go
func (extracts Extracts) Apply(blob json.RawMessage) (map[string]interface{}, error)
```
Apply применить изменения




### <a name="Extracts.Prefix">func</a> (Extracts) [Prefix](https://git.elewise.com/elma365/common/-/tree/develop/pkg/patch/extract.go?s=1395:1456#L57)
``` go
func (extracts Extracts) Prefix(path ...interface{}) Extracts
```
Prefix префиксить путь изменений




### <a name="Extracts.ToStrings">func</a> (Extracts) [ToStrings](https://git.elewise.com/elma365/common/-/tree/develop/pkg/patch/extract.go?s=1661:1706#L68)
``` go
func (extracts Extracts) ToStrings() []string
```
ToStrings - преобразует в массив строк (путей)




## <a name="Path">type</a> [Path](https://git.elewise.com/elma365/common/-/tree/develop/pkg/patch/path.go?s=1286:1306#L69)
``` go
type Path []PathPart
```
Path путь внутри json-документа







### <a name="FromString">func</a> [FromString](https://git.elewise.com/elma365/common/-/tree/develop/pkg/patch/path.go?s=2103:2142#L93)
``` go
func FromString(s string) (Path, error)
```
FromString преобразует строку в json-путь


### <a name="NewPath">func</a> [NewPath](https://git.elewise.com/elma365/common/-/tree/develop/pkg/patch/path.go?s=1371:1410#L72)
``` go
func NewPath(parts ...interface{}) Path
```
NewPath построение пути по кусочкам





### <a name="Path.Get">func</a> (Path) [Get](https://git.elewise.com/elma365/common/-/tree/develop/pkg/patch/path.go?s=3041:3100#L136)
``` go
func (path Path) Get(container *gabs.Container) interface{}
```
Get value by path




### <a name="Path.Prefix">func</a> (Path) [Prefix](https://git.elewise.com/elma365/common/-/tree/develop/pkg/patch/path.go?s=2921:2971#L131)
``` go
func (path Path) Prefix(parts ...interface{}) Path
```
Prefix path with given parts




### <a name="Path.Search">func</a> (Path) [Search](https://git.elewise.com/elma365/common/-/tree/develop/pkg/patch/path.go?s=3505:3571#L154)
``` go
func (path Path) Search(container *gabs.Container) *gabs.Container
```
Search - поиск вложенного контейнера по пути




### <a name="Path.Set">func</a> (Path) [Set](https://git.elewise.com/elma365/common/-/tree/develop/pkg/patch/path.go?s=3219:3291#L145)
``` go
func (path Path) Set(container *gabs.Container, value interface{}) error
```
Set value by path




### <a name="Path.String">func</a> (Path) [String](https://git.elewise.com/elma365/common/-/tree/develop/pkg/patch/path.go?s=2717:2749#L120)
``` go
func (path Path) String() string
```
String возвращает




## <a name="PathPart">type</a> [PathPart](https://git.elewise.com/elma365/common/-/tree/develop/pkg/patch/path.go?s=182:305#L14)
``` go
type PathPart interface {
    Get(*gabs.Container) *gabs.Container
    Set(*gabs.Container, interface{}) error
    String() string
}
```
PathPart кусочек пути в json-документе







### <a name="NewPart">func</a> [NewPart](https://git.elewise.com/elma365/common/-/tree/develop/pkg/patch/path.go?s=345:381#L21)
``` go
func NewPart(x interface{}) PathPart
```
NewPart — конструктор





## <a name="Replace">type</a> [Replace](https://git.elewise.com/elma365/common/-/tree/develop/pkg/patch/replace.go?s=158:212#L12)
``` go
type Replace struct {
    // contains filtered or unexported fields
}

```
Replace изменение JSON-документа







### <a name="NewReplace">func</a> [NewReplace](https://git.elewise.com/elma365/common/-/tree/develop/pkg/patch/replace.go?s=255:318#L18)
``` go
func NewReplace(value interface{}, path ...interface{}) Replace
```
NewReplace — конструктор





### <a name="Replace.Apply">func</a> (Replace) [Apply](https://git.elewise.com/elma365/common/-/tree/develop/pkg/patch/replace.go?s=642:703#L34)
``` go
func (replace Replace) Apply(container *gabs.Container) error
```
Apply применить изменение




### <a name="Replace.Prefix">func</a> (Replace) [Prefix](https://git.elewise.com/elma365/common/-/tree/develop/pkg/patch/replace.go?s=446:505#L26)
``` go
func (replace Replace) Prefix(parts ...interface{}) Replace
```
Prefix префиксить путь изменения




## <a name="Replaces">type</a> [Replaces](https://git.elewise.com/elma365/common/-/tree/develop/pkg/patch/replace.go?s=826:849#L39)
``` go
type Replaces []Replace
```
Replaces набор изменений JSON-документа










### <a name="Replaces.Apply">func</a> (Replaces) [Apply](https://git.elewise.com/elma365/common/-/tree/develop/pkg/patch/replace.go?s=1145:1222#L53)
``` go
func (replaces Replaces) Apply(blob json.RawMessage) (json.RawMessage, error)
```
Apply применить изменения




### <a name="Replaces.ApplyToContainer">func</a> (Replaces) [ApplyToContainer](https://git.elewise.com/elma365/common/-/tree/develop/pkg/patch/replace.go?s=1541:1615#L68)
``` go
func (replaces Replaces) ApplyToContainer(container *gabs.Container) error
```
ApplyToContainer - применить изменения к контейнеру




### <a name="Replaces.Prefix">func</a> (Replaces) [Prefix](https://git.elewise.com/elma365/common/-/tree/develop/pkg/patch/replace.go?s=910:971#L42)
``` go
func (replaces Replaces) Prefix(path ...interface{}) Replaces
```
Prefix префиксить путь изменений









## <a name="pkg-subdirectories">Subdirectories</a>

| Name | Synopsis |
| ---- | -------- |
| [..](..) | |
| [patch_test](patch_test/) |  |


- - -
Generated by [godoc2md](https://github.com/Exa-Networks/godoc2md)
