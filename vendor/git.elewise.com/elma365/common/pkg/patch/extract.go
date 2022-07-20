package patch

import (
	"encoding/json"

	"github.com/Jeffail/gabs"
	"github.com/pkg/errors"
)

// Extract вытаскивает значения по путям
type Extract struct {
	path Path
}

// NewExtract — конструктор
func NewExtract(path ...interface{}) Extract {
	return Extract{
		path: NewPath(path...),
	}
}

// ToReplace - создает из извлекателя заменитель
func (extract Extract) ToReplace(val interface{}) Replace {
	return NewReplace(val, extract.path)
}

// Prefix префиксить путь изменения
func (extract Extract) Prefix(path ...interface{}) Extract {
	return Extract{
		path: extract.path.Prefix(path...),
	}
}

// Apply применить изменение
func (extract Extract) Apply(container *gabs.Container) interface{} {
	return extract.path.Get(container)
}

// String возвращает строковое представление
func (extract Extract) String() string {
	return extract.path.String()
}

// ExtractFromString преобразует строку в Extract
func ExtractFromString(s string) (Extract, error) {
	path, err := FromString(s)
	if err != nil {
		return Extract{}, errors.WithStack(err)
	}
	return Extract{path}, nil
}

// Extracts набор изменений JSON-документа
type Extracts []Extract

// Prefix префиксить путь изменений
func (extracts Extracts) Prefix(path ...interface{}) Extracts {
	res := make(Extracts, len(extracts))

	for i := range extracts {
		res[i] = extracts[i].Prefix(path...)
	}

	return res
}

// ToStrings - преобразует в массив строк (путей)
func (extracts Extracts) ToStrings() []string {
	res := make([]string, len(extracts))
	for i := range extracts {
		res[i] = extracts[i].String()
	}
	return res
}

// Apply применить изменения
func (extracts Extracts) Apply(blob json.RawMessage) (map[string]interface{}, error) {
	container, err := gabs.ParseJSON(blob)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	res := make(map[string]interface{}, len(extracts))
	for _, extract := range extracts {
		res[extract.path.String()] = extract.Apply(container)
	}

	return res, nil
}
