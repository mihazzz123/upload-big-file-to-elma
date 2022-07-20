package validation

import (
	"fmt"
	"path"
)

// Option опции ошибки валидации
type Option interface {
	Apply(Error) Error
}

type withLevel struct {
	level Level
}

// WithLevel установить уровень ошибки
func WithLevel(level Level) Option {
	return withLevel{level}
}

func (wl withLevel) Apply(err Error) Error {
	err.Level = wl.level

	return err
}

type withPath struct {
	path string
}

// WithPath установить путь ошибки
func WithPath(chunks ...interface{}) Option {
	parts := make([]string, len(chunks))
	for i := range chunks {
		parts[i] = fmt.Sprint(chunks[i])
	}

	return withPath{path.Join(parts...)}
}

func (wp withPath) Apply(err Error) Error {
	err.Path = wp.path

	return err
}

type withArg struct {
	key string
	val interface{}
}

// WithArg добавить аргумент
func WithArg(key string, val interface{}) Option {
	return withArg{key, val}
}

func (wa withArg) Apply(err Error) Error {
	if err.Args == nil {
		err.Args = map[string]interface{}{}
	}
	err.Args[wa.key] = wa.val

	return err
}

type withCause struct {
	err error
}

// WithCause добавить стороннюю критическую ошибку
func WithCause(err error) Option {
	return withCause{err}
}

func (wc withCause) Apply(err Error) Error {
	err.Level = LevelCritical
	err.cause = wc.err

	return err
}
