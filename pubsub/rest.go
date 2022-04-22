package pubsub

import (
	"errors"

	"github.com/go-resty/resty/v2"
)

type ConfigRestPublisher struct {
	URL      string
	Username string
	Password string
	Retries  int
}
type Option func(r *RestPublisher)
type FallBackMethod func(topic string, payload interface{})

var defaultFallBackMethod FallBackMethod = func(topic string, payload interface{}) {
}

type RestPublisher struct {
	fallback FallBackMethod
	client   *resty.Client
}

func withBaseConfig(client *resty.Client, cfg *ConfigRestPublisher) *resty.Client {
	return client.
		SetBaseURL(cfg.URL).
		SetBasicAuth(cfg.Username, cfg.Password).
		SetRetryCount(cfg.Retries).
		AddRetryAfterErrorCondition()
}
func WithFallBackMethod(method FallBackMethod) Option {
	return func(r *RestPublisher) {
		r.fallback = method
	}
}
func NewRestPublisher(cfg *ConfigRestPublisher, options ...Option) *RestPublisher {
	client := withBaseConfig(resty.New(), cfg)
	publisher := &RestPublisher{
		fallback: defaultFallBackMethod,
		client:   client,
	}
	for _, option := range options {
		option(publisher)
	}
	return publisher
}

func (rp *RestPublisher) Publish(topic string, payload interface{}) error {
	message := &Message{
		Topic:   topic,
		Payload: payload,
	}
	res, err := rp.client.R().SetBody(message).Post("/publish")

	if res.StatusCode()/200 != 1 || err != nil {
		if rp.fallback != nil {
			rp.fallback(topic, payload)
		}
		if err == nil {
			return errors.New("Invalid status code")
		}
		return err
	}
	return nil
}
