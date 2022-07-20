# uuids
`import "git.elewise.com/elma365/common/pkg/types/uuids"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [type UUIDS](#UUIDS)
  * [func FromStrings(list []string) (UUIDS, error)](#FromStrings)
  * [func (us UUIDS) AsStrings() []string](#UUIDS.AsStrings)
  * [func (us UUIDS) Contains(u uuid.UUID) bool](#UUIDS.Contains)
  * [func (us UUIDS) Diff(x UUIDS) (plus, minus UUIDS)](#UUIDS.Diff)
  * [func (us UUIDS) Equal(x UUIDS) bool](#UUIDS.Equal)
  * [func (us UUIDS) Len() int](#UUIDS.Len)
  * [func (us UUIDS) Less(i, j int) bool](#UUIDS.Less)
  * [func (v UUIDS) MarshalEasyJSON(w *jwriter.Writer)](#UUIDS.MarshalEasyJSON)
  * [func (v UUIDS) MarshalJSON() ([]byte, error)](#UUIDS.MarshalJSON)
  * [func (us *UUIDS) Remove(u uuid.UUID)](#UUIDS.Remove)
  * [func (us *UUIDS) RemoveDuplicates()](#UUIDS.RemoveDuplicates)
  * [func (us *UUIDS) Scan(pSrc interface{}) error](#UUIDS.Scan)
  * [func (us UUIDS) Swap(i, j int)](#UUIDS.Swap)
  * [func (v *UUIDS) UnmarshalEasyJSON(l *jlexer.Lexer)](#UUIDS.UnmarshalEasyJSON)
  * [func (v *UUIDS) UnmarshalJSON(data []byte) error](#UUIDS.UnmarshalJSON)
  * [func (us UUIDS) Value() (driver.Value, error)](#UUIDS.Value)


#### <a name="pkg-files">Package files</a>
[uuids.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/uuids/uuids.go) [uuids_easyjson.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/uuids/uuids_easyjson.go)






## <a name="UUIDS">type</a> [UUIDS](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/uuids/uuids.go?s=310:332#L17)
``` go
type UUIDS []uuid.UUID
```
UUIDS определяет тип - массив uuid-ов и стандартные операции над ними







### <a name="FromStrings">func</a> [FromStrings](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/uuids/uuids.go?s=381:427#L20)
``` go
func FromStrings(list []string) (UUIDS, error)
```
FromStrings creates UUIDS from string slice





### <a name="UUIDS.AsStrings">func</a> (UUIDS) [AsStrings](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/uuids/uuids.go?s=2045:2081#L87)
``` go
func (us UUIDS) AsStrings() []string
```
AsStrings returns slice of UUIDs as slice of strings




### <a name="UUIDS.Contains">func</a> (UUIDS) [Contains](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/uuids/uuids.go?s=2550:2592#L114)
``` go
func (us UUIDS) Contains(u uuid.UUID) bool
```
Contains check array of UUIDS contains specified UUID




### <a name="UUIDS.Diff">func</a> (UUIDS) [Diff](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/uuids/uuids.go?s=3366:3415#L157)
``` go
func (us UUIDS) Diff(x UUIDS) (plus, minus UUIDS)
```
Diff compare UUID array with new UUID array

Return UUIDS that were added as first value
and deleted UUIDS as second value




### <a name="UUIDS.Equal">func</a> (UUIDS) [Equal](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/uuids/uuids.go?s=3030:3065#L137)
``` go
func (us UUIDS) Equal(x UUIDS) bool
```
Equal compare UUID array with new UUID array

Return true if equals




### <a name="UUIDS.Len">func</a> (UUIDS) [Len](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/uuids/uuids.go?s=3927:3952#L195)
``` go
func (us UUIDS) Len() int
```
Len implements sort.Interface




### <a name="UUIDS.Less">func</a> (UUIDS) [Less](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/uuids/uuids.go?s=4105:4140#L201)
``` go
func (us UUIDS) Less(i, j int) bool
```
Less implements sort.Interface




### <a name="UUIDS.MarshalEasyJSON">func</a> (UUIDS) [MarshalEasyJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/uuids/uuids_easyjson.go?s=1627:1676#L74)
``` go
func (v UUIDS) MarshalEasyJSON(w *jwriter.Writer)
```
MarshalEasyJSON supports easyjson.Marshaler interface




### <a name="UUIDS.MarshalJSON">func</a> (UUIDS) [MarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/uuids/uuids_easyjson.go?s=1388:1432#L67)
``` go
func (v UUIDS) MarshalJSON() ([]byte, error)
```
MarshalJSON supports json.Marshaler interface




### <a name="UUIDS.Remove">func</a> (\*UUIDS) [Remove](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/uuids/uuids.go?s=2790:2826#L125)
``` go
func (us *UUIDS) Remove(u uuid.UUID)
```
Remove удаляет первый найденный заданный элемент из списка




### <a name="UUIDS.RemoveDuplicates">func</a> (\*UUIDS) [RemoveDuplicates](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/uuids/uuids.go?s=2250:2285#L98)
``` go
func (us *UUIDS) RemoveDuplicates()
```
RemoveDuplicates remove duplicate values from array




### <a name="UUIDS.Scan">func</a> (\*UUIDS) [Scan](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/uuids/uuids.go?s=679:724#L34)
``` go
func (us *UUIDS) Scan(pSrc interface{}) error
```
Scan implements sql.Scanner interface




### <a name="UUIDS.Swap">func</a> (UUIDS) [Swap](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/uuids/uuids.go?s=4007:4037#L198)
``` go
func (us UUIDS) Swap(i, j int)
```
Swap implements sort.Interface




### <a name="UUIDS.UnmarshalEasyJSON">func</a> (\*UUIDS) [UnmarshalEasyJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/uuids/uuids_easyjson.go?s=2038:2088#L86)
``` go
func (v *UUIDS) UnmarshalEasyJSON(l *jlexer.Lexer)
```
UnmarshalEasyJSON supports easyjson.Unmarshaler interface




### <a name="UUIDS.UnmarshalJSON">func</a> (\*UUIDS) [UnmarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/uuids/uuids_easyjson.go?s=1804:1852#L79)
``` go
func (v *UUIDS) UnmarshalJSON(data []byte) error
```
UnmarshalJSON supports json.Unmarshaler interface




### <a name="UUIDS.Value">func</a> (UUIDS) [Value](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/uuids/uuids.go?s=1670:1715#L72)
``` go
func (us UUIDS) Value() (driver.Value, error)
```
Value implements sql.Valuer interface







- - -
Generated by [godoc2md](https://github.com/Exa-Networks/godoc2md)
