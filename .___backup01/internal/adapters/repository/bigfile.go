package repository

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"upload-big-file-to-elma/internal/domain/entities"
)

type bigfileStorage struct{}

var (
	errDownloadFile = errors.New("error download file by link")
)

func NewBigfileStorage() *bigfileStorage {
	return &bigfileStorage{}
}

// SaveLocal ...
func (b *bigfileStorage) SaveLocal(file *entities.Bigfile) error {

	permissions := 0644
	filename := "tempfile/" + time.Now().Format(time.RFC3339Nano) + file.Name

	if err := ioutil.WriteFile(filename, file.FileBytes, os.FileMode(permissions)); err != nil {
		return err
	}

	file.LocalLink = filename
	file.Size = len(file.FileBytes)

	return nil
}

// DownloadFileByLink ...
func (b *bigfileStorage) DownloadFileByLink(file *entities.Bigfile) error {

	resp, err := http.Get(file.Link)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errDownloadFile
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	file.FileBytes = bytes

	return nil
}

// DeleteLocalTempFile ...
func (b *bigfileStorage) DeleteLocalTempFile(file *entities.Bigfile) error {
	timeOutDeleteFile := time.Hour * 10
	go func() {
		time.Sleep(timeOutDeleteFile)
		if err := os.Remove(file.LocalLink); err != nil {
			fmt.Errorf("DeleteLocalTempFile error: %v", err)
		}
	}()

	return nil
}
