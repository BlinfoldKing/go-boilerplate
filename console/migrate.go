package console

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-boilerplate/helper"
	"go-boilerplate/migration"
	"io/ioutil"
	"strconv"
	"time"
)

var migrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "manage migration",
	Run:   processMigration,
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate database",
	Long:  `This subcommand used to migrate database`,
	Run:   processMigration,
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create migration",
	Long:  `create new migration file on migrations/`,
	Run:   createMigration,
}

func init() {
	migrateCmd.PersistentFlags().Int("step", 0, "maximum migration steps")
	migrateCmd.PersistentFlags().String("direction", "up", "migration direction, up to apply and down to redo")

	createCmd.PersistentFlags().String("label", "new_migration", "migration label")

	migrationCmd.PersistentFlags().Int("step", 0, "maximum migration steps")
	migrationCmd.PersistentFlags().String("direction", "up", "migration direction, up to apply and down to redo")

	migrationCmd.AddCommand(migrateCmd)
	migrationCmd.AddCommand(createCmd)

	Root.AddCommand(migrationCmd)
}

func processMigration(cmd *cobra.Command, args []string) {
	direction := cmd.Flag("direction").Value.String()
	stepStr := cmd.Flag("step").Value.String()
	step, err := strconv.Atoi(stepStr)
	if err != nil {
		helper.Logger.WithField("stepStr", stepStr).Fatal("Failed to parse step to int: ", err)
	}

	n, err := migration.Migrate(direction, step)
	if err != nil {
		helper.Logger.Error(err)
	}

	helper.Logger.Info(fmt.Sprintf("%d migration(s) done", n))
}

func createMigration(cmd *cobra.Command, args []string) {
	name := cmd.Flag("label").Value.String()
	now := time.Now()
	timestamp := helper.FormatDate(now)
	filename := fmt.Sprintf("migration/%s_%s.sql", timestamp, name)
	ioutil.WriteFile(filename,
		[]byte(
			`-- +migrate Up
CREATE TABLE IF NOT EXISTS "new_table" (
    "id" UUID NOT NULL PRIMARY KEY,
);

-- +migrate Down
DROP TABLE IF EXISTS "new_table";`,
		),
		0755)

	helper.Logger.Printf("%s created\n", filename)
}
