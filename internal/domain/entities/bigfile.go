package entities

type Bigfile struct {
	Name      string `json:"name"`
	Link      string `json:"link"`
	Uuid      string
	LocalLink string
	Size      int
	FileBytes []byte
}
