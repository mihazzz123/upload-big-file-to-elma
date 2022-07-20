# validation
`import "git.elewise.com/elma365/common/pkg/errs/validation"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [type Error](#Error)
  * [func NewError(desc string, options ...Option) Error](#NewError)
  * [func (ve Error) Cause() error](#Error.Cause)
  * [func (ve Error) Error() string](#Error.Error)
  * [func (v Error) MarshalEasyJSON(w *jwriter.Writer)](#Error.MarshalEasyJSON)
  * [func (v Error) MarshalJSON() ([]byte, error)](#Error.MarshalJSON)
  * [func (ve Error) Prefix(parts ...string) Error](#Error.Prefix)
  * [func (v *Error) UnmarshalEasyJSON(l *jlexer.Lexer)](#Error.UnmarshalEasyJSON)
  * [func (v *Error) UnmarshalJSON(data []byte) error](#Error.UnmarshalJSON)
* [type Errors](#Errors)
  * [func NewErrors(n int) *Errors](#NewErrors)
  * [func ValidateStruct(ctx context.Context, val interface{}) *Errors](#ValidateStruct)
  * [func (verrs *Errors) Add(err error) *Errors](#Errors.Add)
  * [func (verrs *Errors) AddItem(desc string, options ...Option) *Errors](#Errors.AddItem)
  * [func (verrs *Errors) AddWithPrefix(err error, prefix string) *Errors](#Errors.AddWithPrefix)
  * [func (verrs Errors) Cause() error](#Errors.Cause)
  * [func (verrs Errors) Error() string](#Errors.Error)
  * [func (verrs Errors) HasLevel(level Level) bool](#Errors.HasLevel)
  * [func (verrs Errors) IsCritical() bool](#Errors.IsCritical)
  * [func (verrs Errors) IsEmpty() bool](#Errors.IsEmpty)
  * [func (verrs *Errors) Iter() *ErrorsIterator](#Errors.Iter)
  * [func (verrs Errors) Len() int](#Errors.Len)
  * [func (verrs Errors) Prefix(prefix string) *Errors](#Errors.Prefix)
  * [func (verrs *Errors) Return() error](#Errors.Return)
  * [func (verrs *Errors) WithRecover(fn func()) *Errors](#Errors.WithRecover)
  * [func (verrs Errors) WriteData(w io.Writer) error](#Errors.WriteData)
* [type ErrorsIterator](#ErrorsIterator)
  * [func (verrsi *ErrorsIterator) Next() bool](#ErrorsIterator.Next)
  * [func (verrsi *ErrorsIterator) Value() Error](#ErrorsIterator.Value)
* [type Level](#Level)
  * [func LevelString(s string) (Level, error)](#LevelString)
  * [func LevelValues() []Level](#LevelValues)
  * [func (i Level) IsALevel() bool](#Level.IsALevel)
  * [func (i Level) MarshalJSON() ([]byte, error)](#Level.MarshalJSON)
  * [func (i Level) String() string](#Level.String)
  * [func (i *Level) UnmarshalJSON(data []byte) error](#Level.UnmarshalJSON)
* [type Option](#Option)
  * [func WithArg(key string, val interface{}) Option](#WithArg)
  * [func WithCause(err error) Option](#WithCause)
  * [func WithLevel(level Level) Option](#WithLevel)
  * [func WithPath(chunks ...interface{}) Option](#WithPath)


#### <a name="pkg-files">Package files</a>
[level.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/level.go) [level_string.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/level_string.go) [option.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/option.go) [validation.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/validation.go) [validation_easyjson.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/validation_easyjson.go)



## <a name="pkg-variables">Variables</a>
``` go
var Global *validator.Validate
```
Global — синглтон валидатора

nolint: gochecknoglobals // экземпляр валидатора, на котором регистрируются кастомные проверки




## <a name="Error">type</a> [Error](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/validation.go?s=2373:2584#L78)
``` go
type Error struct {
    Path  string                 `json:"path"`
    Desc  string                 `json:"desc"`
    Level Level                  `json:"level"`
    Args  map[string]interface{} `json:"args"`
    // contains filtered or unexported fields
}

```
Error is an error with path and level

easyjson:json







### <a name="NewError">func</a> [NewError](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/validation.go?s=2610:2661#L87)
``` go
func NewError(desc string, options ...Option) Error
```
NewError constructor





### <a name="Error.Cause">func</a> (Error) [Cause](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/validation.go?s=3204:3233#L119)
``` go
func (ve Error) Cause() error
```
Cause invalid error

Implements: errors.Causer




### <a name="Error.Error">func</a> (Error) [Error](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/validation.go?s=2866:2896#L103)
``` go
func (ve Error) Error() string
```
Error return wrapped error string

Implements: error




### <a name="Error.MarshalEasyJSON">func</a> (Error) [MarshalEasyJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/validation_easyjson.go?s=3261:3310#L154)
``` go
func (v Error) MarshalEasyJSON(w *jwriter.Writer)
```
MarshalEasyJSON supports easyjson.Marshaler interface




### <a name="Error.MarshalJSON">func</a> (Error) [MarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/validation_easyjson.go?s=3018:3062#L147)
``` go
func (v Error) MarshalJSON() ([]byte, error)
```
MarshalJSON supports json.Marshaler interface




### <a name="Error.Prefix">func</a> (Error) [Prefix](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/validation.go?s=2982:3027#L108)
``` go
func (ve Error) Prefix(parts ...string) Error
```
Prefix error with given path




### <a name="Error.UnmarshalEasyJSON">func</a> (\*Error) [UnmarshalEasyJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/validation_easyjson.go?s=3680:3730#L166)
``` go
func (v *Error) UnmarshalEasyJSON(l *jlexer.Lexer)
```
UnmarshalEasyJSON supports easyjson.Unmarshaler interface




### <a name="Error.UnmarshalJSON">func</a> (\*Error) [UnmarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/validation_easyjson.go?s=3442:3490#L159)
``` go
func (v *Error) UnmarshalJSON(data []byte) error
```
UnmarshalJSON supports json.Unmarshaler interface




## <a name="Errors">type</a> [Errors](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/validation.go?s=3367:3386#L128)
``` go
type Errors []Error
```
Errors is a collection of errors of the some struct







### <a name="NewErrors">func</a> [NewErrors](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/validation.go?s=3413:3442#L131)
``` go
func NewErrors(n int) *Errors
```
NewErrors constructor


### <a name="ValidateStruct">func</a> [ValidateStruct](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/validation.go?s=6073:6138#L264)
``` go
func ValidateStruct(ctx context.Context, val interface{}) *Errors
```
ValidateStruct with gopkg.in/go-playground/validator and convert errors to Errors





### <a name="Errors.Add">func</a> (\*Errors) [Add](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/validation.go?s=4294:4337#L182)
``` go
func (verrs *Errors) Add(err error) *Errors
```
Add some error




### <a name="Errors.AddItem">func</a> (\*Errors) [AddItem](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/validation.go?s=4399:4467#L187)
``` go
func (verrs *Errors) AddItem(desc string, options ...Option) *Errors
```
AddItem to list




### <a name="Errors.AddWithPrefix">func</a> (\*Errors) [AddWithPrefix](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/validation.go?s=4581:4649#L192)
``` go
func (verrs *Errors) AddWithPrefix(err error, prefix string) *Errors
```
AddWithPrefix like AddError but prefix validation error(s)




### <a name="Errors.Cause">func</a> (Errors) [Cause](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/validation.go?s=5658:5691#L245)
``` go
func (verrs Errors) Cause() error
```
Cause return critical causer

Implements: errors.Causer




### <a name="Errors.Error">func</a> (Errors) [Error](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/validation.go?s=3554:3588#L140)
``` go
func (verrs Errors) Error() string
```
LevelError return marshaled error

Implements: error




### <a name="Errors.HasLevel">func</a> (Errors) [HasLevel](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/validation.go?s=4130:4176#L171)
``` go
func (verrs Errors) HasLevel(level Level) bool
```
HasLevel check errors for specified error level or less




### <a name="Errors.IsCritical">func</a> (Errors) [IsCritical](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/validation.go?s=3959:3996#L166)
``` go
func (verrs Errors) IsCritical() bool
```
IsCritical error




### <a name="Errors.IsEmpty">func</a> (Errors) [IsEmpty](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/validation.go?s=3875:3909#L161)
``` go
func (verrs Errors) IsEmpty() bool
```
IsEmpty error




### <a name="Errors.Iter">func</a> (\*Errors) [Iter](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/validation.go?s=6931:6974#L299)
``` go
func (verrs *Errors) Iter() *ErrorsIterator
```
Iter return new iterator




### <a name="Errors.Len">func</a> (Errors) [Len](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/validation.go?s=3804:3833#L156)
``` go
func (verrs Errors) Len() int
```
Len return count of errors




### <a name="Errors.Prefix">func</a> (Errors) [Prefix](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/validation.go?s=5006:5055#L213)
``` go
func (verrs Errors) Prefix(prefix string) *Errors
```
Prefix all errors with given path




### <a name="Errors.Return">func</a> (\*Errors) [Return](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/validation.go?s=3679:3714#L147)
``` go
func (verrs *Errors) Return() error
```
Return nil if empty else self




### <a name="Errors.WithRecover">func</a> (\*Errors) [WithRecover](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/validation.go?s=5238:5289#L223)
``` go
func (verrs *Errors) WithRecover(fn func()) *Errors
```
WithRecover execute callback and recover panic to critical error




### <a name="Errors.WriteData">func</a> (Errors) [WriteData](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/validation.go?s=5819:5867#L254)
``` go
func (verrs Errors) WriteData(w io.Writer) error
```
WriteData writes data to writer




## <a name="ErrorsIterator">type</a> [ErrorsIterator](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/validation.go?s=6842:6901#L293)
``` go
type ErrorsIterator struct {
    // contains filtered or unexported fields
}

```
ErrorsIterator implements Next/Value interface










### <a name="ErrorsIterator.Next">func</a> (\*ErrorsIterator) [Next](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/validation.go?s=7111:7152#L307)
``` go
func (verrsi *ErrorsIterator) Next() bool
```
Next update iterator index and return is any next value available




### <a name="ErrorsIterator.Value">func</a> (\*ErrorsIterator) [Value](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/validation.go?s=7289:7332#L317)
``` go
func (verrsi *ErrorsIterator) Value() Error
```
Value return current iterator value




## <a name="Level">type</a> [Level](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/level.go?s=171:186#L6)
``` go
type Level int8
```
Level of validation error


``` go
const (
    // LevelCritical error break follow validation process
    LevelCritical Level = iota
    // LevelError is a standard error shown that data is invalid
    LevelError
    // LevelWarning is a just wraning
    LevelWarning
)
```






### <a name="LevelString">func</a> [LevelString](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/level_string.go?s=731:772#L32)
``` go
func LevelString(s string) (Level, error)
```
LevelString retrieves an enum value from the enum constants string name.
Throws an error if the param is not part of the enum.


### <a name="LevelValues">func</a> [LevelValues](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/level_string.go?s=953:979#L40)
``` go
func LevelValues() []Level
```
LevelValues returns all values of the enum





### <a name="Level.IsALevel">func</a> (Level) [IsALevel](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/level_string.go?s=1098:1128#L45)
``` go
func (i Level) IsALevel() bool
```
IsALevel returns "true" if the value is listed in the enum definition. "false" otherwise




### <a name="Level.MarshalJSON">func</a> (Level) [MarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/level_string.go?s=1283:1327#L55)
``` go
func (i Level) MarshalJSON() ([]byte, error)
```
MarshalJSON implements the json.Marshaler interface for Level




### <a name="Level.String">func</a> (Level) [String](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/level_string.go?s=270:300#L15)
``` go
func (i Level) String() string
```



### <a name="Level.UnmarshalJSON">func</a> (\*Level) [UnmarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/level_string.go?s=1435:1483#L60)
``` go
func (i *Level) UnmarshalJSON(data []byte) error
```
UnmarshalJSON implements the json.Unmarshaler interface for Level




## <a name="Option">type</a> [Option](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/option.go?s=100:145#L9)
``` go
type Option interface {
    Apply(Error) Error
}
```
Option опции ошибки валидации







### <a name="WithArg">func</a> [WithArg](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/option.go?s=859:907#L54)
``` go
func WithArg(key string, val interface{}) Option
```
WithArg добавить аргумент


### <a name="WithCause">func</a> [WithCause](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/option.go?s=1211:1243#L72)
``` go
func WithCause(err error) Option
```
WithCause добавить стороннюю критическую ошибку


### <a name="WithLevel">func</a> [WithLevel](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/option.go?s=249:283#L18)
``` go
func WithLevel(level Level) Option
```
WithLevel установить уровень ошибки


### <a name="WithPath">func</a> [WithPath](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/validation/option.go?s=491:534#L33)
``` go
func WithPath(chunks ...interface{}) Option
```
WithPath установить путь ошибки








- - -
Generated by [godoc2md](https://github.com/Exa-Networks/godoc2md)
