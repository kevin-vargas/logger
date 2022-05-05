package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/kevin-vargas/logger"
)

const (
	contexKeyLogger   = contextKey("logger")
	contextKeyTraceId = contextKey("trace_id")
)

type contextKey string

func (c contextKey) String() string {
	return "request" + string(c)
}

func GetLogger(ctx context.Context) (logger.Logger, bool) {
	l, ok := ctx.Value(contexKeyLogger).(logger.Logger)
	return l, ok
}

func GetTraceId(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(contextKeyTraceId).(string)
	return id, ok
}

func withLoggerCTX(l logger.Logger) optionsCTX {
	return func(ctx context.Context) context.Context {
		ctx = context.WithValue(ctx, contexKeyLogger, l)
		return ctx
	}
}

type optionsCTX func(context.Context) context.Context

func createRequestContext(req *http.Request, options ...optionsCTX) (ctx context.Context) {
	ctx = req.Context()

	id := req.Header.Get(correlation_id_header)

	if id == "" {
		correlationIdLower := strings.ToLower(correlation_id_header)
		id = req.Header.Get(correlationIdLower)
	}

	if id != "" {
		ctx = context.WithValue(ctx, contextKeyTraceId, id)
	}

	for _, option := range options {
		ctx = option(ctx)
	}

	return
}
