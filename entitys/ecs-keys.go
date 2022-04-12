package entitys

// Message
const (
	fieldLabels      = "labels"
	fieldLog         = "log"
	fieldTags        = "tags"
	fieldHttpRequest = "httprequest"
)

// HttpResponse
const (
	fieldHttpResponseBody        = "body"
	fieldHttpResponseBodyContent = "content"
	fieldHttpResponseStatusCode  = "status_code"
	fieldHttpResponse            = "httpresponse"
)

// HttpRequest
const (
	fieldHttpRequestMethod      = "method"
	fieldHttpRequestReferrer    = "referrer"
	fieldHttpRequestBody        = "body"
	fieldHttpRequestBodyContent = "content"
	fieldHttpRequestBodyHeaders = "headers"
)

// Log
const (
	fieldLogLogger = "logger"
	fieldLogLevel  = "level"
)

// Event
const (
	fieldEvent         = "event"
	fieldEventAction   = "action"
	fieldEventCategory = "category"
	fieldEventModule   = "module"
	fieldEventType     = "type"
	fieldEventOriginal = "original"
)

// Trace
const (
	fieldTrace   = "trace"
	fieldTraceId = "id"
)

// Error
const (
	fieldError           = "error"
	fieldErrorMessage    = "message"
	fieldErrorStackTrace = "stack_trace"
	fieldErrorType       = "type"
)

// Default Labels
const (
	fieldLabelApplication = "application"
	fieldLabelService     = "service"
	fieldLabelEnvironment = "environment"
	fieldLabelLibVersion  = "lib_version"
	fieldLabelLibLanguage = "lib_language"
	fieldLabelPodName     = "pod_name"
	fieldLabelNodeName    = "node_name"
)
