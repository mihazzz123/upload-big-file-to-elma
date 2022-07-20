package errs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/pkg/errors"
)

// WithData добавляет информацию к ошибке
func WithData(err error, data interface{}) error {
	return errorWithData{
		orig: err,
		data: data,
	}
}

type errorWithData struct {
	orig error
	data interface{}
}

func (ewd errorWithData) Error() string {
	return ewd.orig.Error()
}

func (ewd errorWithData) Cause() error {
	return ewd.orig
}

func (ewd errorWithData) GetData() interface{} {
	return ewd.data
}

func (ewd errorWithData) WriteData(w io.Writer) error {
	je := json.NewEncoder(w)
	if err := je.Encode(ewd.data); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (ewd errorWithData) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			var buf bytes.Buffer
			_ = ewd.WriteData(&buf)
			_, _ = fmt.Fprintf(s, "%s (%s): %+v", ewd.Error(), buf.String(), ewd.orig)
			return
		}
		fallthrough
	case 's':
		_, _ = io.WriteString(s, ewd.Error())
	case 'q':
		_, _ = fmt.Fprintf(s, "%q", ewd.Error())
	}
}

// WriteDataFromError to writer
//
// We walking up by errors chain until find data or chain is ended.
// If chain ended but data is not found, then error will be writed
// as string or as is if it is valid json. If data has been found it will be writed.
func WriteDataFromError(orig error, w io.Writer) error {
	if orig == nil {
		return nil
	}
	type withData interface {
		WriteData(w io.Writer) error
	}
	type withCause interface {
		Cause() error
	}
	err := orig
	for {
		switch terr := err.(type) {
		case withData:
			return terr.WriteData(w)
		case withCause:
			err = terr.Cause()
		default:
			_, err = w.Write([]byte(orig.Error()))
			return errors.WithStack(err)
		}
	}
}
