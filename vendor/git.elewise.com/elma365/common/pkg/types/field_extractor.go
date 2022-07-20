package types

import (
	"encoding/json"

	"git.elewise.com/elma365/common/pkg/errs"
	"git.elewise.com/elma365/common/pkg/patch"

	"github.com/pkg/errors"
)

// GetExtracts - реализация интерфейса patch.Extractor
func (f Field) GetExtracts(target string) (patch.Extracts, error) {
	extracts := patch.Extracts{}

	baseExtracts := f.getBaseValueExtracts()
	// Без префикса, т.к. это базовые свойства.
	extracts = append(extracts, baseExtracts...)

	viewExtracts, err := f.getViewExtracts(target)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	extracts = append(extracts, viewExtracts.Prefix("view")...)

	dataExtracts, err := f.getDataExtracts(target)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	extracts = append(extracts, dataExtracts.Prefix("data")...)
	defaultsExtracts := f.getDefaultValueExtracts()
	extracts = append(extracts, defaultsExtracts.Prefix("defaultValue")...)

	return extracts, nil
}

func (f Field) getDataExtracts(target string) (patch.Extracts, error) {
	if f.Data == nil {
		return nil, nil
	}
	data, err := f.GetData()
	if err != nil && errors.Cause(err) != errs.NotImplemented {
		return nil, errors.WithStack(err)
	}
	if data == nil {
		return nil, nil
	}
	return patch.CreateExtracts(data, target)
}

func (f Field) getViewExtracts(target string) (patch.Extracts, error) {
	extracts, err := patch.CreateExtracts(f.View, target)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	viewData, err := f.GetViewData()
	if err != nil && errors.Cause(err) != errs.NotImplemented {
		return nil, errors.WithStack(err)
	}
	if viewData != nil {
		viewDataExtracts, err := patch.CreateExtracts(viewData, target)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		prefixedDataExtracts := viewDataExtracts.Prefix("data")
		extracts = append(extracts, prefixedDataExtracts...)
	}

	return extracts, nil
}

func (f Field) getDefaultValueExtracts() patch.Extracts {
	if f.Type == Money && len(f.Default) != 0 && len(f.Default) != len(json.RawMessage("\"\"")) {
		return patch.Extracts{patch.NewExtract("currency")}
	}

	return nil
}

func (f Field) getBaseValueExtracts() patch.Extracts {
	if f.Type == String && f.Formula != "" {
		return patch.Extracts{patch.NewExtract("formula")}
	}

	return nil
}
