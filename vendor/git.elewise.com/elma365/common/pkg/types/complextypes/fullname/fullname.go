package fullname

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

//go:generate ../../../../tooling/bin/easyjson fullname.go

// FullName is a tuple of fullname information
//
// easyjson:json
type FullName struct {
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	MiddleName string `json:"middlename"`
}

// New создает новую структуру на базе переданной строки
func New(s string) FullName {
	s = strings.TrimSpace(s)
	arr := strings.Split(s, " ")

	if len(arr) > 3 {
		arr = arr[:3]
	}

	var fn FullName

	switch len(arr) {
	case 1:
		fn.FirstName = arr[0]

	case 2:
		fn.LastName = arr[0]
		fn.FirstName = arr[1]

	case 3:
		fn.LastName = arr[0]
		fn.FirstName = arr[1]
		fn.MiddleName = arr[2]
	}

	return fn
}

// FromString заполняет структу из строки
func (fn *FullName) FromString(s string) {
	s = strings.TrimSpace(s)
	arr := strings.Split(s, " ")

	if len(arr) > 3 {
		arr = arr[:3]
	}

	switch len(arr) {
	case 1:
		fn.FirstName = arr[0]

	case 2:
		fn.LastName = arr[0]
		fn.FirstName = arr[1]

	case 3:
		fn.LastName = arr[0]
		fn.FirstName = arr[1]
		fn.MiddleName = arr[2]
	}
}

// Scan implements sql.Scanner interface
func (fn *FullName) Scan(pSrc interface{}) error {
	var value FullName

	switch src := pSrc.(type) {
	case nil:
		return nil
	case string:
		data := []byte(src)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
	case []byte:
		if err := json.Unmarshal(src, &value); err != nil {
			return err
		}
	default:
		return fmt.Errorf("FullName.Scan: cannot scan type %T into FullName", pSrc)
	}

	*fn = value
	return nil
}

// Value implements sql.Valuer interface
func (fn FullName) Value() (value driver.Value, err error) {
	val, err := json.Marshal(fn)
	return string(val), errors.WithStack(err)
}

// String returns concatenated full name
func (fn *FullName) String() string {
	return strings.TrimSpace(fmt.Sprintf("%s %s %s", fn.LastName, fn.FirstName, fn.MiddleName))
}
