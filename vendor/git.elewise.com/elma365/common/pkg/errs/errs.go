package errs

import (
	"net/http"

	"go.uber.org/zap"
)

// HandleInternal is a shortcut to log error and return 500 to client
func HandleInternal(l *zap.Logger, w http.ResponseWriter, err error) {
	l.Error(err.Error(), zap.Error(err))
	http.Error(w, Internal.Error(), CodeToHTTPStatus(Internal))
}

// Handle error according to it's type
func Handle(l *zap.Logger, w http.ResponseWriter, err error) {
	code := CodeToHTTPStatus(err)

	if code >= 500 {
		l.Error(err.Error(), zap.Error(err))
		http.Error(w, Internal.Error(), code)
	} else {
		l.Warn(err.Error(), zap.Error(err))
		w.WriteHeader(code)
		_ = WriteDataFromError(err, w)
	}
}
