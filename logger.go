package logger

import (
	"sync"

	"github.com/kevin-vargas/logger/audit"
	"github.com/kevin-vargas/logger/entities"
	"go.uber.org/zap"
)

type Field = zap.Field

type Logger interface {
	Debug(message *entities.Message)
	Info(message *entities.Message)
	Warn(message *entities.Message)
	Error(message *entities.Message)
	Audit(message *audit.Message)
	// Panic(msg string, fields ...Field)
	// Fatal(msg string, fields ...Field)
}

type SantanderLogger struct {
	once        *sync.Once
	auditLogger audit.Client
	logger      *zap.Logger
}

func (l *SantanderLogger) Debug(message *entities.Message) {
	config := entities.EncodeConfig{
		LVL: entities.DebugLevel,
	}
	l.logger.Debug(message.Text, message.Encode(config))
}

func (l *SantanderLogger) Info(message *entities.Message) {
	config := entities.EncodeConfig{
		LVL: entities.InfoLevel,
	}
	l.logger.Info(message.Text, message.Encode(config))
}

func (l *SantanderLogger) Warn(message *entities.Message) {
	config := entities.EncodeConfig{
		LVL: entities.WarnLevel,
	}
	l.logger.Warn(message.Text, message.Encode(config))
}

func (l *SantanderLogger) Error(message *entities.Message) {
	config := entities.EncodeConfig{
		LVL: entities.ErrorLevel,
	}
	l.logger.Error(message.Text, message.Encode(config))
}

func (l *SantanderLogger) Audit(message *audit.Message) {
	l.once.Do(func() {
		if l.auditLogger == nil {
			l.auditLogger = audit.Get()
		}
	})
	l.auditLogger.Audit(message)
}
