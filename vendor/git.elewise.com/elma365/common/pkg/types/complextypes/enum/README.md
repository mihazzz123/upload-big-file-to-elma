# enum
`import "git.elewise.com/elma365/common/pkg/types/complextypes/enum"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [type Enum](#Enum)
  * [func (v Enum) MarshalEasyJSON(w *jwriter.Writer)](#Enum.MarshalEasyJSON)
  * [func (v Enum) MarshalJSON() ([]byte, error)](#Enum.MarshalJSON)
  * [func (v *Enum) UnmarshalEasyJSON(l *jlexer.Lexer)](#Enum.UnmarshalEasyJSON)
  * [func (v *Enum) UnmarshalJSON(data []byte) error](#Enum.UnmarshalJSON)
* [type EnumData](#EnumData)
  * [func (e EnumData) GetItemByCode(code string) (EnumItem, error)](#EnumData.GetItemByCode)
  * [func (v EnumData) MarshalEasyJSON(w *jwriter.Writer)](#EnumData.MarshalEasyJSON)
  * [func (v EnumData) MarshalJSON() ([]byte, error)](#EnumData.MarshalJSON)
  * [func (v *EnumData) UnmarshalEasyJSON(l *jlexer.Lexer)](#EnumData.UnmarshalEasyJSON)
  * [func (v *EnumData) UnmarshalJSON(data []byte) error](#EnumData.UnmarshalJSON)
* [type EnumItem](#EnumItem)
  * [func (v EnumItem) MarshalEasyJSON(w *jwriter.Writer)](#EnumItem.MarshalEasyJSON)
  * [func (v EnumItem) MarshalJSON() ([]byte, error)](#EnumItem.MarshalJSON)
  * [func (v *EnumItem) UnmarshalEasyJSON(l *jlexer.Lexer)](#EnumItem.UnmarshalEasyJSON)
  * [func (v *EnumItem) UnmarshalJSON(data []byte) error](#EnumItem.UnmarshalJSON)


#### <a name="pkg-files">Package files</a>
[data.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/enum/data.go) [data_easyjson.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/enum/data_easyjson.go) [enum.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/enum/enum.go) [enum_easyjson.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/enum/enum_easyjson.go)






## <a name="Enum">type</a> [Enum](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/enum/enum.go?s=138:212#L8)
``` go
type Enum struct {
    Code string `json:"code"`
    Name string `json:"name"`
}

```
Enum - значение для типа Enum

easyjson:json










### <a name="Enum.MarshalEasyJSON">func</a> (Enum) [MarshalEasyJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/enum/enum_easyjson.go?s=1641:1689#L78)
``` go
func (v Enum) MarshalEasyJSON(w *jwriter.Writer)
```
MarshalEasyJSON supports easyjson.Marshaler interface




### <a name="Enum.MarshalJSON">func</a> (Enum) [MarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/enum/enum_easyjson.go?s=1392:1435#L71)
``` go
func (v Enum) MarshalJSON() ([]byte, error)
```
MarshalJSON supports json.Marshaler interface




### <a name="Enum.UnmarshalEasyJSON">func</a> (\*Enum) [UnmarshalEasyJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/enum/enum_easyjson.go?s=2072:2121#L90)
``` go
func (v *Enum) UnmarshalEasyJSON(l *jlexer.Lexer)
```
UnmarshalEasyJSON supports easyjson.Unmarshaler interface




### <a name="Enum.UnmarshalJSON">func</a> (\*Enum) [UnmarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/enum/enum_easyjson.go?s=1828:1875#L83)
``` go
func (v *Enum) UnmarshalJSON(data []byte) error
```
UnmarshalJSON supports json.Unmarshaler interface




## <a name="EnumData">type</a> [EnumData](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/enum/data.go?s=235:311#L15)
``` go
type EnumData struct {
    EnumItems []EnumItem `json:"enumItems" patch:"po"`
}

```
EnumData - структура поля Data

easyjson:json
nolint:golint










### <a name="EnumData.GetItemByCode">func</a> (EnumData) [GetItemByCode](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/enum/data.go?s=376:438#L20)
``` go
func (e EnumData) GetItemByCode(code string) (EnumItem, error)
```
GetItemByCode получить вариант по коду




### <a name="EnumData.MarshalEasyJSON">func</a> (EnumData) [MarshalEasyJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/enum/data_easyjson.go?s=4214:4266#L183)
``` go
func (v EnumData) MarshalEasyJSON(w *jwriter.Writer)
```
MarshalEasyJSON supports easyjson.Marshaler interface




### <a name="EnumData.MarshalJSON">func</a> (EnumData) [MarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/enum/data_easyjson.go?s=3960:4007#L176)
``` go
func (v EnumData) MarshalJSON() ([]byte, error)
```
MarshalJSON supports json.Marshaler interface




### <a name="EnumData.UnmarshalEasyJSON">func</a> (\*EnumData) [UnmarshalEasyJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/enum/data_easyjson.go?s=4655:4708#L195)
``` go
func (v *EnumData) UnmarshalEasyJSON(l *jlexer.Lexer)
```
UnmarshalEasyJSON supports easyjson.Unmarshaler interface




### <a name="EnumData.UnmarshalJSON">func</a> (\*EnumData) [UnmarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/enum/data_easyjson.go?s=4406:4457#L188)
``` go
func (v *EnumData) UnmarshalJSON(data []byte) error
```
UnmarshalJSON supports json.Unmarshaler interface




## <a name="EnumItem">type</a> [EnumItem](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/enum/data.go?s=690:818#L33)
``` go
type EnumItem struct {
    Code    string `json:"code"`
    Name    string `json:"name" patch:"po"`
    Checked bool   `json:"checked"`
}

```
EnumItem - вариант значения поля

easyjson:json
nolint:golint










### <a name="EnumItem.MarshalEasyJSON">func</a> (EnumItem) [MarshalEasyJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/enum/data_easyjson.go?s=1803:1855#L85)
``` go
func (v EnumItem) MarshalEasyJSON(w *jwriter.Writer)
```
MarshalEasyJSON supports easyjson.Marshaler interface




### <a name="EnumItem.MarshalJSON">func</a> (EnumItem) [MarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/enum/data_easyjson.go?s=1550:1597#L78)
``` go
func (v EnumItem) MarshalJSON() ([]byte, error)
```
MarshalJSON supports json.Marshaler interface




### <a name="EnumItem.UnmarshalEasyJSON">func</a> (\*EnumItem) [UnmarshalEasyJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/enum/data_easyjson.go?s=2242:2295#L97)
``` go
func (v *EnumItem) UnmarshalEasyJSON(l *jlexer.Lexer)
```
UnmarshalEasyJSON supports easyjson.Unmarshaler interface




### <a name="EnumItem.UnmarshalJSON">func</a> (\*EnumItem) [UnmarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/enum/data_easyjson.go?s=1994:2045#L90)
``` go
func (v *EnumItem) UnmarshalJSON(data []byte) error
```
UnmarshalJSON supports json.Unmarshaler interface







- - -
Generated by [godoc2md](https://github.com/Exa-Networks/godoc2md)
