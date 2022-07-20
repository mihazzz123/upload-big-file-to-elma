# money
`import "git.elewise.com/elma365/common/pkg/types/complextypes/money"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [type Money](#Money)
  * [func (v Money) MarshalEasyJSON(w *jwriter.Writer)](#Money.MarshalEasyJSON)
  * [func (v Money) MarshalJSON() ([]byte, error)](#Money.MarshalJSON)
  * [func (m Money) String() string](#Money.String)
  * [func (v *Money) UnmarshalEasyJSON(l *jlexer.Lexer)](#Money.UnmarshalEasyJSON)
  * [func (v *Money) UnmarshalJSON(data []byte) error](#Money.UnmarshalJSON)
* [type ViewData](#ViewData)
  * [func (v ViewData) MarshalEasyJSON(w *jwriter.Writer)](#ViewData.MarshalEasyJSON)
  * [func (v ViewData) MarshalJSON() ([]byte, error)](#ViewData.MarshalJSON)
  * [func (v *ViewData) UnmarshalEasyJSON(l *jlexer.Lexer)](#ViewData.UnmarshalEasyJSON)
  * [func (v *ViewData) UnmarshalJSON(data []byte) error](#ViewData.UnmarshalJSON)


#### <a name="pkg-files">Package files</a>
[money.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/money/money.go) [money_easyjson.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/money/money_easyjson.go) [view_data.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/money/view_data.go) [view_data_easyjson.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/money/view_data_easyjson.go)






## <a name="Money">type</a> [Money](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/money/money.go?s=231:327#L15)
``` go
type Money struct {
    C currency.Currency `json:"currency"`
    V int64             `json:"cents"`
}

```
Money is currency cents pair

easyjson:json










### <a name="Money.MarshalEasyJSON">func</a> (Money) [MarshalEasyJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/money/money_easyjson.go?s=1700:1749#L80)
``` go
func (v Money) MarshalEasyJSON(w *jwriter.Writer)
```
MarshalEasyJSON supports easyjson.Marshaler interface




### <a name="Money.MarshalJSON">func</a> (Money) [MarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/money/money_easyjson.go?s=1449:1493#L73)
``` go
func (v Money) MarshalJSON() ([]byte, error)
```
MarshalJSON supports json.Marshaler interface




### <a name="Money.String">func</a> (Money) [String](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/money/money.go?s=450:480#L21)
``` go
func (m Money) String() string
```
String выводит результат вычисления валюты, на базе разряда копеек




### <a name="Money.UnmarshalEasyJSON">func</a> (\*Money) [UnmarshalEasyJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/money/money_easyjson.go?s=2135:2185#L92)
``` go
func (v *Money) UnmarshalEasyJSON(l *jlexer.Lexer)
```
UnmarshalEasyJSON supports easyjson.Unmarshaler interface




### <a name="Money.UnmarshalJSON">func</a> (\*Money) [UnmarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/money/money_easyjson.go?s=1889:1937#L85)
``` go
func (v *Money) UnmarshalJSON(data []byte) error
```
UnmarshalJSON supports json.Unmarshaler interface




## <a name="ViewData">type</a> [ViewData](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/money/view_data.go?s=168:238#L8)
``` go
type ViewData struct {
    Currency string `json:"currency" patch:"po"`
}

```
ViewData информация для отображения поля

easyjson:json










### <a name="ViewData.MarshalEasyJSON">func</a> (ViewData) [MarshalEasyJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/money/view_data_easyjson.go?s=1527:1579#L71)
``` go
func (v ViewData) MarshalEasyJSON(w *jwriter.Writer)
```
MarshalEasyJSON supports easyjson.Marshaler interface




### <a name="ViewData.MarshalJSON">func</a> (ViewData) [MarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/money/view_data_easyjson.go?s=1273:1320#L64)
``` go
func (v ViewData) MarshalJSON() ([]byte, error)
```
MarshalJSON supports json.Marshaler interface




### <a name="ViewData.UnmarshalEasyJSON">func</a> (\*ViewData) [UnmarshalEasyJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/money/view_data_easyjson.go?s=1968:2021#L83)
``` go
func (v *ViewData) UnmarshalEasyJSON(l *jlexer.Lexer)
```
UnmarshalEasyJSON supports easyjson.Unmarshaler interface




### <a name="ViewData.UnmarshalJSON">func</a> (\*ViewData) [UnmarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/money/view_data_easyjson.go?s=1719:1770#L76)
``` go
func (v *ViewData) UnmarshalJSON(data []byte) error
```
UnmarshalJSON supports json.Unmarshaler interface









## <a name="pkg-subdirectories">Subdirectories</a>

| Name | Synopsis |
| ---- | -------- |
| [..](..) | |
| [currency](currency/) |  |


- - -
Generated by [godoc2md](https://github.com/Exa-Networks/godoc2md)
