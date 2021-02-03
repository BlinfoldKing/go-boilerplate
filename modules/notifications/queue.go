package notifications

import (
	"go-boilerplate/adapters"
	"go-boilerplate/entity"
	"go-boilerplate/helper"
	"go-boilerplate/modules/firebase"
	"go-boilerplate/modules/user_device"
)

const topic = "mail"

// Message message for ping
type Message struct {
	Topic   string              `json:"topic"`
	Content entity.Notification `json:"content"`
}

// PublishToQueue publish to ping queue
var PublishToQueue func(msg Message) error

// Queue init ping queue
func Queue(adapters adapters.Adapters) {
	repo := CreatePosgresRepository(adapters.Postgres)
	service := CreateService(repo)

	devicerepo := userdevice.CreatePosgresRepository(adapters.Postgres)
	deviceservice := userdevice.CreateService(devicerepo)
	push := adapters.Nats.NewQueue(
		topic,
		func(data interface{}) {
			msg := data.(Message)

			device, err := deviceservice.GetByUserID(msg.Content.UserID)
			if err != nil {

				helper.Logger.
					WithField("msg", msg).
					Error(err)
			}
			err = firebase.SendToMultipleDevices(adapters.Firebase, []string{device.DeviceToken}, msg.Content)
			if err != nil {
				helper.Logger.
					WithField("msg", msg).
					Error(err)
			}

			_, err = service.CreateNotification(msg.Content.UserID, msg.Content.Title, "", "", "")
		},
		&Message{},
	)

	PublishToQueue = func(msg Message) error {
		return push(msg)
	}
}
