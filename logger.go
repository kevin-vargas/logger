package logger

import (
	"encoding/json"
	"sync"

	"github.com/kevin-vargas/logger/audit"
	"github.com/kevin-vargas/logger/config"
	"github.com/kevin-vargas/logger/entities"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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
	once          *sync.Once
	config        *config.Logger
	auditClient   audit.Client
	logger        *zap.Logger
	defaultLabels entities.Labels
	fallback      audit.FallBackMethod
}

func (l *SantanderLogger) configMessage(lvl entities.Level, message *entities.Message) (string, zapcore.Field) {
	config := entities.EncodeConfig{
		LVL: lvl,
	}
	l.addDefaultFieldsMessage(message)
	return message.Text, message.Encode(config)
}

func (l *SantanderLogger) Debug(message *entities.Message) {
	text, field := l.configMessage(entities.DebugLevel, message)
	l.logger.Debug(text, field)
}

func (l *SantanderLogger) Info(message *entities.Message) {
	text, field := l.configMessage(entities.InfoLevel, message)
	l.logger.Info(text, field)
}

func (l *SantanderLogger) Warn(message *entities.Message) {
	text, field := l.configMessage(entities.WarnLevel, message)
	l.logger.Warn(text, field)
}

func (l *SantanderLogger) Error(message *entities.Message) {
	text, field := l.configMessage(entities.ErrorLevel, message)
	l.logger.Error(text, field)
}

func (l *SantanderLogger) Audit(message *audit.Message) {
	l.once.Do(func() {
		if l.auditClient == nil {
			l.auditClient = audit.New(l.config.Audit)
		}
		if l.fallback == nil {
			l.fallback = defaultAuditFallBack(l)
		}
	})
	err := l.auditClient.Audit(message, l.fallback)
	if err != nil {
		msg := entities.
			NewMessage("On Audit").
			WithError(entities.Error{
				Message: err.Error(),
			})
		l.Error(msg)
	}
}

func defaultAuditFallBack(l Logger) audit.FallBackMethod {
	return func(topic string, payload *audit.Payload) {
		json, _ := json.Marshal(payload)
		trace := entities.Trace{
			ID: payload.CorrelationId,
		}
		msg := entities.NewMessage(string(json)).
			WithTags([]string{"audit"}).
			WithTrace(trace)
		l.Info(msg)
	}
}

func (l *SantanderLogger) addDefaultFieldsMessage(m *entities.Message) *entities.Message {
	if l.config != nil {
		m.
			WithLabels(entities.GetDefaultLabels(l.config))
	}
	if l.defaultLabels != nil {
		m.WithLabels(l.defaultLabels)
	}
	return m
}
