package middleware

import "net/http"

type ResponseObserver struct {
	http.ResponseWriter
	Response    []byte
	Status      int
	Written     int64
	wroteHeader bool
}

func (o *ResponseObserver) Write(p []byte) (n int, err error) {
	if !o.wroteHeader {
		o.WriteHeader(http.StatusOK)
	}
	n, err = o.ResponseWriter.Write(p)
	o.Response = p
	o.Written += int64(n)
	return
}

func (o *ResponseObserver) WriteHeader(code int) {
	o.ResponseWriter.WriteHeader(code)
	if o.wroteHeader {
		return
	}
	o.wroteHeader = true
	o.Status = code
}
