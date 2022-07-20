package models

import uuid "github.com/satori/go.uuid"

// CommonOptions опции diskhelper'а
type CommonOptions interface {
	SetContentType(v string)
}

// DownloadOptionsFunc опции для скачивания файлов
type DownloadOptionsFunc func(options DownloadOptions)

// DownloadOptions опции diskhelper'а
type DownloadOptions interface {
	CommonOptions
	SetFilename(v string)
}

// UploadOptionsFunc опции для загрузки файлов
type UploadOptionsFunc func(options UploadOptions)

// UploadOptions опции diskhelper'а
type UploadOptions interface {
	CommonOptions
	SetSize(v int64)
	SetPredefinedBodyHash(v uuid.UUID)
	SetPersist(v bool)
}
