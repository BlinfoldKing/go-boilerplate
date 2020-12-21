package console

import (
	"go-boilerplate/adapters/firebase"
	"go-boilerplate/entity"
	"go-boilerplate/helper"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var sendToTokenCmd = &cobra.Command{
	Use:   "sendtotoken",
	Short: "send notification to device token",
	Long:  `This subcommand used to send notification to device token`,
	Run:   sendToToken,
}

var subscribeCmd = &cobra.Command{
	Use:   "subscribe",
	Short: "subscribe to topic",
	Long:  `This subcommand used to subscribe the device token to topic`,
	Run:   subscribe,
}

var sendToTopicCmd = &cobra.Command{
	Use:   "sendtotopic",
	Short: "send notification to topic",
	Long:  `This subcommand used to send notification to topic`,
	Run:   sendToTopic,
}

func init() {
	sendToTokenCmd.PersistentFlags().String("token", "token", "destination token")
	sendToTokenCmd.PersistentFlags().String("content", "test", "content of the notification")
	Root.AddCommand(sendToTokenCmd)

	subscribeCmd.PersistentFlags().String("topic", "test", "topic to be subscribed")
	subscribeCmd.PersistentFlags().String("token", "token", "device token that subscribes to topic")
	Root.AddCommand(subscribeCmd)

	sendToTopicCmd.PersistentFlags().String("topic", "test", "destination topic")
	sendToTopicCmd.PersistentFlags().String("content", "test", "content of the notification")
	Root.AddCommand(sendToTopicCmd)
}

func sendToToken(cmd *cobra.Command, args []string) {
	token := []string{cmd.Flag("token").Value.String()}
	data := entity.Notification{
		Body: cmd.Flag("content").Value.String(),
	}

	app, err := firebase.Init()
	if err != nil {
		logrus.Error(err)
	}
	err = helper.SendToMultipleDevices(app, token, data)
	if err != nil {
		logrus.Error(err)
	}
}

func subscribe(cmd *cobra.Command, args []string) {
	topic := cmd.Flag("topic").Value.String()
	token := []string{cmd.Flag("token").Value.String()}
	app, err := firebase.Init()
	if err != nil {
		logrus.Error(err)
	}
	err = helper.SubscribeToTopic(app, topic, token)
	if err != nil {
		logrus.Error(err)
	}
}

func sendToTopic(cmd *cobra.Command, args []string) {
	topic := cmd.Flag("topic").Value.String()
	data := entity.Notification{
		Body: cmd.Flag("content").Value.String(),
	}

	app, err := firebase.Init()
	if err != nil {
		logrus.Error(err)
	}
	err = helper.SendToTopic(app, topic, data)
	if err != nil {
		logrus.Error(err)
	}
}
