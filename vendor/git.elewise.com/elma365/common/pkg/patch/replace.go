package patch

import (
	"encoding/json"
	"fmt"

	"github.com/Jeffail/gabs"
	"github.com/pkg/errors"
)

// Replace изменение JSON-документа
type Replace struct {
	path  Path
	value interface{}
}

// NewReplace — конструктор
func NewReplace(value interface{}, path ...interface{}) Replace {
	return Replace{
		path:  NewPath(path...),
		value: value,
	}
}

// Prefix префиксить путь изменения
func (replace Replace) Prefix(parts ...interface{}) Replace {
	return Replace{
		path:  replace.path.Prefix(parts...),
		value: replace.value,
	}
}

// Apply применить изменение
func (replace Replace) Apply(container *gabs.Container) error {
	return replace.path.Set(container, replace.value)
}

// Replaces набор изменений JSON-документа
type Replaces []Replace

// Prefix префиксить путь изменений
func (replaces Replaces) Prefix(path ...interface{}) Replaces {
	res := make(Replaces, len(replaces))

	for i := range replaces {
		res[i] = replaces[i].Prefix(path...)
	}

	return res
}

// Apply применить изменения
func (replaces Replaces) Apply(blob json.RawMessage) (json.RawMessage, error) {
	container, err := gabs.ParseJSON(blob)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = replaces.ApplyToContainer(container)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return container.Bytes(), nil
}

// ApplyToContainer - применить изменения к контейнеру
func (replaces Replaces) ApplyToContainer(container *gabs.Container) error {
	for _, replace := range replaces {
		if err := replace.Apply(container); err != nil {
			return errors.WithStack(errors.Wrap(
				err,
				fmt.Sprintf("apply %s", replace.path),
			))
		}
	}
	return nil
}
