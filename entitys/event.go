package entities

import (
	"github.com/kevin-vargas/logger/encoder"
	"go.uber.org/zap/zapcore"
)

type Event struct {
	Action   string
	Category []string
	Module   string
	Type     string
	Original string
}

func (event *Event) MarshalLogObject(zapenc zapcore.ObjectEncoder) error {
	enc := encoder.Get(zapenc)
	enc.AddStringValid(fieldEventAction, event.Action)
	enc.AddStringValid(fieldEventModule, event.Module)
	enc.AddStringValid(fieldEventType, event.Type)
	enc.AddStringValid(fieldEventOriginal, event.Original)
	enc.AddStringsValid(fieldEventCategory, event.Category)
	return nil
}
