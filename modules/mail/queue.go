package mail

import (
	"go-boilerplate/adapters"
	"go-boilerplate/helper"
)

const topic = "ping"

// Message message for ping
type Message struct {
	Sender    string `json:"sender"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
	Recipient string `json:"recipient"`
}

// PublishToQueue publish to ping queue
var PublishToQueue func(msg Message) error

// Queue init ping queue
func Queue(adapters adapters.Adapters) {
	var service Service
	// ini service here

	push := adapters.Nats.NewQueue(
		topic,
		func(data interface{}) {
			msg := data.(*Message)
			_, err := service.SendEmail(msg.Sender, msg.Subject, msg.Body, msg.Recipient)
			if err != nil {
				helper.Logger.Error(err)
			}
		},
		&Message{},
	)

	PublishToQueue = func(msg Message) error {
		return push(msg)
	}
}
