package models

//go:generate ../../../tooling/bin/easyjson $GOFILE

import (
	i18n "git.elewise.com/elma365/common/pkg/i18n/models"

	uuid "github.com/satori/go.uuid"
)

// LanguageFile информация о загруженном файле языкового пакета
// easyjson:json
type LanguageFile struct {
	Hash        uuid.UUID `json:"hash"`
	ServiceName string    `json:"serviceName"`
}

// LanguageSettings информация о установленном языковом модуле
// easyjson:json
type LanguageSettings struct {
	LocaleInfo i18n.LocaleInfo `json:"localeInfo"`
	POfiles    []LanguageFile  `json:"poFiles"`
}
