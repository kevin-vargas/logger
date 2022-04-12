package entitys

import (
	"os"
)

type DefaultLabels struct {
	Application string
	Service     string
	Environment string
	LibVersion  string
	LibLanguage string
	PodName     string
	NodeName    string
}

func GetDefaultLabels() (labels Labels) {
	labels = make(map[string]string)
	labels[fieldLabelApplication] = or(os.Getenv(ENV_APPLICATION_NAME), DEFAULT_APPLICATION)
	labels[fieldLabelService] = or(os.Getenv(ENV_LOGGINSERVICE), DEFAULT_SERVICE)
	labels[fieldLabelEnvironment] = or(os.Getenv(ENV_ENVIRONMENT), DEFAULT_ENVIROMENT)
	labels[fieldLabelLibVersion] = LIB_VERSION
	labels[fieldLabelLibLanguage] = LIB_LANGUAGE
	labels[fieldLabelPodName] = or(os.Getenv(ENV_POD), DEFAULT_POD)
	labels[fieldLabelNodeName] = or(os.Getenv(ENV_NODE_NAME), DEFAULT_NODE)
	return
}

func or(str string, defaultStr string) string {
	if str == "" {
		return defaultStr
	}
	return str
}
