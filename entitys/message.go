package entitys

import (
	"github.com/kevin-vargas/logger/encoder"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Field = zap.Field

type EncodeConfig struct {
	LVL Level
}

type Message struct {
	Text         string
	tags         Tags
	labels       Labels
	log          Log
	event        *Event
	trace        *Trace
	err          *Error
	httpRequest  *HTTPRequest
	httpResponse *HTTPResponse
}

func NewMessage(msg string) Message {
	return Message{
		Text:   msg,
		tags:   make(Tags, 0),
		labels: make(Labels),
	}
}
func (message *Message) WithLabels(labels Labels) *Message {
	for key, element := range labels {
		message.labels[key] = element
	}
	return message
}

func (message *Message) WithLoggerInfo(log Log) *Message {
	message.log = log
	return message
}

func (message *Message) WithTags(tags Tags) *Message {
	message.tags = append(message.tags, tags...)
	return message
}

func (message *Message) WithHttpRequest(req HTTPRequest) *Message {
	message.httpRequest = &req
	return message
}

func (message *Message) WithHttpReponse(res HTTPResponse) *Message {
	message.httpResponse = &res
	return message
}

func (message *Message) WithEvent(event Event) *Message {
	message.event = &event
	return message
}

func (message *Message) WithTrace(trace Trace) *Message {
	message.trace = &trace
	return message
}

func (message *Message) WithError(err Error) *Message {
	message.err = &err
	return message
}

func (message *Message) Encode(config EncodeConfig) Field {
	defaultLabels := GetDefaultLabels()
	defaultLog := GetDefaultLog(config.LVL)
	message.WithLabels(defaultLabels).WithLoggerInfo(defaultLog)
	return zap.Inline(message)
}

func (message *Message) MarshalLogObject(zapenc zapcore.ObjectEncoder) error {
	enc := encoder.Get(zapenc)
	enc.AddObject(fieldLabels, &message.labels)
	enc.AddObject(fieldLog, &message.log)
	enc.AddStringsValid(fieldTags, message.tags)
	enc.AddObjectValid(fieldHttpRequest, message.httpRequest)
	enc.AddObjectValid(fieldHttpResponse, message.httpResponse)
	enc.AddObjectValid(fieldEvent, message.event)
	enc.AddObjectValid(fieldTrace, message.trace)
	enc.AddObjectValid(fieldError, message.err)
	return nil
}
