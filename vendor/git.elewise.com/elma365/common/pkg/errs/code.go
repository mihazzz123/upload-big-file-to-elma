package errs

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// List of global error codes
const (
	NotFound    errorCode = "not found"
	Unavailable errorCode = "unavailable"
	// Collision для случая, когда сущность, которую пытается поменять данный запрос,
	// уже как-то изменена (существует по уникальному ключу, дублирует что-то и тд)
	Collision       errorCode = "collision"
	Forbidden       errorCode = "forbidden"
	Unauthorized    errorCode = "unauthorized"
	Internal        errorCode = "internal error"
	InvalidArgument errorCode = "invalid argument"
	NotImplemented  errorCode = "not implemented"
	Timeout         errorCode = "timeout exceeded"
	// Precondition для случая, когда нет нужного состояния для совершения действия (что-то выключено,
	// какие-то настройки не применились, нет необходимых ресурсов) - то, что не зависит от конкретного
	// запроса пользователя, а должно было быть сделано раньше
	Precondition errorCode = "precondition"
	Unknown      errorCode = "unknown"
	// ResourceExhausted ресурс исчерпан - пользовательская квота или например место на диске
	ResourceExhausted errorCode = "resource exhausted"
)

type errorCode string
type protocols struct {
	grpc     codes.Code
	http     int
	logLevel zapcore.Level
}

//nolint: gochecknoglobals // это должна быть константа, но Go так не умеет
var codeToProtocol = map[errorCode]protocols{
	NotFound: {
		codes.NotFound,
		http.StatusNotFound,
		zapcore.DebugLevel,
	},
	Unavailable: {
		codes.Unavailable,
		http.StatusServiceUnavailable,
		zapcore.ErrorLevel,
	},
	Forbidden: {
		codes.PermissionDenied,
		http.StatusForbidden,
		zapcore.DebugLevel,
	},
	Unauthorized: {
		codes.Unauthenticated,
		http.StatusUnauthorized,
		zapcore.DebugLevel,
	},
	Internal: {
		codes.Internal,
		http.StatusInternalServerError,
		zapcore.ErrorLevel,
	},
	InvalidArgument: {
		codes.InvalidArgument,
		http.StatusBadRequest,
		zapcore.DebugLevel,
	},
	NotImplemented: {
		codes.Unimplemented,
		http.StatusNotImplemented,
		zapcore.ErrorLevel,
	},
	Collision: {
		codes.AlreadyExists,
		http.StatusConflict,
		zapcore.DebugLevel,
	},
	Timeout: {
		codes.DeadlineExceeded,
		http.StatusGatewayTimeout,
		zapcore.ErrorLevel,
	},
	Precondition: {
		codes.FailedPrecondition,
		http.StatusPreconditionFailed,
		zapcore.WarnLevel,
	},
	Unknown: {
		codes.Unknown,
		http.StatusBadGateway,
		zapcore.ErrorLevel,
	},
	ResourceExhausted: {
		codes.ResourceExhausted,
		http.StatusPaymentRequired,
		zapcore.WarnLevel,
	},
}

//nolint: gochecknoglobals // это должны быть константы, но Go так не умеет
var (
	grpcToCode = make(map[codes.Code]errorCode, len(codeToProtocol))
	httpToCode = make(map[int]errorCode, len(codeToProtocol))
)

//nolint: gochecknoinits // это должны быть константы, но Go так не умеет
func init() {
	for eCode, proto := range codeToProtocol {
		grpcToCode[proto.grpc] = eCode
		httpToCode[proto.http] = eCode
	}
}

// FromCode принимает строковое представление кода ошибки и пытается сопоставить ей внутренний код ошибки. Если код
// пустой - возвращает nil. Если не совпадает ни с одним внутренним кодом - возвращает Unknown.
func FromCode(code string) error {
	if code == "" {
		return nil
	}

	errCode := errorCode(code)
	if _, ok := codeToProtocol[errCode]; ok {
		return errCode
	}

	return Unknown
}

// FromGRPCError принимает grpc ошибку и пытается сопоставить ей внутренний код ошибки
//
// Если подходящей внутренней ошибки не найдено, то вернёт ошибку как есть.
func FromGRPCError(err error) error {
	if st, ok := status.FromError(err); ok {
		if eCode, ok := grpcToCode[st.Code()]; ok {
			return eCode.New(st.Message())
		}
	}
	return err
}

// FromHTTPResponse принимает http ответ и пытается сопоставить ей внутренний код ошибки
//
// Если подходящей внутренней ошибки не найдено, то вернёт ошибку Unknown.
func FromHTTPResponse(r *http.Response) error {
	if r.StatusCode >= http.StatusOK && r.StatusCode < http.StatusMultipleChoices {
		return nil
	}
	var text string
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		text = err.Error()
	} else {
		text = string(body)
	}

	return FromHTTPStatus(r.StatusCode, text)
}

// FromHTTPStatus принимает http код и описание ошибки и пытается сопоставить ей внутренний код ошибки.
func FromHTTPStatus(statusCode int, reason string) error {
	if eCode, ok := httpToCode[statusCode]; ok {
		return eCode.New(reason)
	}

	return Unknown.New(reason)
}

// Code получить код из ошибки
func Code(err error) error {
	code, ok := errors.Cause(err).(errorCode)
	if !ok {
		code = Unknown
	}
	return code
}

// ErrorLevel возвращает уровень лога ошибки
func ErrorLevel(err error) zapcore.Level {
	code, ok := errors.Cause(err).(errorCode)
	if !ok {
		code = Unknown
	}

	return codeToProtocol[code].logLevel
}

func (ec errorCode) Error() string {
	return string(ec)
}

func (ec errorCode) New(s string) error {
	return errors.Wrap(ec, s)
}

func (ec errorCode) Newf(format string, args ...interface{}) error {
	return errors.Wrapf(ec, format, args...)
}

func (ec errorCode) Wrap(err error) error {
	if err == nil {
		return nil
	}
	err = errors.WithStack(err)

	return errorWithCode{orig: err, code: ec}
}

func (ec errorCode) GRPCCode() codes.Code {
	if code, ok := codeToProtocol[ec]; ok {
		return code.grpc
	}

	return codes.Unknown
}

func (ec errorCode) HTTPStatus() int {
	if code, ok := codeToProtocol[ec]; ok {
		return code.http
	}

	return http.StatusInternalServerError
}

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

func (ewc errorWithCode) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%s: %+v", ewc.code, ewc.orig)
			return
		}
		fallthrough
	case 's':
		_, _ = io.WriteString(s, ewc.Error())
	case 'q':
		fmt.Fprintf(s, "%q", ewc.Error())
	}
}

type withGRPCCode interface {
	GRPCCode() codes.Code
}

// GRPCCodeFromError сопоставляет код возврата GRPC ошибке
//
// Если ошибка сформирована самим gRPC, то возвращает её код, если ошибка образована от константной ошибки,
// то сопоставляет код по таблице. Если извлечь код не удаётся двигается по цепочке Cause. Если не удаётся
// извлечь код по всей цепочке, то возвращает codes.Unknown.
func GRPCCodeFromError(err error) codes.Code {
	type withCause interface {
		Cause() error
	}
	for {
		if s, ok := status.FromError(err); ok {
			return s.Code()
		}
		if ewc, ok := err.(withGRPCCode); ok {
			return ewc.GRPCCode()
		}
		if ewc, ok := err.(withCause); ok {
			err = ewc.Cause()
			continue
		}

		return codes.Unknown
	}
}

// CodeToHTTPStatus пробует извлечь из ошибки константную основу и сопоставить ей HTTP статус
//
// Если статус не найден, то возвращает 500.
func CodeToHTTPStatus(err error) int {
	if err == nil {
		return http.StatusOK
	}
	cause := errors.Cause(err)
	if ec, ok := cause.(interface{ HTTPStatus() int }); ok {
		return ec.HTTPStatus()
	}

	return http.StatusInternalServerError
}
