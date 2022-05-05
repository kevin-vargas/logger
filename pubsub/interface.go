package pubsub

type FallBackMethod = func(topic string, payload interface{})

type Publisher interface {
	Publish(topic string, payload interface{}, fallbacks ...func(topic string, payload interface{})) error
}
