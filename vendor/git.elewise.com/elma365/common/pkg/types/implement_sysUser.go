package types

import (
	"encoding/json"

	"git.elewise.com/elma365/common/pkg/types/complextypes/sysref"
)

// NewSysUserType конструктор
func NewSysUserType() TypeImplement {
	return &sysUserType{}
}

type sysUserType struct {
	defaultType
	canReplaceCheckArrayAndSingle
	canReplaceData
	canReplaceViewData
}

func (s *sysUserType) Validate(value json.RawMessage) (interface{}, error) {
	var res sysref.User
	err := json.Unmarshal(value, &res)
	return res, err
}
