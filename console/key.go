package console

import (
	"go-boilerplate/helper"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var keyCmd = &cobra.Command{
	Use:   "key",
	Short: "generate key",
	Long:  `This subcommand used to generate key`,
	Run:   generateKey,
}

func init() {
	Root.AddCommand(keyCmd)
}

func generateKey(cmd *cobra.Command, args []string) {
	err := helper.GenerateRSAKeyPair(".keys")
	if err != nil {
		logrus.Error(err)
	}
}
