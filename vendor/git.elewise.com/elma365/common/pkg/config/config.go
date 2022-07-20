package config

import (
	"context"
	"encoding/base64"
	"fmt"
	"net"
	"net/url"
	"os"
	"strings"
	"time"

	"git.elewise.com/elma365/common/pkg/connection/constant"
	"git.elewise.com/elma365/common/pkg/edition"
	"git.elewise.com/elma365/common/pkg/envcfg"

	"github.com/davecgh/go-spew/spew" //nolint:depguard // используется для дампа, в продакшене не включен
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/miekg/dns"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/uber/jaeger-client-go"
	jConfig "github.com/uber/jaeger-client-go/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zapgrpc"
	"google.golang.org/grpc/grpclog"
)

// Prefix префикс переменных окружения
const Prefix = "ELMA365_"

// Config of application
type Config struct {
	name  string
	level zap.AtomicLevel

	// Deprecated: Будет удален. Использовать LogLevel.
	Debug        bool
	LogLevel     string `default:"info"` // Уровень логирования в случае Debug=false
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
	PostgresMaxOpenConns int `default:"10"`
	// PostgresMaxIdleConns maximum number of connections in the idle connection pool
	PostgresMaxIdleConns int `default:"-1"`
	// PostgresConnMaxLifetime maximum amount of time a connection may be reused
	PostgresConnMaxLifetime time.Duration `default:"1m"`
	// PostgresConnectionTimeout максимальное время установки соединения с базой данных.
	PostgresConnectionTimeout time.Duration `default:"30s"`

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

	PsqlURL string // default:"postgresql://postgres:postgres@postgres:5432/elma365"

	// URL подключения к MongoDB
	MongoURL string // default:"mongodb://mongo:mongo@mongo:27017/elma365?connectTimeoutMS=10000&serverSelectionTimeoutMS=5000&heartbeatFrequencyMS=30000&authSource=elma365&maxPoolSize=100"
	// MongoDisableTracing отключает трассировку запросов в MongoDB
	MongoDisableTracing bool `default:"false"`

	// URL подключения к redis
	RedisURL string // default:"redis://localhost:6379?ConnectTimeout=5000&IdleTimeOutSecs=180"
	// RedisHealthCheckPeriod период проверки доступности кластера Redis.
	RedisHealthCheckPeriod time.Duration `default:"1m"`
	// Максимальный размер используемой памяти вторым уровнем кеша в байтах. Значение по умолчанию 128 Мб.
	SecondLevelCacheMaxMemoryUsage int64 `default:"134217728"`
	// RedisPubSubPoolSize размер пула подключений для подписок Redis.
	RedisPubSubPoolSize             int64  `default:"100"`
	AmqpURL                         string // default:"amqp://rabbitmq:rabbitmq@host:5672/vhost"
	ClusterRegistryConnectionString string // default:"postgresql://postgres:postgres@postgres:5432/elma365"
	// MultiClusterEnable флаг включения мультикластера.
	MultiClusterEnable bool `default:"false"`

	// PublicDBConnectionString строка подключения к БД содержащей схему "public"
	PublicDBConnectionString string // Если не задано, то используется PsqlURL
	// PublicDBMinOpenConnections минимальное количество одновременно открытых подключений из пула.
	PublicDBMinOpenConnections int32 `default:"1"`
	// PublicDBMaxOpenConnections максимальное количество одновременно открытых подключений из пула.
	PublicDBMaxOpenConnections int32 `default:"10"`
	// PublicDbConnectionTimeout максимальное время установки соединения с базой данных.
	PublicDbConnectionTimeout time.Duration `default:"30s"`
	// PublicDBMaxConnectionLifetime максимальное время жизни подключения из пула.
	PublicDBMaxConnectionLifetime time.Duration `default:"1m"`

	dnsConfig *dns.ClientConfig

	// Edition редакция конкретной инсталляции (в инсталяции у компаний могут быть разные Edition)
	// Внимание!!! Для разграничения функционала и получения Edition конкретной компании использовать md.EditionFromContext(ctx)
	Edition edition.Edition `default:"lite"` // lite - для обратной совместимости
	// Solution Версия сборки
	Solution Solution `default:"saas"`

	// MaxGRPCMessageSize максимальный размер ответов
	MaxGRPCMessageSize int `default:"8388608"` // 1024 * 1024 * 4 дефолтное значение GRPC
	// MaxGRPCBackoffDelay максимальная задержка между попытками установки GRPC соединения
	MaxGRPCBackoffDelay time.Duration `default:"2m"` // дефолтное значение GRPC
	// MaxGRPCConnectionsCount максимальное количество одновременных подключений к сервису.
	MaxGRPCConnectionsCount int `default:"10"`
	// MaxGRPCConcurrentStreamsPerConnection максимальное количество потоков данных на одно подключений к сервису.
	MaxGRPCConcurrentStreamsPerConnection int `default:"100"`

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

	// DisableCommonCache отключать использование кеша
	DisableCommonCache bool `default:"false"`
}

// Interface предназначен для создания расширения стандартного объекта конфигурации
type Interface interface {
	StdConfig() Config
}

// RedisURL - alias добавлен для гибкости работы c Redis
type RedisURL = url.URL

// AmqpURL - alias добавлен для гибкости работы c Amqp
type AmqpURL = url.URL

// PostgresURL - alias
type PostgresURL = url.URL

// MongoURL - alias
type MongoURL = url.URL

// New создать новый конфиг
//
// Создаёт новый конфиг и вычитывает его из переменных окружения. После этого настраивает глобальные логгеры
// и трейсер.
func New(name string) (Config, error) {
	cfg, err := makeStd(name)
	if err != nil {
		return cfg, err
	}
	if err := envcfg.Read(&cfg, envcfg.WithPrefix(Prefix), envcfg.WithDefault(nil)); err != nil {
		return cfg, err
	}

	return cfg, cfg.init()
}

// Custom используется для интерфейса по указателю на расширение объекта конфигурации
type Custom interface {
	Interface
	SetStdConfig(Config)
}

// NewCustom инициализирует расширение объекта конфигурации
func NewCustom(name string, cfg Custom) error {
	std, err := makeStd(name)
	if err != nil {
		return err
	}
	cfg.SetStdConfig(std)

	type Overrider interface {
		GetStdDefaults() map[string]string
	}
	var overrides map[string]string
	if overrider, ok := cfg.(Overrider); ok {
		overrides = overrider.GetStdDefaults()
	}
	if err := envcfg.Read(cfg, envcfg.WithPrefix(Prefix), envcfg.WithDefault(overrides)); err != nil {
		return err
	}

	return cfg.StdConfig().init()
}

func makeStd(name string) (Config, error) {
	cfg := Config{
		name:  name,
		level: zap.NewAtomicLevel(),
	}
	var err error
	cfg.dnsConfig, err = dns.ClientConfigFromFile("/etc/resolv.conf")

	return cfg, errors.WithStack(err)
}

func (cfg Config) init() error {
	cfg.replaceLoggers()

	return cfg.setupJaeger()
}

func (cfg Config) replaceLoggers() {
	// Человеческие временные метки
	enc := zap.NewProductionEncoderConfig()
	enc.TimeKey = "timestamp"
	enc.EncodeTime = zapcore.ISO8601TimeEncoder
	opts := []zap.Option{
		zap.AddCaller(),
	}

	// Уровень логирования
	level := cfg.getLogLevel()
	if level == zap.DebugLevel {
		opts = append(opts, zap.Development())
	}
	cfg.level.SetLevel(level)

	l := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(enc),
		zapcore.Lock(os.Stderr),
		cfg.level,
	), opts...)
	l = l.Named(cfg.FullName())

	zap.ReplaceGlobals(l)
	//nolint // ждём когда реализуют функцилнальность https://github.com/uber-go/zap/pull/538
	grpclog.SetLogger(zapgrpc.NewLogger(zap.L(), zapgrpc.WithDebug()))
}

func (cfg Config) setupJaeger() error {
	// 5775 - порт для UDP agent-zipkin-thrift по умолчанию
	jaddr := fmt.Sprintf("%s:%d", cfg.HostIP, 5775)
	jCfg := jConfig.Configuration{
		Reporter: &jConfig.ReporterConfig{
			LocalAgentHostPort: jaddr,
		},
		Sampler: &jConfig.SamplerConfig{
			Type:  "const",
			Param: 1.0, // sample all traces
		},
	}
	_, err := jCfg.InitGlobalTracer(cfg.Name())

	return errors.WithStack(err)
}

// Name возвращает имя сервиса
func (cfg Config) Name() string {
	return cfg.name
}

// StdConfig возвращает сам объект конфигурации для наследования
func (cfg Config) StdConfig() Config {
	return cfg
}

// FullName возвращает имя сервиса с префиксом приложения (elma365.processor)
func (cfg Config) FullName() string {
	return fmt.Sprintf("elma365.%s", cfg.name)
}

// IsDebug возвращает первоначальную установку конфигурации Debug
func (cfg Config) IsDebug() bool {
	return cfg.getLogLevel() == zap.DebugLevel
}

// EnableDebug включает логирование на уровне debug
func (cfg Config) EnableDebug() {
	cfg.level.SetLevel(zap.DebugLevel)
}

// DisableDebug включает логирование с уровнем по умолчанию
func (cfg Config) DisableDebug() {
	level := cfg.getLogLevel()
	cfg.level.SetLevel(level)
}

// DisableDebug выключает логирование на уровне debug
func (cfg Config) getLogLevel() zapcore.Level {
	if cfg.Debug {
		return zap.DebugLevel
	}

	level := zap.InfoLevel
	if cfg.LogLevel != "" {
		err := level.Set(cfg.LogLevel)
		if err != nil {
			fmt.Printf("invalid log level %q: %s\n", cfg.LogLevel, err)
			// Откат к дефолтному значению
			level = zap.InfoLevel
		}
	}

	return level
}

// GetBind интерфейс для запуска сервера
func (cfg Config) GetBind() string {
	return cfg.Bind
}

// GetCMUXBind использовать механизм шаринга порта между gRPC и HTTP сервисами
func (cfg Config) GetCMUXBind() bool {
	return cfg.CMUXBind
}

// GetGRPCBind интерфейс для запуска grpc сервера
func (cfg Config) GetGRPCBind() string {
	return cfg.GRPCBind
}

// GetHTTPBind интерфейс для запуска http сервера
func (cfg Config) GetHTTPBind() string {
	return cfg.HTTPBind
}

// GetTimeout таймаут соединения для http-интерфейса
func (cfg Config) GetTimeout() time.Duration {
	return cfg.Timeout
}

// PostgresURL возвращает url для подключения к postgres
// Deprecated:  использовать GetPostgresURL после полного перехода на connection string (URL)
func (cfg Config) PostgresURL() (string, error) {
	if cfg.PsqlURL == "" {
		addr, err := cfg.DNSResolve(cfg.PostgresPortName, "tcp", cfg.PostgresSvcName)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf(
			"postgres://%s:%s@%s/%s?sslmode=disable&application_name=%q",
			cfg.PostgresUser, cfg.PostgresPassword, addr, cfg.PostgresDBName, cfg.Name(),
		), nil
	}

	return cfg.PsqlURL, nil
}

// GetPostgresURL - Распарсенная строка подключения к Postgres
// Добавляет в параметры подключения имя приложения
func (cfg Config) GetPostgresURL() (*PostgresURL, error) {
	const applicationName = "application_name"
	psqlURL, err := cfg.PostgresURL()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	pURL, err := url.Parse(psqlURL)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	values := pURL.Query()
	if !values.Has(applicationName) {
		values.Set(applicationName, cfg.Name())
		pURL.RawQuery = values.Encode()
	}

	return pURL, nil
}

// GetClusterRegistryConnectionString возвращает строку подключения к реестру арендаторов и их вычислительных кластеров.
// Если строка подключения в конфигурации не указана, то возвращается строка подключения до основной БД.
func (cfg Config) GetClusterRegistryConnectionString() string {
	if cfg.ClusterRegistryConnectionString != "" {
		return cfg.ClusterRegistryConnectionString
	}

	connStr, err := cfg.GetPostgresURL()
	if err != nil {
		panic(err)
	}

	return connStr.String()
}

// IsClusterRegistryEnabled проверяет включен ли реестр арендаторов и их вычислительных кластеров.
func (cfg Config) IsClusterRegistryEnabled() bool {
	return cfg.MultiClusterEnable
}

// GetPublicDbConnectionString получает строку подключения к базе данных содержащей схему "public".
func (cfg Config) GetPublicDbConnectionString() string {
	if cfg.PublicDBConnectionString != "" {
		return cfg.PublicDBConnectionString
	}

	connStr, err := cfg.GetPostgresURL()
	if err != nil {
		panic(err)
	}

	return connStr.String()
}

// GetPublicDbMinOpenConnections получает минимальное количество одновременно открытых подключений.
func (cfg Config) GetPublicDbMinOpenConnections() int32 {
	return cfg.PublicDBMinOpenConnections
}

// GetPublicDbMaxOpenConnections получает максимальное количество одновременно открытых подключений.
func (cfg Config) GetPublicDbMaxOpenConnections() int32 {
	return cfg.PublicDBMaxOpenConnections
}

// GetPublicDbConnectionTimeout получает максимальное время установки соединения с базой данных.
func (cfg Config) GetPublicDbConnectionTimeout() time.Duration {
	return cfg.PublicDbConnectionTimeout
}

// GetPublicDbMaxConnectionLifetime получает максимальное время в течении которого неиспользуемое подключение может быть переиспользовано.
func (cfg Config) GetPublicDbMaxConnectionLifetime() time.Duration {
	return cfg.PublicDBMaxConnectionLifetime
}

// GetMongoURL - Распарсенная строка подключения к Mongo
// Добавляет/изменяет в параметрах подключения имя приложения на имя конфигурации.
// Параметры по умолчанию:
// connectTimeoutMS          10_000
// serverSelectionTimeoutMS  5_000
// heartbeatFrequencyMS      30_000
// authSource                такой же как и имя базы данных, если имя базы данных указана
// maxPoolSize               100
func (cfg Config) GetMongoURL() (*MongoURL, error) {
	mURL, err := url.Parse(cfg.MongoURL)
	if err != nil {
		return nil, err
	}

	if mURL.Host == "" {
		mURL, err = cfg.buildDefaultMongoURL()
		if err != nil {
			return nil, err
		}
	}

	defaultValues := map[string]string{
		"connectTimeoutMS":         "10000",
		"serverSelectionTimeoutMS": "5000",
		"heartbeatFrequencyMS":     "30000",
		"maxPoolSize":              "100",
		"authSource":               mURL.Path[1:],
	}

	values := mURL.Query()
	values.Set("appName", cfg.Name())

	for k, v := range defaultValues {
		if values.Get(k) == "" {
			values.Set(k, v)
		}
	}

	mURL.RawQuery = values.Encode()
	return mURL, nil
}

// GetMongoDisableTracing возвращает признак отключения трассировки запросов в MongoDB
func (cfg Config) GetMongoDisableTracing() bool {
	return cfg.MongoDisableTracing
}

// buildDefaultMongoURL - получить урл для коннекта к монге
func (cfg Config) buildDefaultMongoURL() (*MongoURL, error) {
	addrs, err := cfg.DNSResolveMulti("mongo", "tcp", "mongo")
	if err != nil {
		return nil, errors.WithStack(err)
	}

	mongoURL, err := url.Parse(fmt.Sprintf("mongodb://%s:%s@%s/%s", cfg.GetMongoUser(), cfg.GetMongoPassword(), strings.Join(addrs, ","), cfg.GetMongoDBName()))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return mongoURL, nil
}

// GetPostgresMaxOpenConns количество подключений к каждой реплике
func (cfg Config) GetPostgresMaxOpenConns() int {
	return cfg.PostgresMaxOpenConns
}

// GetPostgresConnectionTimeout получает максимальное время установки соединения с базой данных.
func (cfg Config) GetPostgresConnectionTimeout() time.Duration {
	return cfg.PostgresConnectionTimeout
}

// GetPostgresMaxIdleConns maximum number of connections in the idle connection pool
func (cfg Config) GetPostgresMaxIdleConns() int {
	return cfg.PostgresMaxIdleConns
}

// GetPostgresConnMaxLifetime maximum amount of time a connection may be reused
func (cfg Config) GetPostgresConnMaxLifetime() time.Duration {
	return cfg.PostgresConnMaxLifetime
}

// GetRabbitmqUser имя пользователя rabbitmq
// Deprecated: use GetRabbitmqURL
func (cfg Config) GetRabbitmqUser() string {
	return cfg.RabbitmqUser
}

// GetRabbitmqPassword пароль rabbitmq
// Deprecated: use GetRabbitmqURL
func (cfg Config) GetRabbitmqPassword() string {
	return cfg.RabbitmqPassword
}

// GetRabbitmqVHName хост rabbitmq
// Deprecated: use GetRabbitmqURL
func (cfg Config) GetRabbitmqVHName() string {
	return cfg.RabbitmqVHName
}

// GetMongoUser имя пользователя mongoDB
// Deprecated: use GetMongoURL
func (cfg Config) GetMongoUser() string {
	return cfg.MongoUser
}

// GetMongoPassword пароль пользователя mongoDB
// Deprecated: use GetMongoURL
func (cfg Config) GetMongoPassword() string {
	return cfg.MongoPassword
}

// GetMongoDBName название коллекции в mongoDB
// Deprecated: use GetMongoURL
func (cfg Config) GetMongoDBName() string {
	return cfg.MongoDBName
}

// GetMongoMaxPoolSize максимальное количество соединений к mongoDB
// Deprecated: use GetMongoURL
func (cfg Config) GetMongoMaxPoolSize() uint64 {
	return cfg.MongoMaxPoolSize
}

// GetEventBusMaxHashCount максимальное количество возможных значений поля hash
// в RoutingKey для шины событий
func (cfg Config) GetEventBusMaxHashCount() int {
	return cfg.EventBusMaxHashCount
}

// GetDefaultLanguage Язык по умолчанию
func (cfg Config) GetDefaultLanguage() string {
	return cfg.DefaultLanguage
}

// GetAllowedLanguages Список доступных языков
func (cfg Config) GetAllowedLanguages() []string {
	return cfg.AllowedLanguages
}

// GetHost получить Host (домен), на котором запущена система
func (cfg Config) GetHost() string {
	if cfg.Host != "" {
		return cfg.Host
	}

	return cfg.Domain + "." + cfg.Zone
}

// GetDomain домен, за который отвечает сервис
// Deprecated: используйте GetHost()
func (cfg Config) GetDomain() string {
	if cfg.Host != "" {
		index := strings.LastIndex(cfg.Host, ".")
		if index == -1 {
			return cfg.Host
		}

		return cfg.Host[:index]
	}

	return cfg.Domain
}

// GetZone зона - ru/com/etc
// Deprecated: Используйте GetHost(), кроме узкие случаев, когда нужно получить только последний уроаень домена
func (cfg Config) GetZone() string {
	if cfg.Host != "" {
		index := strings.LastIndex(cfg.Host, ".")
		if index == -1 {
			return ""
		}
		return cfg.Host[index+1:]
	}

	return cfg.Zone
}

// RabbitmqURL возвращает url для подключения к rabbitmq
// Deprecated: use GetRabbitmqURL
func (cfg Config) RabbitmqURL() (string, error) {
	if cfg.AmqpURL != "" {
		return cfg.AmqpURL, nil
	}

	addr, err := cfg.DNSResolve("amqp", "tcp", "rabbitmq")
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"amqp://%s:%s@%s/%s",
		cfg.RabbitmqUser, cfg.RabbitmqPassword, addr, cfg.RabbitmqVHName,
	), nil
}

// GetRabbitmqURL - получение урл для подключения из конфига к rabbitmq
func (cfg Config) GetRabbitmqURL() (*AmqpURL, error) {
	// nolint
	rURL, err := cfg.RabbitmqURL()
	if err != nil {
		return nil, err
	}

	return url.Parse(rURL)
}

// GetRedisURL - стандартное подключение redis
func (cfg Config) GetRedisURL() (*RedisURL, error) {
	return url.Parse(cfg.RedisURL)
}

// GetServiceURL http-адрес сервиса
func (cfg Config) GetServiceURL(name string) (string, error) {
	addr, err := cfg.DNSResolve("http", "tcp", name)
	return "http://" + addr, err
}

// DNSResolve возвращает резолв адреса в формате Host:Port, Сначала ищет в переменных окружения,
// если не нашел резолвит сначала CNAME. Далее, пытается собрать адрес для известных протоколов. Если не получилось,
// резолвит SRV запись для переданного сервиса. Возвращает только 1 адрес.
func (cfg Config) DNSResolve(service, proto, name string) (string, error) {
	if cfg.Telepresence != "" {
		envAddrs, ok := cfg.addrFromEnv(name, service)
		if ok {
			return envAddrs[0], nil
		}
	}

	// Сформируем url сервиса без использования resolve механизма...
	addr, cname := cfg.dnsCompose(name, service)
	if addr != "" {
		zap.L().Debug("Single DNS composed", zap.String("result", addr))
		return addr, nil
	}

	// ...если не получилось, работаем по старой схеме
	addrs, err := cfg.dnsResolve(service, proto, name, cname)
	if err != nil {
		zap.L().Debug(fmt.Sprintf("DNS lookup for %s failed, fallback to environment variables lookup", name))
		envAddrs, ok := cfg.addrFromEnv(name, service)
		if ok {
			zap.L().Debug(fmt.Sprintf("ENV lookup for %s succeeded", name), zap.String("result", envAddrs[0]))
			return envAddrs[0], nil
		}
		zap.L().Debug(fmt.Sprintf("ENV lookup for %s failed - either %s or %s env variable not found", name, strings.ToUpper(strings.ReplaceAll(fmt.Sprintf("%s_SERVICE_HOST", name), "-", "_")), strings.ToUpper(strings.ReplaceAll(fmt.Sprintf("%s_SERVICE_PORT_%s", name, service), "-", "_"))))
		return "", err
	}

	addr = fmt.Sprintf("%s:%d", name, addrs[0].Port)
	zap.L().Debug("Single DNS resolved", zap.String("result", addr))

	return addr, nil
}

// DNSResolveMulti тоже что DNSResolve, но возвращает весь массив отрезолвленных адресов
func (cfg Config) DNSResolveMulti(service, proto, name string) ([]string, error) {
	if cfg.Telepresence != "" {
		envAddrs, ok := cfg.addrFromEnv(name, service)
		if ok {
			return envAddrs, nil
		}
	}

	addrs, err := cfg.dnsResolve(service, proto, name, "")
	if err != nil {
		zap.L().Debug(fmt.Sprintf("DNS lookup for %s failed, fallback to environment variables lookup", name))
		envAddrs, ok := cfg.addrFromEnv(name, service)
		if ok {
			zap.L().Debug(fmt.Sprintf("ENV lookup for %s succeeded", name), zap.String("result", strings.Join(envAddrs, ",")))
			return envAddrs, nil
		}
		zap.L().Debug(fmt.Sprintf("ENV lookup for %s failed - either %s or %s env variable not found", name, strings.ToUpper(strings.ReplaceAll(fmt.Sprintf("%s_SERVICE_HOST", name), "-", "_")), strings.ToUpper(strings.ReplaceAll(fmt.Sprintf("%s_SERVICE_PORT_%s", name, service), "-", "_"))))
		return nil, err
	}

	addrStr := make([]string, 0, len(addrs))
	for _, ad := range addrs {
		addrStr = append(addrStr, fmt.Sprintf("%s:%d", ad.Target, ad.Port))
	}

	zap.L().Debug("Multi DNS resolved", zap.Strings("result", addrStr))

	return addrStr, nil
}

func (cfg Config) dnsResolve(service, proto, name, cname string) ([]*net.SRV, error) {
	zap.L().Debug("resolving", zap.String("service", service), zap.String("proto", proto), zap.String("name", name))

	if cname == "" {
		var err error
		cname, err = cfg.resolveCname(name)
		if err != nil {
			// если не нашлось CNAME, используем оригинальное имя
			cname = name
		}
	}

	srvCname, addrs, err := net.LookupSRV(service, proto, cname)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if len(addrs) == 0 {
		return nil, errors.Errorf("SRV Lookup for %q service not found", name)
	}

	zap.L().Debug("address resolved",
		zap.String("service", service),
		zap.String("proto", proto),
		zap.String("name", name),
		zap.String("CNAME", cname),
		zap.String("SRV_CNAME", srvCname),
	)

	return addrs, nil
}

func (cfg Config) addrFromEnv(name, service string) ([]string, bool) {
	name = strings.ToUpper(strings.ReplaceAll(name, "-", "_"))
	host, ok := os.LookupEnv(fmt.Sprintf("%s_SERVICE_HOST", name))
	if !ok {
		return nil, false
	}
	portName := strings.ToUpper(strings.ReplaceAll(service, "_", ""))
	port, ok := os.LookupEnv(fmt.Sprintf("%s_SERVICE_PORT_%s", name, strings.ToUpper(portName)))
	if !ok {
		return nil, false
	}

	return []string{net.JoinHostPort(host, port)}, true
}

func (cfg Config) resolveCname(name string) (string, error) {
	cli := new(dns.Client)
	m := new(dns.Msg)
	m.RecursionDesired = true
	for _, searchLine := range cfg.dnsConfig.Search {
		m.SetQuestion(dns.Fqdn(strings.Join([]string{name, searchLine}, ".")), dns.TypeCNAME)
		for _, serv := range cfg.dnsConfig.Servers {
			r, _, err := cli.Exchange(m, fmt.Sprintf("%s:53", serv))
			if err != nil {
				continue
			}
			if len(r.Answer) == 0 {
				continue
			}

			if res, ok := r.Answer[0].(*dns.CNAME); ok {
				return res.Target, nil
			}
		}
	}

	return "", errors.Errorf("No CNAME found for %q", name)
}

// dnsCompose формирует url сервиса без использования resolve механизма
func (cfg Config) dnsCompose(name, service string) (addr, cname string) {
	if !cfg.EnableDNSResolveOptimization {
		return "", ""
	}

	cname, err := cfg.resolveCname(name)
	if err != nil {
		// если не нашлось CNAME, используем оригинальное имя
		cname = name
	}

	switch strings.ToLower(service) {
	case "http":
		addr = fmt.Sprintf("%s%s", cname, cfg.HTTPBind)
	case "grpc":
		addr = fmt.Sprintf("%s%s", cname, cfg.GRPCBind)
	case "redis":
		// Использется при прямом соединении, в случае отсутствия сентинелей и RedisURL
		addr = fmt.Sprintf("%s:%d", cname, constant.RedisDefaultPort)
	case "postgres":
		// Не использует dnsResolve. Подключается через sqlx клиент.
	case "mongo":
		// Не использует dnsResolve. Использует DNSResolveMulti.
	case "rabbit":
		// Не использует dnsResolve. Подключается через amqp клиент.
	}

	return addr, cname
}

// Dialer return closured service dialer
func (cfg Config) Dialer(service, proto, name string) func() (net.Conn, error) {
	return func() (net.Conn, error) {
		addr, err := cfg.DNSResolve(service, proto, name)
		if err != nil {
			return nil, err
		}
		cn, err := net.Dial(proto, addr)

		return cn, errors.WithStack(err)
	}
}

// GRPCTimeoutDialer is a method to pass it in grpc.Dial options
func (cfg Config) GRPCTimeoutDialer(service string, timeout time.Duration) (net.Conn, error) {
	addr, err := cfg.DNSResolve("grpc", "tcp", service)
	if err != nil {
		return nil, err
	}
	cn, err := net.DialTimeout("tcp", addr, timeout)

	return cn, errors.WithStack(err)
}

// GRPCContextDialer is a method to pass it in grpc.Dial options
func (cfg Config) GRPCContextDialer(ctx context.Context, service string) (net.Conn, error) {
	addr, err := cfg.DNSResolve("grpc", "tcp", service)
	if err != nil {
		return nil, err
	}
	cn, err := (&net.Dialer{}).DialContext(ctx, "tcp", addr)

	return cn, errors.WithStack(err)
}

// GetMaxGRPCMessageSize получить максимальный размер ответов
func (cfg Config) GetMaxGRPCMessageSize() int {
	return cfg.MaxGRPCMessageSize
}

// GetMaxGRPCBackoffDelay получить максимальную задержку между попытками установки GRPC соединения
func (cfg Config) GetMaxGRPCBackoffDelay() time.Duration {
	return cfg.MaxGRPCBackoffDelay
}

// GetMaxConnectionCount возвращает максимальное количество подключений к сервису.
func (cfg Config) GetMaxConnectionCount() int {
	return cfg.MaxGRPCConnectionsCount
}

// GetMaxConcurrentStreamsPerConnection возвращает максимальное количество потоков данных на одно подключение к сервису.
func (cfg Config) GetMaxConcurrentStreamsPerConnection() int {
	return cfg.MaxGRPCConcurrentStreamsPerConnection
}

// GetScriptCompileTTL ограничение на время транспиляции сценария
func (cfg Config) GetScriptCompileTTL() time.Duration {
	return cfg.ScriptCompileTTL
}

// GetScriptExecuteTTL ограничение на время исполнения сценария
func (cfg Config) GetScriptExecuteTTL() time.Duration {
	return cfg.ScriptExecuteTTL
}

// GetScriptQueueTTL ограничения на время ожидания в очереди
func (cfg Config) GetScriptQueueTTL() time.Duration {
	return cfg.ScriptQueueTTL
}

// Dump если включена соответствующая опция, то пишет дамп объекта в base64 в лог
func (cfg Config) Dump(ctx context.Context, name string, obj ...interface{}) {
	if !cfg.EnableDump || !cfg.level.Enabled(zap.DebugLevel) {
		return
	}
	log := zap.L().Named(name)
	log = LoggerWithTraceID(ctx, log)
	dump := spew.Sdump(obj...)
	b64 := base64.StdEncoding.EncodeToString([]byte(dump))
	log.Debug(b64)
}

// TraceIDFromContext extract trace and span ids if span started or empty string otherwise
func TraceIDFromContext(ctx context.Context) (traceID, spanID string) {
	if sp := opentracing.SpanFromContext(ctx); sp != nil {
		if sc, ok := sp.Context().(jaeger.SpanContext); ok {
			return sc.TraceID().String(), sc.SpanID().String()
		}
	}
	return "", ""
}

// LoggerWithTraceID добавляет поля TraceID и SpanID к полям логгера
func LoggerWithTraceID(ctx context.Context, logger *zap.Logger) *zap.Logger {
	traceID, spanID := TraceIDFromContext(ctx)
	if traceID != "" {
		logger = logger.With(zap.String(grpc_opentracing.TagTraceId, traceID))
	}
	if spanID != "" {
		logger = logger.With(zap.String(grpc_opentracing.TagSpanId, spanID))
	}

	return logger
}
