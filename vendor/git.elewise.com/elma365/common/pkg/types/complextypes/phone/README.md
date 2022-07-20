# phone
`import "git.elewise.com/elma365/common/pkg/types/complextypes/phone"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [type Phone](#Phone)
  * [func (v Phone) MarshalEasyJSON(w *jwriter.Writer)](#Phone.MarshalEasyJSON)
  * [func (v Phone) MarshalJSON() ([]byte, error)](#Phone.MarshalJSON)
  * [func (p *Phone) Scan(pSrc interface{}) error](#Phone.Scan)
  * [func (p *Phone) String() string](#Phone.String)
  * [func (v *Phone) UnmarshalEasyJSON(l *jlexer.Lexer)](#Phone.UnmarshalEasyJSON)
  * [func (v *Phone) UnmarshalJSON(data []byte) error](#Phone.UnmarshalJSON)
  * [func (p Phone) Value() (value driver.Value, err error)](#Phone.Value)
* [type Phones](#Phones)
  * [func (v Phones) MarshalEasyJSON(w *jwriter.Writer)](#Phones.MarshalEasyJSON)
  * [func (v Phones) MarshalJSON() ([]byte, error)](#Phones.MarshalJSON)
  * [func (ps *Phones) Scan(pSrc interface{}) error](#Phones.Scan)
  * [func (v *Phones) UnmarshalEasyJSON(l *jlexer.Lexer)](#Phones.UnmarshalEasyJSON)
  * [func (v *Phones) UnmarshalJSON(data []byte) error](#Phones.UnmarshalJSON)
  * [func (ps Phones) Value() (value driver.Value, err error)](#Phones.Value)
* [type Type](#Type)
* [type ViewData](#ViewData)


#### <a name="pkg-files">Package files</a>
[phone.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/phone/phone.go) [phone_easyjson.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/phone/phone_easyjson.go) [view_data.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/phone/view_data.go)






## <a name="Phone">type</a> [Phone](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/phone/phone.go?s=258:838#L15)
``` go
type Phone struct {
    Tel     string `json:"tel"`               // локальная часть, действует в пределах страны
    Context string `json:"context,omitempty"` // Глобальная часть, например код страны или домен (для SIP) (в RFC называется phone-context)
    Ext     string `json:"ext,omitempty"`     // добавочный номер
    Type    string `json:"type"`              // тип номера, это уже не RFC а наша приблуда - Домашний/мобильный/etc
}

```
Phone is a tuple of phone information

See <a href="https://www.ietf.org/rfc/rfc3966.txt">https://www.ietf.org/rfc/rfc3966.txt</a> for details of format
easyjson:json










### <a name="Phone.MarshalEasyJSON">func</a> (Phone) [MarshalEasyJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/phone/phone_easyjson.go?s=3743:3792#L158)
``` go
func (v Phone) MarshalEasyJSON(w *jwriter.Writer)
```
MarshalEasyJSON supports easyjson.Marshaler interface




### <a name="Phone.MarshalJSON">func</a> (Phone) [MarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/phone/phone_easyjson.go?s=3491:3535#L151)
``` go
func (v Phone) MarshalJSON() ([]byte, error)
```
MarshalJSON supports json.Marshaler interface




### <a name="Phone.Scan">func</a> (\*Phone) [Scan](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/phone/phone.go?s=881:925#L23)
``` go
func (p *Phone) Scan(pSrc interface{}) error
```
Scan implements sql.Scanner interface




### <a name="Phone.String">func</a> (\*Phone) [String](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/phone/phone.go?s=1369:1400#L49)
``` go
func (p *Phone) String() string
```
String implements fmt.Stringer interface




### <a name="Phone.UnmarshalEasyJSON">func</a> (\*Phone) [UnmarshalEasyJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/phone/phone_easyjson.go?s=4180:4230#L170)
``` go
func (v *Phone) UnmarshalEasyJSON(l *jlexer.Lexer)
```
UnmarshalEasyJSON supports easyjson.Unmarshaler interface




### <a name="Phone.UnmarshalJSON">func</a> (\*Phone) [UnmarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/phone/phone_easyjson.go?s=3933:3981#L163)
``` go
func (v *Phone) UnmarshalJSON(data []byte) error
```
UnmarshalJSON supports json.Unmarshaler interface




### <a name="Phone.Value">func</a> (Phone) [Value](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/phone/phone.go?s=1241:1295#L44)
``` go
func (p Phone) Value() (value driver.Value, err error)
```
Value implements sql.Valuer interface




## <a name="Phones">type</a> [Phones](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/phone/phone.go?s=1484:1503#L55)
``` go
type Phones []Phone
```
Phones - array of phone structs
easyjson:json










### <a name="Phones.MarshalEasyJSON">func</a> (Phones) [MarshalEasyJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/phone/phone_easyjson.go?s=1559:1609#L71)
``` go
func (v Phones) MarshalEasyJSON(w *jwriter.Writer)
```
MarshalEasyJSON supports easyjson.Marshaler interface




### <a name="Phones.MarshalJSON">func</a> (Phones) [MarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/phone/phone_easyjson.go?s=1307:1352#L64)
``` go
func (v Phones) MarshalJSON() ([]byte, error)
```
MarshalJSON supports json.Marshaler interface




### <a name="Phones.Scan">func</a> (\*Phones) [Scan](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/phone/phone.go?s=1546:1592#L58)
``` go
func (ps *Phones) Scan(pSrc interface{}) error
```
Scan implements sql.Scanner interface




### <a name="Phones.UnmarshalEasyJSON">func</a> (\*Phones) [UnmarshalEasyJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/phone/phone_easyjson.go?s=1996:2047#L83)
``` go
func (v *Phones) UnmarshalEasyJSON(l *jlexer.Lexer)
```
UnmarshalEasyJSON supports easyjson.Unmarshaler interface




### <a name="Phones.UnmarshalJSON">func</a> (\*Phones) [UnmarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/phone/phone_easyjson.go?s=1749:1798#L76)
``` go
func (v *Phones) UnmarshalJSON(data []byte) error
```
UnmarshalJSON supports json.Unmarshaler interface




### <a name="Phones.Value">func</a> (Phones) [Value](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/phone/phone.go?s=1912:1968#L79)
``` go
func (ps Phones) Value() (value driver.Value, err error)
```
Value implements sql.Valuer interface




## <a name="Type">type</a> [Type](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/phone/view_data.go?s=47:63#L4)
``` go
type Type string
```
Type тип телефона


``` go
const (
    TypeHome    Type = "home"
    TypeWork    Type = "work"
    TypeMobile  Type = "mobile"
    TypeMain    Type = "main"
    TypeHomeFax Type = "home-fax"
    TypeWorkFax Type = "work-fax"
    TypePager   Type = "page"
)
```
Типы телефонов










## <a name="ViewData">type</a> [ViewData](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/phone/view_data.go?s=371:430#L18)
``` go
type ViewData struct {
    Type Type `json:"type,omitempty"`
}

```
ViewData описание представления типа













- - -
Generated by [godoc2md](https://github.com/Exa-Networks/godoc2md)
