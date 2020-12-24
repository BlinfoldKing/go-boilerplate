package ping

import (
	"go-boilerplate/adapters"
	"go-boilerplate/helper"
)

const topic = "ping"

// Message message for ping
type Message struct {
	Name string `json:"name"`
}

// PublishToQueue publish to ping queue
var PublishToQueue func(msg Message) error

// Queue init ping queue
func Queue(adapters adapters.Adapters) {
	push := adapters.Nats.NewQueue(
		topic,
		func(data interface{}) {
			msg := data.(*Message)
			helper.Logger.WithField("hello", msg).Info("PINGQUEUE")
		},
		&Message{},
	)

	PublishToQueue = func(msg Message) error {
		return push(msg)
	}
}
