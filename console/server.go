package console

import (
	"go-boilerplate/server"

	"github.com/spf13/cobra"
)

var serve = &cobra.Command{
	Use:   "serve",
	Short: "run server",
	Long:  `This subcommand start the server`,
	Run:   run,
}

func init() {
	Root.AddCommand(serve)
}

func run(_ *cobra.Command, _ []string) {
	server.New().Listen()
}
