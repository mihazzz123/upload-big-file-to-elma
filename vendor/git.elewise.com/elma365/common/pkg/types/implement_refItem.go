package types

import (
	"encoding/json"

	"git.elewise.com/elma365/common/pkg/types/complextypes/refitem"
)

// NewRefItemType конструтор
func NewRefItemType() TypeImplement {
	return &refItemType{}
}

type refItemType struct {
	defaultType
	canReplaceCheckArrayAndSingle
	canReplaceData
	canReplaceViewData
}

func (r *refItemType) Validate(value json.RawMessage) (interface{}, error) {
	var res refitem.RefItem
	err := json.Unmarshal(value, &res)

	if err != nil {
		var arrRes []refitem.RefItem
		err = json.Unmarshal(value, &arrRes)

		return arrRes, err
	}

	return res, err
}
