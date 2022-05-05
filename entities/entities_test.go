package entities

import (
	"encoding/json"
	"testing"

	"github.com/kevin-vargas/logger/config"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func Test_Default_Labels(t *testing.T) {
	cfg := &config.Logger{}
	labels := GetDefaultLabels(cfg)
	cases := []struct {
		desc   string
		key    string
		expect string
	}{
		{
			"Version",
			fieldLabelLibVersion,
			LIB_VERSION,
		},
		{
			"Language",
			fieldLabelLibLanguage,
			LIB_LANGUAGE,
		},
	}

	for _, tt := range cases {
		t.Run(tt.desc, func(t *testing.T) {
			assert.Equal(t, tt.expect, labels[tt.key])
		})
	}
}

func Test_Default_log(t *testing.T) {
	cases := []struct {
		desc   string
		level  Level
		expect Log
	}{
		{
			"Debug",
			DebugLevel,
			Log{
				LOG_LOGGER,
				"debug",
			},
		},
		{
			"Info",
			InfoLevel,
			Log{
				LOG_LOGGER,
				"info",
			},
		},
		{
			"Warn",
			WarnLevel,
			Log{
				LOG_LOGGER,
				"warn",
			},
		},
	}
	for _, tt := range cases {
		t.Run(tt.desc, func(t *testing.T) {
			result := GetDefaultLog(tt.level)
			assert.Equal(t, tt.expect, result)
		})
	}
}

// entities
func Test_Entity_Error(t *testing.T) {

	expect := `{"message":"Message_test","stack_trace":"StackTrace_test","type":"Type_test"}`

	result := getJsonFromLog(t, &err)

	assert.JSONEq(t, expect, result)
}

func Test_Entity_Event(t *testing.T) {

	expect := `{"action":"action","category":["category1","category2"],"module":"module","original":"original","type":"type"}`

	result := getJsonFromLog(t, &event)

	assert.JSONEq(t, expect, result)
}

func Test_Entity_HTTPRequest(t *testing.T) {

	expect := `{"body":{"content":"testing","headers":["Authorization=secret"]},"method":"GET","referrer":"/ping"}`

	result := getJsonFromLog(t, &httpRequest)

	assert.JSONEq(t, expect, result)
}

func Test_Entity_HTTPResponse(t *testing.T) {

	expect := `{"body":{"content":"testing"},"status_code":201}`

	result := getJsonFromLog(t, &httpResponse)

	assert.JSONEq(t, expect, result)
}

func Test_Entity_Labels(t *testing.T) {

	expect := `{"foo":"bar"}`

	result := getJsonFromLog(t, &labels)

	assert.JSONEq(t, expect, result)
}

func Test_Entity_Trace(t *testing.T) {

	expect := `{"id":"id_trace"}`

	result := getJsonFromLog(t, &trace)

	assert.JSONEq(t, expect, result)
}

func Test_Entity_Log(t *testing.T) {

	expect := `{"level":"level_test","logger":"logger_test"}`

	result := getJsonFromLog(t, &log)

	assert.JSONEq(t, expect, result)
}

func Test_Entity_Message(t *testing.T) {

	expect := `{"error":{"message":"Message_test","stack_trace":"StackTrace_test","type":"Type_test"},"event":{"action":"action","category":["category1","category2"],"module":"module","original":"original","type":"type"},"httprequest":{"body":{"content":"testing","headers":["Authorization=secret"]},"method":"GET","referrer":"/ping"},"httpresponse":{"body":{"content":"testing"},"status_code":201},"labels":{"foo":"bar"},"log":{"level":"level_test","logger":"logger_test"},"tags":["tag1","tag2"],"trace":{"id":"id_trace"}}`
	msg := NewMessage("test").
		WithLabels(labels).
		WithLoggerInfo(log).
		WithTags(tags).
		WithHttpRequest(httpRequest).
		WithHttpReponse(httpResponse).
		WithEvent(event).
		WithTrace(trace).
		WithError(err)

	result := getJsonFromLog(t, msg)

	assert.JSONEq(t, expect, result)
}

func Test_Entity_Message_Default(t *testing.T) {

	expect := `{"error":{"message":"Message_test","stack_trace":"StackTrace_test","type":"Type_test"},"event":{"action":"action","category":["category1","category2"],"module":"module","original":"original","type":"type"},"httprequest":{"body":{"content":"testing","headers":["Authorization=secret"]},"method":"GET","referrer":"/ping"},"httpresponse":{"body":{"content":"testing"},"status_code":201},"labels":{"foo":"bar"},"log":{"level":"info","logger":"santander-logger_(uber-go/zap)"},"tags":["tag1","tag2"],"trace":{"id":"id_trace"}}`
	field := NewMessage("test").
		WithLabels(labels).
		WithLoggerInfo(log).
		WithTags(tags).
		WithHttpRequest(httpRequest).
		WithHttpReponse(httpResponse).
		WithEvent(event).
		WithTrace(trace).
		WithError(err).
		Encode(encondeConfig)

	result := getJsonFromField(t, field)

	assert.JSONEq(t, expect, result)
}

func getJsonFromLog(t *testing.T, val zapcore.ObjectMarshaler) string {
	return getJsonFromField(t, zap.Inline(val))
}

func getJsonFromField(t *testing.T, field zapcore.Field) string {
	observedZapCore, observedLogs := observer.New(zap.InfoLevel)
	observedLogger := zap.New(observedZapCore)
	observedLogger.Info("test", field)
	allLogs := observedLogs.All()
	result, errJson := json.Marshal(allLogs[0].ContextMap())
	if errJson != nil {
		t.FailNow()
	}
	return string(result)
}

var log = Log{
	Logger: "logger_test",
	Level:  "level_test",
}
var trace = Trace{
	ID: "id_trace",
}
var labels = Labels{
	"foo": "bar",
}
var httpResponse = HTTPResponse{
	StatusCode: 201,
	Body: &HTTPResponseBody{
		Content: []byte("testing"),
	},
}
var httpRequest = HTTPRequest{
	Method:   "GET",
	Referrer: "/ping",
	Body: &HTTPRequestBody{
		Content: []byte("testing"),
		Headers: Headers{
			"Authorization": []string{"Bearer Token"},
		},
	},
}
var event = Event{
	Action:   "action",
	Category: []string{"category1", "category2"},
	Module:   "module",
	Type:     "type",
	Original: "original",
}
var tags = Tags{"tag1", "tag2"}

var err = Error{
	Message:    "Message_test",
	StackTrace: "StackTrace_test",
	Type:       "Type_test",
}

var encondeConfig = EncodeConfig{
	LVL: InfoLevel,
}
