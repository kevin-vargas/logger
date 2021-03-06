package encoder

import "go.uber.org/zap/zapcore"

type BaseEncoder interface {
	zapcore.ObjectEncoder
}

type ObjectEncoder interface {
	BaseEncoder
	AddStrings(key string, value []string)
	AddStringsValid(key string, value []string)
	AddBytesValid(key string, value []byte)
	AddObjectValid(key string, marshaler zapcore.ObjectMarshaler) error
	AddStringValid(key string, value string)
}
