package pubsub

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Post(t *testing.T) {
	ts := CreateServer(t, http.StatusOK)
	defer ts.Close()
	cfg := &ConfigRestPublisher{
		URL: ts.URL,
	}
	pubsub := NewRestPublisher(cfg)
	err := pubsub.Publish("test", 1)
	assert.Nil(t, err)
}

func Test_Post_Error(t *testing.T) {
	ts := CreateServer(t, http.StatusInternalServerError)
	defer ts.Close()
	cfg := &ConfigRestPublisher{
		URL: ts.URL,
	}
	pubsub := NewRestPublisher(cfg)
	err := pubsub.Publish("test", 1)
	assert.NotNil(t, err)
}

func CreateServer(t *testing.T, statusCode int) *httptest.Server {
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
