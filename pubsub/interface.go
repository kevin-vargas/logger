package pubsub

type Pubisher interface {
	Publish(topic string, payload interface{}) error
}
