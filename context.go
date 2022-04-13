package logger

import "context"

const (
	logCtxKey = logContextKey("_santanderLogger")
)

type logContextKey string

func WithLogger(ctx context.Context, l Logger) context.Context {
	return context.WithValue(ctx, logCtxKey, l)
}
