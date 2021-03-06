package audit

type Payload struct {
	Nup            string      `json:"nup"`
	CorrelationId  string      `json:"correlationId"`
	SessionId      string      `json:"sessionId"`
	Type           Type        `json:"type"`
	AdditionalData interface{} `json:"additionalData"`
}

type Message struct {
	Topic   string
	Payload Payload
}
