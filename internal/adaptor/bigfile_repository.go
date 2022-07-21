package adaptor

import (
	"github.com/mihazzz123/upload-big-file-to-elma/internal/model"
	"io/ioutil"
	"net/http"
)

type bigfileRepository struct {
}

// NewBigfileRepository ...
func NewBigfileRepository() *bigfileRepository {
	return &bigfileRepository{}
}

func SendFileToElma(bf *model.Bigfile) error {

	return nil
}

// DownloadByLink ...
func DownloadByLink(bf *model.Bigfile) error {
	resp, err := http.Get(bf.Link)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
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
