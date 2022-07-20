package i18n

//go:generate ../../tooling/bin/easyjson $GOFILE

import (
	"fmt"
	"strings"
)

// EntityPOContext PO-контекст перевода сущности
// easyjson:json
type EntityPOContext struct {
	Name   string           `json:"name"`             // название сущности
	ID     string           `json:"id,omitempty"`     // идентификатор сущности
	Parent *EntityPOContext `json:"parent,omitempty"` // родительский контекст
}

// String превращает PO-контекст в строку
func (ctxt EntityPOContext) String() string {
	// вычисляем корневую сущность
	var rootCtxt *EntityPOContext
	rootName := "_"
	rootID := "_"

	if ctxt.Parent != nil {
		rootCtxt = &ctxt
		for rootCtxt.Parent != nil {
			rootCtxt = rootCtxt.Parent
		}
	}
	if rootCtxt != nil {
		rootName = rootCtxt.Name
		if strings.TrimSpace(rootName) == "" {
			rootName = "_"
		}
		rootID = rootCtxt.ID
		if strings.TrimSpace(rootID) == "" {
			rootID = "_"
		}
	}

	name := ctxt.Name
	if strings.TrimSpace(name) == "" {
		name = "_"
	}

	// !!! генерацию хеша пока отрубил, есть подозрение, что контекст как таковой вообще будет не нужен
	//
	// // сериализуем в json (можно использовать любой алгоритм сериализации)
	// s, _ := json.Marshal(ctxt)
	// // берем hash от полученной json-ки
	// // #nosec
	// hash := md5.Sum(s)
	// // генерируем и возвращаем po-контекст
	// return fmt.Sprintf("/entities/%s/%s@%s/hash:%x", rootName, rootID, name, hash)

	return fmt.Sprintf("/entities/%s/%s@%s", rootName, rootID, name)
}
