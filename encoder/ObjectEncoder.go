package encoder

import (
	"reflect"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ObjectEncoder struct {
	zapcore.ObjectEncoder
}

func Get(enc zapcore.ObjectEncoder) ObjectEncoder {
	return ObjectEncoder{
		enc,
	}
}

func (enc *ObjectEncoder) AddStrings(key string, value []string) {
	zap.Strings(key, value).AddTo(enc)
}

func (enc *ObjectEncoder) AddStringsValid(key string, value []string) {
	if len(value) > 0 {
		enc.AddStrings(key, value)
	}
}
func (enc *ObjectEncoder) AddBytesValid(key string, value []byte) {
	if len(value) > 0 {
		enc.AddByteString(key, value)
	}
}

// TODO: remove use of reflection
func (enc *ObjectEncoder) AddObjectValid(key string, marshaler zapcore.ObjectMarshaler) {
	if !reflect.ValueOf(marshaler).IsNil() {
		enc.AddObject(key, marshaler)
	}
}

func (enc *ObjectEncoder) AddStringValid(key string, value string) {
	if value != "" {
		enc.AddString(key, value)
	}
}
