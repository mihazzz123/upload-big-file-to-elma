package account

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

//go:generate ../../../../tooling/bin/easyjson account.go

// Account is a tuple of account information
//
// easyjson:json
type Account struct {
	Login string `json:"login"`
	Type  string `json:"type"`
}

// Scan implements sql.Scanner interface
func (a *Account) Scan(pSrc interface{}) error {
	var value Account

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
		return fmt.Errorf("Account.Scan: cannot scan type %T into Account", pSrc)
	}

	*a = value
	return nil
}

// Value implements sql.Valuer interface
func (a Account) Value() (value driver.Value, err error) {
	val, err := json.Marshal(a)
	return string(val), errors.WithStack(err)
}

// String implements fmt.Stringer interface
func (a *Account) String() string {
	return a.Login
}

// Accounts - array of accounts structs
// easyjson:json
type Accounts []Account

// Scan implements sql.Scanner interface
func (as *Accounts) Scan(pSrc interface{}) error {
	var value Accounts

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
		return fmt.Errorf("Accounts.Scan: cannot scan type %T into Accounts", pSrc)
	}

	*as = value
	return nil
}

// Value implements sql.Valuer interface
func (as Accounts) Value() (value driver.Value, err error) {
	val, err := json.Marshal(as)
	return string(val), errors.WithStack(err)
}
