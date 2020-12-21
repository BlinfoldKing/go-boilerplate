package helper

import (
	"context"
	"fmt"
	"go-boilerplate/entity"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
)

// SendToMultipleDevices sends the notifications to all device tokens listed
func SendToMultipleDevices(app *firebase.App, deviceTokens []string, data entity.Notification) error {
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		return err
	}
	notification := &messaging.MulticastMessage{
		Data: map[string]string{
			"title":    data.Title,
			"subtitle": data.Subtitle,
			"url_link": data.URLLink,
			"body":     data.Body,
		},
		Tokens: deviceTokens,
	}
	response, err := client.SendMulticast(ctx, notification)
	if err != nil {
		return err
	}
	fmt.Println("Successfully sent notification:", response)
	return nil
}

// SendToTopic sends the notification to all users that subscirbed to the topic
func SendToTopic(app *firebase.App, topic string, data entity.Notification) error {
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		return err
	}
	notification := &messaging.Message{
		Data: map[string]string{
			"title":    data.Title,
			"subtitle": data.Subtitle,
			"url_link": data.URLLink,
			"body":     data.Body,
		},
		Topic: topic,
	}
	response, err := client.Send(ctx, notification)
	if err != nil {
		return err
	}
	fmt.Println("Successfully sent notification:", response)
	return nil
}

// SubscribeToTopic subscribes all devices with the tokens to the topic
func SubscribeToTopic(app *firebase.App, topic string, deviceTokens []string) error {
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		return err
	}

	response, err := client.SubscribeToTopic(ctx, deviceTokens, topic)
	if err != nil {
		return err
	}
	fmt.Println(response.SuccessCount, "token were subscribed successfully")
	return nil
}

// UnsubscribeFromTopic unsubscribes all devices with the tokens to the topic
func UnsubscribeFromTopic(app *firebase.App, topic string, deviceTokens []string) error {
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		return err
	}

	response, err := client.UnsubscribeFromTopic(ctx, deviceTokens, topic)
	if err != nil {
		return err
	}
	fmt.Println(response.SuccessCount, "token were unsubscribed successfully")
	return nil
}
