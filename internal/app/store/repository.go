package store

import (
	"upload-big-file-to-elma/internal/app/model"
)

// BigFileRepository ...
type BigFileRepository interface {
	SaveLocal(file *model.Bigfile) error
	DownloadFileByLink(file *model.Bigfile) error
	DeleteLocalTempFile(file *model.Bigfile) error
}
