package console

import (
	"fmt"
	"go-boilerplate/migration"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate database",
	Long:  `This subcommand used to migrate database`,
	Run:   processMigration,
}

func init() {
	migrateCmd.PersistentFlags().Int("step", 0, "maximum migration steps")
	migrateCmd.PersistentFlags().String("direction", "up", "migration direction")
	Root.AddCommand(migrateCmd)
}

func processMigration(cmd *cobra.Command, args []string) {
	direction := cmd.Flag("direction").Value.String()
	stepStr := cmd.Flag("step").Value.String()
	step, err := strconv.Atoi(stepStr)
	if err != nil {
		logrus.WithField("stepStr", stepStr).Fatal("Failed to parse step to int: ", err)
	}

	n, err := migration.Migrate(direction, step)
	if err != nil {
		logrus.Error(err)
	}

	logrus.Info(fmt.Sprintf("%d migration(s) done", n))
}
