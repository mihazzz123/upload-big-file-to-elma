package util

import (
	"encoding/base64"
	"net/http"
	"net/http/httputil"

	"go.uber.org/zap"
)

// DebugTransport http transport с логированием запроса и ответа
type DebugTransport struct {
	Transport http.RoundTripper
	Logger    *zap.Logger
}

// RoundTrip Реализует http.RoundTripper
func (t DebugTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.Transport == nil {
		t.Transport = http.DefaultTransport
	}

	t.dumpRequest(req)
	resp, err := t.Transport.RoundTrip(req)
	if err != nil {
		return resp, err
	}
	t.dumpResponse(resp)

	return resp, err
}

func (t DebugTransport) dumpResponse(resp *http.Response) {
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		t.logger().Error(err.Error())
	}
	t.logger().Debug(base64.StdEncoding.EncodeToString(dump))
}

func (t DebugTransport) dumpRequest(req *http.Request) {
	dump, err := httputil.DumpRequest(req, true)
	if err != nil {
		t.logger().Error(err.Error())
	}
	t.logger().Debug(base64.StdEncoding.EncodeToString(dump))
}

func (t DebugTransport) logger() *zap.Logger {
	if t.Logger == nil {
		return zap.L()
	}

	return t.Logger
}
