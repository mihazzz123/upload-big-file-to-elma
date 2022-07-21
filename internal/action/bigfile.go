package action

import (
	"github.com/mihazzz123/upload-big-file-to-elma/internal/action/di"
	"github.com/mihazzz123/upload-big-file-to-elma/internal/adaptor"
	"github.com/mihazzz123/upload-big-file-to-elma/internal/model"
)

func CreateBigfile(bf *model.Bigfile, di adaptor.DIContainer) error {
	bf, err := model.NewBigfile(bf)
	if err != nil {
		return err
	}

	if err = di DownloadByLink(); err != nil {
		return err
	}

	if err = adaptor.SendFileToElma(bf); err != nil {
		return err
	}

	return nil
}
