package notifications

import (
	"go-boilerplate/adapters"
	"go-boilerplate/helper"
	"go-boilerplate/modules/firebase"
	"go-boilerplate/modules/user_device"
)

const topic = "mail"

// Message message for ping
type Message struct {
	UserID   string `json:"user_id"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	URLLink  string `json:"url_link"`
	Body     string `json:"body"`
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

			device, err := deviceservice.GetByUserID(msg.UserID)
			if err != nil {

				helper.Logger.
					WithField("msg", msg).
					Error(err)
			}

			notif, err := service.CreateNotification(msg.UserID, msg.Title, msg.Subtitle, msg.URLLink, msg.Body)
			if err != nil {

				helper.Logger.
					WithField("msg", msg).
					Error(err)
			}

			err = firebase.SendToMultipleDevices(adapters.Firebase, []string{device.DeviceToken}, notif)
			if err != nil {
				helper.Logger.
					WithField("msg", msg).
					Error(err)
			}

		},
		&Message{},
	)

	PublishToQueue = func(msg Message) error {
		return push(msg)
	}
}
