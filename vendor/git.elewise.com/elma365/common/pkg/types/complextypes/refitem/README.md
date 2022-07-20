# refitem
`import "git.elewise.com/elma365/common/pkg/types/complextypes/refitem"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [type RefItem](#RefItem)
  * [func (v RefItem) MarshalEasyJSON(w *jwriter.Writer)](#RefItem.MarshalEasyJSON)
  * [func (v RefItem) MarshalJSON() ([]byte, error)](#RefItem.MarshalJSON)
  * [func (ri RefItem) String() string](#RefItem.String)
  * [func (v *RefItem) UnmarshalEasyJSON(l *jlexer.Lexer)](#RefItem.UnmarshalEasyJSON)
  * [func (v *RefItem) UnmarshalJSON(data []byte) error](#RefItem.UnmarshalJSON)


#### <a name="pkg-files">Package files</a>
[refitem.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/refitem/refitem.go) [refitem_easyjson.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/refitem/refitem_easyjson.go)






## <a name="RefItem">type</a> [RefItem](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/refitem/refitem.go?s=291:655#L16)
``` go
type RefItem struct {
    Namespace namespace.Namespace `json:"namespace" validate:"required"` // Раздел справочника элемента
    Code      string              `json:"code"      validate:"required"` // Код справочника элемента
    ID        uuid.UUID           `json:"id"        validate:"required"` // ИД элемента
}

```
RefItem ссылка на элемент произвольного приложения

easyjson:json










### <a name="RefItem.MarshalEasyJSON">func</a> (RefItem) [MarshalEasyJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/refitem/refitem_easyjson.go?s=1961:2012#L88)
``` go
func (v RefItem) MarshalEasyJSON(w *jwriter.Writer)
```
MarshalEasyJSON supports easyjson.Marshaler interface




### <a name="RefItem.MarshalJSON">func</a> (RefItem) [MarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/refitem/refitem_easyjson.go?s=1706:1752#L81)
``` go
func (v RefItem) MarshalJSON() ([]byte, error)
```
MarshalJSON supports json.Marshaler interface




### <a name="RefItem.String">func</a> (RefItem) [String](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/refitem/refitem.go?s=657:690#L22)
``` go
func (ri RefItem) String() string
```



### <a name="RefItem.UnmarshalEasyJSON">func</a> (\*RefItem) [UnmarshalEasyJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/refitem/refitem_easyjson.go?s=2404:2456#L100)
``` go
func (v *RefItem) UnmarshalEasyJSON(l *jlexer.Lexer)
```
UnmarshalEasyJSON supports easyjson.Unmarshaler interface




### <a name="RefItem.UnmarshalJSON">func</a> (\*RefItem) [UnmarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/refitem/refitem_easyjson.go?s=2154:2204#L93)
``` go
func (v *RefItem) UnmarshalJSON(data []byte) error
```
UnmarshalJSON supports json.Unmarshaler interface







- - -
Generated by [godoc2md](https://github.com/Exa-Networks/godoc2md)