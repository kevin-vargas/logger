package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_OR(t *testing.T) {
	testCase := []struct {
		desc        string
		firstValue  string
		secondValue string
		expected    string
	}{
		{
			"Only first value",
			"first",
			"",
			"first",
		},
		{
			"Only second value",
			"",
			"second",
			"second",
		},
		{
			"Both values",
			"first",
			"second",
			"first",
		},
	}
	for _, tt := range testCase {
		t.Run(tt.desc, func(t *testing.T) {
			result := OR(tt.firstValue, tt.secondValue)
			assert.Equal(t, tt.expected, result)
		})
	}
}
