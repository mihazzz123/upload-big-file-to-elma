# i18n
`import "git.elewise.com/elma365/common/pkg/i18n"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [func ContextWithLang(ctx context.Context, lang string) context.Context](#ContextWithLang)
* [func LangFromContext(ctx context.Context) string](#LangFromContext)
* [func LoadFromTC(tcURL, path, serviceName string) error](#LoadFromTC)
* [func Patch(localesDir, defaultLang string) error](#Patch)
* [type Config](#Config)
* [type I18n](#I18n)
  * [func New(config Config, localesDir string, opts ...Option) (I18n, error)](#New)
  * [func (i18n I18n) GetAvailableLangs() []string](#I18n.GetAvailableLangs)
  * [func (i18n I18n) GetLangDir(ctx context.Context) string](#I18n.GetLangDir)
  * [func (i18n I18n) IsExistTranslate(ctx context.Context, str string) bool](#I18n.IsExistTranslate)
  * [func (i18n I18n) LangFromContext(ctx context.Context) string](#I18n.LangFromContext)
  * [func (i18n I18n) TranslateString(ctx context.Context, str string, args ...interface{}) string](#I18n.TranslateString)
  * [func (i18n I18n) TranslateTemplate(ctx context.Context, src string, data interface{}) (string, error)](#I18n.TranslateTemplate)
* [type Option](#Option)
  * [func WithDomain(domain string) Option](#WithDomain)
* [type Patcher](#Patcher)
  * [func NewPatcher(localesDir, defaultLang string) *Patcher](#NewPatcher)
  * [func (i18n Patcher) PatchJSONFiles() error](#Patcher.PatchJSONFiles)


#### <a name="pkg-files">Package files</a>
[i18n.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/i18n/i18n.go) [metrics.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/i18n/metrics.go) [patcher.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/i18n/patcher.go) [tc.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/i18n/tc.go) [tc_easyjson.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/i18n/tc_easyjson.go)


## <a name="pkg-constants">Constants</a>
``` go
const DefaultLocalesPath = "data/locales"
```
DefaultLocalesPath - рекомендуемый относительный путь до папки с локализациями

``` go
const DefaultTCUrl = "http://tc.elma-bpm.com/api/elma365"
```
DefaultTCUrl - адрес проекта elma365 в TC




## <a name="ContextWithLang">func</a> [ContextWithLang](https://git.elewise.com/elma365/common/-/tree/develop/pkg/i18n/i18n.go?s=490:560#L26)
``` go
func ContextWithLang(ctx context.Context, lang string) context.Context
```
ContextWithLang добавляет язык запроса в контекст



## <a name="LangFromContext">func</a> [LangFromContext](https://git.elewise.com/elma365/common/-/tree/develop/pkg/i18n/i18n.go?s=4330:4378#L137)
``` go
func LangFromContext(ctx context.Context) string
```
LangFromContext извлекает язык из контекста, если языка нет в контексте - возвращает пустой



## <a name="LoadFromTC">func</a> [LoadFromTC](https://git.elewise.com/elma365/common/-/tree/develop/pkg/i18n/tc.go?s=1488:1542#L47)
``` go
func LoadFromTC(tcURL, path, serviceName string) error
```
LoadFromTC загружает локализации для сервиса из TC
Параметр tcURL - путь к TC (можно использовать константу DefaultTCUrl для проекта elma365)
Параметр path - относительный путь до папки с языками (можно использовать константу DefaultLocalesPath)
Параметр serviceName - имя сервиса (имя папки в архиве с локализацией)

Пример использования для загрузки локализаций сервиса:
```
func main() {


	 if len(os.Args) < 2 {
		panic(errors.New("Service name is required as 1st argument"))
	 }
	 serviceName := os.Args[1]
	 if err := i18n.LoadFromTC(i18n.DefaultTCUrl, i18n.DefaultLocalesPath, serviceName); err != nil {
		panic(err)
	 }

}
```



## <a name="Patch">func</a> [Patch](https://git.elewise.com/elma365/common/-/tree/develop/pkg/i18n/patcher.go?s=408:456#L21)
``` go
func Patch(localesDir, defaultLang string) error
```
Patch - Перевести все ресурсы. Ресурсы берутся из папки языка по умолчанию




## <a name="Config">type</a> [Config](https://git.elewise.com/elma365/common/-/tree/develop/pkg/i18n/i18n.go?s=298:384#L18)
``` go
type Config interface {
    GetDefaultLanguage() string
    GetAllowedLanguages() []string
}
```
Config Параметры конфигурации для i18n










## <a name="I18n">type</a> [I18n](https://git.elewise.com/elma365/common/-/tree/develop/pkg/i18n/i18n.go?s=1054:1176#L49)
``` go
type I18n struct {
    // contains filtered or unexported fields
}

```
I18n сервис переводов







### <a name="New">func</a> [New](https://git.elewise.com/elma365/common/-/tree/develop/pkg/i18n/i18n.go?s=2628:2700#L86)
``` go
func New(config Config, localesDir string, opts ...Option) (I18n, error)
```
New загружает все доступные локали

Для корректной работы загрузки локалей необходимо в корне проекта добавить папку data, внутри которой создать папку
locales со следующей структурой:
```
/data
/data/locales
/data/locales/ru-RU
/data/locales/ru-RU/default.po
```

Внимание! По умолчанию читаются только файлы default.po. Остальные файлы можно использовать для других целей.

Для того, чтобы файлы добавились в docker-контейнер, необходимо прописать в Dockerfile следующее:
```
...
COPY data /data
...
WORKDIR "/"
...
```
То есть надо скопировать папку data в корень контейнера, а также рабочей директорией установить корень.

В качестве второго аргумента необходимо передать путь до папки locales.

Для обобщения работы с переводами также рекомендуется отражать в ключах подстановки, то есть:
```
msgid "some@key(%d)"
msgstr "Сообщение №%d"
```





### <a name="I18n.GetAvailableLangs">func</a> (I18n) [GetAvailableLangs](https://git.elewise.com/elma365/common/-/tree/develop/pkg/i18n/i18n.go?s=6383:6428#L196)
``` go
func (i18n I18n) GetAvailableLangs() []string
```
GetAvailableLangs возвращает коды доступных языков




### <a name="I18n.GetLangDir">func</a> (I18n) [GetLangDir](https://git.elewise.com/elma365/common/-/tree/develop/pkg/i18n/i18n.go?s=5058:5113#L156)
``` go
func (i18n I18n) GetLangDir(ctx context.Context) string
```
GetLangDir - возвращает папку, где хранятся локализованные ресурсы для текущего контекста




### <a name="I18n.IsExistTranslate">func</a> (I18n) [IsExistTranslate](https://git.elewise.com/elma365/common/-/tree/develop/pkg/i18n/i18n.go?s=6177:6248#L191)
``` go
func (i18n I18n) IsExistTranslate(ctx context.Context, str string) bool
```
IsExistTranslate проверка существование ключа




### <a name="I18n.LangFromContext">func</a> (I18n) [LangFromContext](https://git.elewise.com/elma365/common/-/tree/develop/pkg/i18n/i18n.go?s=4651:4711#L143)
``` go
func (i18n I18n) LangFromContext(ctx context.Context) string
```
LangFromContext извлекает язык из контекста, если языка нет в контексте или среди известных, возвращает язык по умолчанию




### <a name="I18n.TranslateString">func</a> (I18n) [TranslateString](https://git.elewise.com/elma365/common/-/tree/develop/pkg/i18n/i18n.go?s=5953:6046#L186)
``` go
func (i18n I18n) TranslateString(ctx context.Context, str string, args ...interface{}) string
```
TranslateString локализует ключ




### <a name="I18n.TranslateTemplate">func</a> (I18n) [TranslateTemplate](https://git.elewise.com/elma365/common/-/tree/develop/pkg/i18n/i18n.go?s=5444:5545#L166)
``` go
func (i18n I18n) TranslateTemplate(ctx context.Context, src string, data interface{}) (string, error)
```
TranslateTemplate рендерит текстовый шаблон с функцией `translate`

Внутри шаблона можно использовать конструкции вида
```
{{ translate `some@key(%d)` .Count }}
```




## <a name="Option">type</a> [Option](https://git.elewise.com/elma365/common/-/tree/develop/pkg/i18n/i18n.go?s=685:724#L31)
``` go
type Option interface {
    // contains filtered or unexported methods
}
```
Option опции создания сервиса переводов







### <a name="WithDomain">func</a> [WithDomain](https://git.elewise.com/elma365/common/-/tree/develop/pkg/i18n/i18n.go?s=906:943#L42)
``` go
func WithDomain(domain string) Option
```
WithDomain читать переводы из файлов с переданным именем





## <a name="Patcher">type</a> [Patcher](https://git.elewise.com/elma365/common/-/tree/develop/pkg/i18n/patcher.go?s=209:272#L15)
``` go
type Patcher struct {
    // contains filtered or unexported fields
}

```
Patcher - автопереводчик ресурсов







### <a name="NewPatcher">func</a> [NewPatcher](https://git.elewise.com/elma365/common/-/tree/develop/pkg/i18n/patcher.go?s=662:718#L30)
``` go
func NewPatcher(localesDir, defaultLang string) *Patcher
```
NewPatcher - новый автопереводчик





### <a name="Patcher.PatchJSONFiles">func</a> (Patcher) [PatchJSONFiles](https://git.elewise.com/elma365/common/-/tree/develop/pkg/i18n/patcher.go?s=829:871#L38)
``` go
func (i18n Patcher) PatchJSONFiles() error
```
PatchJSONFiles - перевести json файлы









## <a name="pkg-subdirectories">Subdirectories</a>

| Name | Synopsis |
| ---- | -------- |
| [..](..) | |
| [test_data](test_data/) |  |
| [test_data/locales](test_data/locales/) |  |


- - -
Generated by [godoc2md](https://github.com/Exa-Networks/godoc2md)
