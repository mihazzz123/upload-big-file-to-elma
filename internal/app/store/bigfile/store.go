package bigfile

import "github.com/mihazzz123/upload-big-file-to-elma/internal/app/store"

type Store struct {
	id                string
	BigFileRepository *BigFileRepository
}

// New ...
func New() *Store {
	return &Store{
		id: "hi",
	}
}

// Bigfile ...
func (s *Store) Bigfile() store.BigFileRepository {
	if s.BigFileRepository != nil {
		return s.BigFileRepository
	}
	s.BigFileRepository = &BigFileRepository{
		store: s,
	}
	return s.BigFileRepository
}
