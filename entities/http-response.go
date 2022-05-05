package entities

import (
	"github.com/kevin-vargas/logger/encoder"
	"go.uber.org/zap/zapcore"
)

type HTTPResponseBody struct {
	Content []byte
}

func (body *HTTPResponseBody) MarshalLogObject(zapenc zapcore.ObjectEncoder) error {
	enc := encoder.Get(zapenc)
	enc.AddBytesValid(fieldHttpResponseBodyContent, body.Content)
	return nil
}

type HTTPResponse struct {
	StatusCode int64
	Body       *HTTPResponseBody
}

func (res *HTTPResponse) MarshalLogObject(zapenc zapcore.ObjectEncoder) error {
	enc := encoder.Get(zapenc)
	enc.AddInt64(fieldHttpResponseStatusCode, res.StatusCode)
	return enc.AddObjectValid(fieldHttpResponseBody, res.Body)
}
