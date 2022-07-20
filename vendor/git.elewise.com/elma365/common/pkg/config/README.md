# config
`import "git.elewise.com/elma365/common/pkg/config"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [func LoggerWithTraceID(ctx context.Context, logger *zap.Logger) *zap.Logger](#LoggerWithTraceID)
* [func NewCustom(name string, cfg Custom) error](#NewCustom)
* [func TraceIDFromContext(ctx context.Context) (traceID, spanID string)](#TraceIDFromContext)
* [type AmqpURL](#AmqpURL)
* [type Config](#Config)
  * [func New(name string) (Config, error)](#New)
  * [func (cfg Config) DNSResolve(service, proto, name string) (string, error)](#Config.DNSResolve)
  * [func (cfg Config) DNSResolveMulti(service, proto, name string) ([]string, error)](#Config.DNSResolveMulti)
  * [func (cfg Config) Dialer(service, proto, name string) func() (net.Conn, error)](#Config.Dialer)
  * [func (cfg Config) DisableDebug()](#Config.DisableDebug)
  * [func (cfg Config) Dump(ctx context.Context, name string, obj ...interface{})](#Config.Dump)
  * [func (cfg Config) EnableDebug()](#Config.EnableDebug)
  * [func (cfg Config) FullName() string](#Config.FullName)
  * [func (cfg Config) GRPCContextDialer(ctx context.Context, service string) (net.Conn, error)](#Config.GRPCContextDialer)
  * [func (cfg Config) GRPCTimeoutDialer(service string, timeout time.Duration) (net.Conn, error)](#Config.GRPCTimeoutDialer)
  * [func (cfg Config) GetAllowedLanguages() []string](#Config.GetAllowedLanguages)
  * [func (cfg Config) GetBind() string](#Config.GetBind)
  * [func (cfg Config) GetCMUXBind() bool](#Config.GetCMUXBind)
  * [func (cfg Config) GetDefaultLanguage() string](#Config.GetDefaultLanguage)
  * [func (cfg Config) GetDomain() string](#Config.GetDomain)
  * [func (cfg Config) GetEventBusMaxHashCount() int](#Config.GetEventBusMaxHashCount)
  * [func (cfg Config) GetGRPCBind() string](#Config.GetGRPCBind)
  * [func (cfg Config) GetHTTPBind() string](#Config.GetHTTPBind)
  * [func (cfg Config) GetHost() string](#Config.GetHost)
  * [func (cfg Config) GetMaxGRPCMessageSize() int](#Config.GetMaxGRPCMessageSize)
  * [func (cfg Config) GetMongoDBName() string](#Config.GetMongoDBName)
  * [func (cfg Config) GetMongoMaxPoolSize() uint64](#Config.GetMongoMaxPoolSize)
  * [func (cfg Config) GetMongoPassword() string](#Config.GetMongoPassword)
  * [func (cfg Config) GetMongoURL() (*MongoURL, error)](#Config.GetMongoURL)
  * [func (cfg Config) GetMongoUser() string](#Config.GetMongoUser)
  * [func (cfg Config) GetPostgresConnMaxLifetime() time.Duration](#Config.GetPostgresConnMaxLifetime)
  * [func (cfg Config) GetPostgresMaxIdleConns() int](#Config.GetPostgresMaxIdleConns)
  * [func (cfg Config) GetPostgresMaxOpenConns() int](#Config.GetPostgresMaxOpenConns)
  * [func (cfg Config) GetPostgresURL() (*PostgresURL, error)](#Config.GetPostgresURL)
  * [func (cfg Config) GetRabbitmqPassword() string](#Config.GetRabbitmqPassword)
  * [func (cfg Config) GetRabbitmqURL() (*AmqpURL, error)](#Config.GetRabbitmqURL)
  * [func (cfg Config) GetRabbitmqUser() string](#Config.GetRabbitmqUser)
  * [func (cfg Config) GetRabbitmqVHName() string](#Config.GetRabbitmqVHName)
  * [func (cfg Config) GetRedisURL() (*RedisURL, error)](#Config.GetRedisURL)
  * [func (cfg Config) GetScriptCompileTTL() time.Duration](#Config.GetScriptCompileTTL)
  * [func (cfg Config) GetScriptExecuteTTL() time.Duration](#Config.GetScriptExecuteTTL)
  * [func (cfg Config) GetScriptQueueTTL() time.Duration](#Config.GetScriptQueueTTL)
  * [func (cfg Config) GetServiceURL(name string) (string, error)](#Config.GetServiceURL)
  * [func (cfg Config) GetTimeout() time.Duration](#Config.GetTimeout)
  * [func (cfg Config) GetZone() string](#Config.GetZone)
  * [func (cfg Config) IsDebug() bool](#Config.IsDebug)
  * [func (cfg Config) Name() string](#Config.Name)
  * [func (cfg Config) PostgresURL() (string, error)](#Config.PostgresURL)
  * [func (cfg Config) RabbitmqURL() (string, error)](#Config.RabbitmqURL)
  * [func (cfg Config) StdConfig() Config](#Config.StdConfig)
* [type Custom](#Custom)
* [type Interface](#Interface)
* [type MongoURL](#MongoURL)
* [type PostgresURL](#PostgresURL)
* [type RedisURL](#RedisURL)
* [type Solution](#Solution)
  * [func SolutionString(s string) (Solution, error)](#SolutionString)
  * [func SolutionValues() []Solution](#SolutionValues)
  * [func (i Solution) IsASolution() bool](#Solution.IsASolution)
  * [func (i Solution) MarshalJSON() ([]byte, error)](#Solution.MarshalJSON)
  * [func (i Solution) MarshalText() ([]byte, error)](#Solution.MarshalText)
  * [func (i Solution) String() string](#Solution.String)
  * [func (i *Solution) UnmarshalJSON(data []byte) error](#Solution.UnmarshalJSON)
  * [func (i *Solution) UnmarshalText(text []byte) error](#Solution.UnmarshalText)

#### <a name="pkg-examples">Examples</a>
* [Config](#example-config)
* [NewCustom](#example-newcustom)

#### <a name="pkg-files">Package files</a>
[config.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go) [solution.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/solution.go) [solution_string.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/solution_string.go)


## <a name="pkg-constants">Constants</a>
``` go
const Prefix = "ELMA365_"
```
Prefix префикс переменных окружения




## <a name="LoggerWithTraceID">func</a> [LoggerWithTraceID](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=26694:26769#L793)
``` go
func LoggerWithTraceID(ctx context.Context, logger *zap.Logger) *zap.Logger
```
LoggerWithTraceID добавляет поля TraceID и SpanID к полям логгера



## <a name="NewCustom">func</a> [NewCustom](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=7175:7220#L171)
``` go
func NewCustom(name string, cfg Custom) error
```
NewCustom инициализирует расширение объекта конфигурации


##### Example NewCustom:
Расширение стандартного объекта настроек

``` go
var cfg Config

if err := config.NewCustom("my service", &cfg); err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
}
```


## <a name="TraceIDFromContext">func</a> [TraceIDFromContext](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=26336:26405#L783)
``` go
func TraceIDFromContext(ctx context.Context) (traceID, spanID string)
```
TraceIDFromContext extract trace and span ids if span started or empty string otherwise




## <a name="AmqpURL">type</a> [AmqpURL](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=6219:6241#L140)
``` go
type AmqpURL = url.URL
```
AmqpURL - alias добавлен для гибкости работы c Amqp










## <a name="Config">type</a> [Config](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=889:5832#L34)
``` go
type Config struct {
    Debug        bool
    EnableDump   bool
    Bind         string        `default:":8000"`
    HTTPBind     string        `default:":3000"`
    GRPCBind     string        `default:":5000"`
    CMUXBind     bool          `default:"true"`
    Timeout      time.Duration `default:"1m"`
    HostIP       string
    Telepresence string `envcfg:"TELEPRESENCE_CONTAINER"`

    // DefaultLanguage Язык по умолчанию
    DefaultLanguage string `default:"ru-RU"`
    // AllowedLanguages Список доступных языков через запятую
    AllowedLanguages []string `default:"ru-RU"`

    // Host домен/ip, на котором запущена elma365. Замена Host + "." + Zone
    Host string
    // Deprecated: Domain домен, за который отвечает сервис. Используйте Host
    Domain string `default:"elma365"`
    // Deprecated: Zone зона - ru/com/etc. Используйте Host
    Zone string `default:"ru"`

    // Postgres config
    // Deprecated: use GetPostgresURL
    PostgresUser string `default:"postgres"`
    // Deprecated: use GetPostgresURL
    PostgresPassword string `default:"postgres"`
    // Deprecated: use GetPostgresURL
    PostgresDBName string `default:"elma365"`
    // Deprecated: use GetPostgresURL
    PostgresSvcName string `default:"postgres"`
    // Deprecated: use GetPostgresURL
    PostgresPortName string `default:"postgres"`
    // PostgresMaxOpenConns количество подключений к каждой реплике. По умолчанию, не ограничено.
    PostgresMaxOpenConns int `default:"-1"`
    // PostgresMaxIdleConns maximum number of connections in the idle connection pool
    PostgresMaxIdleConns int `default:"-1"`
    // PostgresConnMaxLifetime maximum amount of time a connection may be reused
    PostgresConnMaxLifetime time.Duration `default:"5m"`

    // RabbitMQ config
    // Deprecated: use GetRabbitmqURL
    RabbitmqUser string `default:"rabbitmq"`
    // Deprecated: use GetRabbitmqURL
    RabbitmqPassword string `default:"rabbitmq"`
    // Deprecated: use GetRabbitmqURL
    RabbitmqVHName string `default:"elma365"`

    // Mongo config
    // Deprecated: use GetMongoURL
    MongoUser string `default:"mongo"`
    // Deprecated: use GetMongoURL
    MongoPassword string `default:"mongo"`
    // Deprecated: use GetMongoURL
    MongoDBName string `default:"elma365"`
    // Deprecated: use GetMongoURL
    MongoMaxPoolSize uint64 `default:"100"`

    // EventBus config
    EventBusMaxHashCount int `default:"128"`

    PsqlURL  string // default:"postgresql://postgres:postgres@postgres:5432/elma365"
    MongoURL string // default:"mongodb://mongo:mongo@mongo:27017/elma365?connectTimeoutMS=10000&serverSelectionTimeoutMS=5000&heartbeatFrequencyMS=30000&authSource=elma365&maxPoolSize=100"
    // URL подключения к redis
    RedisURL string // default:"redis://localhost:6379?ConnectTimeout=5000&IdleTimeOutSecs=180"
    AmqpURL  string // default:"amqp://rabbitmq:rabbitmq@host:5672/vhost"

    // Edition редакция конкретной инсталляции (в инсталяции у компаний могут быть разные Edition)
    // Внимание!!! Для разграничения функционала и получения Edition конкретной компании использовать md.EditionFromContext(ctx)
    Edition edition.Edition `default:"lite"` // lite - для обратной совместимости
    // Solution Версия сборки
    Solution Solution `default:"saas"`
    // MaxGRPCMessageSize максимальный размер ответов
    MaxGRPCMessageSize int `default:"8388608"` // 1024 * 1024 * 4 дефолтное значение GRPC
    // EnableDNSResolveOptimization отключает использование dnsResolve для http, grpc и redis сервисов
    EnableDNSResolveOptimization bool
    // SkipSSLVerify игнорировать ошибки сертификата в http клиенте
    SkipSSLVerify bool `default:"false"`

    // ScriptCompileTTL ограничение на время транспиляции сценария
    ScriptCompileTTL time.Duration `default:"10m"`
    // ScriptExecuteTTL ограничение на время исполнения сценария
    ScriptExecuteTTL time.Duration `default:"10m"`
    // ScriptQueueTTL ограничения на время ожидания в очереди
    //
    // Так как коммуникация с сервисом worker использует RabbitMQRPC, то сообщение может какое-то время
    // находиться в очереди до обработки. Если в течении указанного ограничения сообщение не будет вычитано из очереди,
    // оно будет уничтожено. Клиент получит ошибку таймаута.
    ScriptQueueTTL time.Duration `default:"10m"`
    // contains filtered or unexported fields
}

```
Config of application



##### Example Config:
В начале работы сервиса необходимо получить все его настройки

``` go
cfg, err := config.New("my service")
if err != nil {
    log.Fatal(err)
}
fmt.Println(cfg.Name())
// Output: my service
```

Output:

```
my service
```





### <a name="New">func</a> [New](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=6610:6647#L152)
``` go
func New(name string) (Config, error)
```
New создать новый конфиг

Создаёт новый конфиг и вычитывает его из переменных окружения. После этого настраивает глобальные логгеры
и трейсер.





### <a name="Config.DNSResolve">func</a> (Config) [DNSResolve](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=18319:18392#L548)
``` go
func (cfg Config) DNSResolve(service, proto, name string) (string, error)
```
DNSResolve возвращает резолв адреса в формате Host:Port, Сначала ищет в переменных окружения,
если не нашел резолвит сначала CNAME. Далее, пытается собрать адрес для известных протоколов. Если не получилось,
резолвит SRV запись для переданного сервиса. Возвращает только 1 адрес.




### <a name="Config.DNSResolveMulti">func</a> (Config) [DNSResolveMulti](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=19825:19905#L583)
``` go
func (cfg Config) DNSResolveMulti(service, proto, name string) ([]string, error)
```
DNSResolveMulti тоже что DNSResolve, но возвращает весь массив отрезолвленных адресов




### <a name="Config.Dialer">func</a> (Config) [Dialer](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=24085:24163#L716)
``` go
func (cfg Config) Dialer(service, proto, name string) func() (net.Conn, error)
```
Dialer return closured service dialer




### <a name="Config.DisableDebug">func</a> (Config) [DisableDebug](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=10042:10074#L279)
``` go
func (cfg Config) DisableDebug()
```
DisableDebug выключает логирование на уровне debug




### <a name="Config.Dump">func</a> (Config) [Dump](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=25928:26004#L771)
``` go
func (cfg Config) Dump(ctx context.Context, name string, obj ...interface{})
```
Dump если включена соответствующая опция, то пишет дамп объекта в base64 в лог




### <a name="Config.EnableDebug">func</a> (Config) [EnableDebug](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=9887:9918#L274)
``` go
func (cfg Config) EnableDebug()
```
EnableDebug включает логирование на уровне debug




### <a name="Config.FullName">func</a> (Config) [FullName](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=9556:9591#L264)
``` go
func (cfg Config) FullName() string
```
FullName возвращает имя сервиса с префиксом приложения (elma365.processor)




### <a name="Config.GRPCContextDialer">func</a> (Config) [GRPCContextDialer](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=24775:24865#L740)
``` go
func (cfg Config) GRPCContextDialer(ctx context.Context, service string) (net.Conn, error)
```
GRPCContextDialer is a method to pass it in grpc.Dial options




### <a name="Config.GRPCTimeoutDialer">func</a> (Config) [GRPCTimeoutDialer](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=24436:24528#L729)
``` go
func (cfg Config) GRPCTimeoutDialer(service string, timeout time.Duration) (net.Conn, error)
```
GRPCTimeoutDialer is a method to pass it in grpc.Dial options




### <a name="Config.GetAllowedLanguages">func</a> (Config) [GetAllowedLanguages](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=15731:15779#L463)
``` go
func (cfg Config) GetAllowedLanguages() []string
```
GetAllowedLanguages Список доступных языков




### <a name="Config.GetBind">func</a> (Config) [GetBind](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=10182:10216#L284)
``` go
func (cfg Config) GetBind() string
```
GetBind интерфейс для запуска сервера




### <a name="Config.GetCMUXBind">func</a> (Config) [GetCMUXBind](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=10365:10401#L289)
``` go
func (cfg Config) GetCMUXBind() bool
```
GetCMUXBind использовать механизм шаринга порта между gRPC и HTTP сервисами




### <a name="Config.GetDefaultLanguage">func</a> (Config) [GetDefaultLanguage](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=15584:15629#L458)
``` go
func (cfg Config) GetDefaultLanguage() string
```
GetDefaultLanguage Язык по умолчанию




### <a name="Config.GetDomain">func</a> (Config) [GetDomain](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=16155:16191#L478)
``` go
func (cfg Config) GetDomain() string
```
GetDomain домен, за который отвечает сервис
Deprecated: используйте GetHost()




### <a name="Config.GetEventBusMaxHashCount">func</a> (Config) [GetEventBusMaxHashCount](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=15443:15490#L453)
``` go
func (cfg Config) GetEventBusMaxHashCount() int
```
GetEventBusMaxHashCount максимальное количество возможных значений поля hash
в RoutingKey для шины событий




### <a name="Config.GetGRPCBind">func</a> (Config) [GetGRPCBind](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=10504:10542#L294)
``` go
func (cfg Config) GetGRPCBind() string
```
GetGRPCBind интерфейс для запуска grpc сервера




### <a name="Config.GetHTTPBind">func</a> (Config) [GetHTTPBind](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=10645:10683#L299)
``` go
func (cfg Config) GetHTTPBind() string
```
GetHTTPBind интерфейс для запуска http сервера




### <a name="Config.GetHost">func</a> (Config) [GetHost](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=15913:15947#L468)
``` go
func (cfg Config) GetHost() string
```
GetHost получить Host (домен), на котором запущена система




### <a name="Config.GetMaxGRPCMessageSize">func</a> (Config) [GetMaxGRPCMessageSize](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=25150:25195#L751)
``` go
func (cfg Config) GetMaxGRPCMessageSize() int
```
GetMaxGRPCMessageSize получить максимальный размер ответов




### <a name="Config.GetMongoDBName">func</a> (Config) [GetMongoDBName](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=14988:15029#L441)
``` go
func (cfg Config) GetMongoDBName() string
```
GetMongoDBName название коллекции в mongoDB
Deprecated: use GetMongoURL




### <a name="Config.GetMongoMaxPoolSize">func</a> (Config) [GetMongoMaxPoolSize](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=15191:15237#L447)
``` go
func (cfg Config) GetMongoMaxPoolSize() uint64
```
GetMongoMaxPoolSize максимальное количество соединений к mongoDB
Deprecated: use GetMongoURL




### <a name="Config.GetMongoPassword">func</a> (Config) [GetMongoPassword](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=14817:14860#L435)
``` go
func (cfg Config) GetMongoPassword() string
```
GetMongoPassword пароль пользователя mongoDB
Deprecated: use GetMongoURL




### <a name="Config.GetMongoURL">func</a> (Config) [GetMongoURL](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=12360:12410#L345)
``` go
func (cfg Config) GetMongoURL() (*MongoURL, error)
```
GetMongoURL - Распарсенная строка подключения к Mongo
Добавляет/изменяет в параметрах подключения имя приложения на имя конфигурации.
Параметры по умолчанию:
connectTimeoutMS          10_000
serverSelectionTimeoutMS  5_000
heartbeatFrequencyMS      30_000
authSource                такой же как и имя базы данных, если имя базы данных указана
maxPoolSize               100




### <a name="Config.GetMongoUser">func</a> (Config) [GetMongoUser](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=14653:14692#L429)
``` go
func (cfg Config) GetMongoUser() string
```
GetMongoUser имя пользователя mongoDB
Deprecated: use GetMongoURL




### <a name="Config.GetPostgresConnMaxLifetime">func</a> (Config) [GetPostgresConnMaxLifetime](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=13987:14047#L405)
``` go
func (cfg Config) GetPostgresConnMaxLifetime() time.Duration
```
GetPostgresConnMaxLifetime maximum amount of time a connection may be reused




### <a name="Config.GetPostgresMaxIdleConns">func</a> (Config) [GetPostgresMaxIdleConns](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=13821:13868#L400)
``` go
func (cfg Config) GetPostgresMaxIdleConns() int
```
GetPostgresMaxIdleConns maximum number of connections in the idle connection pool




### <a name="Config.GetPostgresMaxOpenConns">func</a> (Config) [GetPostgresMaxOpenConns](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=13650:13697#L395)
``` go
func (cfg Config) GetPostgresMaxOpenConns() int
```
GetPostgresMaxOpenConns количество подключений к каждой реплике




### <a name="Config.GetPostgresURL">func</a> (Config) [GetPostgresURL](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=11633:11689#L328)
``` go
func (cfg Config) GetPostgresURL() (*PostgresURL, error)
```
GetPostgresURL - Распарсенная строка подключения к Postgres
Добавляет в параметры подключения имя приложения




### <a name="Config.GetRabbitmqPassword">func</a> (Config) [GetRabbitmqPassword](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=14335:14381#L417)
``` go
func (cfg Config) GetRabbitmqPassword() string
```
GetRabbitmqPassword пароль rabbitmq
Deprecated: use GetRabbitmqURL




### <a name="Config.GetRabbitmqURL">func</a> (Config) [GetRabbitmqURL](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=17309:17361#L524)
``` go
func (cfg Config) GetRabbitmqURL() (*AmqpURL, error)
```
GetRabbitmqURL - получение урл для подключения из конфига к rabbitmq




### <a name="Config.GetRabbitmqUser">func</a> (Config) [GetRabbitmqUser](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=14183:14225#L411)
``` go
func (cfg Config) GetRabbitmqUser() string
```
GetRabbitmqUser имя пользователя rabbitmq
Deprecated: use GetRabbitmqURL




### <a name="Config.GetRabbitmqVHName">func</a> (Config) [GetRabbitmqVHName](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=14489:14533#L423)
``` go
func (cfg Config) GetRabbitmqVHName() string
```
GetRabbitmqVHName хост rabbitmq
Deprecated: use GetRabbitmqURL




### <a name="Config.GetRedisURL">func</a> (Config) [GetRedisURL](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=17542:17592#L535)
``` go
func (cfg Config) GetRedisURL() (*RedisURL, error)
```
GetRedisURL - стандартное подключение redis




### <a name="Config.GetScriptCompileTTL">func</a> (Config) [GetScriptCompileTTL](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=25336:25389#L756)
``` go
func (cfg Config) GetScriptCompileTTL() time.Duration
```
GetScriptCompileTTL ограничение на время транспиляции сценария




### <a name="Config.GetScriptExecuteTTL">func</a> (Config) [GetScriptExecuteTTL](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=25524:25577#L761)
``` go
func (cfg Config) GetScriptExecuteTTL() time.Duration
```
GetScriptExecuteTTL ограничение на время исполнения сценария




### <a name="Config.GetScriptQueueTTL">func</a> (Config) [GetScriptQueueTTL](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=25707:25758#L766)
``` go
func (cfg Config) GetScriptQueueTTL() time.Duration
```
GetScriptQueueTTL ограничения на время ожидания в очереди




### <a name="Config.GetServiceURL">func</a> (Config) [GetServiceURL](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=17678:17738#L540)
``` go
func (cfg Config) GetServiceURL(name string) (string, error)
```
GetServiceURL http-адрес сервиса




### <a name="Config.GetTimeout">func</a> (Config) [GetTimeout](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=10793:10837#L304)
``` go
func (cfg Config) GetTimeout() time.Duration
```
GetTimeout таймаут соединения для http-интерфейса




### <a name="Config.GetZone">func</a> (Config) [GetZone](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=16573:16607#L493)
``` go
func (cfg Config) GetZone() string
```
GetZone зона - ru/com/etc
Deprecated: Используйте GetHost(), кроме узкие случаев, когда нужно получить только последний уроаень домена




### <a name="Config.IsDebug">func</a> (Config) [IsDebug](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=9752:9784#L269)
``` go
func (cfg Config) IsDebug() bool
```
IsDebug возвращает первоначальную установку конфигурации Debug




### <a name="Config.Name">func</a> (Config) [Name](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=9219:9250#L254)
``` go
func (cfg Config) Name() string
```
Name возвращает имя сервиса




### <a name="Config.PostgresURL">func</a> (Config) [PostgresURL](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=11073:11120#L310)
``` go
func (cfg Config) PostgresURL() (string, error)
```
PostgresURL возвращает url для подключения к postgres
Deprecated:  использовать GetPostgresURL после полного перехода на connection string (URL)




### <a name="Config.RabbitmqURL">func</a> (Config) [RabbitmqURL](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=16879:16926#L507)
``` go
func (cfg Config) RabbitmqURL() (string, error)
```
RabbitmqURL возвращает url для подключения к rabbitmq
Deprecated: use GetRabbitmqURL




### <a name="Config.StdConfig">func</a> (Config) [StdConfig](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=9384:9420#L259)
``` go
func (cfg Config) StdConfig() Config
```
StdConfig возвращает сам объект конфигурации для наследования




## <a name="Custom">type</a> [Custom](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=7012:7070#L165)
``` go
type Custom interface {
    Interface
    SetStdConfig(Config)
}
```
Custom используется для интерфейса по указателю на расширение объекта конфигурации










## <a name="Interface">type</a> [Interface](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=5982:6030#L132)
``` go
type Interface interface {
    StdConfig() Config
}
```
Interface предназначен для создания расширения стандартного объекта конфигурации










## <a name="MongoURL">type</a> [MongoURL](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=6314:6337#L146)
``` go
type MongoURL = url.URL
```
MongoURL - alias










## <a name="PostgresURL">type</a> [PostgresURL](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=6266:6292#L143)
``` go
type PostgresURL = url.URL
```
PostgresURL - alias










## <a name="RedisURL">type</a> [RedisURL](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/config.go?s=6114:6137#L137)
``` go
type RedisURL = url.URL
```
RedisURL - alias добавлен для гибкости работы c Redis










## <a name="Solution">type</a> [Solution](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/solution.go?s=268:286#L12)
``` go
type Solution int8
```
Solution редакция конкретной инсталляции


``` go
const (
    // Saas в облаке
    Saas Solution = iota
    // Onpremise на сервере клиента
    Onpremise
)
```






### <a name="SolutionString">func</a> [SolutionString](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/solution_string.go?s=732:779#L30)
``` go
func SolutionString(s string) (Solution, error)
```
SolutionString retrieves an enum value from the enum constants string name.
Throws an error if the param is not part of the enum.


### <a name="SolutionValues">func</a> [SolutionValues](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/solution_string.go?s=969:1001#L38)
``` go
func SolutionValues() []Solution
```
SolutionValues returns all values of the enum





### <a name="Solution.IsASolution">func</a> (Solution) [IsASolution](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/solution_string.go?s=1126:1162#L43)
``` go
func (i Solution) IsASolution() bool
```
IsASolution returns "true" if the value is listed in the enum definition. "false" otherwise




### <a name="Solution.MarshalJSON">func</a> (Solution) [MarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/solution.go?s=474:521#L22)
``` go
func (i Solution) MarshalJSON() ([]byte, error)
```
MarshalJSON implements the json.Marshaler interface for Solution




### <a name="Solution.MarshalText">func</a> (Solution) [MarshalText](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/solution.go?s=971:1018#L39)
``` go
func (i Solution) MarshalText() ([]byte, error)
```
MarshalText implements the encoding.TextMarshaler interface for Solution




### <a name="Solution.String">func</a> (Solution) [String](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/solution_string.go?s=257:290#L14)
``` go
func (i Solution) String() string
```



### <a name="Solution.UnmarshalJSON">func</a> (\*Solution) [UnmarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/solution.go?s=632:683#L27)
``` go
func (i *Solution) UnmarshalJSON(data []byte) error
```
UnmarshalJSON implements the json.Unmarshaler interface for Solution




### <a name="Solution.UnmarshalText">func</a> (\*Solution) [UnmarshalText](https://git.elewise.com/elma365/common/-/tree/develop/pkg/config/solution.go?s=1136:1187#L44)
``` go
func (i *Solution) UnmarshalText(text []byte) error
```
UnmarshalText implements the encoding.TextUnmarshaler interface for Solution







- - -
Generated by [godoc2md](https://github.com/Exa-Networks/godoc2md)
