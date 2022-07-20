package permissions

//go:generate ../../tooling/bin/easyjson -no_std_marshalers $GOFILE

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"git.elewise.com/elma365/common/pkg/types/complextypes/refitem"
	"git.elewise.com/elma365/common/pkg/types/uuids"

	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

// Group группа владельца доступа
//
// easyjson:json
type Group struct {
	ID   uuid.UUID `json:"id"   validate:"required"`
	Type GroupType `json:"type" validate:"required"`
}

// Permission отдельный доступ, связанный с определённой группой пользователей и набором разрешений
//
// easyjson:json
type Permission struct {
	Group        Group   `json:"group"`
	OrgunitField *string `json:"orgunitField,omitempty"` // код поля (свойства Приложения) типа "Пользователь", "Группа" или "Орг. структура" на основе значений которого будет вычисляться доступ для текущего пользователя
	Types        Access  `json:"types"`
	Inherited    bool    `json:"inherited"` // флаг - это наследованная пермиссия, или родная
}

// Scan implements sql.Scanner interface
func (pn *Permission) Scan(pSrc interface{}) error {
	switch src := pSrc.(type) {
	case string:
		data := []byte(src)
		if err := json.Unmarshal(data, &pn); err != nil {
			return err
		}
	case []byte:
		if err := json.Unmarshal(src, &pn); err != nil {
			return err
		}
	default:
		return fmt.Errorf("Permission.Scan: cannot scan type %T into Permission", pSrc)
	}

	return nil
}

// Value implements sql.Valuer interface
func (pn Permission) Value() (driver.Value, error) {
	val, err := json.Marshal(pn)
	return string(val), errors.WithStack(err)
}

// IsValid - валидация прав
func (pn *Permission) IsValid() bool {
	return (pn.Group.ID != uuid.Nil || pn.OrgunitField != nil) && (pn.Group.ID == uuid.Nil || pn.OrgunitField == nil)
}

// Permissions набор доступов к объекту
//
// easyjson:json
type Permissions struct {
	InheritParent bool             `json:"inheritParent"`     // глобальный флаг, наследовать или нет родительские пермиссии
	Values        []Permission     `json:"values"`            // список пермиссий
	Timestamp     int64            `json:"timestamp"`         // дата изменения пермишшенов, в unixtime
	RefItem       *refitem.RefItem `json:"refItem,omitempty"` // ссылка на элемент приложения, учитываем его права
}

// Inherited возвращает только отнаследованные пермиссии
func (p Permissions) Inherited() []Permission {
	var inhPerms []Permission
	for _, v := range p.Values {
		if v.Inherited {
			inhPerms = append(inhPerms, v)
		}
	}

	return inhPerms
}

// NonInherited возвращает только родные пермиссии (не наследованные)
func (p Permissions) NonInherited() []Permission {
	var inhPerms []Permission
	for _, v := range p.Values {
		if !v.Inherited {
			inhPerms = append(inhPerms, v)
		}
	}

	return inhPerms
}

// AsInherited returns all permissions as inherited
func (p Permissions) AsInherited() []Permission {
	var inh []Permission

	var tmpMap = make(map[uuid.UUID]Permission)

	for _, v := range p.Values {
		val, ok := tmpMap[v.Group.ID]
		if ok {
			// если уже есть такая пермиссия, надо слить разрешения
			val.Types |= v.Types
			continue
		}

		if v.Inherited {
			tmpMap[v.Group.ID] = v
		} else {
			tmpMap[v.Group.ID] = Permission{
				Group:        v.Group,
				OrgunitField: v.OrgunitField,
				Types:        v.Types,
				Inherited:    true,
			}
		}
	}

	for _, pr := range tmpMap {
		inh = append(inh, pr)
	}

	return inh
}

// GetIDsByAtom — возвращает идентификаторы групп, имеющих указанный атом в правах
func (p Permissions) GetIDsByAtom(atom AccessAtom) uuids.UUIDS {
	var res uuids.UUIDS

	for _, perm := range p.Values {
		if (p.InheritParent || !perm.Inherited) && perm.Types.Has(atom) {
			res = append(res, perm.Group.ID)
		}
	}

	return res
}

// Add добавляет ненаследованные права в список
//
// Если права для группы уже были в списке и были не наследованы,
// то новые права просто добавляются к имеющимся.
// Иначе добавляется новая строка.
//
// Добавляемые наследованные права откидываются.
func (p Permissions) Add(plus []Permission) Permissions {
	values := make([]Permission, len(p.Values))
	copy(values, p.Values)
	m := make(map[Group]int)
	for i, x := range values {
		if !x.Inherited {
			m[x.Group] = i
		}
	}
	for _, x := range plus {
		if !x.Inherited {
			if i, ok := m[x.Group]; ok {
				values[i].Types |= x.Types
			} else {
				values = append(values, x)
				m[x.Group] = len(values) - 1
			}
		}
	}
	p.Values = values
	return p
}

// Squeeze объединяет записи пермишенов по совокупности признаков UUID, Inherited и OrgunitField
func (p Permissions) Squeeze() Permissions {
	type adaptiveKey struct {
		uuid         uuid.UUID
		inherited    bool
		orgunitField *string
	}
	m := make(map[adaptiveKey]int)
	items := []Permission{}
	for _, item := range p.Values {
		if idx, ok := m[adaptiveKey{item.Group.ID, item.Inherited, item.OrgunitField}]; ok {
			items[idx].Types |= item.Types
		} else {
			items = append(items, item)
			m[adaptiveKey{item.Group.ID, item.Inherited, item.OrgunitField}] = len(items) - 1
		}
	}
	p.Values = items
	return p
}

// Delete удаляет ненаследованные права из списка
//
// Наследованные права не затрагиваются.
func (p Permissions) Delete(minus []Permission) Permissions {
	values := make([]Permission, len(p.Values))
	copy(values, p.Values)
	m := make(map[Group]int)
	for i, x := range values {
		if !x.Inherited {
			m[x.Group] = i
		}
	}
	for _, x := range minus {
		if !x.Inherited {
			if i, ok := m[x.Group]; ok {
				values[i].Types &^= x.Types
			}
		}
	}
	p.Values = values
	return p
}

// Scan implements sql.Scanner interface
func (p *Permissions) Scan(pSrc interface{}) error {
	switch src := pSrc.(type) {
	case string:
		data := []byte(src)
		if err := json.Unmarshal(data, &p); err != nil {
			return err
		}
	case []byte:
		if err := json.Unmarshal(src, &p); err != nil {
			return err
		}
	default:

		return fmt.Errorf("Permissions.Scan: cannot scan type %T into Permissions", pSrc)
	}

	return nil
}

// Value implements sql.Valuer interface
func (p Permissions) Value() (driver.Value, error) {
	val, err := json.Marshal(p)
	return string(val), errors.WithStack(err)
}

// MarshalJSON вызов easyjson маршаллера с флагами
func (p Permissions) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{Flags: jwriter.NilSliceAsEmpty}
	p.MarshalEasyJSON(&w)
	return w.Buffer.BuildBytes(), w.Error
}

// UnmarshalJSON вызов анмаршаллера с инициализацией пустого слайса пермишенов
func (p *Permissions) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	p.UnmarshalEasyJSON(&r)
	if p.Values == nil {
		p.Values = []Permission{}
	}
	return r.Error()
}

// методы для протобаф

// Marshal marshaler interface
func (p Permissions) Marshal() ([]byte, error) {
	return json.Marshal(p)
}

// MarshalTo protobuf marshaler interface
func (p *Permissions) MarshalTo(data []byte) (n int, err error) {
	d, err := json.Marshal(p)
	if err != nil {
		return 0, err
	}
	return copy(data, d), nil
}

// Unmarshal unmarshaler interface
func (p *Permissions) Unmarshal(data []byte) error {
	return json.Unmarshal(data, p)
}

// Size - size for protobuf
func (p *Permissions) Size() int {
	if p == nil {
		return 0
	}

	d, _ := json.Marshal(p)

	return len(d)
}

// IsValid - валидация списка прав
func (p *Permissions) IsValid() bool {
	for _, value := range p.Values {
		if !value.IsValid() {
			return false
		}
	}
	return true
}
