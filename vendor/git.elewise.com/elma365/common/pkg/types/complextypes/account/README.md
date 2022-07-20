# account
`import "git.elewise.com/elma365/common/pkg/types/complextypes/account"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [type Account](#Account)
  * [func (v Account) MarshalEasyJSON(w *jwriter.Writer)](#Account.MarshalEasyJSON)
  * [func (v Account) MarshalJSON() ([]byte, error)](#Account.MarshalJSON)
  * [func (a *Account) Scan(pSrc interface{}) error](#Account.Scan)
  * [func (a *Account) String() string](#Account.String)
  * [func (v *Account) UnmarshalEasyJSON(l *jlexer.Lexer)](#Account.UnmarshalEasyJSON)
  * [func (v *Account) UnmarshalJSON(data []byte) error](#Account.UnmarshalJSON)
  * [func (a Account) Value() (value driver.Value, err error)](#Account.Value)
* [type Accounts](#Accounts)
  * [func (v Accounts) MarshalEasyJSON(w *jwriter.Writer)](#Accounts.MarshalEasyJSON)
  * [func (v Accounts) MarshalJSON() ([]byte, error)](#Accounts.MarshalJSON)
  * [func (as *Accounts) Scan(pSrc interface{}) error](#Accounts.Scan)
  * [func (v *Accounts) UnmarshalEasyJSON(l *jlexer.Lexer)](#Accounts.UnmarshalEasyJSON)
  * [func (v *Accounts) UnmarshalJSON(data []byte) error](#Accounts.UnmarshalJSON)
  * [func (as Accounts) Value() (value driver.Value, err error)](#Accounts.Value)
* [type Type](#Type)
* [type ViewData](#ViewData)


#### <a name="pkg-files">Package files</a>
[account.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/account/account.go) [account_easyjson.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/account/account_easyjson.go) [view_data.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/account/view_data.go)






## <a name="Account">type</a> [Account](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/account/account.go?s=200:280#L14)
``` go
type Account struct {
    Login string `json:"login"`
    Type  string `json:"type"`
}

```
Account is a tuple of account information

easyjson:json










### <a name="Account.MarshalEasyJSON">func</a> (Account) [MarshalEasyJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/account/account_easyjson.go?s=3459:3510#L144)
``` go
func (v Account) MarshalEasyJSON(w *jwriter.Writer)
```
MarshalEasyJSON supports easyjson.Marshaler interface




### <a name="Account.MarshalJSON">func</a> (Account) [MarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/account/account_easyjson.go?s=3203:3249#L137)
``` go
func (v Account) MarshalJSON() ([]byte, error)
```
MarshalJSON supports json.Marshaler interface




### <a name="Account.Scan">func</a> (\*Account) [Scan](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/account/account.go?s=323:369#L20)
``` go
func (a *Account) Scan(pSrc interface{}) error
```
Scan implements sql.Scanner interface




### <a name="Account.String">func</a> (\*Account) [String](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/account/account.go?s=821:854#L46)
``` go
func (a *Account) String() string
```
String implements fmt.Stringer interface




### <a name="Account.UnmarshalEasyJSON">func</a> (\*Account) [UnmarshalEasyJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/account/account_easyjson.go?s=3904:3956#L156)
``` go
func (v *Account) UnmarshalEasyJSON(l *jlexer.Lexer)
```
UnmarshalEasyJSON supports easyjson.Unmarshaler interface




### <a name="Account.UnmarshalJSON">func</a> (\*Account) [UnmarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/account/account_easyjson.go?s=3653:3703#L149)
``` go
func (v *Account) UnmarshalJSON(data []byte) error
```
UnmarshalJSON supports json.Unmarshaler interface




### <a name="Account.Value">func</a> (Account) [Value](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/account/account.go?s=691:747#L41)
``` go
func (a Account) Value() (value driver.Value, err error)
```
Value implements sql.Valuer interface




## <a name="Accounts">type</a> [Accounts](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/account/account.go?s=933:956#L52)
``` go
type Accounts []Account
```
Accounts - array of accounts structs
easyjson:json










### <a name="Accounts.MarshalEasyJSON">func</a> (Accounts) [MarshalEasyJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/account/account_easyjson.go?s=1579:1631#L71)
``` go
func (v Accounts) MarshalEasyJSON(w *jwriter.Writer)
```
MarshalEasyJSON supports easyjson.Marshaler interface




### <a name="Accounts.MarshalJSON">func</a> (Accounts) [MarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/account/account_easyjson.go?s=1323:1370#L64)
``` go
func (v Accounts) MarshalJSON() ([]byte, error)
```
MarshalJSON supports json.Marshaler interface




### <a name="Accounts.Scan">func</a> (\*Accounts) [Scan](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/account/account.go?s=999:1047#L55)
``` go
func (as *Accounts) Scan(pSrc interface{}) error
```
Scan implements sql.Scanner interface




### <a name="Accounts.UnmarshalEasyJSON">func</a> (\*Accounts) [UnmarshalEasyJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/account/account_easyjson.go?s=2024:2077#L83)
``` go
func (v *Accounts) UnmarshalEasyJSON(l *jlexer.Lexer)
```
UnmarshalEasyJSON supports easyjson.Unmarshaler interface




### <a name="Accounts.UnmarshalJSON">func</a> (\*Accounts) [UnmarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/account/account_easyjson.go?s=1773:1824#L76)
``` go
func (v *Accounts) UnmarshalJSON(data []byte) error
```
UnmarshalJSON supports json.Unmarshaler interface




### <a name="Accounts.Value">func</a> (Accounts) [Value](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/account/account.go?s=1373:1431#L76)
``` go
func (as Accounts) Value() (value driver.Value, err error)
```
Value implements sql.Valuer interface




## <a name="Type">type</a> [Type](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/account/view_data.go?s=55:71#L4)
``` go
type Type string
```
Type тип мессенджера


``` go
const (
    TypeTelegram       Type = "telegram"
    TypeWhatsapp       Type = "whatsapp"
    TypeInstagram      Type = "instagram"
    TypeViber          Type = "viber"
    TypeFacebook       Type = "facebook"
    TypeVkontakte      Type = "vkontakte"
    TypeSkype          Type = "skype"
    TypeGoogleTalk     Type = "google-talk"
    TypeGoogleHangouts Type = "google-hangouts"
    TypeWeChat         Type = "we-chat"
    TypeSnapchat       Type = "snapchat"
    TypeLine           Type = "line"
)
```
Типы мессенджеров










## <a name="ViewData">type</a> [ViewData](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/account/view_data.go?s=660:719#L23)
``` go
type ViewData struct {
    Type Type `json:"type,omitempty"`
}

```
ViewData описание представления поля типа Account













- - -
Generated by [godoc2md](https://github.com/Exa-Networks/godoc2md)
