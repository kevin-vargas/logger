package logger

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	"github.com/kevin-vargas/logger/entitys"
	"github.com/stretchr/testify/assert"
	"gitlab.ar.bsch/stlibs/golang/logger"
	"go.uber.org/zap/zapcore"
)

var loggTwo, _ = NewLogger()
var res = entitys.HTTPResponse{
	StatusCode: 201,
	Body: &entitys.HTTPResponseBody{
		Content: []byte("testing"),
	},
}
var req = entitys.HTTPRequest{
	Method:   "GET",
	Referrer: "/ping",
	Body: &entitys.HTTPRequestBody{
		Content: []byte("testing"),
		Headers: entitys.Headers{
			"Host":          []string{"www.host.com"},
			"Content-Type":  []string{"application/json"},
			"Authorization": []string{"Bearer Token"},
		},
	},
}
var erre = entitys.Error{
	Message:    "_message",
	Type:       "_type",
	StackTrace: "_stack_trace",
}
var trace = entitys.Trace{
	ID: "id_trace",
}
var event = entitys.Event{
	Action:   "action",
	Category: []string{"category1", "category2"},
	Module:   "module",
	Type:     "type",
	Original: "original",
}
var tags = entitys.Tags{"tag1", "tag2", "tag3"}

var msg = entitys.NewMessage("msg").
	WithHttpReponse(res).
	WithHttpRequest(req).
	WithError(erre).
	WithEvent(event).
	WithTrace(trace).
	WithTags(tags)

func Test_WithoutLogger(t *testing.T) {
	expected := []byte(`{"@timestamp":0,"message":"msg","labels":{"service":"service","environment":"enviroment","lib_version":"0.0.1","lib_language":"golang","pod_name":"pod_name","node_name":"node_name","application":"application"},"log":{"logger":"santander-logger_(uber-go/zap)","level":"info"},"tags":["tag1","tag2","tag3"],"httprequest":{"method":"GET","referrer":"/ping","body":{"content":"testing","headers":["Host=www.host.com","Content-Type=application/json","Authorization=secret"]}},"httpresponse":{"status_code":201,"body":{"content":"testing"}},"event":{"action":"action","module":"module","type":"type","original":"original","category":["category1","category2"]},"trace":{"id":"id_trace"},"error":{"message":"_message","stack_trace":"_stack_trace","type":"_type"}}`)
	buf := &bytes.Buffer{}
	logger, err := NewLogger(WithIoWriter(buf))
	if err != nil {
		t.Error(err)
	}

	logger.Info(*msg)
	result := sanitizeTimestamp(buf.Bytes())
	assert.Equal(t, string(expected), string(result))
}

func BenchmarkLogOne(b *testing.B) {
	loggerOne, _ := logger.NewDefaultLogger()
	for i := 0; i <= b.N; i++ {
		loggerOne.Info("test")
	}
}
func BenchmarkLogTwo(b *testing.B) {
	loggerTwo, _ := NewLogger()
	var msg = entitys.NewMessage("test")
	for i := 0; i <= b.N; i++ {
		loggerTwo.Info(*msg)
	}
}
func BenchmarkLogOneComplete(b *testing.B) {
	loggerOne, _ := logger.NewDefaultLogger()
	fields := []zapcore.Field{

		logger.String(FieldServiceName, "logger.golden"),

		// error
		logger.String(FieldErrorMessage, "_message"),
		logger.String(FieldStackTrace, "_stack_trace"),
		logger.String(FieldErrorType, "_type"),
		// event
		logger.String(FieldEventAction, "action"),
		logger.String(FieldEventCategory, "category"),
		logger.String(FieldEventModule, "module"),
		logger.String(FieldEventType, "type"),
		logger.String(FieldEventOriginal, "original"),
		// trace
		logger.String(FieldTraceID, "id_trace"),
		// request
		logger.String(FieldHTTPRequestBodyContent, "testing"),
		logger.String(FieldHTTPRequestMethod, "GET"),
		logger.Any(FieldHTTPRequestBodyHeaders, []http.Header{{ // tiene que ser http.header no slice
			"Host":          []string{"www.host.com"},
			"Content-Type":  []string{"application/json"},
			"Authorization": []string{"Bearer Token"},
		}}),
		logger.String(FieldHTTPRequestReferrer, "/ping"),
		// response
		logger.String(FieldHTTPResponseBodyContent, "testing"),
		logger.String(FieldHTTPResponseStatusCode, fmt.Sprintf("%v", 201)), // tendria que ser int
		logger.Tags([]string{"tag1", "tag2", "tag3"}),
	}
	for i := 0; i <= b.N; i++ {
		loggerOne.Info("test", fields...)
	}
}
func BenchmarkLogTwoComplete(b *testing.B) {
	var res = entitys.HTTPResponse{
		StatusCode: 201,
		Body: &entitys.HTTPResponseBody{
			Content: []byte("testing"),
		},
	}
	var req = entitys.HTTPRequest{
		Method:   "GET",
		Referrer: "/ping",
		Body: &entitys.HTTPRequestBody{
			Content: []byte("testing"),
			Headers: entitys.Headers{
				"Host":          []string{"www.host.com"},
				"Content-Type":  []string{"application/json"},
				"Authorization": []string{"Bearer Token"},
			},
		},
	}
	var erre = entitys.Error{
		Message:    "_message",
		Type:       "_type",
		StackTrace: "_stack_trace",
	}
	var trace = entitys.Trace{
		ID: "id_trace",
	}
	var event = entitys.Event{
		Action:   "action",
		Category: []string{"category1", "category2"},
		Module:   "module",
		Type:     "type",
		Original: "original",
	}
	var tags = entitys.Tags{"tag1", "tag2", "tag3"}

	var msg = entitys.NewMessage("msg").
		WithHttpReponse(res).
		WithHttpRequest(req).
		WithError(erre).
		WithEvent(event).
		WithTrace(trace).
		WithTags(tags)

	loggerTwo, _ := NewLogger()
	for i := 0; i <= b.N; i++ {
		loggerTwo.Info(*msg)
	}
}
