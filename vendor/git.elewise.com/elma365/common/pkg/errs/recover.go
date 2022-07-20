package errs

import (
	"fmt"

	"github.com/pkg/errors"
)

type withStack interface {
	StackTrace() errors.StackTrace
}

// WithRecover отлавливает паники внутри себя и преобразовывает в ошибку
func WithRecover(fn func()) (err error) {
	defer func() {
		var ok bool

		rec := recover()
		if rec != nil {
			if err, ok = rec.(error); !ok {
				err = Internal.New(fmt.Sprint(rec))
			} else if _, ok := err.(withStack); !ok {
				err = errors.WithStack(err)
			}
		}
	}()

	fn()

	return
}
