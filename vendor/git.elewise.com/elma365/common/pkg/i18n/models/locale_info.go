package models

//go:generate ../../../tooling/bin/easyjson $GOFILE

// Языковые модули

// LocaleInfo описание локали
// easyjson:json
type LocaleInfo struct {
	Name string `json:"name"`
	Code string `json:"code"`
}
