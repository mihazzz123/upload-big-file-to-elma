package models

import (
	"sync"

	"github.com/leonelquinteros/gotext"
)

// Translators типизированная sync.Map для работы с трансляторами
type Translators struct {
	sync.Map
}

// Load типизированный Load sync.Map
func (t *Translators) Load(key string) (gotext.Translator, bool) {
	if translator, ok := t.Map.Load(key); ok {
		return translator.(gotext.Translator), ok
	}
	return nil, false
}

// LoadOrStore типизированный LoadOrStore sync.Map
func (t *Translators) LoadOrStore(key string, translator gotext.Translator) (gotext.Translator, bool) {
	actual, loaded := t.Map.LoadOrStore(key, translator)
	if actual != nil {
		return actual.(gotext.Translator), loaded
	}
	return nil, loaded
}

// Store типизированный Store sync.Map
func (t *Translators) Store(key string, translator gotext.Translator) {
	t.Map.Store(key, translator)
}

// Range типизированный Range sync.Map
func (t *Translators) Range(f func(key string, value gotext.Translator) bool) {
	t.Map.Range(func(key interface{}, value interface{}) bool {
		var translator gotext.Translator
		if value != nil {
			translator = value.(gotext.Translator)
		}
		return f(key.(string), translator)
	})
}
