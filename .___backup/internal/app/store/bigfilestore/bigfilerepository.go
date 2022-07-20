package bigfilestore

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"upload-big-file-to-elma/___backup/internal/app/model"
)

var (
	errDownloadFile = errors.New("error download file by link")
)

type BigFileRepository struct {
	store *Store
}

// SaveLocal ...
func (b *BigFileRepository) SaveLocal(file *model.Bigfile) error {

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
func (b *BigFileRepository) DownloadFileByLink(file *model.Bigfile) error {

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
func (b *BigFileRepository) DeleteLocalTempFile(file *model.Bigfile) error {
	timeOutDeleteFile := time.Hour * 10
	go func() {
		time.Sleep(timeOutDeleteFile)
		if err := os.Remove(file.LocalLink); err != nil {
			fmt.Errorf("DeleteLocalTempFile error: %v", err)
		}
	}()

	return nil
}
