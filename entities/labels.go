package entities

import (
	"github.com/kevin-vargas/logger/encoder"
	"go.uber.org/zap/zapcore"
)

type Labels map[string]string

func (labels *Labels) MarshalLogObject(zapenc zapcore.ObjectEncoder) error {
	enc := encoder.Get(zapenc)
	for key, element := range *labels {
		enc.AddStringValid(key, element)
	}
	return nil
}
