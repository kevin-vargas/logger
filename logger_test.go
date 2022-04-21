package logger

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/kevin-vargas/logger/entities"
	"github.com/stretchr/testify/assert"
)

const (
	expect_simple   = `{"@timestamp":0,"message":"test","labels":{"lib_version":"0.0.1","lib_language":"golang","pod_name":"pod_name","node_name":"node_name","application":"application","service":"service","environment":"enviroment"},"log":{"logger":"santander-logger_(uber-go/zap)","level":"%s"}}`
	expect_tags     = `{"@timestamp":0,"message":"test","labels":{"service":"service","environment":"enviroment","lib_version":"0.0.1","lib_language":"golang","pod_name":"pod_name","node_name":"node_name","application":"application"},"log":{"logger":"santander-logger_(uber-go/zap)","level":"%s"},"tags":["tag1","tag2","tag3"]}`
	expect_event    = `{"@timestamp":0,"message":"test","labels":{"lib_version":"0.0.1","lib_language":"golang","pod_name":"pod_name","node_name":"node_name","application":"application","service":"service","environment":"enviroment"},"log":{"logger":"santander-logger_(uber-go/zap)","level":"%s"},"event":{"action":"action","module":"module","type":"type","original":"original","category":["category1","category2"]}}`
	expect_trace    = `{"@timestamp":0,"message":"test","labels":{"application":"application","service":"service","environment":"enviroment","lib_version":"0.0.1","lib_language":"golang","pod_name":"pod_name","node_name":"node_name"},"log":{"logger":"santander-logger_(uber-go/zap)","level":"%s"},"trace":{"id":"id_trace"}}`
	expect_error    = `{"@timestamp":0,"message":"test","labels":{"application":"application","service":"service","environment":"enviroment","lib_version":"0.0.1","lib_language":"golang","pod_name":"pod_name","node_name":"node_name"},"log":{"logger":"santander-logger_(uber-go/zap)","level":"%s"},"error":{"message":"_message","stack_trace":"_stack_trace","type":"_type"}}`
	expect_request  = `{"@timestamp":0,"message":"test","labels":{"node_name":"node_name","application":"application","service":"service","environment":"enviroment","lib_version":"0.0.1","lib_language":"golang","pod_name":"pod_name"},"log":{"logger":"santander-logger_(uber-go/zap)","level":"%s"},"httprequest":{"method":"GET","referrer":"/ping","body":{"content":"testing","headers":["Authorization=secret"]}}}`
	expect_response = `{"@timestamp":0,"message":"test","labels":{"environment":"enviroment","lib_version":"0.0.1","lib_language":"golang","pod_name":"pod_name","node_name":"node_name","application":"application","service":"service"},"log":{"logger":"santander-logger_(uber-go/zap)","level":"%s"},"httpresponse":{"status_code":201,"body":{"content":"testing"}}}`
	expect_labels   = `{"@timestamp":0,"message":"test","labels":{"service":"service","environment":"enviroment","lib_version":"0.0.1","lib_language":"golang","pod_name":"pod_name","node_name":"node_name","foo":"bar","application":"application"},"log":{"logger":"santander-logger_(uber-go/zap)","level":"%s"}}`
	expect_complete = `{"@timestamp":0,"message":"test","labels":{"application":"application","service":"service","environment":"enviroment","lib_version":"0.0.1","lib_language":"golang","pod_name":"pod_name","node_name":"node_name","foo":"bar"},"log":{"logger":"santander-logger_(uber-go/zap)","level":"%s"},"tags":["tag1","tag2","tag3"],"httprequest":{"method":"GET","referrer":"/ping","body":{"content":"testing","headers":["Authorization=secret"]}},"httpresponse":{"status_code":201,"body":{"content":"testing"}},"event":{"action":"action","module":"module","type":"type","original":"original","category":["category1","category2"]},"trace":{"id":"id_trace"},"error":{"message":"_message","stack_trace":"_stack_trace","type":"_type"}}`
)

const (
	default_msg = "test"
)

var getExpected = func(level string, base string) string {
	return fmt.Sprintf(base, level)
}
var labels = entities.Labels{
	"foo": "bar",
}
var res = entities.HTTPResponse{
	StatusCode: 201,
	Body: &entities.HTTPResponseBody{
		Content: []byte("testing"),
	},
}
var req = entities.HTTPRequest{
	Method:   "GET",
	Referrer: "/ping",
	Body: &entities.HTTPRequestBody{
		Content: []byte("testing"),
		Headers: entities.Headers{
			"Authorization": []string{"Bearer Token"},
		},
	},
}
var erre = entities.Error{
	Message:    "_message",
	Type:       "_type",
	StackTrace: "_stack_trace",
}
var trace = entities.Trace{
	ID: "id_trace",
}
var event = entities.Event{
	Action:   "action",
	Category: []string{"category1", "category2"},
	Module:   "module",
	Type:     "type",
	Original: "original",
}
var tags = entities.Tags{"tag1", "tag2", "tag3"}

var msg = entities.NewMessage("msg").
	WithHttpReponse(res).
	WithHttpRequest(req).
	WithError(erre).
	WithEvent(event).
	WithTrace(trace).
	WithTags(tags)

type MethodLogger func(*entities.Message)

func Test_Logger(t *testing.T) {

	buf := &bytes.Buffer{}
	logger, err := NewLogger(WithIoWriter(buf))
	if err != nil {
		t.Error(err)
	}
	reset := func() {
		buf.Reset()
	}
	methods := []struct {
		MethodLogger
		level string
	}{
		{
			logger.Info,
			"info",
		},
		{
			logger.Error,
			"error",
		},
		{
			logger.Warn,
			"warn",
		},
		{
			logger.Debug,
			"debug",
		},
	}

	testCases := []struct {
		desc         string
		expectedBase string
		msg          *entities.Message
	}{
		{
			"Simple",
			expect_simple,
			entities.NewMessage(default_msg),
		},
		{
			"With Tags",
			expect_tags,
			entities.NewMessage(default_msg).WithTags(tags),
		},
		{
			"With Event",
			expect_event,
			entities.NewMessage(default_msg).WithEvent(event),
		},
		{
			"With Trace",
			expect_trace,
			entities.NewMessage(default_msg).WithTrace(trace),
		},
		{
			"With Error",
			expect_error,
			entities.NewMessage(default_msg).WithError(erre),
		},
		{
			"With Request",
			expect_request,
			entities.NewMessage(default_msg).WithHttpRequest(req),
		},
		{
			"With Response",
			expect_response,
			entities.NewMessage(default_msg).WithHttpReponse(res),
		},
		{
			"With Labels",
			expect_labels,
			entities.NewMessage(default_msg).WithLabels(labels),
		},
		{
			"Complete",
			expect_complete,
			entities.
				NewMessage(default_msg).
				WithError(erre).
				WithEvent(event).
				WithHttpReponse(res).
				WithHttpRequest(req).
				WithLabels(labels).
				WithTags(tags).
				WithTrace(trace),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			for _, methodt := range methods {
				t.Run(
					methodt.level,
					func(t *testing.T) {

						t.Cleanup(reset)
						expected := getExpected(methodt.level, tt.expectedBase)

						// act
						methodt.MethodLogger(tt.msg)
						result := sanitize(buf.Bytes())

						// asert
						assert.JSONEq(t, expected, result)
					})
			}
		})
	}
}
