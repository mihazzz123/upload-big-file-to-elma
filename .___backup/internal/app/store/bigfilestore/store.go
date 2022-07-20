package bigfilestore

import (
	_ "github.com/lib/pq" // ...
	"upload-big-file-to-elma/___backup/internal/app/store"
)

type Store struct {
	BigFileRepository *BigFileRepository
}

// New ...
func New() *Store {
	return &Store{}
}

// BigFile ...
func (s *Store) BigFile() store.BigFileRepository {
	if s.BigFileRepository != nil {
		return s.BigFileRepository
	}
	s.BigFileRepository = &BigFileRepository{
		store: s,
	}
	return s.BigFileRepository
}
