
# ⚙️ Logging and tracing Lib

  

Table of contents:

- [Logging and tracing Lib](#️-logging-and-tracing-lib)

- [Summary](#summary)

- [Library structure](#library-structure)

- [Usage](#usage)

- [Library Setup](#-library-setup)

- [Logging](#-logging)

- [Automatic request response log setup](#-automatic-request-response-log-setup)

- [Logging stuff](#logging-stuff)

- [Building a message](#building-a-message)

- [Audit](#Audit)


## Summary

  

This is a logging library coded in golang and intended to standardize the *__way__*, *__structure__* and *__output__* of the application and audit logs.

As logging engine it is using [zap](https://github.com/uber-go/zap).

The default behavior of the logs is to log to stdout. There will be an additional process that will collect this logs and post them to an elastic search for visualization and analysis as per the following.

  

## Library structure

  

The project is structured in the following way:

```

|-- audit

|-- config

|-- encoder

|-- entities

|-- loggertest

|-- middleware

|-- pubsub

|-- strings

|-- variables

```

Brief project structure summary:

  

- audit: `Audit logging definition`.

- config: `Config entity definition`.

- encoder: `Custom encoder with validations`.

- entities: `Domain entities used in the library`.

- loggertest: `Utils for test logs`.

- middleware: `Http interceptors implemented for logging purposes`.

- pubsub: `Pubsub definition and implementation`.

- strings: `String Utilities used within the library`.

- variables: `Process environment variables validations`.  


## Usage

  

### Library Setup

---

In order to use the library we need to add it to our application via 
```
go get github.com/kevin-vargas/logger
```
#### Library Configuration Properties
In order to finish the basic setup of the library we need to specify the configuration properties to construct the logger. For proporuse this we are going to instantiate a configuration object:
```go
	cfg := config.
	New("test_app", "test_service", "test_env").
	WithAudit("url", "user", "pass").
	WithEnvironment("pod_name", "node_name")
```
### Logging
The recommended way to use the logger is with a config instance
```go
	var cfg *config.Logger
	log, err := logger.New(logger.WithConfig(cfg))
```
We can set custom output for our logs
```go
	// example not using stdout
	buf := &bytes.Buffer{}
	log, err := logger.New(logger.WithIoWriter(buf))
```
To set default labels per logger instance we can use the next option
```go
	// example not using stdout
	var labels entities.Labels
	log, err := logger.New(logger.WithLabels(labels))
```

---
### Automatic request response log setup

This library considers the possibility to automatically log the http requests and responses applying a middleware,

---

> ⚠️ **This will obviously add an overhead to EVERY request**: Be very careful here and only use it if you absolutely need to log requests and responses anyhow!

---

In order to setup the automatic requests & responses logger you need to configure it in you application configuration file and specify the interceptors as the following description illustrates:

First we need to add the needed require/import statement:

```go
import "github.com/kevin-vargas/logger/middleware"
```

Once the require/import statement was added we will need to proceed with get an instance of the logging handler we can pass options

  

```go

var cfg *config.Logger
loggingHandler, err := middleware.NewLoggingHandler(cfg)

// Handler applied to '/ping' endpoint on request and response
mux.HandleFunc("/ping", loggingHandler.Handle(PingHandler))

// Logging apply to all endpoints
server := http.Server{Addr: ":9001", Handler: loggingHandler.Logging(mux)}
server.ListenAndServe()

```

With the option patter we can pass custom behavior

```go

// Setup custom behaviour on response log
func WithResponseLogging(rl ResponseLogger) Option
// Setup custom behaviour on request log
func WithRequestLogging(rl RequestLogger) Option
// Setup a custom logger
func WithLogger(logger logger.Logger) Option]

var l logger.Logger
var rpl middleware.ResponseLogger
var rl middleware.RequestLogger

withLogger := middleware.WithLogger(l)
withResponseLogging := middleware.WithResponseLogging(rpl)
withRequestLogging := middleware.WithRequestLogging(rpl)

loggingHandler, err := middleware.NewLoggingHandler(cfg, withLogger, withResponseLogging, withRequestLogging)


```
  

### Logging stuff

  

So we already configured the library to use the logger, now the fun part.

  

For example, in the controller where you need to log data you will start by adding the required references with the following require/import statements:

  

```go

import "github.com/kevin-vargas/logger"

```

After including the references from the previous require/import statement we can start performing the logs through the logger object like this:

```go
// code to execute in order to perform an application logging as info severity
var cfg *config.Logger
const log = logger.New(cfg);
var msg *entities.Message
log.info(msg)

```

  

#### Building a message

  

In order to build the log message you must use the Message object. This entity will be populated using the following methods:

Instance:

```go

func  NewMessage(msg string) *Message;

```

  

Chain methods:

```go

func (message *Message) WithTags(tags Tags) *Message

func (message *Message) WithEvent(event Event) *Message

func (message *Message) WithTrace(trace Trace) *Message

func (message *Message) WithLoggerInfo(log Log) *Message

func (message *Message) WithLabels(labels Labels) *Message

func (message *Message) WithError(err Error) *Message
  

// these chain methods are to be used exclusively by the request interceptor methods

func (message *Message) WithHttpRequest(req HTTPRequest) *Message

func (message *Message) WithHttpReponse(res HTTPResponse) *Message

```
Example of building a msg
```go
var erre = entities.Error{
	Message: "_message",
	Type: "_type",
	StackTrace: "_stack_trace",
}
var trace = entities.Trace{
	ID: "id_trace",
}
var event = entities.Event{
	Action: "action",
	Category: []string{"category1", "category2"},
	Module: "module",
	Type: "type",
	Original: "original",
}
var tags = entities.Tags{"tag1", "tag2", "tag3"}
var msg = entities.NewMessage("msg").
WithError(erre).
WithEvent(event).
WithTrace(trace).
WithTags(tags)
```

### Audit
We can use a custom audit client when making logging instance
```go
var cfg *config.Logger
var auditClientCustom audit.Client 
log, err := logger.New(logger.WithConfig(cfg),logger.WithAuditClient(auditClientCustom))

```
if we use the default logger then we will use the default audit or we can specify it as follows
```go
var cfg *config.Logger
var defaultClientAudit audit.Client = audit.Get(cfg.Audit)
log, err := logger.New(logger.WithConfig(cfg), logger.WithAuditClient(defaultClientAudit))

```
Same as
```go
var cfg *config.Logger
log := logger.New(logger.WithConfig(cfg))

```
Now we can use the audit method with a audit message
```go
var msg *audit.Message = &audit.Message{
	Topic: "topic_test",
	Payload: audit.Payload{
	Type: audit.DATABASE_REQUEST,
	Nup: "nup",
	CorrelationId: "correlation_id",
	},
}
log.Audit(msg)
```
### Fallback

To override the default behavior when auditing fails, we need to pass this custom behavior when we create the instance of logging, we use the Option patter for this.

```go
var cfg *config.Logger
var fallback fallback audit.FallBackMethod

withFallBack := logger.WithFallback(fallback)
withConfig := logger.WithConfig(cfg)

log := logger.New(withConfig, withFallBack)

```
