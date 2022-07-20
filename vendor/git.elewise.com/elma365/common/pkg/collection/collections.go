package collection

import (
	"errors"

	"git.elewise.com/elma365/common/pkg/errs"
)

// Collections тип для массива коллекций
type Collections []Collection

// GetCollectionByAlias выбирает коллекцию по алиасу из массива коллекций
func (c Collections) GetCollectionByAlias(alias string) (Collection, error) {
	if alias == "" {
		return c[0], nil
	}
	for _, coll := range c {
		if alias == coll.Alias {
			return coll, nil
		}
	}

	return Collection{}, errors.New("collection not found")
}

// ValidateNotColumned выдаёт ошибку, если хотя бы у одной коллекции из массива поле Columned установлено в true
func (c Collections) ValidateNotColumned() error {
	for _, coll := range c {
		if coll.Columned {
			return errs.NotImplemented.New("collection is columned")
		}
	}
	return nil
}
