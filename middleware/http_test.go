package middleware

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kevin-vargas/logger"
	"github.com/kevin-vargas/logger/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var getReqRes = func(method string) (*http.Request, *httptest.ResponseRecorder) {
	reader := bytes.NewReader([]byte(`sample`))
	req := httptest.NewRequest(method, "https://www.some-domain.com", reader)
	req.Header.Add("header-key", "header-value")
	res := httptest.NewRecorder()
	return req, res
}

var emptyResponse ResponseLogger = func(l logger.Logger, r *ResponseObserver) {
}

var emptyRequest RequestLogger = func(l logger.Logger, r *http.Request) error {
	return nil
}

var testHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bar"))
})

func Test_Middleware(t *testing.T) {
	mockLogger := &mocks.Logger{}
	mockLogger.On("Info", mock.Anything).Return(nil)
	handler := NewLoggingHandler(WithLogger(mockLogger))
	req, res := getReqRes(http.MethodPost)
	handler.Handle(testHandler).ServeHTTP(res, req)
	mockLogger.AssertNumberOfCalls(t, "Info", 2)
}

func Test_Middleware_Response(t *testing.T) {
	mockLogger := &mocks.Logger{}
	mockLogger.On("Info", mock.Anything).Return(nil)
	handler := NewLoggingHandler(WithLogger(mockLogger), WithRequestLogging(emptyRequest))
	req, res := getReqRes(http.MethodPost)
	handler.Handle(testHandler).ServeHTTP(res, req)

	mockLogger.AssertNumberOfCalls(t, "Info", 1)
}

func Test_Middleware_Request(t *testing.T) {
	mockLogger := &mocks.Logger{}
	mockLogger.On("Info", mock.Anything).Return(nil)
	handler := NewLoggingHandler(WithLogger(mockLogger), WithResponseLogging(emptyResponse))
	req, res := getReqRes(http.MethodPost)
	handler.Handle(testHandler).ServeHTTP(res, req)

	mockLogger.AssertNumberOfCalls(t, "Info", 1)
}

func Test_Middleware_Response_MSG(t *testing.T) {
	mockLogger := &mocks.Logger{}
	mockLogger.On("Info", mock.Anything).Return(nil)
	handler := NewLoggingHandler(WithLogger(mockLogger), WithRequestLogging(emptyRequest))
	req, res := getReqRes(http.MethodPost)
	handler.Handle(testHandler).ServeHTTP(res, req)
	assert.NotNil(t, mockLogger.Msg)
	assert.Equal(t, mockLogger.Msg.Text, "Response")
}

func Test_Middleware_Request_MSG(t *testing.T) {
	mockLogger := &mocks.Logger{}
	mockLogger.On("Info", mock.Anything).Return(nil)
	handler := NewLoggingHandler(WithLogger(mockLogger), WithResponseLogging(emptyResponse))
	req, res := getReqRes(http.MethodPost)
	handler.Handle(testHandler).ServeHTTP(res, req)
	assert.NotNil(t, mockLogger.Msg)
	assert.Equal(t, mockLogger.Msg.Text, "Request")
}
