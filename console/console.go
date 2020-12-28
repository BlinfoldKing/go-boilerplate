package console

import (
	"os"

	"github.com/spf13/cobra"
	"go-boilerplate/helper"
)

// Root base root cli
var Root = &cobra.Command{
	Use:   "cobra-example",
	Short: "An example of cobra",
	Long:  "use `help` to get started",
}

// Execute run Root
func Execute() {
	err := Root.Execute()
	if err != nil {
		helper.Logger.Error(err)
		os.Exit(1)
	}
}
