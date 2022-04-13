package middleware

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/kevin-vargas/logger"
	"github.com/kevin-vargas/logger/entitys"
)

type (
	RequestLogger  func(l logger.Logger, r *http.Request) error
	ResponseLogger func(l logger.Logger, r *ResponseObserver)
	LoggerBuilder  func(r *http.Request) logger.Logger
)

type Option func(*LoggingHandler)

//WithResponseLogging allows to set a custom behavior for logging the outgoing Response
func WithResponseLogging(rl ResponseLogger) Option {
	return func(args *LoggingHandler) {
		args.respLog = rl
	}
}

//WithRequestLogging allows to set a custom behaviour for logging the incoming request
func WithRequestLogging(rl RequestLogger) Option {
	return func(args *LoggingHandler) {
		args.reqLog = rl
	}
}

//WithLogger allows to build a new logger from the current http.Request
func WithLogger(builder LoggerBuilder) Option {
	return func(args *LoggingHandler) {
		args.builder = builder
	}
}

type LoggingHandler struct {
	builder LoggerBuilder
	reqLog  RequestLogger
	respLog ResponseLogger
}

func NewLoggingHandler(options ...Option) *LoggingHandler {
	handler := &LoggingHandler{
		reqLog:  defaultLogRequest,
		respLog: defaultLogResponse,
		builder: defaultLoggerBuilder,
	}

	for _, opt := range options {
		opt(handler)
	}

	return handler
}
func (h *LoggingHandler) Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := h.builder(r)

		ctx := logger.WithLogger(r.Context(), l)
		r = r.Clone(ctx)

		o := &ResponseObserver{ResponseWriter: w}

		err := h.reqLog(l, r)
		if err != nil {
			// TODO: log with err
			// l.Error("LoggingHandler: An error occurred during request logging %v", logger.Err(err))
		}

		next.ServeHTTP(o, r)
		h.respLog(l, o)

	})
}

//Handle returns a wrapped handler to log http.Request/http.Response with the provided logger
// said logger wil be propagated through the request context
func (h *LoggingHandler) Handle(next http.HandlerFunc) http.HandlerFunc {
	withLogging := h.Logging(next)
	return withLogging.ServeHTTP
}

func defaultLogRequest(l logger.Logger, r *http.Request) (err error) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	r.Body.Close()
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	req := entitys.HTTPRequest{
		Method:   r.Method,
		Referrer: r.Referer(),
		Body: &entitys.HTTPRequestBody{
			Content: bodyBytes,
			Headers: entitys.Headers(r.Header),
		},
	}
	msg := entitys.NewMessage("request").WithHttpRequest(req)
	l.Info(*msg)
	return
}

func defaultLogResponse(l logger.Logger, o *ResponseObserver) {
	res := entitys.HTTPResponse{
		StatusCode: int64(o.Status),
		Body: &entitys.HTTPResponseBody{
			Content: o.Response,
		},
	}
	msg := entitys.NewMessage("Response").WithHttpReponse(res)
	l.Info(*msg)
}

//defaultLoggerBuilder builds a default logger and adds it to the current context
func defaultLoggerBuilder(r *http.Request) logger.Logger {
	l, err := logger.NewLogger()
	if err != nil {
		return nil
	}

	return l
}
