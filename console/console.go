package console

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Root = &cobra.Command{
	Use:   "cobra-example",
	Short: "An example of cobra",
	Long:  "use `help` to get started",
}

func Execute() {
	err := Root.Execute()
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}
