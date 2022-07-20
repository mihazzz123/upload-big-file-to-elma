package collection

import (
	"git.elewise.com/elma365/common/pkg/permissions"

	uuid "github.com/satori/go.uuid"
)

const (
	// DiskCode системная коллекция файлов
	DiskCode string = "disk_files"
	// GroupsCode системная коллекция групп
	GroupsCode string = "groups"
	// TasksCode системная коллекция задач
	TasksCode string = "tasks"
	// UsersCode системная коллекция пользователей
	UsersCode string = "users"

	// ProcessInstanceCodePrefix префикс кода коллеуции экземпляра процесса
	ProcessInstanceCodePrefix = "_process_"

	// CollectionIDCol - название колонки, в которой хранится идентификатор элемента
	CollectionIDCol = "id"
	// CollectionBodyCol — название колонки, в которой хранится тело элемента
	CollectionBodyCol = "body"
	// CollectionPermissionsCol — название колонки, в которой хранятся разрешения объекта
	CollectionPermissionsCol = "permissions"
	// CollectionReadCol — название колонки, в которой хранится развёрнутый список групп, имеющих право на чтение
	CollectionReadCol = "read"
	// CollectionInheritCol — название колонки, в которой хранится признак наследования разрешений объекта
	CollectionInheritCol = "inherit"

	// CollectionType - поле типа коллекции
	CollectionType = "type"

	// FieldWithAlias - название поля фильтра, содержащего алиас коллекции.
	// Префикс # сделан для однозначной идентификации поля среди других пользовательских полей фильтра.
	// Наличие этого поля в фильтре указывает на то, что фильтр будет применён для конкретной коллекции с алиасом, как в фильтре.
	// Для table-фильтра поле #alias добавляется в список существующих полей.
	// Для fts-фильтра поле #alias добавляется после поля const.
	// Для других видов фильтра поле #alias добавляется после поля field.
	FieldWithAlias = "#alias"
)

// nolint:gochecknoglobals // Это константы
var (
	// GroupUserProfiles определяет ИД группы "Внешние пользователи".
	GroupUserProfiles = uuid.NewV5(uuid.NamespaceOID, "system:user_profiles") // f25906e4-41c3-5a89-8ec2-06648dd1f614
	// DefaultPermissionsForPortalUser права на коллекцию по умолчанию для портального пользователя,
	// если явно не прописаны в коллекции collection.AccessTypeNone.
	DefaultPermissionsForPortalUser = permissions.Permissions{
		Values: []permissions.Permission{
			{
				Group: permissions.Group{
					ID:   GroupUserProfiles,
					Type: permissions.GroupTypeGroup,
				},
				Types: permissions.EmptyAccess,
			},
		},
	}
)
