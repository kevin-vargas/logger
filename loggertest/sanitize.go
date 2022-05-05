package loggertest

import (
	"fmt"
	"strings"
)

var token = fmt.Sprintf("\"%v\":", "@timestamp")

func sanitizeTimestamp(raw []byte) []byte {
	hits := 0
	for i, elem := range raw {
		if elem == token[hits] {
			if hits == len(token)-1 {
				j := i + 1
				start := j
				for ('0' <= raw[j] && raw[j] <= '9') || raw[j] == '.' {
					if j == start {
						raw[j] = '0'
					} else {
						raw[j] = 0
					}
					j++
				}
				return raw
			}
			hits++
		} else {
			hits = 0
		}
	}
	return raw
}

const (
	null_char = "\x00"
	empty     = ""
)

func Sanitize(raw []byte) string {

	raw = sanitizeTimestamp(raw)

	str := string(raw)

	return strings.Replace(str, null_char, empty, -1)
}
