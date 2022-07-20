package phone

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

//go:generate ../../../../tooling/bin/easyjson phone.go

// Phone is a tuple of phone information
//
// See https://www.ietf.org/rfc/rfc3966.txt for details of format
// easyjson:json
type Phone struct {
	Tel     string `json:"tel"`               // локальная часть, действует в пределах страны
	Context string `json:"context,omitempty"` // Глобальная часть, например код страны или домен (для SIP) (в RFC называется phone-context)
	Ext     string `json:"ext,omitempty"`     // добавочный номер
	Type    string `json:"type"`              // тип номера, это уже не RFC а наша приблуда - Домашний/мобильный/etc
}

// Scan implements sql.Scanner interface
func (p *Phone) Scan(pSrc interface{}) error {
	var value Phone

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
		return fmt.Errorf("Phone.Scan: cannot scan type %T into Phone", pSrc)
	}

	*p = value
	return nil
}

// Value implements sql.Valuer interface
func (p Phone) Value() (value driver.Value, err error) {
	val, err := json.Marshal(p)
	return string(val), errors.WithStack(err)
}

// String implements fmt.Stringer interface
func (p *Phone) String() string {
	return p.Context + p.Tel
}

// Phones - array of phone structs
// easyjson:json
type Phones []Phone

// Scan implements sql.Scanner interface
func (ps *Phones) Scan(pSrc interface{}) error {
	var value Phones

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
		return fmt.Errorf("Phones.Scan: cannot scan type %T into Phones", pSrc)
	}

	*ps = value
	return nil
}

// Value implements sql.Valuer interface
func (ps Phones) Value() (value driver.Value, err error) {
	val, err := json.Marshal(ps)
	return string(val), errors.WithStack(err)
}
