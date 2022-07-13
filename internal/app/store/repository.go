package store

import "github.com/mihazzz123/upload-big-file-to-elma/internal/app/model"

// BigFileRepository ...
type BigFileRepository interface {
	Create(bigfile *model.Bigfile) error
}
