package syscollection

import "encoding/json"

//go:generate ../../../../tooling/bin/easyjson data.go

// Data - детализация данных для этого типа
//
// ВНИМАНИЕ! Эта структура должна быть идентична интерфейсу с фронта app/types/schema/collection/collection-field-data.ts
//
// easyjson:json
type Data struct {
	Namespace       string            `json:"namespace"        validate:"required"` // Раздел связанного справочника
	Code            string            `json:"code"             validate:"required"` // Код связанного справочника
	LinkedFieldCode string            `json:"linkedFieldCode"`                      // Код связанного поля. Если не установлено, то нет и связанного поля.
	IsDependent     bool              `json:"isDependent"`                          // Признак завимой связи
	Bindings        []json.RawMessage `json:"bindings"`                             // Отображение полей для создания
	Filter          json.RawMessage   `json:"filter"`                               // Фильтрация вывода приложений
}

func (d Data) String() string {
	result := d.Namespace + ":" + d.Code
	if len(d.LinkedFieldCode) > 0 {
		result = result + ":" + d.LinkedFieldCode
	}
	return result
}
