package logger

import (
	"bytes"
	"testing"

	"github.com/kevin-vargas/logger/entities"
	"github.com/stretchr/testify/assert"
)

func Test_Logger(t *testing.T) {
	default_msg := "test"

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
