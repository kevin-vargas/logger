package pubsub

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kevin-vargas/logger/pubsub/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Post(t *testing.T) {
	ts := createServer(t, http.StatusOK)
	defer ts.Close()
	cfg := &ConfigRestPublisher{
		URL: ts.URL,
	}
	pubsub := NewRestPublisher(cfg)
	err := pubsub.Publish("test", 1)
	assert.Nil(t, err)
}

func Test_Post_Error(t *testing.T) {
	ts := createServer(t, http.StatusInternalServerError)
	defer ts.Close()
	cfg := &ConfigRestPublisher{
		URL: ts.URL,
	}
	pubsub := NewRestPublisher(cfg)
	err := pubsub.Publish("test", 1)
	assert.NotNil(t, err)
}

func Test_Post_Error_FallBack(t *testing.T) {
	ts := createServer(t, http.StatusInternalServerError)
	defer ts.Close()
	cfg := &ConfigRestPublisher{
		URL: ts.URL,
	}
	mockFallBack := mocks.Fallback{}
	mockFallBack.On("Method", "test", mock.Anything).Return(nil)
	pubsub := NewRestPublisher(cfg)
	err := pubsub.Publish("test", 1, mockFallBack.Method)
	assert.NotNil(t, err)
	mockFallBack.AssertNumberOfCalls(t, "Method", 1)
}

func createServer(t *testing.T, statusCode int) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			if r.URL.Path == "/publish" {
				w.WriteHeader(statusCode)
				_, _ = w.Write([]byte(`test`))
			}
			return
		}
	}))

	return ts
}
