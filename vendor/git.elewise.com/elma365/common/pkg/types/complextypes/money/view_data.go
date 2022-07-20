package money

//go:generate ../../../../tooling/bin/easyjson view_data.go

// ViewData информация для отображения поля
//
// easyjson:json
type ViewData struct {
	Currency string `json:"currency" patch:"po"`
}
