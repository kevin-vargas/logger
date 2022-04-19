package encoder

import (
	"testing"

	"github.com/kevin-vargas/logger/encoder/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap/zapcore"
)

func Test_Get(t *testing.T) {
	mockBaseEncoder := &mocks.BaseEncoder{}
	instance := Get(mockBaseEncoder)
	assert.NotNil(t, instance)
}

func Test_Encoder_Add_Strings(t *testing.T) {
	field := "test"
	mockBaseEncoder := &mocks.BaseEncoder{}
	mockBaseEncoder.On("AddArray", field, mock.Anything).Return(nil)
	instance := &objectEncoder{
		mockBaseEncoder,
	}
	values := []string{"test", "test2", "test3"}
	instance.AddStrings(field, values)
	mockBaseEncoder.AssertNumberOfCalls(t, "AddArray", 1)
}

func Test_Encoder_Add_Strings_Valid(t *testing.T) {
	field := "test"
	testCase := []struct {
		desc   string
		value  []string
		expect int
	}{
		{
			desc:   "valid",
			value:  []string{"t", "t2"},
			expect: 1,
		},
		{
			desc:   "not valid",
			value:  []string{},
			expect: 0,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.desc, func(t *testing.T) {
			mockBaseEncoder := &mocks.BaseEncoder{}
			mockBaseEncoder.On("AddArray", field, mock.Anything).Return(nil)
			instance := &objectEncoder{
				mockBaseEncoder,
			}
			instance.AddStringsValid(field, tt.value)
			mockBaseEncoder.AssertNumberOfCalls(t, "AddArray", tt.expect)
		})
	}
}

func Test_Encoder_Add_Bytes_Valid(t *testing.T) {
	field := "test"
	testCase := []struct {
		desc   string
		value  []byte
		expect int
	}{
		{
			desc:   "valid",
			value:  []byte("t"),
			expect: 1,
		},
		{
			desc:   "not valid",
			value:  []byte{},
			expect: 0,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.desc, func(t *testing.T) {
			mockBaseEncoder := &mocks.BaseEncoder{}
			mockBaseEncoder.On("AddByteString", field, mock.Anything).Return(nil)
			instance := &objectEncoder{
				mockBaseEncoder,
			}
			instance.AddBytesValid(field, tt.value)
			mockBaseEncoder.AssertNumberOfCalls(t, "AddByteString", tt.expect)
		})
	}
}

func Test_Encoder_Add_Object_Valid(t *testing.T) {
	var forTest *mocks.ObjectMarshaler = nil
	field := "test"
	testCase := []struct {
		desc   string
		value  zapcore.ObjectMarshaler
		expect int
	}{
		{
			desc:   "valid",
			value:  &mocks.ObjectMarshaler{},
			expect: 1,
		},
		{
			desc:   "not valid",
			value:  forTest,
			expect: 0,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.desc, func(t *testing.T) {
			mockBaseEncoder := &mocks.BaseEncoder{}
			mockBaseEncoder.On("AddObject", field, mock.Anything).Return(nil)
			instance := &objectEncoder{
				mockBaseEncoder,
			}
			instance.AddObjectValid(field, tt.value)
			mockBaseEncoder.AssertNumberOfCalls(t, "AddObject", tt.expect)
		})
	}
}

func Test_Encoder_Add_String_Valid(t *testing.T) {
	field := "test"
	testCase := []struct {
		desc   string
		value  string
		expect int
	}{
		{
			desc:   "valid",
			value:  "testing",
			expect: 1,
		},
		{
			desc:   "not valid",
			value:  "",
			expect: 0,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.desc, func(t *testing.T) {
			mockBaseEncoder := &mocks.BaseEncoder{}
			mockBaseEncoder.On("AddString", field, mock.Anything).Return(nil)
			instance := &objectEncoder{
				mockBaseEncoder,
			}
			instance.AddStringValid(field, tt.value)
			mockBaseEncoder.AssertNumberOfCalls(t, "AddString", tt.expect)
		})
	}
}
