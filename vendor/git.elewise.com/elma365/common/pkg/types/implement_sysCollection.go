package types

import (
	"encoding/json"

	"github.com/pkg/errors"

	"git.elewise.com/elma365/common/pkg/types/complextypes/syscollection"
)

// NewSysCollectionType конструктор
func NewSysCollectionType() TypeImplement {
	return &sysCollectionType{}
}

type sysCollectionType struct {
	defaultType
	canReplaceCheckArrayAndSingle
	canReplaceViewData
}

func (s *sysCollectionType) Validate(value json.RawMessage) (interface{}, error) {
	var res string
	err := json.Unmarshal(value, &res)

	return res, err
}

func (s *sysCollectionType) UnmarshalData(rawData json.RawMessage) (interface{}, error) {
	return s.unmarshalData(rawData)
}

func (s *sysCollectionType) unmarshalData(rawData json.RawMessage) (syscollection.Data, error) {
	data := syscollection.Data{}
	if len(rawData) == 0 {
		return data, nil
	}
	err := json.Unmarshal(rawData, &data)
	if err != nil {
		return data, errors.WithStack(err)
	}
	return data, nil
}

func (s *sysCollectionType) CanReplaceData(source, candidate json.RawMessage) (bool, error) {
	sourceData, err := s.unmarshalData(source)
	if err != nil {
		return false, errors.WithStack(err)
	}
	candidateData, err := s.unmarshalData(candidate)
	if err != nil {
		return false, errors.WithStack(err)
	}

	if sourceData.Namespace != candidateData.Namespace || sourceData.Code != candidateData.Code {
		return false, nil
	}
	return true, nil
}
