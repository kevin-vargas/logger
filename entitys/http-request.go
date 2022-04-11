package entitys

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	fieldHttpRequestMethod      = "method"
	fieldHttpRequestReferrer    = "referrer"
	fieldHttpRequestBody        = "body"
	fieldHttpRequestBodyContent = "content"
	fieldHttpRequestBodyHeaders = "headers"
)

type HTTPRequestBody struct {
	Content []byte
	Headers []string
}

func (body *HTTPRequestBody) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddByteString(fieldHttpRequestBodyContent, body.Content)
	zap.Strings(fieldHttpRequestBodyHeaders, body.Headers).AddTo(enc)
	return nil
}

type HTTPRequest struct {
	Method   string
	Referrer string
	Body     HTTPRequestBody
}

func (req *HTTPRequest) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString(fieldHttpRequestMethod, req.Method)
	enc.AddString(fieldHttpRequestReferrer, req.Referrer)
	enc.AddObject(fieldHttpRequestBody, &req.Body)
	return nil
}
