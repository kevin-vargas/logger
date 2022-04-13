package entitys

import (
	"os"
	"sync"
)

var instance Labels
var once sync.Once

func GetDefaultLabels() Labels {
	once.Do(func() {
		labels := make(Labels)
		labels[fieldLabelApplication] = or(os.Getenv(ENV_APPLICATION_NAME), DEFAULT_APPLICATION)
		labels[fieldLabelService] = or(os.Getenv(ENV_LOGGINSERVICE), DEFAULT_SERVICE)
		labels[fieldLabelEnvironment] = or(os.Getenv(ENV_ENVIRONMENT), DEFAULT_ENVIROMENT)
		labels[fieldLabelLibVersion] = LIB_VERSION
		labels[fieldLabelLibLanguage] = LIB_LANGUAGE
		labels[fieldLabelPodName] = or(os.Getenv(ENV_POD), DEFAULT_POD)
		labels[fieldLabelNodeName] = or(os.Getenv(ENV_NODE_NAME), DEFAULT_NODE)
		instance = labels
	})
	return instance
}

func or(str string, defaultStr string) string {
	if str == "" {
		return defaultStr
	}
	return str
}
