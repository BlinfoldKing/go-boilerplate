package nats

import (
	"encoding/json"
	"go-boilerplate/config"
	"go-boilerplate/helper"

	"github.com/nats-io/nats.go"
)

// Nats nats client
type Nats struct {
	*nats.Conn
}

// Init create new nats instance
func Init() (*Nats, error) {
	nc, err := nats.Connect(config.NATSURI())
	if err != nil {
		return nil, err
	}

	return &Nats{nc}, nil
}

// Push alias for publish function
type Push = func(T interface{}) error

// Worker alias for worker function
type Worker = func(T interface{})

// NewQueue create new queue return publisher
func (queue *Nats) NewQueue(topic string, worker Worker, messageType interface{}) Push {
	queue.Conn.Subscribe(topic, func(m *nats.Msg) {
		err := json.Unmarshal(m.Data, &messageType)
		if err != nil {
			helper.Logger.Error(err)
		}

		worker(messageType)
	})

	return func(T interface{}) error {
		data, err := json.Marshal(T)
		if err != nil {
			return err
		}
		return queue.Conn.Publish(topic, data)
	}
}
