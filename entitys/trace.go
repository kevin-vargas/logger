package entitys

import "go.uber.org/zap/zapcore"

type Trace struct {
	ID string
}

func (trace *Trace) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString(fieldTraceId, trace.ID)
	return nil
}
