package pubsub

import "github.com/go-resty/resty/v2"

type logger struct {
}

func (l *logger) Errorf(format string, v ...interface{}) {
}
func (l *logger) Warnf(format string, v ...interface{}) {
}
func (l *logger) Debugf(format string, v ...interface{}) {
}

func discardLog() resty.Logger {
	return &logger{}
}
