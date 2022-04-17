package pubsub

import "github.com/go-resty/resty/v2"

type ConfigRestPublisher struct {
	URL      string
	Username string
	Password string
	Retrys   int
}

type RestPublisher struct {
	client *resty.Client
}

func withBaseConfig(client *resty.Client, cfg *ConfigRestPublisher) *resty.Client {
	return client.
		SetBaseURL(cfg.URL).
		SetBasicAuth(cfg.Username, cfg.Password).
		SetRetryCount(cfg.Retrys)
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
	_, err := rp.client.R().SetBody(message).Post("/publish")
	if err != nil {
		return err
	}
	return nil
}
