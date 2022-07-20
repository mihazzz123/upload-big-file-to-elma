package i18n

//go:generate mkdir -p mock
//go:generate ../../tooling/bin/minimock -g -s .go -o=mock -i DiskService,SettingsService

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"path/filepath"
	"sync"
	"text/template"
	"time"

	diskservice "git.elewise.com/elma365/common/pkg/diskservice/models"
	"git.elewise.com/elma365/common/pkg/errs"
	"git.elewise.com/elma365/common/pkg/i18n/models"
	sm "git.elewise.com/elma365/common/pkg/settings/models"
	"git.elewise.com/elma365/common/pkg/syncutils"

	"github.com/leonelquinteros/gotext"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

// Config Параметры конфигурации для i18n
type Config interface {
	GetDefaultLanguage() string
	GetAllowedLanguages() []string
}

// SettingsService интерфейс сервиса settings
type SettingsService interface {
	GetLanguageSettingsList(ctx context.Context) ([]sm.LanguageSettings, error)
	GetLanguageSettings(ctx context.Context, locale string) (sm.LanguageSettings, error)
}

// DiskService интерфейс сервиса
type DiskService interface {
	DownloadFileByBodyHash(ctx context.Context, bodyHash uuid.UUID, options ...diskservice.DownloadOptionsFunc) (body io.ReadCloser, size int64, err error)
	DownloadFileByFilePath(ctx context.Context, dirID uuid.UUID, path string) (body io.ReadCloser, size int64, err error)
}

type ctxKey struct{}

// ContextWithLang добавляет язык запроса в контекст
func ContextWithLang(ctx context.Context, lang string) context.Context {
	return context.WithValue(ctx, ctxKey{}, lang)
}

// LangFromContext извлекает язык из контекста, если языка нет в контексте - возвращает пустой
func LangFromContext(ctx context.Context) string {
	lang, _ := ctx.Value(ctxKey{}).(string)
	return lang
}

type companyLangCtxKey struct{}

// ContextWithCompanyLang устанавливает в контекст язык компании
func ContextWithCompanyLang(ctx context.Context, lang string) context.Context {
	return context.WithValue(ctx, companyLangCtxKey{}, lang)
}

// CompanyLangFromContext извлекает язык компании из контекста
func CompanyLangFromContext(ctx context.Context) string {
	lang, _ := ctx.Value(companyLangCtxKey{}).(string)
	return lang
}

// Option опции создания сервиса переводов
type Option interface {
	apply(*I18n)
}

type optionFunc func(*I18n)

func (f optionFunc) apply(i18n *I18n) {
	f(i18n)
}

// WithDomain читать переводы из файлов с переданным именем
func WithDomain(domain string) Option {
	return optionFunc(func(i18n *I18n) {
		i18n.domain = domain
	})
}

// WithCompanyLocales включает возможность использования локалей установленных с помощью языковых пакетов
func WithCompanyLocales(
	serviceName string,
	settingsService SettingsService,
	diskService DiskService,
	fnCompanyFromCtx func(ctx context.Context) string,
) Option {
	return optionFunc(func(i18n *I18n) {
		i18n.serviceName = serviceName
		i18n.diskService = diskService
		i18n.settingsService = settingsService
		i18n.fnCompanyFromCtx = fnCompanyFromCtx
	})
}

// I18n сервис переводов
type I18n struct {
	domain      string
	defaultLang string
	localesDir  string

	translators *models.Translators

	entityPOLoader            EntityPOLoader
	entityTranslators         *sync.Map
	entityTranslatorsCacheTTL time.Duration
	fnCompanyFromCtx          func(ctx context.Context) string

	serviceName              string
	diskService              DiskService
	settingsService          SettingsService
	languageSettingsCache    *sync.Map
	languageSettingsCacheTTL time.Duration

	mu      *sync.Mutex
	ttlRand *rand.Rand
}

// New загружает все доступные локали
//
// Для корректной работы загрузки локалей необходимо в корне проекта добавить папку data, внутри которой создать папку
// locales со следующей структурой:
// ```
// /data
// /data/locales
// /data/locales/ru-RU
// /data/locales/ru-RU/default.po
// ```
//
// Внимание! По умолчанию читаются только файлы default.po. Остальные файлы можно использовать для других целей.
//
// Для того, чтобы файлы добавились в docker-контейнер, необходимо прописать в Dockerfile следующее:
// ```
// ...
// COPY data /data
// ...
// WORKDIR "/"
// ...
// ```
// То есть надо скопировать папку data в корень контейнера, а также рабочей директорией установить корень.
//
// В качестве второго аргумента необходимо передать путь до папки locales.
//
// Для обобщения работы с переводами также рекомендуется отражать в ключах подстановки, то есть:
// ```
// msgid "some@key(%d)"
// msgstr "Сообщение №%d"
// ```
func New(config Config, localesDir string, opts ...Option) (I18n, error) {
	i18n := I18n{
		defaultLang: config.GetDefaultLanguage(),
		translators: new(models.Translators),
		localesDir:  localesDir,
		domain:      "default",

		entityTranslators:         new(sync.Map),
		entityTranslatorsCacheTTL: defaultEntityTranslatorsCacheTTL,

		languageSettingsCache:    new(sync.Map),
		languageSettingsCacheTTL: defaultLanguageSettingsCacheTTL,

		mu: &sync.Mutex{},
		//nolint:gosec,gochecknoglobals // тут не нужна криптографическая надежность и это приватная переменная
		ttlRand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}

	for _, opt := range opts {
		opt.apply(&i18n)
	}

	files, err := filepath.Glob(filepath.Join(localesDir, "*", fmt.Sprintf("%s.%s", i18n.domain, "po")))
	if err != nil {
		return I18n{}, errors.WithStack(err)
	}

	allowedLangsMap := make(map[string]interface{}, len(config.GetAllowedLanguages()))
	for _, lang := range config.GetAllowedLanguages() {
		allowedLangsMap[lang] = nil
	}

	for _, file := range files {
		buf, err := ioutil.ReadFile(file) //nolint:gosec // файл задается developer'ом
		if err != nil {
			return I18n{}, errors.WithStack(err)
		}
		lang := filepath.Base(filepath.Dir(file))
		if _, ok := allowedLangsMap[lang]; !ok {
			// Загружаем только разрешенные языки
			continue
		}
		translator := gotext.NewPoTranslator()
		translator.Parse(buf)
		i18n.translators.Store(lang, translator)
	}

	if _, loaded := i18n.translators.LoadOrStore(i18n.defaultLang, gotext.NewPoTranslator()); !loaded {
		zap.L().Warn("Not found translator for default language", zap.String("Error", fmt.Sprintf("%s resources is not found in %s", i18n.defaultLang, localesDir)))
		labels := prometheus.Labels{
			"language":   i18n.defaultLang,
			"localesDir": localesDir,
		}
		i18nTranslatorsResourceCounter.With(labels).Inc()
	}

	return i18n, nil
}

// LangFromContext извлекает язык из контекста, если языка нет в контексте или среди известных, возвращает язык по умолчанию
// ( зачем нужен этот метод когда в пакете "i18n" уже есть функция с точно таким же названием? )
func (i18n I18n) LangFromContext(ctx context.Context) string {
	lang, ok := ctx.Value(ctxKey{}).(string)
	if !ok || lang == "" {
		return i18n.defaultLang
	}

	if _, ok = i18n.translators.Load(lang); !ok {
		return i18n.defaultLang
	}
	return lang
}

// GetLangDir - возвращает папку, где хранятся локализованные ресурсы для текущего контекста
func (i18n I18n) GetLangDir(ctx context.Context) string {
	return filepath.Join(i18n.localesDir, i18n.LangFromContext(ctx))
}

// TranslateTemplate рендерит текстовый шаблон с функцией `translate`
//
// Внутри шаблона можно использовать конструкции вида
// ```
// {{ translate `some@key(%d)` .Count }}
// ```
func (i18n I18n) TranslateTemplate(ctx context.Context, src string, data interface{}) (string, error) {
	base := template.New("").
		Funcs(template.FuncMap{
			"translate": func(str string, vars ...interface{}) string {
				return i18n.TranslateString(ctx, str, vars...)
			},
		})
	tmpl, err := base.Parse(src)
	if err != nil {
		return "", errors.WithStack(err)
	}
	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, data); err != nil {
		return "", errors.WithStack(err)
	}

	return buf.String(), nil
}

// TranslateString локализует ключ
func (i18n I18n) TranslateString(ctx context.Context, str string, args ...interface{}) string {
	translator, err := i18n.getTranslator(ctx)
	if err != nil {
		zap.L().Error("translate string", zap.Error(err))
		return str
	}
	return translator.Get(str, args...)
}

// IsExistTranslate проверка существование ключа
func (i18n I18n) IsExistTranslate(ctx context.Context, str string) bool {
	return str != i18n.TranslateString(ctx, str)
}

// GetAvailableLangs возвращает коды доступных языков
func (i18n I18n) GetAvailableLangs() []string {
	langs := make([]string, 0)
	i18n.translators.Range(func(key string, value gotext.Translator) bool {
		langs = append(langs, key)
		return true
	})
	return langs
}

func (i18n I18n) getTranslator(ctx context.Context) (gotext.Translator, error) {
	lang := LangFromContext(ctx)
	if lang != "" {
		if translator, ok := i18n.translators.Load(lang); ok {
			return translator, nil
		}
		translator, err := i18n.tryGetTranslator(ctx, lang)
		if err == nil {
			return translator, nil
		}
		switch errors.Cause(err) {
		case errs.Precondition, errs.NotFound:
			zap.L().Warn("get translator", zap.Error(err))
		default:
			return nil, err
		}
	}

	lang = CompanyLangFromContext(ctx)
	if translator, ok := i18n.translators.Load(lang); ok {
		return translator, nil
	}

	lang = i18n.defaultLang
	if translator, ok := i18n.translators.Load(lang); ok {
		return translator, nil
	}
	return nil, errs.NotFound.New("translator not found")
}

func (i18n I18n) tryGetTranslator(ctx context.Context, lang string) (gotext.Translator, error) {
	hash, err := i18n.getOrLoadLanguageHash(ctx, lang)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if translator, ok := i18n.translators.Load(hash.String()); ok {
		return translator, nil
	}

	return i18n.loadTranslator(ctx, hash)
}

func (i18n I18n) getOrLoadLanguageHash(ctx context.Context, locale string) (uuid.UUID, error) {
	var hash uuid.UUID
	if i18n.fnCompanyFromCtx == nil {
		return hash, errs.Precondition.New(`function "fnCompanyFromCtx" is undefined`)
	}

	ttl := i18n.getInexactTTL(defaultLanguageSettingsCacheTTL)
	key := fmt.Sprintf("%s/%s", i18n.fnCompanyFromCtx(ctx), locale)
	value, _ := i18n.languageSettingsCache.LoadOrStore(
		key,
		syncutils.NewPromise(
			func(ctx context.Context) (interface{}, error) {
				return i18n.getLanguageHash(ctx, locale)
			},
			syncutils.WithTTL(ttl),
		),
	)

	promise, ok := value.(*syncutils.Promise)
	if !ok {
		return hash, errors.WithStack(errors.New("fail to get language settings"))
	}
	res, err := promise.Await(ctx)
	if err != nil {
		return hash, errors.WithStack(err)
	}
	hash, ok = res.(uuid.UUID)
	if !ok {
		return hash, errors.WithStack(errors.New("fail to get language settings"))
	}
	return hash, nil
}

func (i18n I18n) getLanguageHash(ctx context.Context, locale string) (uuid.UUID, error) {
	if i18n.settingsService == nil {
		return uuid.Nil, errs.Precondition.New("settings services not inited")
	}
	if i18n.serviceName == "" {
		return uuid.Nil, errs.Precondition.New("service name is empty")
	}

	languageSettings, err := i18n.settingsService.GetLanguageSettings(ctx, locale)
	if err != nil {
		return uuid.Nil, errors.WithStack(err)
	}
	for _, poFile := range languageSettings.POfiles {
		if poFile.ServiceName == i18n.serviceName {
			return poFile.Hash, nil
		}
	}
	return uuid.Nil, errs.NotFound.Newf("not found lang hash for locale \"%s\"", locale)
}

func (i18n I18n) loadTranslator(ctx context.Context, hash uuid.UUID) (gotext.Translator, error) {
	if i18n.diskService == nil {
		return nil, errs.Precondition.New("disk service not inited")
	}
	body, _, err := i18n.diskService.DownloadFileByBodyHash(ctx, hash)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = body.Close()
	}()

	buf, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}

	newTranslator := gotext.NewPoTranslator()
	newTranslator.Parse(buf)

	i18n.translators.Store(hash.String(), newTranslator)

	return newTranslator, nil
}

// возвращает ttl со случайным смещением +/- 25% от базового значения
// это нужно чтобы более или менее размазать протухание кэша по времени
func (i18n I18n) getInexactTTL(ttlValue time.Duration) time.Duration {
	n25 := int(float64(ttlValue) / 4)
	i18n.mu.Lock()
	defer i18n.mu.Unlock()
	d := (i18n.ttlRand.Intn(n25) - n25/2) * 2
	return ttlValue + time.Duration(d)
}
