package adaptor

import (
	"bytes"
	"context"
	"fmt"
	"github.com/mihazzz123/upload-big-file-to-elma/internal/config"
	"github.com/mihazzz123/upload-big-file-to-elma/internal/model"
	"go.uber.org/zap"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"time"
)

type bigfileRepository struct {
}

// NewBigfileRepository ...
func NewBigfileRepository() *bigfileRepository {
	return &bigfileRepository{}
}

// SendFileToElma ...
func (r bigfileRepository) SendFileToElma(ctx context.Context, bf *model.Bigfile, cfgElma *config.Config) error {

	//go func(bf *model.Bigfile) {
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
			zap.L().Fatal("multiPartWriter", zap.Error(err))
			return err
		}

		_, err = fw.Write(bf.FileBytes[from:to])
		if err != nil {
			zap.L().Fatal("fw.Write", zap.Error(err))
			return err
		}
		multiPartWriter.Close()

		url := fmt.Sprintf("%s%s", cfgElma.GetElmaURL(), bf.Uuid)
		request, err := http.NewRequest("POST", url, &req)
		if err != nil {
			zap.L().Fatal("NewRequest", zap.Error(err))
			return err
		}

		request.Header.Set("Content-Type", multiPartWriter.FormDataContentType())
		request.Header.Add("X-Token", cfgElma.Token)
		request.Header.Add("Content-Range", contentRange)

		timeout := time.Duration(300) * time.Second

		client := &http.Client{
			Timeout: timeout,
		}
		response, err := client.Do(request)
		if err != nil {
			zap.L().Fatal("client.Do", zap.Error(err))
			return err
		}

		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			zap.L().Fatal("ioutil.ReadAll", zap.Error(err))
			return err
		}

		if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusPartialContent {
			zap.L().Fatal(fmt.Sprintf("request failed with status: %s, body: %s", response.Status, string(body)), zap.Error(err))
			return err
		}

		zap.L().Info("part of the file has been uploaded successfully",
			zap.String("from-to/full", contentRange),
			zap.String("name", bf.Name),
			zap.String("uuid", bf.Uuid.String()),
			zap.Int("size", bf.Size),
			zap.String("link", bf.Link),
		)

	}

	zap.L().Info("the file has been uploaded successfully",
		zap.String("name", bf.Name),
		zap.String("uuid", bf.Uuid.String()),
		zap.Int("size", bf.Size),
		zap.String("link", bf.Link),
	)

	return nil
}

// DownloadByLink ...
func (r bigfileRepository) DownloadByLink(bf *model.Bigfile) error {

	resp, err := http.Get(bf.Link)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		zap.L().Error(fmt.Sprintf("error download file by link. code: %d status: %s", resp.StatusCode, resp.Status), zap.Error(err))
		return err
	}

	zap.L().Info("file download beginning",
		zap.String("name", bf.Name),
		zap.String("uuid", fmt.Sprintln(bf.Uuid)),
		zap.String("link", bf.Link),
	)

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	bf.FileBytes = bytes
	bf.Size = len(bytes)

	zap.L().Info("file download complete",
		zap.String("name", bf.Name),
		zap.String("uuid", fmt.Sprintln(bf.Uuid)),
		zap.String("size", fmt.Sprintln(bf.Size)),
		zap.String("link", bf.Link),
	)
	return nil
}
