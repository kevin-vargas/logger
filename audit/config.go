package audit

import (
	"github.com/kevin-vargas/logger/config"
	"github.com/kevin-vargas/logger/pubsub"
)

func New(c *config.Audit) Client {
	return &client{
		publisher: pubsub.NewRestPublisher(toRestPubSubConfig(c)),
	}
}

func toRestPubSubConfig(config *config.Audit) *pubsub.ConfigRestPublisher {
	return &pubsub.ConfigRestPublisher{
		Username: config.Username,
		Password: config.Password,
		URL:      config.URL,
		Retries:  RETRIES,
	}
}
