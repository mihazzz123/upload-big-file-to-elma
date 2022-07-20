package types

import (
	"encoding/json"

	"git.elewise.com/elma365/common/pkg/types/complextypes/link"
)

// NewLinkType конструктор
func NewLinkType() TypeImplement {
	return &linkType{}
}

type linkType struct {
	defaultType
	canReplaceCheckArrayAndSingle
	canReplaceData
	canReplaceViewData
}

func (l *linkType) Validate(value json.RawMessage) (interface{}, error) {
	var res link.Link
	err := json.Unmarshal(value, &res)
	return res, err
}
