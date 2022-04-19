package variables

import (
	"errors"
	"os"
	"strings"

	lstrings "github.com/kevin-vargas/logger/strings"
)

// TODO: Cache env variables?
func OR(envVar string, defaultStr string) string {
	return lstrings.OR(os.Getenv(envVar), defaultStr)
}

func ORError(envVar string) (str string, err error) {
	result := os.Getenv(envVar)
	if result == "" {
		var sb strings.Builder
		sb.WriteString("Not found env variable: ")
		sb.WriteString(envVar)
		err = errors.New(sb.String())
	}
	return result, err
}

func ORPanic(envVar string) (str string) {
	result := os.Getenv(envVar)
	if result == "" {
		var sb strings.Builder
		sb.WriteString("Not found env variable: ")
		sb.WriteString(envVar)
		panic(sb.String())
	}
	return result
}
