package email

import (
	"context"
	"fmt"

	"git.elewise.com/elma365/common/pkg/errs/validation"
)

//go:generate ../../../../tooling/bin/easyjson email.go

// Email is a tuple of email information
//
// easyjson:json
type Email struct {
	Type  string `json:"type"`
	Email string `json:"email" validate:"email"`
}

// Validate email field type
func (e Email) Validate() error {
	errs := validation.ValidateStruct(context.Background(), e)
	if !errs.IsEmpty() {
		return fmt.Errorf("invalid EMAIL <%s>", e.Email)
	}

	return nil
}

// String implements fmt.Stringer interface
func (e Email) String() string {
	return e.Email
}
