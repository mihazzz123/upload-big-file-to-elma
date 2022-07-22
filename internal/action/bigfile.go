package action

import (
	"context"
	"github.com/mihazzz123/upload-big-file-to-elma/internal/action/di"
	"github.com/mihazzz123/upload-big-file-to-elma/internal/model"
)

func CreateBigfile(bf *model.Bigfile, di di.Container) (*model.Bigfile, error) {
	bigfile, err := model.NewBigfile(bf)
	if err != nil {
		return nil, err
	}
	return bigfile, nil
}

func DownloadByLink(bf *model.Bigfile, di di.Container) (*model.Bigfile, error) {
	if err := di.GetBigfileRepository().DownloadByLink(bf); err != nil {
		return nil, err
	}
	return bf, nil
}

func SendFileToElma(ctx context.Context, bf *model.Bigfile, di di.Container) error {
	return di.GetBigfileRepository().SendFileToElma(ctx, bf, di.GetConfig())
}
