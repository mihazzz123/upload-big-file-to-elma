package types

import (
	"encoding/json"

	"git.elewise.com/elma365/common/pkg/types/complextypes/file"
)

// NewFileType конструктор
func NewFileType() TypeImplement {
	return &fileType{}
}

type fileType struct {
	defaultType
	canReplaceCheckArrayAndSingle
	canReplaceData
	canReplaceViewData
}

func (f *fileType) Validate(value json.RawMessage) (interface{}, error) {
	var res file.File
	err := json.Unmarshal(value, &res)
	if err != nil {
		var tmpRes file.FileTmp
		err = json.Unmarshal(value, &tmpRes)
		return tmpRes, err
	}
	return res, err
}
