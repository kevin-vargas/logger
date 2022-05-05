package middleware

import (
	"bytes"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kevin-vargas/logger"
	"github.com/kevin-vargas/logger/config"
	"github.com/kevin-vargas/logger/loggertest"
	"github.com/kevin-vargas/logger/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Middleware(t *testing.T) {
	mockLogger := &mocks.Logger{}
	mockLogger.On("Info", mock.Anything).Return(nil)
	handler, _ := NewLoggingHandler(testConfig, withLogger(mockLogger))
	req, res := getReqRes(http.MethodPost)
	handler.Handle(testHandler).ServeHTTP(res, req)
	mockLogger.AssertNumberOfCalls(t, "Info", 2)
}

func Test_Middleware_Response(t *testing.T) {
	mockLogger := &mocks.Logger{}
	mockLogger.On("Info", mock.Anything).Return(nil)
	handler, _ := NewLoggingHandler(testConfig, withLogger(mockLogger), WithRequestLogging(emptyRequest))
	req, res := getReqRes(http.MethodPost)
	handler.Handle(testHandler).ServeHTTP(res, req)

	mockLogger.AssertNumberOfCalls(t, "Info", 1)
}

func Test_Middleware_Request(t *testing.T) {
	mockLogger := &mocks.Logger{}
	mockLogger.On("Info", mock.Anything).Return(nil)
	handler, _ := NewLoggingHandler(testConfig, withLogger(mockLogger), WithResponseLogging(emptyResponse))
	req, res := getReqRes(http.MethodPost)
	handler.Handle(testHandler).ServeHTTP(res, req)

	mockLogger.AssertNumberOfCalls(t, "Info", 1)
}

func Test_Middleware_Response_MSG(t *testing.T) {
	mockLogger := &mocks.Logger{}
	mockLogger.On("Info", mock.Anything).Return(nil)
	handler, _ := NewLoggingHandler(testConfig, withLogger(mockLogger), WithRequestLogging(emptyRequest))
	req, res := getReqRes(http.MethodPost)
	handler.Handle(testHandler).ServeHTTP(res, req)
	assert.NotNil(t, mockLogger.Msg)
	assert.Equal(t, mockLogger.Msg.Text, "Response")
}

func Test_Middleware_Request_MSG(t *testing.T) {
	mockLogger := &mocks.Logger{}
	mockLogger.On("Info", mock.Anything).Return(nil)
	handler, _ := NewLoggingHandler(testConfig, withLogger(mockLogger), WithResponseLogging(emptyResponse))
	req, res := getReqRes(http.MethodPost)
	handler.Handle(testHandler).ServeHTTP(res, req)
	assert.NotNil(t, mockLogger.Msg)
	assert.Equal(t, mockLogger.Msg.Text, "Request")
}

func Test_Middleware_Request_Trace_Id(t *testing.T) {
	buf := &bytes.Buffer{}
	log, err := logger.New(logger.WithIoWriter(buf))
	if err != nil {
		t.FailNow()
	}
	handler, _ := NewLoggingHandler(testConfig, withLogger(log), WithResponseLogging(emptyResponse))
	req, res := getReqRes(http.MethodPost)
	handler.Handle(testHandler).ServeHTTP(res, req)
	result := loggertest.Sanitize(buf.Bytes())
	assert.JSONEq(t, expected_trace_id, result)
}

func Test_Middleware_Request_Error(t *testing.T) {
	mockLogger := &mocks.Logger{}
	mockLogger.On("Error", mock.Anything).Return(nil)
	handler, _ := NewLoggingHandler(testConfig, withLogger(mockLogger), WithResponseLogging(emptyResponse))
	req, res := getReqResError(http.MethodPost)
	handler.Handle(testHandler).ServeHTTP(res, req)

	mockLogger.AssertNumberOfCalls(t, "Error", 1)
}

const (
	expected_trace_id = `{"@timestamp":0,"message":"Request","labels":{},"log":{"logger":"santander-logger_(uber-go/zap)","level":"info"},"httprequest":{"method":"POST","referrer":"","body":{"content":"sample","headers":["X-San-Correlationid=trace_id"]}},"trace":{"id":"trace_id"}}`
)

var testConfig = &config.Logger{}

var getReqRes = func(method string) (*http.Request, *httptest.ResponseRecorder) {
	reader := bytes.NewReader([]byte(`sample`))
	req := httptest.NewRequest(method, "https://www.some-domain.com", reader)
	req.Header.Add(correlation_id_header, "trace_id")
	res := httptest.NewRecorder()
	return req, res
}

type errReader int

func (errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test")
}

var getReqResError = func(method string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, "https://www.some-domain.com", errReader(0))
	res := httptest.NewRecorder()
	return req, res
}

var emptyResponse ResponseLogger = func(l logger.Logger, r *ResponseObserver) {
}

var emptyRequest RequestLogger = func(l logger.Logger, r *http.Request) error {
	return nil
}

var testHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("bar"))
	if err != nil {
		log.Fatal(err.Error())
	}
})

var withLogger = func(l logger.Logger) Option {
	return WithLoggerBuilder(func(c *config.Logger) logger.Logger {
		return l
	})
}
