package logger

import "fmt"

var token = fmt.Sprintf("\"%v\":", fieldTimestamp)

// to timestamp asserts
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
