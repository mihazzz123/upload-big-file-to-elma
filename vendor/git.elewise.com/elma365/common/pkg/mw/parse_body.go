package mw

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"

	"git.elewise.com/elma365/common/pkg/errs"
	"git.elewise.com/elma365/common/pkg/errs/validation"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type parseBodyCtxKey struct{}

// ParseBody получает из тела запроса данные и валидирует их согласно тэгам валидации структуры bodyInstance
//
// В качестве валидатора используется gopkg.in/go-playground/validator.v9
func ParseBody(bodyInstance interface{}) func(next http.Handler) http.Handler {
	bodyType := reflect.TypeOf(bodyInstance)
	if bodyType.Kind() != reflect.Struct {
		panic("bodyInstance must be a struct (not pointer)")
	}

	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			l := ctxzap.Extract(ctx)

			newObjectInstance := reflect.New(bodyType)

			jd := json.NewDecoder(r.Body)
			if err := jd.Decode(newObjectInstance.Interface()); err != nil {
				l.Debug("unmarshal", zap.Error(err))
				http.Error(w, errors.Wrap(err, "unmarshal").Error(), http.StatusBadRequest)
				return
			}

			verrs := validation.ValidateStruct(ctx, newObjectInstance.Interface())
			if verrs != nil && len(*verrs) > 0 {
				w.Header().Set("Content-Type", "application/json")
				errs.Handle(l, w, verrs)
				return
			}

			ctx = contextWithParsedBody(ctx, newObjectInstance.Elem().Interface())
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}

// ParsedBodyFromContext получает контекст с разобранным и отвалидированным телом запроса
//
// Т.к. метод общий, то возвращает он пустой интерфейс, хотя реально в нем лежит ссылка на объект нужного типа.
// Пример использования: req := commonmw.ParsedBodyFromContext(ctx).(createMessageRequest)
func ParsedBodyFromContext(ctx context.Context) interface{} {
	return ctx.Value(parseBodyCtxKey{})
}

func contextWithParsedBody(ctx context.Context, value interface{}) context.Context {
	return context.WithValue(ctx, parseBodyCtxKey{}, value)
}
