# md
`import "git.elewise.com/elma365/common/pkg/md"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [func AddKV(ctx context.Context, key, value string) context.Context](#AddKV)
* [func CompanyAliasFromContext(ctx context.Context) string](#CompanyAliasFromContext)
* [func CompanyFromContext(ctx context.Context) string](#CompanyFromContext)
* [func CompanyNameFromContext(ctx context.Context) string](#CompanyNameFromContext)
* [func ContextWithCompany(ctx context.Context, company string) context.Context](#ContextWithCompany)
* [func ContextWithCompanyAlias(ctx context.Context, alias string) context.Context](#ContextWithCompanyAlias)
* [func ContextWithEdition(ctx context.Context, edition pkgEdition.Edition) context.Context](#ContextWithEdition)
* [func ContextWithGateway(ctx context.Context, gateway string) context.Context](#ContextWithGateway)
* [func ContextWithIsAdmin(ctx context.Context, isAdmin bool) context.Context](#ContextWithIsAdmin)
* [func ContextWithIsPortalUser(ctx context.Context, isPortal bool) context.Context](#ContextWithIsPortalUser)
* [func ContextWithTimestamp(ctx context.Context, ts time.Time) context.Context](#ContextWithTimestamp)
* [func ContextWithUserCompanyID(ctx context.Context, userCompanyID uuid.UUID) context.Context](#ContextWithUserCompanyID)
* [func ContextWithUserID(ctx context.Context, userID uuid.UUID) context.Context](#ContextWithUserID)
* [func Copy(ctx context.Context) context.Context](#Copy)
* [func EditionFromContext(ctx context.Context) pkgEdition.Edition](#EditionFromContext)
* [func ExtractGateways(ctx context.Context, key string, headers IncomingHeaders) context.Context](#ExtractGateways)
* [func ExtractKV(ctx context.Context, prefix string, headers IncomingHeaders) context.Context](#ExtractKV)
* [func GatewaysFromContext(ctx context.Context) []string](#GatewaysFromContext)
* [func HasKV(ctx context.Context, key, value string) bool](#HasKV)
* [func InjectGateways(ctx context.Context, key string, headers OutgoingHeaders)](#InjectGateways)
* [func InjectKV(ctx context.Context, prefix string, headers OutgoingHeaders)](#InjectKV)
* [func IsAdminFromContext(ctx context.Context) bool](#IsAdminFromContext)
* [func IsPortalUserFromContext(ctx context.Context) bool](#IsPortalUserFromContext)
* [func ListV(ctx context.Context, key string) (res []string)](#ListV)
* [func PeakV(ctx context.Context, key string) (string, bool)](#PeakV)
* [func RangeKV(ctx context.Context, visitor func(key string, value string) bool)](#RangeKV)
* [func TimestampFromContext(ctx context.Context) time.Time](#TimestampFromContext)
* [func TryCompanyFromContext(ctx context.Context) (string, bool)](#TryCompanyFromContext)
* [func TryEditionFromContext(ctx context.Context) (pkgEdition.Edition, bool)](#TryEditionFromContext)
* [func TryUserIDFromContext(ctx context.Context) (uuid.UUID, bool)](#TryUserIDFromContext)
* [func UserCompanyIDFromContext(ctx context.Context) uuid.UUID](#UserCompanyIDFromContext)
* [func UserIDFromContext(ctx context.Context) uuid.UUID](#UserIDFromContext)
* [type IncomingHeaders](#IncomingHeaders)
* [type OutgoingHeaders](#OutgoingHeaders)


#### <a name="pkg-files">Package files</a>
[copy.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/md/copy.go) [headers.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/md/headers.go) [kv.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/md/kv.go) [md.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/md/md.go)





## <a name="AddKV">func</a> [AddKV](https://git.elewise.com/elma365/common/-/tree/develop/pkg/md/kv.go?s=613:679#L23)
``` go
func AddKV(ctx context.Context, key, value string) context.Context
```
AddKV добавить строковую пару ключ-значение в контекст

Ключи приводятся к нижнему регистру. Ключу соответствует список значений,
поэтому при добавлении нескольких значений по одному ключу они все будут
сохранены в контексте.



## <a name="CompanyAliasFromContext">func</a> [CompanyAliasFromContext](https://git.elewise.com/elma365/common/-/tree/develop/pkg/md/md.go?s=1253:1309#L39)
``` go
func CompanyAliasFromContext(ctx context.Context) string
```
CompanyAliasFromContext извлечь алиас компании из контекста



## <a name="CompanyFromContext">func</a> [CompanyFromContext](https://git.elewise.com/elma365/common/-/tree/develop/pkg/md/md.go?s=480:531#L20)
``` go
func CompanyFromContext(ctx context.Context) string
```
CompanyFromContext извлекает имя компании из контекста



## <a name="CompanyNameFromContext">func</a> [CompanyNameFromContext](https://git.elewise.com/elma365/common/-/tree/develop/pkg/md/md.go?s=1774:1829#L51)
``` go
func CompanyNameFromContext(ctx context.Context) string
```
CompanyNameFromContext извлечь финальное имя компании из контекста
Если у компании есть Alias, то извлекает его, если нет - возвращает сгенерированное имя компании (CompanyFromContext)
которое используется в PG схемах



## <a name="ContextWithCompany">func</a> [ContextWithCompany](https://git.elewise.com/elma365/common/-/tree/develop/pkg/md/md.go?s=252:328#L15)
``` go
func ContextWithCompany(ctx context.Context, company string) context.Context
```
ContextWithCompany сохраняет имя компании в контексте



## <a name="ContextWithCompanyAlias">func</a> [ContextWithCompanyAlias](https://git.elewise.com/elma365/common/-/tree/develop/pkg/md/md.go?s=1014:1093#L34)
``` go
func ContextWithCompanyAlias(ctx context.Context, alias string) context.Context
```
ContextWithCompanyAlias сохранить алиса компании (человеческое название) в контексте



## <a name="ContextWithEdition">func</a> [ContextWithEdition](https://git.elewise.com/elma365/common/-/tree/develop/pkg/md/md.go?s=5156:5244#L142)
``` go
func ContextWithEdition(ctx context.Context, edition pkgEdition.Edition) context.Context
```
ContextWithEdition сохраняет edition компании в контексте



## <a name="ContextWithGateway">func</a> [ContextWithGateway](https://git.elewise.com/elma365/common/-/tree/develop/pkg/md/md.go?s=6023:6099#L165)
``` go
func ContextWithGateway(ctx context.Context, gateway string) context.Context
```
ContextWithGateway добавляет запись о точке входа/проксирования в контекст



## <a name="ContextWithIsAdmin">func</a> [ContextWithIsAdmin](https://git.elewise.com/elma365/common/-/tree/develop/pkg/md/md.go?s=2945:3019#L90)
``` go
func ContextWithIsAdmin(ctx context.Context, isAdmin bool) context.Context
```
ContextWithIsAdmin сохраняет признак того, что пользователь администратор



## <a name="ContextWithIsPortalUser">func</a> [ContextWithIsPortalUser](https://git.elewise.com/elma365/common/-/tree/develop/pkg/md/md.go?s=3509:3589#L104)
``` go
func ContextWithIsPortalUser(ctx context.Context, isPortal bool) context.Context
```
ContextWithIsPortalUser сохраняет признак того, что пользователь портальный



## <a name="ContextWithTimestamp">func</a> [ContextWithTimestamp](https://git.elewise.com/elma365/common/-/tree/develop/pkg/md/md.go?s=4063:4139#L118)
``` go
func ContextWithTimestamp(ctx context.Context, ts time.Time) context.Context
```
ContextWithTimestamp создает новый контекст с временной меткой



## <a name="ContextWithUserCompanyID">func</a> [ContextWithUserCompanyID](https://git.elewise.com/elma365/common/-/tree/develop/pkg/md/md.go?s=4590:4681#L130)
``` go
func ContextWithUserCompanyID(ctx context.Context, userCompanyID uuid.UUID) context.Context
```
ContextWithUserCompanyID сохраняет id компании с которой связан текущий пользователь в контексте



## <a name="ContextWithUserID">func</a> [ContextWithUserID](https://git.elewise.com/elma365/common/-/tree/develop/pkg/md/md.go?s=2186:2263#L71)
``` go
func ContextWithUserID(ctx context.Context, userID uuid.UUID) context.Context
```
ContextWithUserID сохраняет id пользователя в контексте



## <a name="Copy">func</a> [Copy](https://git.elewise.com/elma365/common/-/tree/develop/pkg/md/copy.go?s=293:339#L13)
``` go
func Copy(ctx context.Context) context.Context
```
Copy возвращает новый контекст с информацией из текущего



## <a name="EditionFromContext">func</a> [EditionFromContext](https://git.elewise.com/elma365/common/-/tree/develop/pkg/md/md.go?s=5397:5460#L147)
``` go
func EditionFromContext(ctx context.Context) pkgEdition.Edition
```
EditionFromContext извлекает edition компании из контекста



## <a name="ExtractGateways">func</a> [ExtractGateways](https://git.elewise.com/elma365/common/-/tree/develop/pkg/md/md.go?s=7148:7242#L191)
``` go
func ExtractGateways(ctx context.Context, key string, headers IncomingHeaders) context.Context
```
ExtractGateways извлечь шлюзы из заголовков, записанные методом InjectGateways



## <a name="ExtractKV">func</a> [ExtractKV](https://git.elewise.com/elma365/common/-/tree/develop/pkg/md/kv.go?s=2944:3035#L91)
``` go
func ExtractKV(ctx context.Context, prefix string, headers IncomingHeaders) context.Context
```
ExtractKV извлечь пары из заголовков, записанные методом InjectKV



## <a name="GatewaysFromContext">func</a> [GatewaysFromContext](https://git.elewise.com/elma365/common/-/tree/develop/pkg/md/md.go?s=6400:6454#L172)
``` go
func GatewaysFromContext(ctx context.Context) []string
```
GatewaysFromContext извлекает последовательность шлюзов и прокси запроса из контекста



## <a name="HasKV">func</a> [HasKV](https://git.elewise.com/elma365/common/-/tree/develop/pkg/md/kv.go?s=1446:1501#L48)
``` go
func HasKV(ctx context.Context, key, value string) bool
```
HasKV проверить наличие значения в списке по ключу



## <a name="InjectGateways">func</a> [InjectGateways](https://git.elewise.com/elma365/common/-/tree/develop/pkg/md/md.go?s=6859:6936#L184)
``` go
func InjectGateways(ctx context.Context, key string, headers OutgoingHeaders)
```
InjectGateways записать шлюзы в заголовки для последующего извлечения методом ExtractGateways



## <a name="InjectKV">func</a> [InjectKV](https://git.elewise.com/elma365/common/-/tree/develop/pkg/md/kv.go?s=2622:2696#L84)
``` go
func InjectKV(ctx context.Context, prefix string, headers OutgoingHeaders)
```
InjectKV записать пары в заголовки для последующего извлечения методом ExtractKV



## <a name="IsAdminFromContext">func</a> [IsAdminFromContext](https://git.elewise.com/elma365/common/-/tree/develop/pkg/md/md.go?s=3231:3280#L95)
``` go
func IsAdminFromContext(ctx context.Context) bool
```
IsAdminFromContext извлекает из контекста признак того, что пользователь администратор



## <a name="IsPortalUserFromContext">func</a> [IsPortalUserFromContext](https://git.elewise.com/elma365/common/-/tree/develop/pkg/md/md.go?s=3806:3860#L109)
``` go
func IsPortalUserFromContext(ctx context.Context) bool
```
IsPortalUserFromContext извлекает из контекста признак того, что пользователь портальный



## <a name="ListV">func</a> [ListV](https://git.elewise.com/elma365/common/-/tree/develop/pkg/md/kv.go?s=1010:1068#L34)
``` go
func ListV(ctx context.Context, key string) (res []string)
```
ListV извлечь массив значений по ключу



## <a name="PeakV">func</a> [PeakV](https://git.elewise.com/elma365/common/-/tree/develop/pkg/md/kv.go?s=1794:1852#L59)
``` go
func PeakV(ctx context.Context, key string) (string, bool)
```
PeakV получить первое значение из списка по ключу



## <a name="RangeKV">func</a> [RangeKV](https://git.elewise.com/elma365/common/-/tree/develop/pkg/md/kv.go?s=2276:2354#L76)
``` go
func RangeKV(ctx context.Context, visitor func(key string, value string) bool)
```
RangeKV обойти список всех пар в контексте

Если visitor возвращает false, итерирование прекращается.



## <a name="TimestampFromContext">func</a> [TimestampFromContext](https://git.elewise.com/elma365/common/-/tree/develop/pkg/md/md.go?s=4290:4346#L123)
``` go
func TimestampFromContext(ctx context.Context) time.Time
```
TimestampFromContext извлекает временную метку из контекста



## <a name="TryCompanyFromContext">func</a> [TryCompanyFromContext](https://git.elewise.com/elma365/common/-/tree/develop/pkg/md/md.go?s=699:761#L25)
``` go
func TryCompanyFromContext(ctx context.Context) (string, bool)
```
TryCompanyFromContext извлекает имя компании из контекста, если она есть



## <a name="TryEditionFromContext">func</a> [TryEditionFromContext](https://git.elewise.com/elma365/common/-/tree/develop/pkg/md/md.go?s=5639:5713#L152)
``` go
func TryEditionFromContext(ctx context.Context) (pkgEdition.Edition, bool)
```
TryEditionFromContext извлекает edition компании из контекста, если он есть



## <a name="TryUserIDFromContext">func</a> [TryUserIDFromContext](https://git.elewise.com/elma365/common/-/tree/develop/pkg/md/md.go?s=2648:2712#L81)
``` go
func TryUserIDFromContext(ctx context.Context) (uuid.UUID, bool)
```
TryUserIDFromContext извлекает id пользователя из контекста если оно там есть



## <a name="UserCompanyIDFromContext">func</a> [UserCompanyIDFromContext](https://git.elewise.com/elma365/common/-/tree/develop/pkg/md/md.go?s=4920:4980#L135)
``` go
func UserCompanyIDFromContext(ctx context.Context) uuid.UUID
```
UserCompanyIDFromContext извлекате из контекста ид компании с которой связан текущий пользователь



## <a name="UserIDFromContext">func</a> [UserIDFromContext](https://git.elewise.com/elma365/common/-/tree/develop/pkg/md/md.go?s=2416:2469#L76)
``` go
func UserIDFromContext(ctx context.Context) uuid.UUID
```
UserIDFromContext извлекает id пользователя из контекста




## <a name="IncomingHeaders">type</a> [IncomingHeaders](https://git.elewise.com/elma365/common/-/tree/develop/pkg/md/headers.go?s=326:441#L9)
``` go
type IncomingHeaders interface {
    Values(key string) []string
    Range(func(key string, values []string) bool) bool
}
```
IncomingHeaders интерфейс заголовков из которых можно извлечь метаданные










## <a name="OutgoingHeaders">type</a> [OutgoingHeaders](https://git.elewise.com/elma365/common/-/tree/develop/pkg/md/headers.go?s=140:198#L4)
``` go
type OutgoingHeaders interface {
    Add(key, value string)
}
```
OutgoingHeaders интерфейс заголовков в которые можно сохранить метаданные













- - -
Generated by [godoc2md](https://github.com/Exa-Networks/godoc2md)
