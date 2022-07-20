package types

import (
	"encoding/json"

	"git.elewise.com/elma365/common/pkg/types/complextypes/sysref"
)

// NewSysOSNodeType конструктор
func NewSysOSNodeType() TypeImplement {
	return &sysOSNodeType{}
}

type sysOSNodeType struct {
	defaultType
	canReplaceCheckArrayAndSingle
	canReplaceData
	canReplaceViewData
}

func (s *sysOSNodeType) Validate(value json.RawMessage) (interface{}, error) {
	var res sysref.OSNode
	err := json.Unmarshal(value, &res)
	return res, err
}
