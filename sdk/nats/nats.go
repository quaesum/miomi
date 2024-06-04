package nats

import (
	"fmt"
	na "github.com/nats-io/nats.go"
	"madmax/services/mail/entity"
)

type Request struct {
	Code string
	To   string
	entity.CommonData
	Meta string
}
type Response struct {
	Code  int
	Error string
}

func NewClient(url string) (*na.Conn, error) {
	nc, err := na.Connect(url)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize NATS client: %w", err)
	}

	return nc, nil
}
