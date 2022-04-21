package variables

import (
	"errors"
	"os"
	"strings"

	lstrings "github.com/kevin-vargas/logger/strings"
)

var getEnv = os.Getenv

// TODO: Cache env variables?
func OR(envVar string, defaultStr string) string {
	return lstrings.OR(getEnv(envVar), defaultStr)
}

func ORError(envVar string) (str string, err error) {
	result := getEnv(envVar)
	if result == "" {
		var sb strings.Builder
		sb.WriteString("Not found env variable: ")
		sb.WriteString(envVar)
		err = errors.New(sb.String())
	}
	return result, err
}

func ORPanic(envVar string) (str string) {
	result := getEnv(envVar)
	if result == "" {
		var sb strings.Builder
		sb.WriteString("Not found env variable: ")
		sb.WriteString(envVar)
		panic(sb.String())
	}
	return result
}
