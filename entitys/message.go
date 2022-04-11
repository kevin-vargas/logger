package entitys

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Field = zap.Field

type EncodeConfig struct {
	LVL Level
}

const (
	fieldLabels      = "labels"
	fieldLog         = "log"
	fieldTags        = "tags"
	fieldHttpRequest = "httprequest"
)

type Message struct {
	Text        string
	tags        Tags
	labels      Labels
	log         Log
	httpRequest *HTTPRequest
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

func (message *Message) Encode(config EncodeConfig) Field {
	defaultLabels := GetDefaultLabels()
	defaultLog := GetDefaultLog(config.LVL)
	message.WithLabels(defaultLabels).WithLoggerInfo(defaultLog)
	return zap.Inline(message)
}

func (message *Message) MarshalLogObject(zapenc zapcore.ObjectEncoder) error {
	enc := GetEncoder(zapenc)
	enc.AddObject(fieldLabels, &message.labels)
	enc.AddObject(fieldLog, &message.log)
	enc.AddStringsValid(fieldTags, message.tags)
	enc.AddObjectValid(fieldHttpRequest, message.httpRequest)
	return nil
}
