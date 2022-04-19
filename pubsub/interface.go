package pubsub

type Publisher interface {
	Publish(topic string, payload interface{}) error
}
