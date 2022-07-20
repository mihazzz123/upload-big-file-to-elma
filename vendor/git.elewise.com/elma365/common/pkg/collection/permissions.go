package collection

import (
	"context"

	"git.elewise.com/elma365/common/pkg/md"
	"git.elewise.com/elma365/common/pkg/types"

	"strings"
)

// ReducePermissionsForExternalUsers ограничивает видимость для внешних пользователей.
func ReducePermissionsForExternalUsers(
	ctx context.Context,
	coll *Collection,
) *Collection {
	// Для внешних пользователей, если явно права на коллекцию не заданы
	// ограничиваем ее видимость для них (урезаем CRUD), кроме системных коллекций,
	// которые нужны для корректной работы системы
	if md.IsPortalUserFromContext(ctx) &&
		coll.AccessType == AccessTypeNone &&
		!(coll.Namespace.IsSystem() ||
			strings.HasPrefix(coll.Code, "_process_") ||
			strings.Contains(coll.Code, "@"+types.DirectoryFieldCode)) {
		coll.AccessType = AccessTypeCollection
		coll.Permissions = DefaultPermissionsForPortalUser
	}
	return coll
}
