package audit

import (
	"bytes"
)

type Type int8

const (
	BUSINESS Type = iota - 1
	HTTP_REQUEST
	HTTP_RESPONSE
	DATABASE_REQUEST
	DATABASE_RESPONSE
)

var types map[Type]string = map[Type]string{
	BUSINESS:          "BUSINESS",
	HTTP_REQUEST:      "HTTP_REQUEST",
	HTTP_RESPONSE:     "HTTP_RESPONSE",
	DATABASE_REQUEST:  "DATABASE_REQUEST",
	DATABASE_RESPONSE: "DATABASE_RESPONSE",
}

func (t Type) String() string {
	return types[t]
}

func (t Type) MarshalJSON() ([]byte, error) {
	var buffer bytes.Buffer
	buffer.WriteByte('"')
	buffer.WriteString(t.String())
	buffer.WriteByte('"')
	return buffer.Bytes(), nil
}
