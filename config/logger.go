package config

type Audit struct {
	URL      string
	Username string
	Password string
}
type Logger struct {
	ApplicationName string
	ServiceName     string
	Environment     string
	PodName         string
	NodeName        string
	Audit           *Audit
}

func New(applicationName string, serviceName string, environment string) *Logger {
	return &Logger{
		ApplicationName: applicationName,
		ServiceName:     serviceName,
		Environment:     environment,
	}
}

func (c *Logger) WithAudit(url string, username string, password string) *Logger {
	auditConfig := &Audit{
		URL:      url,
		Username: username,
		Password: password,
	}
	c.Audit = auditConfig
	return c
}

func (c *Logger) WithEnvironment(podName string, nodeName string) *Logger {
	c.NodeName = nodeName
	c.PodName = podName
	return c
}
