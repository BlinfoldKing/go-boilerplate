package nats

import (
	"encoding/json"
	"go-boilerplate/config"
	"go-boilerplate/helper"

	"github.com/nats-io/stan.go"
)

// Nats nats client
type Nats struct {
	stan.Conn
}

// Init create new nats instance
func Init() (*Nats, error) {
	nc, err := stan.Connect(
		config.NATSCLUSTERID(),
		config.NATSCLIENTID(),
		stan.NatsURL(config.NATSURI()),
	)

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
	queue.Conn.Subscribe(topic, func(m *stan.Msg) {
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
