package audit

import (
	"github.com/kevin-vargas/logger/pubsub"
)

type FallBackMethod func(topic string, payload *Payload)

func (f FallBackMethod) toPubslih() pubsub.FallBackMethod {
	return func(topic string, payload interface{}) {
		payloadAudit := payload.(*Payload)
		f(topic, payloadAudit)
	}
}

type Client interface {
	Audit(message *Message, fallbacks ...FallBackMethod) error
}
type client struct {
	defaultTopic string
	publisher    pubsub.Publisher
}

func (c *client) Audit(message *Message, fallbacks ...FallBackMethod) (err error) {
	topic := message.Topic
	fallbacksToPublish := make([]pubsub.FallBackMethod, len(fallbacks))
	for i, f := range fallbacks {
		fallbacksToPublish[i] = f.toPubslih()
	}
	errPublish := c.publisher.Publish(topic, &message.Payload, fallbacksToPublish...)
	if errPublish != nil {
		return errPublish
	}
	return
}
