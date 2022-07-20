package status

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

//go:generate ../../../../tooling/bin/easyjson status.go

// Status is a tuple of satus information
//
// easyjson:json
type Status struct {
	Status uint32 `json:"status"`
	Order  uint32 `json:"order"`
}

// Scan implements sql.Scanner interface
func (s *Status) Scan(pSrc interface{}) error {
	var value Status

	switch src := pSrc.(type) {
	case nil:
		return nil
	case string:
		data := []byte(src)
		if err := json.Unmarshal(data, &value); err != nil {
			return errors.WithStack(err)
		}
	case []byte:
		if err := json.Unmarshal(src, &value); err != nil {
			return errors.WithStack(err)
		}
	default:
		return fmt.Errorf("Status.Scan: cannot scan type %T into Status", pSrc)
	}

	*s = value
	return nil
}

// Value implements sql.Valuer interface
func (s Status) Value() (value driver.Value, err error) {
	val, err := json.Marshal(s)
	return string(val), errors.WithStack(err)
}
