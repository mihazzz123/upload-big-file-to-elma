package types

import (
	"encoding/json"

	"git.elewise.com/elma365/common/pkg/types/complextypes/role"
)

// RoleType тип Роль
type RoleType struct {
	defaultType
	canReplaceCheckArrayAndSingle
	canReplaceData
	canReplaceViewData
}

// NewRoleType конструктор
func NewRoleType() TypeImplement {
	return &RoleType{}
}

// Validate Валидация
func (s *RoleType) Validate(value json.RawMessage) (interface{}, error) {
	var res role.Role
	err := json.Unmarshal(value, &res)
	return res, err
}
