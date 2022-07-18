package entities

type Bigfile struct {
	Name      string `json:"name"`
	Link      string `json:"link"`
	LocalLink string
	Size      int
	FileBytes []byte
}
