package pubsub

type Message struct {
	Topic   string `json:"topic"`
	Payload interface{}
}
