package types

import (
	"encoding/json"

	"git.elewise.com/elma365/common/pkg/types/complextypes/email"

	"github.com/pkg/errors"
)

// NewEmailType конструктор
func NewEmailType() TypeImplement {
	return &emailType{}
}

type emailType struct {
	defaultType
	canReplaceCheckArrayAndSingle
	canReplaceData
	canReplaceViewDataCheckType
}

func (e *emailType) Validate(value json.RawMessage) (interface{}, error) {
	var res email.Email
	err := json.Unmarshal(value, &res)
	if err != nil {
		return nil, err
	}
	if res.Email == "" {
		return res, nil
	}

	err = res.Validate()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return res, nil
}

func (e *emailType) UnmarshalViewData(rawData json.RawMessage) (interface{}, error) {
	res := email.ViewData{}
	if len(rawData) == 0 {
		return res, nil
	}
	err := json.Unmarshal(rawData, &res)
	return res, errors.WithStack(err)
}
