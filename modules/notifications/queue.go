package notifications

import (
	"go-boilerplate/adapters"
	"go-boilerplate/helper"
	"go-boilerplate/modules/firebase"
	"go-boilerplate/modules/user_device"
)

const topic = "notifications"

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
			msg := data.(*Message)

			devices, err := deviceservice.GetByUserID(msg.UserID)
			if err != nil {
				helper.Logger.
					WithField("msg", msg).
					Error(err)
			}

			tokens := make([]string, 0)
			for _, device := range devices {
				tokens = append(tokens, device.DeviceToken)
			}

			tokens = []string{"cCr3j_BnRI-FfsSTf14J4r:APA91bHxFfPciKEZ3M-o8PKQHnLIkZhBwEsAeFTa8qJeqIlELAn8zqUy8LUajgkQLsqQcseLPos_oZnKq9VUskasybQOa-bOQ7bSKWjrAwOwZAJ8JTFdXhVr4lh2-EhY1ZUrzA5GybTY", "ceL8CJ4DTs6aFtY5GifMps:APA91bH4OQ9-9LVDoXkqd69R6nW-h3og-JKfoxRyVs0pj-SXZl3xmr1XHqdmd3rtfhzoXjrwOU2D8r-p_bnFrFfwBjeqTIw9MmGHsxZh29dDk-1KfnSdJQdqbwd3wT6Hle-4S0o6bow6"}

			notif, err := service.CreateNotification(msg.UserID, msg.Title, msg.Subtitle, msg.URLLink, msg.Body)
			if err != nil {

				helper.Logger.
					WithField("msg", msg).
					Error(err)
			}

			err = firebase.SendToMultipleDevices(adapters.Firebase, tokens, notif)
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
