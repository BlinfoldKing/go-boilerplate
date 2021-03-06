package migration

import (
	"database/sql"
	migrate "github.com/rubenv/sql-migrate"
	"go-boilerplate/config"
)

// Migrate run migration files
// direction = up, will apply migration
// direction = down, will redo migration
// step represent how many files to apply/redo, apply/redo all if empty
func Migrate(direction string, step int) (int, error) {
	migrations := &migrate.FileMigrationSource{
		Dir: "./migration",
	}

	migrate.SetTable("schema_migrations")

	db, err := sql.Open("postgres", config.DBCONFIG())
	if err != nil {
		return 0, err
	}

	if direction == "down" {
		return migrate.ExecMax(db, "postgres", migrations, migrate.Down, step)
	}

	return migrate.ExecMax(db, "postgres", migrations, migrate.Up, step)

}
