package audit

import (
	"github.com/kevin-vargas/logger/pubsub"
	"github.com/kevin-vargas/logger/strings"
)

type Client struct {
	defaultTopic string
	publisher    pubsub.Pubisher
}

// if we set a default topic it will override actual message topic
func (c *Client) Audit(message *Message) {
	topic := strings.OR(c.defaultTopic, message.Topic)
	c.publisher.Publish(topic, &message.Payload)
}
