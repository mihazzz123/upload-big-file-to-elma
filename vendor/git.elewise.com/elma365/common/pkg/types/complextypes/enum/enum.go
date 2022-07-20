package enum

//go:generate ../../../../tooling/bin/easyjson enum.go

// Enum - значение для типа Enum
//
// easyjson:json
type Enum struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
