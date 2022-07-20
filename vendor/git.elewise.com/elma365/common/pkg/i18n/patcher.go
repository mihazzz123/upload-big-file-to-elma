package i18n

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"git.elewise.com/elma365/common/pkg/patch"

	"github.com/pkg/errors"
)

// Patcher - автопереводчик ресурсов
type Patcher struct {
	localesDir  string
	defaultLang string
	formatting  bool
}

// Patch - переводит все ресурсы. Ресурсы берутся из папки языка по умолчанию
func Patch(localesDir, defaultLang string) error {
	return makePatch(localesDir, defaultLang, false)
}

// PatchAndFormatting - переводит и форматирует все ресурсы. Ресурсы берутся из папки языка по умолчанию
func PatchAndFormatting(localesDir, defaultLang string) error {
	return makePatch(localesDir, defaultLang, true)
}

func makePatch(localesDir, defaultLang string, formatting bool) error {
	patcher := NewPatcher(localesDir, defaultLang)
	patcher.formatting = formatting
	if err := patcher.PatchJSONFiles(); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// NewPatcher - новый автопереводчик
func NewPatcher(localesDir, defaultLang string) *Patcher {
	return &Patcher{
		localesDir,
		defaultLang,
		false,
	}
}

// PatchJSONFiles - перевести json файлы
func (i18n Patcher) PatchJSONFiles() error {
	jsonsToPatch, err := filepath.Glob(filepath.Join(i18n.localesDir, i18n.defaultLang, "*.json"))
	if err != nil {
		panic(err)
	}
	for i := range jsonsToPatch {
		if err := i18n.patchJSONFile(jsonsToPatch[i]); err != nil {
			return err
		}
	}
	return nil
}

func (i18n Patcher) patchJSONFile(jsonFilePath string) error {
	jsonFileName := filepath.Base(jsonFilePath)
	jsonFile, err := ioutil.ReadFile(jsonFilePath) //nolint:gosec // задается разработчиком
	if err != nil {
		return errors.WithStack(err)
	}
	poFileName := jsonFileName[0:len(jsonFileName)-len("json")] + "po"
	poFilesPaths, err := filepath.Glob(filepath.Join(i18n.localesDir, "*", poFileName))
	if err != nil {
		return errors.WithStack(err)
	}
	for j := range poFilesPaths {
		poFile, err := ioutil.ReadFile(poFilesPaths[j])
		if err != nil {
			return errors.WithStack(err)
		}
		poDir := filepath.Dir(poFilesPaths[j])
		lang := filepath.Base(poDir)
		if lang == i18n.defaultLang {
			continue
		}
		translatedJSONFile, err := patch.ReplacePO(json.RawMessage(jsonFile), poFile)
		if err != nil {
			return errors.WithStack(err)
		}

		if i18n.formatting {
			var formatted bytes.Buffer
			err = json.Indent(&formatted, translatedJSONFile, "", " ")
			if err != nil {
				return errors.WithStack(err)
			}
			translatedJSONFile = formatted.Bytes()
		}

		err = ioutil.WriteFile(filepath.Join(poDir, jsonFileName), translatedJSONFile, os.FileMode(0666))
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}
