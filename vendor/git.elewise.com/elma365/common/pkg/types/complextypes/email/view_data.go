package email

// Type тип почтового адреса
type Type string

// Типы почтовых адресов
const (
	TypeHome   Type = "home"
	TypeWork   Type = "work"
	TypeMobile Type = "mobile"
	TypeMain   Type = "main"
)

// ViewData описание представления поля типа EMAIL
type ViewData struct {
	Type Type `json:"type"`
}
