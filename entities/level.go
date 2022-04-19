package entities

import "github.com/kevin-vargas/logger/strings"

type Level int8

const (
	DebugLevel Level = iota - 1
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	FatalLevel
)

var levels map[Level]string = map[Level]string{
	DebugLevel: "debug",
	InfoLevel:  "info",
	WarnLevel:  "warn",
	ErrorLevel: "error",
	PanicLevel: "panic",
	FatalLevel: "fatal",
}

func (l Level) String() string {
	return strings.OR(levels[l], "UNKNOWN")
}
