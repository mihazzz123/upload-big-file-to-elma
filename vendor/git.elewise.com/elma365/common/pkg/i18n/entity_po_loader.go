package i18n

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"

	"git.elewise.com/elma365/common/pkg/errs"
	"git.elewise.com/elma365/common/pkg/namespace"
)

const (
	// POFilesDirNameID папка для переводов
	POFilesDirNameID = "67698d44-0a99-52c2-86d4-e0084affc24c" // _po_files
)

// EntityPOLoader загрузчик po-файлов (файлов с переводами)
type EntityPOLoader interface {
	// GetPO получить переводы из раздела для текущей локали в контексте
	GetPO(ctx context.Context, ns namespace.Namespace, locale string) (io.Reader, error)
}

// NewEntityPOLoader конструктор для POLoader
func NewEntityPOLoader(diskSrv DiskService, httpClient *http.Client) EntityPOLoader {
	return poLoader{
		diskSrv:    diskSrv,
		httpClient: httpClient,
	}
}

// httpClient - лишняя зависимость нужно убирать во всех сервисах
type poLoader struct {
	diskSrv    DiskService
	httpClient *http.Client
}

// nolint: gocyclo // все нормально тут
func (pl poLoader) GetPO(ctx context.Context, ns namespace.Namespace, locale string) (io.Reader, error) {
	if locale == "" {
		return nil, nil
	}
	res := &bytes.Buffer{}
	var data []byte

	nsDirID := ns.GetDirID()
	if uuid.Equal(nsDirID, uuid.Nil) {
		return nil, errors.WithStack(
			errs.NotFound.Newf(`unknown or incorrect namespace "%s"`, ns),
		)
	}
	body, _, err := pl.diskSrv.DownloadFileByFilePath(
		ctx,
		nsDirID,
		fmt.Sprintf("%s/%s.po", POFilesDirNameID, locale),
	)
	if errors.Cause(err) == errs.NotFound {
		err = nil
	}
	if err != nil {
		zap.L().Warn("cant generate download link", zap.String("reason", err.Error()))
		return res, nil
	}
	defer func() {
		if body != nil {
			_ = body.Close()
		}
	}()
	if body != nil {
		data, err = ioutil.ReadAll(body)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		_, err = res.Write(data)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}

	// рекурсивно получаем переводы из родительского namespace
	if ns == namespace.Global || ns == namespace.System {
		return res, nil
	}
	r, err := pl.GetPO(ctx, ns.GetParent(), locale)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if r != nil {
		data, err = ioutil.ReadAll(r)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		_, err = res.Write(data)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}

	return res, nil
}
