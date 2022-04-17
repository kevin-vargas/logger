package audit

import (
	"sync"

	"github.com/kevin-vargas/logger/pubsub"
	"github.com/kevin-vargas/logger/variables"
)

var once sync.Once
var instance *Client

// TODO: check default config audit
func Get() *Client {
	once.Do(func() {
		instance = &Client{
			defaultTopic: variables.OR(env_topic, ""),
			publisher: pubsub.NewRestPublisher(&pubsub.ConfigRestPublisher{
				URL:      variables.ORPanic(env_url),
				Username: variables.ORPanic(env_user),
				Password: variables.ORPanic(env_password),
				Retrys:   RETRYS,
			}),
		}
	})
	return instance
}
