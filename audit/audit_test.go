package audit

import (
	"errors"
	"fmt"
	"testing"

	"github.com/kevin-vargas/logger/config"
	"github.com/kevin-vargas/logger/pubsub/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Audit_Message(t *testing.T) {
	testCase := []struct {
		desc         string
		topic        string
		defaultTopic string
		expect       string
	}{
		{
			desc:   "With out default",
			topic:  "topic",
			expect: "topic",
		},
		{
			desc:         "With Default",
			defaultTopic: "default",
			topic:        "topic",
			expect:       "default",
		},
	}
	for _, tt := range testCase {
		t.Run(tt.desc, func(t *testing.T) {
			mockPublisher := &mocks.Publisher{}
			mockPublisher.On("Publish", mock.AnythingOfType("string"), mock.Anything).Return(nil)
			msg := &Message{
				Topic: tt.topic,
				Payload: Payload{
					Nup:           "nup",
					CorrelationId: "correlation_id",
					SessionId:     "session_id",
					Type:          BUSINESS,
				},
			}
			client := &client{
				defaultTopic: tt.defaultTopic,
				publisher:    mockPublisher,
			}
			err := client.Audit(msg)
			if err != nil {
				t.Fail()
			}
			mockPublisher.AssertNumberOfCalls(t, "Publish", 1)
		})
	}

}
func Test_Publish_With_error(t *testing.T) {
	mockPublisher := &mocks.Publisher{}
	mockPublisher.On("Publish", mock.AnythingOfType("string"), mock.Anything).Return(errors.New("test"))
	msg := &Message{}
	client := &client{
		publisher: mockPublisher,
	}
	err := client.Audit(msg)
	mockPublisher.AssertNumberOfCalls(t, "Publish", 1)
	assert.Error(t, err)
}
func Test_Type_Marshal(t *testing.T) {
	testCase := []struct {
		actual Type
		expect []byte
	}{
		{
			actual: BUSINESS,
			expect: []byte("\"BUSINESS\""),
		},
		{
			actual: HTTP_REQUEST,
			expect: []byte("\"HTTP_REQUEST\""),
		},
	}
	for _, tt := range testCase {
		desc := fmt.Sprintf("%s, marshal", tt.actual)
		t.Run(desc, func(t *testing.T) {
			result, err := tt.actual.MarshalJSON()
			if err != nil {
				t.Fail()
			}
			assert.Equal(t, tt.expect, result)
		})
	}

}

func Test_toRestPubSubConfig(t *testing.T) {
	cfg := &config.Audit{}

	newConfig := toRestPubSubConfig(cfg)

	expect := RETRIES
	assert.Equal(t, expect, newConfig.Retries)
}
