package collection

//go:generate ../../tooling/bin/easyjson -lower_camel_case $GOFILE

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

// ItemIdentifier описывает интерфейс c методом получения ИД элемента коллекции
type ItemIdentifier interface {
	GetID() uuid.UUID
}

// Item описывает основные системные поля элемента коллекции
//
// easyjson:json
type Item struct {
	ID        uuid.UUID  `db:"id"   json:"__id"`
	CreatedAt time.Time  `db:"-"    json:"__createdAt"`
	CreatedBy uuid.UUID  `db:"-"    json:"__createdBy"`
	UpdatedAt time.Time  `db:"-"    json:"__updatedAt"`
	UpdatedBy uuid.UUID  `db:"-"    json:"__updatedBy"`
	DeletedAt *time.Time `db:"-"    json:"__deletedAt"`
	Version   uint32     `db:"-"    json:"__version"`
	Name      string     `db:"-"    json:"__name"`
}

// GetID реализует интерфейс collection.ItemIdentifier
func (item Item) GetID() uuid.UUID {
	return item.ID
}

// Scan implements sql.Scanner interface
func (item *Item) Scan(pSrc interface{}) error {
	switch src := pSrc.(type) {
	case []byte:
		err := json.Unmarshal(src, item)
		if err != nil {
			return errors.WithStack(err)
		}

	default:
		return fmt.Errorf("Item.Scan: cannot scan type %T into Item", pSrc)
	}

	return nil
}

// BodyItem описывает основные системные поля элемента коллекции, и хранит остальные его поля в "сыром" виде.
// При сериализации в json, возвращает только "сырые значения".
// Реализует интерфейс collection.ItemIdentifier.
type BodyItem struct {
	Item
	Body json.RawMessage `db:"body" json:"body"`
}

// NewBodyItem создает экземпляр BodyItem
func NewBodyItem(itemJSON json.RawMessage) (*BodyItem, error) {
	baseItem := Item{}
	if err := json.Unmarshal(itemJSON, &baseItem); err != nil {
		return nil, errors.WithStack(err)
	}

	baseBodyItem := BodyItem{
		Item: baseItem,
		Body: make(json.RawMessage, len(itemJSON)),
	}

	copy(baseBodyItem.Body, itemJSON)

	return &baseBodyItem, nil
}

// GetID реализует интерфейс collection.ItemIdentifier
func (item BodyItem) GetID() uuid.UUID {
	return item.ID
}

// MarshalJSON реализует интерфейс json.Marshaler
func (item BodyItem) MarshalJSON() ([]byte, error) {
	data, err := json.Marshal(item.Body)
	if err != nil {
		return nil, fmt.Errorf("BodyItem json serializer failed: %v", err)
	}

	return data, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (item *BodyItem) UnmarshalJSON(data []byte) error {
	newItem, err := NewBodyItem(data)
	if err != nil {
		return errors.WithStack(err)
	}

	*item = *newItem

	return nil
}

// Scan implements sql.Scanner interface
func (item *BodyItem) Scan(pSrc interface{}) error {
	switch src := pSrc.(type) {
	case []byte:
		newItem, err := NewBodyItem(src)
		if err != nil {
			return errors.WithStack(err)
		}
		*item = *newItem

	default:
		return fmt.Errorf("BodyItem.Scan: cannot scan type %T into BodyItem", pSrc)
	}

	return nil
}
