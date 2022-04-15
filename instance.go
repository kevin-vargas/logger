package logger

import "sync"

var instance Logger
var once sync.Once

// Get global singleton logger
func Get() Logger {
	once.Do(func() {
		newInstance, err := NewLogger()
		if err != nil {
			panic(err)
		}
		instance = newInstance
	})
	return instance
}
