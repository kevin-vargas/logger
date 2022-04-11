package entitys

import "go.uber.org/zap/zapcore"

type Labels map[string]string

func (labels *Labels) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	for key, element := range *labels {
		enc.AddString(key, element)
	}
	return nil
}
