package bigfilestore

import "upload-big-file-to-elma/internal/app/model"

type BigFileRepository struct {
	store *Store
}

func (b *BigFileRepository) SaveLocal(file *model.Bigfile) error {
	return nil
}
