package middleware

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/kevin-vargas/logger"
	"github.com/kevin-vargas/logger/entities"
)

type (
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

func WithLogger(logger logger.Logger) Option {
	return func(args *LoggingHandler) {
		args.instance = logger
	}
}

type LoggingHandler struct {
	instance logger.Logger
	reqLog   RequestLogger
	respLog  ResponseLogger
}

func NewLoggingHandler(options ...Option) *LoggingHandler {
	handler := &LoggingHandler{
		reqLog:   defaultLogRequest,
		respLog:  defaultLogResponse,
		instance: logger.Get(),
	}

	for _, opt := range options {
		opt(handler)
	}

	return handler
}
func (h *LoggingHandler) Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		loggerInstance := h.instance

		observer := &ResponseObserver{ResponseWriter: writer}

		err := h.reqLog(loggerInstance, req)
		if err != nil {
			// TODO: log with err
			// l.Error("LoggingHandler: An error occurred during request logging %v", logger.Err(err))
		}

		next.ServeHTTP(observer, req)
		h.respLog(loggerInstance, observer)

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
	req := entities.HTTPRequest{
		Method:   r.Method,
		Referrer: r.Referer(),
		Body: &entities.HTTPRequestBody{
			Content: bodyBytes,
			Headers: entities.Headers(r.Header),
		},
	}
	msg := entities.NewMessage("Request").WithHttpRequest(req)
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
