package uuids

//go:generate ../../../tooling/bin/easyjson $GOFILE

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"sort"
	"strings"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

// UUIDS определяет тип - массив uuid-ов и стандартные операции над ними
//easyjson:json
type UUIDS []uuid.UUID

// FromStrings creates UUIDS from string slice
func FromStrings(list []string) (UUIDS, error) {
	uuids := make([]uuid.UUID, 0, len(list))
	for _, item := range list {
		id, err := uuid.FromString(item)
		if err != nil {
			return nil, err
		}
		uuids = append(uuids, id)
	}

	return UUIDS(uuids), nil
}

// Scan implements sql.Scanner interface
func (us *UUIDS) Scan(pSrc interface{}) error {
	var byteSrc []byte
	switch src := pSrc.(type) {
	case nil:
		*us = UUIDS{}
		return nil
	case string:
		byteSrc = []byte(src)
	case []byte:
		byteSrc = src
	default:
		return fmt.Errorf("can`t scan type %T into UUIDS. Column should have []UUID type", pSrc)
	}

	if len(byteSrc) == 0 {
		return fmt.Errorf("can`t scan type %T into UUIDS. Column should have []UUID type", pSrc)
	}

	byteSrc = bytes.TrimLeft(byteSrc, "{")
	byteSrc = bytes.TrimRight(byteSrc, "}")
	parts := bytes.Split(byteSrc, []byte(","))

	uuidsList := make([]uuid.UUID, 0, len(parts))
	for _, part := range parts {
		if len(part) == 0 {
			continue
		}
		// strconv.Unquote не умеет строки вида 'строка'. В одинарных кавычках у нее может быть только литерал.
		part = bytes.Trim(part, ` '"`)
		uuidItem, err := uuid.FromString(string(part))
		if err != nil {
			return errors.WithStack(err)
		}

		uuidsList = append(uuidsList, uuidItem)
	}
	*us = uuidsList

	return nil
}

// Value implements sql.Valuer interface
func (us UUIDS) Value() (driver.Value, error) {
	var builder strings.Builder
	_ = builder.WriteByte('{')
	for index, u := range us {
		_, _ = builder.WriteString(fmt.Sprintf(`"%s"`, u.String()))
		if index < len(us)-1 {
			_ = builder.WriteByte(',')
		}
	}
	_ = builder.WriteByte('}')

	return builder.String(), nil
}

// AsStrings returns slice of UUIDs as slice of strings
func (us UUIDS) AsStrings() []string {
	res := make([]string, 0, len(us))

	for _, u := range us {
		res = append(res, u.String())
	}

	return res
}

// RemoveDuplicates remove duplicate values from array
func (us *UUIDS) RemoveDuplicates() {
	tmp := make(map[uuid.UUID]struct{}, len(*us))

	for _, v := range *us {
		tmp[v] = struct{}{}
	}

	list := make([]uuid.UUID, 0, len(tmp))
	for v := range tmp {
		list = append(list, v)
	}

	*us = list
}

// Contains check array of UUIDS contains specified UUID
func (us UUIDS) Contains(u uuid.UUID) bool {
	for _, v := range us {
		if uuid.Equal(v, u) {
			return true
		}
	}

	return false
}

// Remove удаляет первый найденный заданный элемент из списка
func (us *UUIDS) Remove(u uuid.UUID) {
	for index, id := range *us {
		if uuid.Equal(u, id) {
			*us = append((*us)[:index], (*us)[index+1:]...)
			break
		}
	}
}

// Equal compare UUID array with new UUID array
//
// Return true if equals
func (us UUIDS) Equal(x UUIDS) bool {
	if len(us) != len(x) {
		return false
	}
	sort.Sort(us)
	sort.Sort(x)
	i := 0
	for i < len(us) {
		if us[i] != x[i] {
			return false
		}
		i++
	}
	return true
}

// Diff compare UUID array with new UUID array
//
// Return UUIDS that were added as first value
// and deleted UUIDS as second value
func (us UUIDS) Diff(x UUIDS) (plus, minus UUIDS) {
	sort.Sort(us)
	sort.Sort(x)

	i, j := 0, 0

	for i < len(us) || j < len(x) {
		if j >= len(x) {
			minus = append(minus, us[i])
			i++
			continue
		}

		if i >= len(us) {
			plus = append(plus, x[j])
			j++
			continue
		}

		switch {
		case us[i] == x[j]:
			i++
			j++
		case bytes.Compare(us[i].Bytes(), x[j].Bytes()) > 0:
			plus = append(plus, x[j])
			j++
		default:
			minus = append(minus, us[i])
			i++
		}
	}

	return plus, minus
}

// methods for sort.Interface

// Len implements sort.Interface
func (us UUIDS) Len() int { return len(us) }

// Swap implements sort.Interface
func (us UUIDS) Swap(i, j int) { us[i], us[j] = us[j], us[i] }

// Less implements sort.Interface
func (us UUIDS) Less(i, j int) bool { return bytes.Compare(us[i].Bytes(), us[j].Bytes()) < 0 }
