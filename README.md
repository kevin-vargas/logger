
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

|-- encoder

|-- entitys

|-- middleware

|-- pubsub

|-- strings

|-- variables

```

Brief project structure summary:

  

- audit: `Audit logging definition`.

- encoder: `Custom encoder with validations`.

- entitys: `Domain entities used in the library`.

- middleware: `Http interceptors implemented for logging purposes`.

- pubsub: `Pubsub definition and implementation`.

- strings: `String Utilities used within the library`.

- variables: `Process environment variables validations`.  


## Usage

  

### Library Setup

---

In order to use the library we need to add it to our application via `go get github.com/kevin-vargas/logger` and configure it following these criteria:

We use the next env variables for logging:
|Name |Required|
|--|--|
| APPLICATION_NAME |NO|
|LOGGING_SERVICE_NAME|NO|
|ENVIRONMENT|NO|
|MY_POD_NAME|NO|
|MY_NODE_NAME|NO|
|AUDITS_ADAPTER_PUSH_URL|only for audit logging|
|AUDITS_TOPIC_NAME|only for audit logging|
|AUDITS_AUTHENTICATION_USER|only for audit logging|
|AUDITS_AUTHENTICATION_PASSWORD|only for audit logging|



### Logging
The recommended way to use the logger is to use the default global instance
```go
	log := logger.Get()
```
Or we can instantiate one with options we need
```go
	// example not using stdout
	buf := &bytes.Buffer{}
	log, err := logger.NewLogger(logger.WithIoWriter(buf))
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

// by default we use the global logger

loggingHandler := middleware.NewLoggingHandler()

// Handler applied to '/ping' endpoint on request and response
mux.HandleFunc("/ping", loggingHandler.Handle(PingHandler))

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

const log = logger.get();
var msg *entitys.Message
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
var erre = entitys.Error{
	Message: "_message",
	Type: "_type",
	StackTrace: "_stack_trace",
}
var trace = entitys.Trace{
	ID: "id_trace",
}
var event = entitys.Event{
	Action: "action",
	Category: []string{"category1", "category2"},
	Module: "module",
	Type: "type",
	Original: "original",
}
var tags = entitys.Tags{"tag1", "tag2", "tag3"}
var msg = entitys.NewMessage("msg").
WithError(erre).
WithEvent(event).
WithTrace(trace).
WithTags(tags)
```

### Audit
We can use a custom audit client when making custom logging instance
```go
var auditClientCustom audit.Client 
log, err := logger.NewLogger(logger.WithAuditClient(auditClientCustom))

```
if we use the default logger then we will use the default audit or we can specify it as follows
```go
var defaultClientAudit audit.Client = audit.Get()
log, err := logger.NewLogger(logger.WithAuditClient(defaultClientAudit))

```
Same as
```go
log := logger.Get()

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
