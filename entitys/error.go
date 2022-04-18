package entities

import (
	"github.com/kevin-vargas/logger/encoder"
	"go.uber.org/zap/zapcore"
)

type Error struct {
	Message    string
	StackTrace string
	Type       string
}

func (err *Error) MarshalLogObject(zapenc zapcore.ObjectEncoder) error {
	enc := encoder.Get(zapenc)
	enc.AddStringValid(fieldErrorMessage, err.Message)
	enc.AddStringValid(fieldErrorStackTrace, err.StackTrace)
	enc.AddStringValid(fieldErrorType, err.Type)
	return nil
}
