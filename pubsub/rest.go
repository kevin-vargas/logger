package pubsub

import (
	"errors"
	"time"

	"github.com/go-resty/resty/v2"
)

type ConfigRestPublisher struct {
	URL      string
	Username string
	Password string
	Retries  int
}

type Option func(r *RestPublisher)

type RestPublisher struct {
	client *resty.Client
}

func withBaseConfig(client *resty.Client, cfg *ConfigRestPublisher) *resty.Client {
	return client.
		SetBaseURL(cfg.URL).
		SetBasicAuth(cfg.Username, cfg.Password).
		SetRetryCount(cfg.Retries).
		SetDisableWarn(true).
		SetLogger(discardLog()).
		SetTimeout(timeout_request * time.Millisecond)
}

func NewRestPublisher(cfg *ConfigRestPublisher, options ...Option) *RestPublisher {
	client := withBaseConfig(resty.New(), cfg)
	publisher := &RestPublisher{
		client: client,
	}
	for _, option := range options {
		option(publisher)
	}
	return publisher
}

func (rp *RestPublisher) Publish(topic string, payload interface{}, fallbacks ...FallBackMethod) error {
	message := &Message{
		Topic:   topic,
		Payload: payload,
	}
	res, err := rp.client.R().SetBody(message).Post("/publish")

	if res.StatusCode()/200 != 1 || err != nil {
		for _, fallback := range fallbacks {
			fallback(topic, payload)
		}
		if err == nil {
			return errors.New("Invalid status code")
		}
		return err
	}
	return nil
}
