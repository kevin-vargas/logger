package loggertest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sanitize(t *testing.T) {
	actual := []byte(`{"a":"b","@timestamp":123456789.123456789,"ff":{"a":"b"}}`)
	expect := `{"a":"b","@timestamp":0,"ff":{"a":"b"}}`
	result := Sanitize(actual)
	assert.Equal(t, expect, result)
}
