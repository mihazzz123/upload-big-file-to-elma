package sqldb

import (
	"fmt"

	"git.elewise.com/elma365/common/pkg/errs"

	"github.com/pkg/errors"
)

const (
	// ErrNoRows ошибка возникающая когда запрос не возвращает результата.
	ErrNoRows errorCode = "no rows in result set"
	// ErrUniqueViolation ошибка возникающая когда происходит нарушение уникальности данных.
	ErrUniqueViolation errorCode = "unique violation"
	// ErrNotNullViolation ошибка возникающая когда происходит нарушение условия NOT NULL.
	ErrNotNullViolation errorCode = "not null violation"
	// ErrUndefinedTable ошибка возникающая когда таблица запроса не определена.
	ErrUndefinedTable errorCode = "undefined table"
)

// errorCode код ошибки БД.
type errorCode string

func (e errorCode) Error() string {
	return string(e)
}

func (e errorCode) Wrap(err error) error {
	if err == nil {
		return nil
	}
	err = errors.WithStack(err)

	return errorWithCode{orig: err, code: e}
}

// errorWithCode ошибка содержащая код ошибки.
type errorWithCode struct {
	orig error
	code errorCode
}

func (ewc errorWithCode) Cause() error {
	return ewc.code
}

func (ewc errorWithCode) Error() string {
	return fmt.Sprintf("%s: %s", ewc.code.Error(), ewc.orig.Error())
}

// FromDBError преобразует ошибки БД во внутренний формат common/pkg/errs.
func FromDBError(dbErr error) error {
	if dbErr == nil {
		return nil
	}
	if errors.Cause(dbErr) == ErrNoRows {
		return errs.NotFound.Wrap(dbErr)
	}
	if errors.Cause(dbErr) == ErrUniqueViolation {
		return errs.Collision.Wrap(dbErr)
	}

	return dbErr
}
