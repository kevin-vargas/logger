package variables

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var mockedOsGetEnv = func(key string) string {
	keys := map[string]string{
		"foo": "bar",
	}
	return keys[key]
}

func reset() {
	getEnv = os.Getenv
}
func Test_OR(t *testing.T) {
	t.Cleanup(reset)

	getEnv = mockedOsGetEnv

	testCase := []struct {
		desc       string
		env        string
		defaultStr string
		expected   string
	}{
		{
			"Missing env var",
			"missing",
			"default",
			"default",
		},
		{
			"With env var",
			"foo",
			"default",
			"bar",
		},
	}
	for _, tt := range testCase {
		t.Run(tt.desc, func(t *testing.T) {
			result := OR(tt.env, tt.defaultStr)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func Test_ORError(t *testing.T) {
	t.Cleanup(reset)

	getEnv = mockedOsGetEnv

	testCase := []struct {
		desc     string
		env      string
		errStr   string
		expected string
	}{
		{
			"Missing env var",
			"missing",
			"Not found env variable: missing",
			"",
		},
		{
			"With env var",
			"foo",
			"",
			"bar",
		},
	}
	for _, tt := range testCase {
		t.Run(tt.desc, func(t *testing.T) {
			result, err := ORError(tt.env)
			if err != nil {
				assert.EqualError(t, err, tt.errStr)
			} else {
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func Test_ORPanic(t *testing.T) {
	t.Cleanup(reset)

	getEnv = mockedOsGetEnv

	testCase := []struct {
		desc     string
		env      string
		panic    bool
		expected string
	}{
		{
			"Missing env var",
			"missing",
			true,
			"",
		},
		{
			"With env var",
			"foo",
			false,
			"bar",
		},
	}
	for _, tt := range testCase {
		t.Run(tt.desc, func(t *testing.T) {
			if tt.panic {
				assert.Panics(t, func() {
					ORPanic(tt.env)
				})
			} else {
				result := ORPanic(tt.env)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
