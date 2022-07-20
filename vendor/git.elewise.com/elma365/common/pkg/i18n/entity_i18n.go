package i18n

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"sync"
	"time"

	"github.com/leonelquinteros/gotext"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"git.elewise.com/elma365/common/pkg/namespace"
	"git.elewise.com/elma365/common/pkg/syncutils"
)

var (
	//nolint:gochecknoglobals // приватная переменная
	defaultEntityTranslatorsCacheTTL = time.Duration(5) * time.Minute // по умолчанию кэшируем на 5 минут
	//nolint:gochecknoglobals // приватная переменная
	defaultLanguageSettingsCacheTTL = time.Duration(5) * time.Minute // по умолчанию кэшируем на 5 минут
)

// LocalizableEntity локализуемая сущность
type LocalizableEntity interface {

	// ExtractPO Извлекает переводы из сущности
	ExtractPO(ctxt *EntityPOContext) ([]byte, error)

	// ApplyTranslation применяет переводы к сущности
	ApplyTranslation(poCtxt *EntityPOContext, translator gotext.Translator)
}

// WithEntityPOLoader добавить загрузчик переводов для сущностей
func WithEntityPOLoader(loader EntityPOLoader) Option {
	return optionFunc(func(i18n *I18n) {
		i18n.entityPOLoader = loader
	})
}

// WithEntityTranslatorsCashTTL установить время жизни кэша переводов сущностей
func WithEntityTranslatorsCashTTL(ttl time.Duration) Option {
	return optionFunc(func(i18n *I18n) {
		i18n.entityTranslatorsCacheTTL = ttl
	})
}

// WithFnCompanyFromContext установить функцию, которая будет извлекать компанию из контекста
func WithFnCompanyFromContext(fn func(ctx context.Context) string) Option {
	return optionFunc(func(i18n *I18n) {
		i18n.fnCompanyFromCtx = fn
	})
}

// TranslateEntities применяет переводы к списку локализуемых сущностей сущностям
// nolint:gocyclo,funlen // if-ов многовато, но они нужны :-)
func (i18n I18n) TranslateEntities(
	ctx context.Context, entities []LocalizableEntity, ns *namespace.Namespace,
) error {
	userLocale := LangFromContext(ctx)
	if userLocale == "" {
		return nil
	}

	companyLocale := CompanyLangFromContext(ctx)
	// если язык пользователя совпадает с языком компании, то переводить ничего не нужно
	if companyLocale == userLocale {
		return nil
	}

	// формируем список namespace-ов
	entityNss := namespace.Namespaces{}
	if ns == nil {
		for i := range entities {
			if entities[i] == nil {
				continue
			}
			entityNs := i18n.getEntityNamespace(entities[i])
			if entityNs.String() == "" {
				zap.L().Debug("entity has empty namepsace")
				continue
			}
			entityNss = append(entityNss, entityNs)
		}
	} else {
		entityNss = append(entityNss, *ns)
	}
	if len(entityNss) == 0 {
		return nil
	}

	// получаем "трансляторы" для каждого namespace-а
	translators, err := i18n.getOrLoadEntityTranslatorList(ctx, entityNss, userLocale)
	if err != nil {
		return errors.WithStack(err)
	}

	// применяем переводы к каждой сущности
	for i := range entities {
		if entities[i] == nil {
			continue
		}
		var entityNs namespace.Namespace
		if ns != nil {
			entityNs = *ns
		} else {
			entityNs = i18n.getEntityNamespace(entities[i])
			if entityNs.String() == "" {
				continue
			}
		}
		translator := translators[entityNs]
		if translator != nil {
			entities[i].ApplyTranslation(nil, translator)
		} else {
			zap.L().Debug("translator is undefined")
		}
	}

	return nil
}

// TranslateEntity применяет переводы к локализуемой сущности
func (i18n I18n) TranslateEntity(
	ctx context.Context, entity LocalizableEntity, ns *namespace.Namespace,
) error {
	if entity == nil {
		return nil
	}

	userLocale := LangFromContext(ctx)
	if userLocale == "" {
		return nil
	}

	companyLocale := CompanyLangFromContext(ctx)
	// если язык пользователя совпадает с языком компании, то переводить ничего не нужно
	if companyLocale == userLocale {
		return nil
	}

	var entityNs namespace.Namespace
	if ns == nil {
		entityNs = i18n.getEntityNamespace(entity)
	} else {
		entityNs = *ns
	}
	if entityNs.String() == "" {
		zap.L().Debug("entity has empty namepsace")
		return nil
	}

	translator, err := i18n.getOrLoadEntityTranslator(ctx, entityNs, userLocale)
	if err != nil {
		return errors.WithStack(err)
	}
	if translator != nil {
		entity.ApplyTranslation(nil, translator)
	}

	return nil
}

type entityWithNamespace interface {
	GetNamespace() namespace.Namespace
}

// GetPO возвращает переводы (po-файл) из раздела для переданной локали
func (i18n I18n) GetPO(ctx context.Context, ns namespace.Namespace, locale string) ([]byte, error) {
	const entityPOLoaderIsNilError = "entityPOLoader is not initialized"

	if i18n.entityPOLoader == nil {
		return nil, errors.Errorf(entityPOLoaderIsNilError)
	}

	r, err := i18n.entityPOLoader.GetPO(ctx, ns, locale)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	buf, err := io.ReadAll(r)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return buf, nil
}

func (i18n I18n) loadEntityTranslator(
	ctx context.Context, ns namespace.Namespace, locale string,
) (gotext.Translator, error) {
	zap.L().Debug("loading entity translator: start",
		zap.String("namespace", ns.String()), zap.String("locale", locale),
	)

	if locale == "" {
		zap.L().Debug("load entity translator: locale is undefined",
			zap.String("namespace", ns.String()), zap.String("locale", locale),
		)
		return nil, nil
	}

	if i18n.entityPOLoader == nil {
		zap.L().Debug("load entity translator: entityPOLoader is undefined",
			zap.String("namespace", ns.String()), zap.String("locale", locale),
		)
		return nil, nil
	}

	var r io.Reader
	var err error

	r, err = i18n.entityPOLoader.GetPO(ctx, ns, locale)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if len(buf) == 0 {
		zap.L().Debug("load entity translator: po-file not found (or empty)",
			zap.String("namespace", ns.String()), zap.String("locale", locale),
		)
		return nil, nil
	}
	translator := gotext.NewPoTranslator()
	translator.Parse(buf)

	zap.L().Debug("loading entity translator: finish",
		zap.String("namespace", ns.String()), zap.String("locale", locale),
	)
	return translator, nil
}

func (i18n I18n) getOrLoadEntityTranslator(
	ctx context.Context, ns namespace.Namespace, locale string,
) (gotext.Translator, error) {
	if i18n.fnCompanyFromCtx == nil {
		zap.L().Warn(`Function "fnCompanyFromCtx" is undefined`)
		return nil, nil
	}

	entityTranslatorTotalCallsCountMetric.WithLabelValues(fmt.Sprintf("%s/%s", ns, locale)).Inc()

	ttl := i18n.getInexactTTL(i18n.entityTranslatorsCacheTTL)
	key := fmt.Sprintf("%s/%s/%s", i18n.fnCompanyFromCtx(ctx), ns, locale)

	translatorsProvider, _ := i18n.entityTranslators.LoadOrStore(
		key,
		syncutils.NewPromise(
			func(ctx context.Context) (interface{}, error) {
				entityTranslatorStorageCallsCountMetric.WithLabelValues(fmt.Sprintf("%s/%s", ns, locale)).Inc()
				return i18n.loadEntityTranslator(ctx, ns, locale)
			},
			syncutils.WithTTL(ttl),
		),
	)

	promise, ok := translatorsProvider.(*syncutils.Promise)
	if !ok {
		return nil, errors.WithStack(errors.New("fail to get entity translator"))
	}
	res, err := promise.Await(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	t, ok := res.(gotext.Translator)
	if !ok {
		return nil, nil
	}
	return t, nil
}

// nolint:gocyclo // много комментариев, все должно просто читаться
func (i18n I18n) getOrLoadEntityTranslatorList(
	ctx context.Context, nss namespace.Namespaces, locale string,
) (map[namespace.Namespace]gotext.Translator, error) {
	res := map[namespace.Namespace]gotext.Translator{}

	// создаем список уникальных namespace-ов
	uniqueNss := map[namespace.Namespace]struct{}{}
	for _, ns := range nss {
		if ns.String() == "" {
			continue
		}
		uniqueNss[ns] = struct{}{}
	}
	if len(uniqueNss) == 0 {
		return res, nil
	}

	// запускам получение траслаторов в 5 потоков
	workerPool := make(chan struct{}, 5)
	errGroup := errgroup.Group{}
	var locker sync.Mutex
	for ns := range uniqueNss {
		_ns := ns
		workerPool <- struct{}{}
		errGroup.Go(func() error {
			defer func() {
				<-workerPool
			}()
			t, _err := i18n.getOrLoadEntityTranslator(ctx, _ns, locale)
			if _err != nil {
				return errors.WithStack(_err)
			}
			locker.Lock()
			res[_ns] = t
			locker.Unlock()
			return nil
		})
	}
	err := errGroup.Wait()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return res, nil
}

func (i18n I18n) getEntityNamespace(entity LocalizableEntity) namespace.Namespace {
	if entity == nil {
		return namespace.Namespace("")
	}

	v, ok := entity.(entityWithNamespace)
	if !ok {
		v, ok = interface{}(&entity).(entityWithNamespace)
		if !ok {
			zap.L().Debug(`entity not implement interface "entityWithNamespace"`)
			return namespace.Namespace("")
		}
	}
	return v.GetNamespace()
}
