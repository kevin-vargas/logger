package logger

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/kevin-vargas/logger/entitys"
)

var loggTwo, _ = NewLogger()

var result = []byte{}

func Test_WithoutLogger(t *testing.T) {
	buf := &bytes.Buffer{}
	logger, err := NewLogger(WithIoWriter(buf))
	if err != nil {
		t.Error(err)
	}
	res := entitys.HTTPResponse{
		StatusCode: 201,
		Body: &entitys.HTTPResponseBody{
			Content: []byte("testing"),
		},
	}
	req := entitys.HTTPRequest{
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
	erre := entitys.Error{
		Message:    "_message",
		Type:       "_type",
		StackTrace: "_stack_trace",
	}
	trace := entitys.Trace{
		ID: "id_trace",
	}
	event := entitys.Event{
		Action:   "action",
		Category: []string{"category1", "category2"},
		Module:   "module",
		Type:     "type",
		Original: "original",
	}
	tags := entitys.Tags{"tag1", "tag2", "tag3"}

	var msg = entitys.NewMessage("msg")

	msg.
		WithHttpReponse(res).
		WithHttpRequest(req).
		WithError(erre).
		WithEvent(event).
		WithTrace(trace).
		WithTags(tags)

	logger.Info(msg)
	result := sanitizeTimestamp(buf.Bytes())
	fmt.Println(string(result))
}
