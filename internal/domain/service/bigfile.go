package service

import "upload-big-file-to-elma/internal/domain/entities"

// BigfileStorage ...
type BigfileStorage interface {
	SaveLocal(file *entities.Bigfile) error
	DownloadFileByLink(file *entities.Bigfile) error
	DeleteLocalTempFile(file *entities.Bigfile) error
}

type bigfileService struct {
	storage BigfileStorage
}

func NewBigfileService(storage BigfileStorage) *bigfileService {
	return &bigfileService{
		storage: storage,
	}
}

// SaveLocal ...
func (s *bigfileService) SaveLocal(file *entities.Bigfile) error {
	return s.storage.SaveLocal(file)
}

// DownloadFileByLink ...
func (s *bigfileService) DownloadFileByLink(file *entities.Bigfile) error {
	return s.storage.DownloadFileByLink(file)
}

// DeleteLocalTempFile ...
func (s *bigfileService) DeleteLocalTempFile(file *entities.Bigfile) error {
	return s.storage.DownloadFileByLink(file)
}
