package validation

//go:generate ../../../tooling/bin/easyjson -lower_camel_case -omit_empty $GOFILE

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"path"
	"reflect"
	"regexp"
	"strings"

	"git.elewise.com/elma365/common/pkg/errs"

	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
)

// Global — синглтон валидатора
//
//nolint: gochecknoglobals // экземпляр валидатора, на котором регистрируются кастомные проверки
var Global *validator.Validate

//nolint: gochecknoinits // экземпляр валидатора, на котором регистрируются кастомные проверки
func init() {
	Global = validator.New()
	Global.RegisterTagNameFunc(func(field reflect.StructField) string {
		values := field.Tag.Get("json")
		name := strings.Split(values, ",")[0]
		if name == "-" {
			name = ""
		}

		return name
	})

	codeRE := regexp.MustCompile(`^[A-Za-z_][A-Za-z0-9\-_.]+$`)

	_ = Global.RegisterValidation("code", func(fl validator.FieldLevel) bool {
		return fl.Field().String() == "" || codeRE.MatchString(fl.Field().String())
	})

	userCodeRE := regexp.MustCompile(`^[A-Za-z][A-Za-z0-9\-_.]+$`)

	_ = Global.RegisterValidation("userCode", func(fl validator.FieldLevel) bool {
		return fl.Field().String() == "" || userCodeRE.MatchString(fl.Field().String())
	})

	extendedCodeRE := regexp.MustCompile(`^[A-Za-z_][A-Za-z0-9\-_]+([.][A-Za-z_][A-Za-z0-9\-_]+)?$`)

	_ = Global.RegisterValidation("extendedCode", func(fl validator.FieldLevel) bool {
		return fl.Field().String() == "" || extendedCodeRE.MatchString(fl.Field().String())
	})

	rguCodeRE := regexp.MustCompile(`^global$|^system$|^\d+$`)

	_ = Global.RegisterValidation("rguCode", func(fl validator.FieldLevel) bool {
		return fl.Field().String() == "" || rguCodeRE.MatchString(fl.Field().String())
	})

	fieldCodeRE := regexp.MustCompile(`^[A-Za-z_][A-Za-z0-9\-_]+$`)
	_ = Global.RegisterValidation("fieldCode", func(fl validator.FieldLevel) bool {
		return fl.Field().String() == "" || fieldCodeRE.MatchString(fl.Field().String())
	})

	pathRE := regexp.MustCompile(`^(/[^/\x00]*)+/?$`)

	_ = Global.RegisterValidation("path", func(fl validator.FieldLevel) bool {
		return pathRE.MatchString(fl.Field().String())
	})
}

// Error is an error with path and level
//
// easyjson:json
type Error struct {
	Path  string                 `json:"path"`
	Desc  string                 `json:"desc"`
	Level Level                  `json:"level"`
	Args  map[string]interface{} `json:"args"`
	cause error
}

// NewError constructor
func NewError(desc string, options ...Option) Error {
	err := Error{
		Path:  "/",
		Desc:  desc,
		Level: LevelError,
	}
	for _, opt := range options {
		err = opt.Apply(err)
	}

	return err
}

// Error return wrapped error string
//
// Implements: error
func (ve Error) Error() string {
	return fmt.Sprintf("%s: %s", ve.Path, ve.Desc)
}

// Prefix error with given path
func (ve Error) Prefix(parts ...string) Error {
	parts = append([]string{"/"}, parts...)
	parts = append(parts, ve.Path)
	ve.Path = path.Join(parts...)

	return ve
}

// Cause invalid error
//
// Implements: errors.Causer
func (ve Error) Cause() error {
	if ve.cause != nil {
		return ve.cause
	}

	return errs.InvalidArgument
}

// Errors is a collection of errors of the some struct
type Errors []Error

// NewErrors constructor
func NewErrors(n int) *Errors {
	verrs := make(Errors, 0, n)

	return &verrs
}

// LevelError return marshaled error
//
// Implements: error
func (verrs Errors) Error() string {
	res, _ := json.Marshal(verrs)

	return string(res)
}

// Return nil if empty else self
func (verrs *Errors) Return() error {
	if len(*verrs) == 0 {
		return nil
	}

	return verrs
}

// Len return count of errors
func (verrs Errors) Len() int {
	return len(verrs)
}

// IsEmpty error
func (verrs Errors) IsEmpty() bool {
	return len(verrs) == 0
}

// IsCritical error
func (verrs Errors) IsCritical() bool {
	return len(verrs) > 0 && verrs[len(verrs)-1].Level == LevelCritical
}

// HasLevel check errors for specified error level or less
func (verrs Errors) HasLevel(level Level) bool {
	for _, verr := range verrs {
		if verr.Level <= level {
			return true
		}
	}

	return false
}

// Add some error
func (verrs *Errors) Add(err error) *Errors {
	return verrs.AddWithPrefix(err, "")
}

// AddItem to list
func (verrs *Errors) AddItem(desc string, options ...Option) *Errors {
	return verrs.Add(NewError(desc, options...))
}

// AddWithPrefix like AddError but prefix validation error(s)
func (verrs *Errors) AddWithPrefix(err error, prefix string) *Errors {
	switch verr := err.(type) {
	case nil:
		// pass

	case Error:
		*verrs = append(*verrs, verr.Prefix(prefix))

	case *Errors:
		if verr != nil {
			*verrs = append(*verrs, *verr.Prefix(prefix)...)
		}

	default:
		*verrs = append(*verrs, NewError(err.Error(), WithPath(prefix), WithCause(err)))
	}

	return verrs
}

// Prefix all errors with given path
func (verrs Errors) Prefix(prefix string) *Errors {
	res := make(Errors, len(verrs))
	for i := range verrs {
		res[i] = verrs[i].Prefix(prefix)
	}

	return &res
}

// WithRecover execute callback and recover panic to critical error
func (verrs *Errors) WithRecover(fn func()) *Errors {
	func() {
		defer func() {
			if res := recover(); res != nil {
				if err, ok := res.(error); ok {
					_ = verrs.Add(err)
					(*verrs)[len(*verrs)-1].Level = LevelCritical
				} else {
					_ = verrs.AddItem(fmt.Sprint(res), WithLevel(LevelCritical))
				}
			}
		}()

		fn()
	}()

	return verrs
}

// Cause return critical causer
//
// Implements: errors.Causer
func (verrs Errors) Cause() error {
	if verrs.IsCritical() {
		return verrs[len(verrs)-1]
	}

	return errs.InvalidArgument
}

// WriteData writes data to writer
func (verrs Errors) WriteData(w io.Writer) error {
	je := json.NewEncoder(w)
	if err := je.Encode(verrs); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// ValidateStruct with gopkg.in/go-playground/validator and convert errors to Errors
func ValidateStruct(ctx context.Context, val interface{}) *Errors {
	res := NewErrors(0)
	err := Global.StructCtx(ctx, val)
	if err == nil {
		return res
	}
	verrs, ok := err.(validator.ValidationErrors)
	if !ok {
		return res.Add(errors.WithStack(err))
	}
	res = NewErrors(len(verrs))
	for _, verr := range verrs {
		_ = res.Add(NewError(verr.ActualTag(), WithPath(namespaceToPath(verr.Namespace()))))
	}

	return res
}

var indexRE = regexp.MustCompile(`\[([^\]]+)\]`)

func namespaceToPath(ns string) string {
	indexNormalized := indexRE.ReplaceAllString(ns, ".$1")
	chunks := strings.Split(indexNormalized, ".")
	chunks[0] = "" // Remove type name and ensure that start with /

	return strings.Join(chunks, "/")
}

// ErrorsIterator implements Next/Value interface
type ErrorsIterator struct {
	records Errors
	index   int
}

// Iter return new iterator
func (verrs *Errors) Iter() *ErrorsIterator {
	return &ErrorsIterator{
		records: *verrs,
		index:   -1,
	}
}

// Next update iterator index and return is any next value available
func (verrsi *ErrorsIterator) Next() bool {
	if verrsi.index+1 >= len(verrsi.records) {
		return false
	}
	verrsi.index++

	return true
}

// Value return current iterator value
func (verrsi *ErrorsIterator) Value() Error {
	return verrsi.records[verrsi.index]
}
