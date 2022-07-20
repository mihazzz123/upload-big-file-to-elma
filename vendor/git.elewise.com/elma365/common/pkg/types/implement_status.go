package types

import (
	"encoding/json"

	"git.elewise.com/elma365/common/pkg/types/complextypes/status"
)

// NewStatusType конструктор
func NewStatusType() TypeImplement {
	return &statusType{}
}

type statusType struct {
	defaultType
	canReplaceCheckArrayAndSingle
	canReplaceData
	canReplaceViewData
}

func (s *statusType) Validate(value json.RawMessage) (interface{}, error) {
	var res status.Status
	err := json.Unmarshal(value, &res)

	return res, err
}
