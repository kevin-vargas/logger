package entities

import (
	"github.com/kevin-vargas/logger/config"
)

func GetDefaultLabels(c *config.Logger) (labels Labels) {
	labels = make(Labels)
	labels[fieldLabelApplication] = c.ApplicationName
	labels[fieldLabelService] = c.ServiceName
	labels[fieldLabelEnvironment] = c.Environment
	labels[fieldLabelLibVersion] = LIB_VERSION
	labels[fieldLabelLibLanguage] = LIB_LANGUAGE
	labels[fieldLabelPodName] = c.PodName
	labels[fieldLabelNodeName] = c.NodeName
	return
}
