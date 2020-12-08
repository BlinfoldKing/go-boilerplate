package migration

import (
	"database/sql"
	"go-boilerplate/config"

	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

func Migrate(direction string, step int) (int, error) {
	migrations := &migrate.FileMigrationSource{
		Dir: "./migration",
	}

	migrate.SetTable("schema_migrations")

	db, err := sql.Open("postgres", config.DB_URL())
	if err != nil {
		return 0, err
	}

	if direction == "down" {
		return migrate.ExecMax(db, "postgres", migrations, migrate.Down, step)
	} else {
		return migrate.ExecMax(db, "postgres", migrations, migrate.Up, step)
	}

}
