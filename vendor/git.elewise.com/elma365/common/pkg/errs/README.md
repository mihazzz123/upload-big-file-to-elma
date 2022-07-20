# errs
`import "git.elewise.com/elma365/common/pkg/errs"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [func Code(err error) error](#Code)
* [func CodeToHTTPStatus(err error) int](#CodeToHTTPStatus)
* [func ErrorLevel(err error) zapcore.Level](#ErrorLevel)
* [func FromGRPCError(err error) error](#FromGRPCError)
* [func FromHTTPResponse(r *http.Response) error](#FromHTTPResponse)
* [func GRPCCodeFromError(err error) codes.Code](#GRPCCodeFromError)
* [func Handle(l *zap.Logger, w http.ResponseWriter, err error)](#Handle)
* [func HandleInternal(l *zap.Logger, w http.ResponseWriter, err error)](#HandleInternal)
* [func PanicIfError(err error, msg string, args ...interface{})](#PanicIfError)
* [func WithData(err error, data interface{}) error](#WithData)
* [func WithRecover(fn func()) (err error)](#WithRecover)
* [func WriteDataFromError(orig error, w io.Writer) error](#WriteDataFromError)
* [type ConstantError](#ConstantError)
  * [func (ce ConstantError) Error() string](#ConstantError.Error)

#### <a name="pkg-examples">Examples</a>
* [ConstantError](#example-constanterror)
* [WithRecover](#example-withrecover)

#### <a name="pkg-files">Package files</a>
[code.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/code.go) [const.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/const.go) [data.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/data.go) [errs.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/errs.go) [panic_if_error.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/panic_if_error.go) [recover.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/recover.go)


## <a name="pkg-constants">Constants</a>
``` go
const (
    NotFound    errorCode = "not found"
    Unavailable errorCode = "unavailable"
    // Collision для случая, когда сущность, которую пытается поменять данный запрос,
    // уже как-то изменена (существует по уникальному ключу, дублирует что-то и тд)
    Collision       errorCode = "collision"
    Forbidden       errorCode = "forbidden"
    Unauthorized    errorCode = "unauthorized"
    Internal        errorCode = "internal error"
    InvalidArgument errorCode = "invalid argument"
    NotImplemented  errorCode = "not implemented"
    Timeout         errorCode = "timeout exceeded"
    // Precondition для случая, когда нет нужного состояния для совершения действия (что-то выключено,
    // какие-то настройки не применились, нет необходимых ресурсов) - то, что не зависит от конкретного
    // запроса пользователя, а должно было быть сделано раньше
    Precondition errorCode = "precondition"
    Unknown      errorCode = "unknown"
    // ResourceExhausted ресурс исчерпан - пользовательская квота или например место на диске
    ResourceExhausted errorCode = "resource exhausted"
)
```
List of global error codes




## <a name="Code">func</a> [Code](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/code.go?s=4655:4681#L156)
``` go
func Code(err error) error
```
Code получить код из ошибки



## <a name="CodeToHTTPStatus">func</a> [CodeToHTTPStatus](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/code.go?s=7387:7423#L271)
``` go
func CodeToHTTPStatus(err error) int
```
CodeToHTTPStatus пробует извлечь из ошибки константную основу и сопоставить ей HTTP статус

Если статус не найден, то возвращает 500.



## <a name="ErrorLevel">func</a> [ErrorLevel](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/code.go?s=4845:4885#L165)
``` go
func ErrorLevel(err error) zapcore.Level
```
ErrorLevel возвращает уровень лога ошибки



## <a name="FromGRPCError">func</a> [FromGRPCError](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/code.go?s=3770:3805#L125)
``` go
func FromGRPCError(err error) error
```
FromGRPCError принимает grpc ошибку и пытается сопоставить ей внутренний код ошибки

Если подходящей внутренней ошибки не найдено, то вернёт ошибку как есть.



## <a name="FromHTTPResponse">func</a> [FromHTTPResponse](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/code.go?s=4227:4272#L137)
``` go
func FromHTTPResponse(r *http.Response) error
```
FromHTTPResponse принимает http ответ и пытается сопоставить ей внутренний код ошибки

Если подходящей внутренней ошибки не найдено, то вернёт ошибку Unknown.



## <a name="GRPCCodeFromError">func</a> [GRPCCodeFromError](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/code.go?s=6815:6859#L248)
``` go
func GRPCCodeFromError(err error) codes.Code
```
GRPCCodeFromError сопоставляет код возврата GRPC ошибке

Если ошибка сформирована самим gRPC, то возвращает её код, если ошибка образована от константной ошибки,
то сопоставляет код по таблице. Если извлечь код не удаётся двигается по цепочке Cause. Если не удаётся
извлечь код по всей цепочке, то возвращает codes.Unknown.



## <a name="Handle">func</a> [Handle](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/errs.go?s=340:400#L16)
``` go
func Handle(l *zap.Logger, w http.ResponseWriter, err error)
```
Handle error according to it's type



## <a name="HandleInternal">func</a> [HandleInternal](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/errs.go?s=128:196#L10)
``` go
func HandleInternal(l *zap.Logger, w http.ResponseWriter, err error)
```
HandleInternal is a shortcut to log error and return 500 to client



## <a name="PanicIfError">func</a> [PanicIfError](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/panic_if_error.go?s=208:269#L6)
``` go
func PanicIfError(err error, msg string, args ...interface{})
```
PanicIfError - если есть ошибка, то бросается паника. Дополнительно можно передать сообщение



## <a name="WithData">func</a> [WithData](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/data.go?s=159:207#L13)
``` go
func WithData(err error, data interface{}) error
```
WithData добавляет информацию к ошибке



## <a name="WithRecover">func</a> [WithRecover](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/recover.go?s=244:283#L14)
``` go
func WithRecover(fn func()) (err error)
```
WithRecover отлавливает паники внутри себя и преобразовывает в ошибку


##### Example WithRecover:
При создании внутренней логики бывает неудобно постоянно обрабатывать ошибку,
тогда можно внутри использовать паники, но отлавливать для внешнего интерфейса.

``` go
err := WithRecover(func() {
    panic(NotFound)
})

if err != nil {
    fmt.Println(err.Error())
}
// Output: not found
```

Output:

```
not found
```


## <a name="WriteDataFromError">func</a> [WriteDataFromError](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/data.go?s=1330:1384#L68)
``` go
func WriteDataFromError(orig error, w io.Writer) error
```
WriteDataFromError to writer

We walking up by errors chain until find data or chain is ended.
If chain ended but data is not found, then error will be writed
as string or as is if it is valid json. If data has been found it will be writed.




## <a name="ConstantError">type</a> [ConstantError](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/const.go?s=145:170#L4)
``` go
type ConstantError string
```
ConstantError может использоваться для создания обрабатываемых исключений



##### Example ConstantError:
Иногда необходимо обрабатывать определённые ошибки как стандартное поведение системы.
В этом случае можно использовать константные ошибки, при этом можно обрачивать их `errors.WithMessage`
после чего доступаться до оригинальной ошибки с помощью `errors.Cause`.

``` go
const MyException ConstantError = "my exception"

MyFunction := func() error {
    return errors.WithMessage(MyException, "something went wrong")
}

err := MyFunction()
switch errors.Cause(err) {
case nil:
    // pass
case MyException:
    fmt.Println(err.Error())
default:
    panic(err)
}
// Output: something went wrong: my exception
```

Output:

```
something went wrong: my exception
```








### <a name="ConstantError.Error">func</a> (ConstantError) [Error](https://git.elewise.com/elma365/common/-/tree/develop/pkg/errs/const.go?s=263:301#L9)
``` go
func (ce ConstantError) Error() string
```
Error возвращает значение как строку

Implements: error









## <a name="pkg-subdirectories">Subdirectories</a>

| Name | Synopsis |
| ---- | -------- |
| [..](..) | |
| [validation](validation/) |  |


- - -
Generated by [godoc2md](https://github.com/Exa-Networks/godoc2md)
