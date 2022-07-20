package i18n

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

//nolint: gochecknoglobals // метрики объявляются так
var (
	i18nTranslatorsResourceCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "i18n",
		Subsystem: "translators",
		Name:      "resources",
		Help:      "A counter of not found translation resources",
	}, []string{"language", "localesDir"})

	// метрика количества всех запросов для получения "транлятора" сущностей
	// nolint gochecknoglobals
	entityTranslatorTotalCallsCountMetric = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "elma365",
			Subsystem: "translators",
			Name:      "entity_translator_total_calls_count",
			Help:      "Amount of total calls for getting entity translator",
		},
		[]string{"entity_translator"},
	)

	// метрика количества запросов в хранилище (БД, файлы и т.п.) для получения "транлятора" сущностей
	// nolint gochecknoglobals
	entityTranslatorStorageCallsCountMetric = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "elma365",
			Subsystem: "translators",
			Name:      "entity_translator_storage_calls_count",
			Help:      "Amount of calls to storage for getting entity translator",
		},
		[]string{"entity_translator"},
	)
)
