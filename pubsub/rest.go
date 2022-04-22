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

type RestPublisher struct {
	client *resty.Client
}

func withBaseConfig(client *resty.Client, cfg *ConfigRestPublisher) *resty.Client {
	return client.
		SetBaseURL(cfg.URL).
		SetBasicAuth(cfg.Username, cfg.Password).
		SetRetryCount(cfg.Retries).
		AddRetryAfterErrorCondition()
	/*
		AddRetryCondition(
			func(r *resty.Response, err error) bool {
				return (r.StatusCode() / 200) != 1
			},
		)
	*/
}

func NewRestPublisher(cfg *ConfigRestPublisher) *RestPublisher {
	client := withBaseConfig(resty.New(), cfg)
	return &RestPublisher{
		client: client,
	}
}

func (rp *RestPublisher) Publish(topic string, payload interface{}) error {
	message := &Message{
		Topic:   topic,
		Payload: payload,
	}
	res, err := rp.client.R().SetBody(message).Post("/publish")
	if err != nil {
		return err
	}
	if res.StatusCode()/200 != 1 {
		return errors.New("Invalid status code")
	}
	return nil
}
