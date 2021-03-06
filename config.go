package logger

import (
	"io"
	"sync"

	"github.com/kevin-vargas/logger/audit"
	"github.com/kevin-vargas/logger/config"
	"github.com/kevin-vargas/logger/entities"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	fieldMessage   = "message"
	fieldTimestamp = "@timestamp"
)

func buildConfig() zap.Config {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"stdout"}
	cfg.ErrorOutputPaths = []string{"stdout"}
	cfg.EncoderConfig.MessageKey = fieldMessage
	cfg.EncoderConfig.TimeKey = fieldTimestamp
	cfg.EncoderConfig.LevelKey = "" // Omit it, we will generate it on our own (it conflicts with ObjectEncoder)
	cfg.DisableStacktrace = true    // Omit automatic stacktraces, these will be emitted by the recovery middleware
	cfg.DisableCaller = true
	return cfg
}

type Option func(*SantanderLogger)

func WithIoWriter(w io.Writer) Option {
	return func(santanderLogger *SantanderLogger) {
		ws := zapcore.AddSync(w)

		jsonEncoder := zapcore.NewJSONEncoder(buildConfig().EncoderConfig)
		core := zapcore.NewCore(jsonEncoder, ws, zap.DebugLevel)

		santanderLogger.logger = zap.New(core)
	}
}

func WithAuditClient(client audit.Client) Option {
	return func(santanderLogger *SantanderLogger) {
		santanderLogger.auditClient = client
	}
}

func WithConfig(config *config.Logger) Option {
	return func(santanderLogger *SantanderLogger) {
		santanderLogger.config = config
	}
}

func WithFallBack(fallback audit.FallBackMethod) Option {
	return func(santanderLogger *SantanderLogger) {
		santanderLogger.fallback = fallback
	}
}

func WithLabels(labels entities.Labels) Option {
	return func(santanderLogger *SantanderLogger) {
		santanderLogger.defaultLabels = labels
	}
}

func New(options ...Option) (Logger, error) {
	cfg := buildConfig()
	zapLogger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	logger := &SantanderLogger{
		once:   &sync.Once{},
		logger: zapLogger,
	}

	for _, opt := range options {
		opt(logger)
	}

	return logger, nil
}
