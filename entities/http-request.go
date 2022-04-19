package entities

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/kevin-vargas/logger/encoder"
	"go.uber.org/zap/zapcore"
)

var secretEcsKeysMap = map[string]string{
	"x-authorization":      secretPlaceholder,
	"authorization":        secretPlaceholder,
	"cookie":               secretPlaceholder,
	"cookies":              secretPlaceholder,
	"x-san-iatx-user-pass": secretPlaceholder,
}

const (
	secretPlaceholder = "secret"
)

type Headers http.Header

func sanitize(headers *Headers) []string {
	plainHeaders := []string{}
	headerValue := *headers
	for key := range headerValue {
		if value, found := secretEcsKeysMap[strings.ToLower(key)]; found {
			headerValue[key] = []string{value}
		}
		plainHeaders = append(plainHeaders, fmt.Sprintf("%s=%s", key, strings.Join(headerValue[key], ",")))
	}
	return plainHeaders
}

type HTTPRequestBody struct {
	Content []byte
	Headers Headers
}

func (body *HTTPRequestBody) MarshalLogObject(zapenc zapcore.ObjectEncoder) error {
	enc := encoder.Get(zapenc)
	enc.AddBytesValid(fieldHttpRequestBodyContent, body.Content)
	enc.AddStrings(fieldHttpRequestBodyHeaders, sanitize(&body.Headers))
	return nil
}

type HTTPRequest struct {
	Method   string
	Referrer string
	Body     *HTTPRequestBody
}

func (req *HTTPRequest) MarshalLogObject(zapenc zapcore.ObjectEncoder) error {
	enc := encoder.Get(zapenc)
	enc.AddString(fieldHttpRequestMethod, req.Method)
	enc.AddString(fieldHttpRequestReferrer, req.Referrer)
	enc.AddObjectValid(fieldHttpRequestBody, req.Body)
	return nil
}
