package file

// File is a inner hash of file body
type File string

// String implements fmt.Stringer interface
func (f File) String() string {
	return string(f)
}

// FileTmp defines temp file structure
// nolint:golint
type FileTmp struct {
	ID   string `json:"id"`
	Hash string `json:"hash"`
	Name string `json:"name"`
}

// String implements fmt.Stringer interface
func (f FileTmp) String() string {
	return f.Name
}
