package pubsub

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

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
