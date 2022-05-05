package middleware

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/kevin-vargas/logger"
	"github.com/kevin-vargas/logger/config"
	"github.com/kevin-vargas/logger/entities"
)

type (
	LoggerBuilder  func(c *config.Logger) logger.Logger
	RequestLogger  func(l logger.Logger, r *http.Request) error
	ResponseLogger func(l logger.Logger, r *ResponseObserver)
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

func WithLoggerBuilder(lb LoggerBuilder) Option {
	return func(args *LoggingHandler) {
		args.builder = lb
	}
}

type LoggingHandler struct {
	config  *config.Logger
	builder LoggerBuilder
	reqLog  RequestLogger
	respLog ResponseLogger
}

func NewLoggingHandler(c *config.Logger, options ...Option) (*LoggingHandler, error) {

	handler := &LoggingHandler{
		config:  c,
		reqLog:  defaultLogRequest,
		respLog: defaultLogResponse,
		builder: defaultLoggerBuilder,
	}

	for _, opt := range options {
		opt(handler)
	}

	return handler, nil
}
func (h *LoggingHandler) Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		l := h.builder(h.config)

		ctx := createRequestContext(req, withLoggerCTX(l))

		req = req.Clone(ctx)

		observer := &ResponseObserver{ResponseWriter: writer}

		err := h.reqLog(l, req)

		if err != nil {
			msg := fmt.Sprintf("LoggingHandler: An error occurred during request logging %s", err.Error())
			l.Error(entities.NewMessage(msg))
		}

		next.ServeHTTP(observer, req)
		h.respLog(l, observer)

	})
}

// Handle returns a wrapped handler to log http.Request/http.Response with the provided logger
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
	req := entities.HTTPRequest{
		Method:   r.Method,
		Referrer: r.Referer(),
		Body: &entities.HTTPRequestBody{
			Content: bodyBytes,
			Headers: entities.Headers(r.Header),
		},
	}
	msg := entities.NewMessage("Request").WithHttpRequest(req)

	ctx := r.Context()

	if id, ok := GetTraceId(ctx); ok {
		msg.WithTrace(entities.Trace{
			ID: id,
		})
	}

	l.Info(msg)
	return
}

func defaultLogResponse(l logger.Logger, o *ResponseObserver) {
	res := entities.HTTPResponse{
		StatusCode: int64(o.Status),
		Body: &entities.HTTPResponseBody{
			Content: o.Response,
		},
	}
	msg := entities.NewMessage("Response").WithHttpReponse(res)
	l.Info(msg)
}

func defaultLoggerBuilder(c *config.Logger) logger.Logger {
	withConfig := logger.WithConfig(c)
	l, err := logger.New(withConfig)
	if err != nil {
		return nil
	}
	return l
}
