package boolean

//go:generate ../../../../tooling/bin/easyjson view_data.go

// ViewData информация для отображения поля
//
// easyjson:json
type ViewData struct {
	YesValue string `json:"yesValue" patch:"po"`
	NoValue  string `json:"noValue" patch:"po"`
}
