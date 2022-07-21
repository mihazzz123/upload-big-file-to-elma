package action

import (
	"github.com/mihazzz123/upload-big-file-to-elma/internal/action/di"
	"github.com/mihazzz123/upload-big-file-to-elma/internal/model"
)

func CreateBigfile(bf *model.Bigfile, di di.Container) error {
	bigfile, err := model.NewBigfile(bf)
	if err != nil {
		return err
	}

	if err = di.GetBigfileRepository().DownloadByLink(bigfile); err != nil {
		return err
	}

	if err = di.GetBigfileRepository().SendFileToElma(bigfile, di.GetConfig()); err != nil {
		return err
	}

	return nil
}
