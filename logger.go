package logger

import (
	"github.com/kevin-vargas/logger/audit"
	"github.com/kevin-vargas/logger/entitys"
	"go.uber.org/zap"
)

type Field = zap.Field

type Logger interface {
	Debug(message *entitys.Message)
	Info(message *entitys.Message)
	Warn(message *entitys.Message)
	Error(message *entitys.Message)
	// Panic(msg string, fields ...Field)
	// Fatal(msg string, fields ...Field)
}

type SantanderLogger struct {
	auditLogger audit.Client
	logger      *zap.Logger
}

func (l *SantanderLogger) Debug(message *entitys.Message) {
	config := entitys.EncodeConfig{
		LVL: entitys.DebugLevel,
	}
	l.logger.Debug(message.Text, message.Encode(config))
}

func (l *SantanderLogger) Info(message *entitys.Message) {
	config := entitys.EncodeConfig{
		LVL: entitys.InfoLevel,
	}
	l.logger.Info(message.Text, message.Encode(config))
}

func (l *SantanderLogger) Warn(message *entitys.Message) {
	config := entitys.EncodeConfig{
		LVL: entitys.WarnLevel,
	}
	l.logger.Warn(message.Text, message.Encode(config))
}

func (l *SantanderLogger) Error(message *entitys.Message) {
	config := entitys.EncodeConfig{
		LVL: entitys.ErrorLevel,
	}
	l.logger.Error(message.Text, message.Encode(config))
}

// TODO: safe concurrent audit
func (l *SantanderLogger) Audit(message *audit.Message) {
	if l.auditLogger == nil {
		l.auditLogger = audit.Get()
	}
	l.auditLogger.Audit(message)
}
