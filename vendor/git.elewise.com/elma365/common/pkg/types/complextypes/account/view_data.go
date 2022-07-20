package account

// Type тип мессенджера
type Type string

// Типы мессенджеров
const (
	TypeTelegram       Type = "telegram"
	TypeWhatsapp       Type = "whatsapp"
	TypeInstagram      Type = "instagram"
	TypeViber          Type = "viber"
	TypeFacebook       Type = "facebook"
	TypeVkontakte      Type = "vkontakte"
	TypeSkype          Type = "skype"
	TypeGoogleTalk     Type = "google-talk"
	TypeGoogleHangouts Type = "google-hangouts"
	TypeWeChat         Type = "we-chat"
	TypeSnapchat       Type = "snapchat"
	TypeLine           Type = "line"
)

// ViewData описание представления поля типа Account
type ViewData struct {
	Type Type `json:"type,omitempty"`
}
