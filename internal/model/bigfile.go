package model

import (
	uuid "github.com/satori/go.uuid"
)

// Bigfile ...
type Bigfile struct {
	Uuid      uuid.UUID
	Name      string `json:"name"`
	Link      string `json:"link"`
	Size      int
	FileBytes []byte
}

// NewBigfile ...
func NewBigfile(bigfile *Bigfile) (*Bigfile, error) {
	return &Bigfile{
		Uuid: uuid.NewV4(),
		Name: bigfile.Name,
		Link: bigfile.Link,
	}, nil
}
