package errs

import "github.com/pkg/errors"

// PanicIfError - если есть ошибка, то бросается паника. Дополнительно можно передать сообщение
func PanicIfError(err error, msg string, args ...interface{}) {
	if err != nil {
		panic(errors.Wrapf(err, msg, args...))
	}
}
