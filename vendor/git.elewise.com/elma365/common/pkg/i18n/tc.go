package i18n

//go:generate ../../tooling/bin/easyjson $GOFILE

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

// easyjson:json
type locale struct {
	Name string `json:"name"`
}

// DefaultTCUrl - адрес проекта elma365 в TC
const DefaultTCUrl = "http://tc.elma-bpm.com/api/elma365"

// DefaultLocalesPath - рекомендуемый относительный путь до папки с локализациями
const DefaultLocalesPath = "data/locales"

// LoadFromTC загружает локализации для сервиса из TC
// Параметр tcURL - путь к TC (можно использовать константу DefaultTCUrl для проекта elma365)
// Параметр path - относительный путь до папки с языками (можно использовать константу DefaultLocalesPath)
// Параметр serviceName - имя сервиса (имя папки в архиве с локализацией)
//
// Пример использования для загрузки локализаций сервиса:
// ```
// func main() {
//	 if len(os.Args) < 2 {
// 		panic(errors.New("Service name is required as 1st argument"))
//	 }
//	 serviceName := os.Args[1]
//	 if err := i18n.LoadFromTC(i18n.DefaultTCUrl, i18n.DefaultLocalesPath, serviceName); err != nil {
// 		panic(err)
//	 }
// }
// ```
func LoadFromTC(tcURL, path, serviceName string) error {
	response, err := http.Get(tcURL) //nolint:gosec // не используется в production
	if err != nil {
		return errors.Wrap(err, "Unable to load locales from TC")
	}
	defer func() {
		_ = response.Body.Close()
	}()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return errors.Wrap(err, "Unable to read locales from TC")
	}
	var locales []locale
	err = json.Unmarshal(content, &locales)
	if err != nil {
		return errors.Wrap(err, "Unable to parse locales from TC")
	}
	for _, localeItem := range locales {
		if strings.ToLower(localeItem.Name) != "ru-ru" {
			languageReader, err := loadLanguage(tcURL, localeItem.Name)
			if err != nil {
				return errors.Wrapf(err, "Unable to load language: %s", localeItem.Name)
			}
			err = unpackLanguage(languageReader, path, serviceName, localeItem.Name)
			if err != nil {
				return errors.Wrapf(err, "Unable to unpack language: %s", localeItem.Name)
			}
		}
	}
	return nil
}

func loadLanguage(tcURL, lang string) (*zip.Reader, error) {
	// Загружаем архив с локализацией
	langURL := tcURL + "/" + lang
	response, err := http.Get(langURL) //nolint:gosec
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to load locale from url: %s", langURL)
	}
	defer func() {
		_ = response.Body.Close()
	}()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to read locale from response")
	}
	zipReader, err := zip.NewReader(bytes.NewReader(content), int64(len(content)))
	if err != nil {
		return nil, errors.Wrap(err, "Unable to create zip reader from response")
	}
	return zipReader, nil
}

func unpackLanguage(languageReader *zip.Reader, path, serviceName, lang string) error {
	// Создаём папку языка
	fullPath := filepath.Join(path, lang)
	if err := createDirIfNotExist(fullPath); err != nil {
		return err
	}

	for _, zipFile := range languageReader.File {
		if strings.Index(zipFile.Name, serviceName+"/") == 0 {
			if err := unpackFile(zipFile, serviceName, fullPath); err != nil {
				return errors.Wrapf(err, "Unable to unpack file: %s", zipFile.Name)
			}
		}
	}
	return nil
}

func unpackFile(zipFile *zip.File, serviceName, fullPath string) error {
	relativePath := zipFile.Name[len(serviceName)+1:]
	if relativePath == "" || relativePath[len(relativePath)-3:] == ".mo" {
		// Корневая папка или mo-файл
		return nil
	}
	if relativePath[len(relativePath)-1:] == "/" {
		// Папка
		err := createDirIfNotExist(filepath.Join(fullPath, relativePath[0:len(relativePath)-1]))
		if err != nil {
			return err
		}
		return nil
	}

	// Файл
	fileReader, err := zipFile.Open()
	if err != nil {
		return errors.Wrap(err, "Unable to open zip file")
	}
	defer func() {
		_ = fileReader.Close()
	}()
	osFile, err := createFile(filepath.Join(fullPath, relativePath))
	if err != nil {
		return err
	}
	defer func() {
		_ = osFile.Close()
	}()
	//nolint:gosec // мы доверяем источнику
	if _, err := io.Copy(osFile, fileReader); err != nil {
		return errors.Wrap(err, "Unable to write file")
	}
	return nil
}

func dirOrFileExist(fullPath string) (bool, error) {
	_, err := os.Stat(fullPath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func createDirIfNotExist(fullPath string) error {
	dirExist, err := dirOrFileExist(fullPath)
	if err != nil {
		return err
	}
	if !dirExist {
		if err := os.Mkdir(fullPath, os.ModePerm); err != nil {
			return errors.Wrapf(err, "Unable to make directory: %s", fullPath)
		}
	}
	return nil
}

func createFile(fullPath string) (*os.File, error) {
	// удаляем файл если он существует
	fileExist, err := dirOrFileExist(fullPath)
	if err != nil {
		return nil, err
	}
	if fileExist {
		if err = os.Remove(fullPath); err != nil {
			return nil, errors.Wrapf(err, "Unable to delete file \"%s\"", fullPath)
		}
	}

	osFile, err := os.Create(fullPath)
	if err != nil {
		if osFile != nil {
			_ = osFile.Close()
		}
		return nil, errors.Wrapf(err, "Unable to create file \"%s\"", fullPath)
	}
	return osFile, nil
}
