package logger

import "go.uber.org/zap"

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

func NewLogger() (Logger, error) {
	cfg := buildConfig()

	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}
	return &SantanderLogger{
		logger: logger,
	}, nil
}
