package logger

import (
	"fmt"
	"strings"

	"github.com/kevin-vargas/logger/entities"
)

var token = fmt.Sprintf("\"%v\":", fieldTimestamp)

func sanitizeTimestamp(raw []byte) []byte {
	hits := 0
	for i, elem := range raw {
		if elem == token[hits] {
			if hits == len(token)-1 {
				j := i + 1
				start := j
				for ('0' <= raw[j] && raw[j] <= '9') || raw[j] == '.' {
					if j == start {
						raw[j] = '0'
					} else {
						raw[j] = 0
					}
					j++
				}
				return raw
			}
			hits++
		} else {
			hits = 0
		}
	}
	return raw
}

const (
	NULL_CHAR = "\x00"
	EMPTY     = ""
)

func sanitize(raw []byte) string {

	raw = sanitizeTimestamp(raw)

	str := string(raw)

	return strings.Replace(str, NULL_CHAR, EMPTY, -1)
}

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
