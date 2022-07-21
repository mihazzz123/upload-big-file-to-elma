package adaptor

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/mihazzz123/upload-big-file-to-elma/internal/config"
	"github.com/mihazzz123/upload-big-file-to-elma/internal/model"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

type bigfileRepository struct {
}

var (
	errDownloadFile = errors.New("error download file by link")
)

// NewBigfileRepository ...
func NewBigfileRepository() *bigfileRepository {
	return &bigfileRepository{}
}

// SendFileToElma ...
func (r bigfileRepository) SendFileToElma(bf *model.Bigfile, cfgElma *config.Config) error {

	go func(bf *model.Bigfile) {
		partSize := cfgElma.DefaultPartSize
		defaultPartSize := partSize * 1024 * 1024
		partCount := bf.Size/defaultPartSize + 1

		for i := 0; i < partCount; i++ {
			full := bf.Size
			from := defaultPartSize * i
			to := defaultPartSize * (i + 1)
			if to > full {
				to = full
			}

			req := bytes.Buffer{}
			contentRange := fmt.Sprintf("bytes %d-%d/%d", from, to, full)

			multiPartWriter := multipart.NewWriter(&req)
			fw, err := multiPartWriter.CreateFormFile("file", bf.Name)
			if err != nil {
				return
			}

			defer multiPartWriter.Close()

			_, err = fw.Write(bf.FileBytes[from:to])
			if err != nil {
				return
			}

			url := fmt.Sprintf("%s%s", cfgElma.GetElmaURL(), bf.Uuid)
			request, err := http.NewRequest("POST", url, &req)
			if err != nil {
				return
			}

			request.Header.Set("Content-Type", multiPartWriter.FormDataContentType())
			request.Header.Add("X-Token", cfgElma.Token)
			request.Header.Add("Content-Range", contentRange)

			client := &http.Client{}
			response, err := client.Do(request)
			if err != nil {
				return
			}

			defer response.Body.Close()

			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				return
			}

			if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusPartialContent {
				fmt.Errorf("request failed with status: %s, body: %s", response.Status, string(body))
				return
			}

		}

	}(bf)

	return nil
}

// DownloadByLink ...
func (r bigfileRepository) DownloadByLink(bf *model.Bigfile) error {
	resp, err := http.Get(bf.Link)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errDownloadFile
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	bf.FileBytes = bytes
	bf.Size = len(bytes)

	return nil
}
