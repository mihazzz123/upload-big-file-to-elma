package phone

// Type тип телефона
type Type string

// Типы телефонов
const (
	TypeHome    Type = "home"
	TypeWork    Type = "work"
	TypeMobile  Type = "mobile"
	TypeMain    Type = "main"
	TypeHomeFax Type = "home-fax"
	TypeWorkFax Type = "work-fax"
	TypePager   Type = "page"
)

// ViewData описание представления типа
type ViewData struct {
	Type Type `json:"type,omitempty"`
}
