package entities

import "go.uber.org/zap/zapcore"

type Log struct {
	Logger string
	Level  string
}

func (log *Log) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString(fieldLogLogger, log.Logger)
	enc.AddString(fieldLogLevel, log.Level)
	return nil
}
