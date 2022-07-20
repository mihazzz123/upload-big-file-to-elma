package collection

//go:generate ../../tooling/bin/easyjson -lower_camel_case $GOFILE
//go:generate ../../tooling/bin/easylocalizer $GOFILE

import (
	"time"

	"git.elewise.com/elma365/common/pkg/namespace"
	"git.elewise.com/elma365/common/pkg/permissions"
	"git.elewise.com/elma365/common/pkg/types"

	uuid "github.com/satori/go.uuid"
)

// Collection описание коллекции
//
// easyjson:json
// localizer:collection
type Collection struct {
	ID          uuid.UUID               `db:"id"             json:"__id"`
	Alias       string                  `db:"alias"          json:"alias" localizer:"value"`
	Namespace   namespace.Namespace     `db:"namespace"      json:"namespace" localizer:"id"`
	Name        string                  `db:"name"           json:"name" localizer:"value"`
	Code        string                  `db:"code"           json:"code" localizer:"id"`
	Type        Type                    `db:"type"           json:"type"`
	Fields      types.Fields            `db:"fields"         json:"fields"`
	AccessType  AccessType              `db:"access_type"    json:"accessType"`
	Subordinate bool                    `db:"subordinate"    json:"subordinate"` // Разрешать доступ к элементам по иерархии оргструктуры
	Permissions permissions.Permissions `db:"permissions"    json:"permissions"`
	Indices     *bool                   `db:"indices"        json:"indices"` // Флаг влюченных индексов коллекции
	CreatedAt   time.Time               `db:"created_at"     json:"__createdAt"`
	CreatedBy   uuid.UUID               `db:"created_by"     json:"__createdBy"`
	UpdatedAt   time.Time               `db:"updated_at"     json:"__updatedAt"`
	UpdatedBy   uuid.UUID               `db:"updated_by"     json:"__updatedBy"`
	DeletedAt   *time.Time              `db:"deleted_at"     json:"__deletedAt"`
	Columned    bool                    `db:"-"              json:"-"               default:"false"`
}

// GetNamespace имплементирует метод интерфейса connection.CollectionMetadata
func (c *Collection) GetNamespace() namespace.Namespace {
	return c.Namespace
}

// GetCode имплементирует метод интерфейса connection.CollectionMetadata
func (c *Collection) GetCode() string {
	return c.Code
}

// GetFields имплементирует метод интерфейса connection.CollectionMetadata
func (c *Collection) GetFields() types.Fields {
	return c.Fields
}

// Metadata - метаданные коллекции (раздел, код, поля)
type Metadata interface {
	GetNamespace() namespace.Namespace
	GetCode() string
	GetFields() types.Fields
}

// SimpleMetadata - дефолтная реализация интерфейса метаданные коллекции (раздел, код, поля)
type SimpleMetadata struct {
	Namespace namespace.Namespace
	Code      string
	Fields    types.Fields
}

// GetNamespace - Collection namespace
func (md SimpleMetadata) GetNamespace() namespace.Namespace {
	return md.Namespace
}

// GetCode - Collection code
func (md SimpleMetadata) GetCode() string {
	return md.Code
}

// GetFields - Collection fields
func (md SimpleMetadata) GetFields() types.Fields {
	return md.Fields
}
