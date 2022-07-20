package role

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

//go:generate ../../../../tooling/bin/easyjson role.go

// Role is a tuple of satus information
//
// easyjson:json
type Role struct {
	Type string `json:"type"`
	Code string `json:"code"`
}

// Scan implements sql.Scanner interface
func (s *Role) Scan(pSrc interface{}) error {
	var value Role

	switch src := pSrc.(type) {
	case nil:
		return nil

	case []byte:
		if err := json.Unmarshal(src, &value); err != nil {
			return err
		}

	default:
		return fmt.Errorf("Role.Scan: cannot scan type %T into Role", pSrc)
	}

	*s = value
	return nil
}

// Value implements sql.Valuer interface
func (s Role) Value() (value driver.Value, err error) {
	return json.Marshal(s)
}
