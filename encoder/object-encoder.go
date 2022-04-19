package encoder

import (
	"reflect"

	"go.uber.org/zap"
)

type objectEncoder struct {
	BaseEncoder
}

func Get(enc BaseEncoder) ObjectEncoder {
	return &objectEncoder{
		enc,
	}
}

func (enc *objectEncoder) AddStrings(key string, value []string) {
	zap.Strings(key, value).AddTo(enc)
}

func (enc *objectEncoder) AddStringsValid(key string, value []string) {
	if len(value) > 0 {
		enc.AddStrings(key, value)
	}
}
func (enc *objectEncoder) AddBytesValid(key string, value []byte) {
	if len(value) > 0 {
		enc.AddByteString(key, value)
	}
}

// TODO: remove use of reflection
func (enc *objectEncoder) AddObjectValid(key string, marshaler ObjectMarshaler) {
	if !reflect.ValueOf(marshaler).IsNil() {
		enc.AddObject(key, marshaler)
	}
}

func (enc *objectEncoder) AddStringValid(key string, value string) {
	if value != "" {
		enc.AddString(key, value)
	}
}
