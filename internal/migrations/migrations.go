package migrations

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/viper"
	"os"
)

func RunDbMigrations(db *sql.DB) {
	migrationsPath := viper.GetString("database.migrations")
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		fmt.Println("Error while initing PG driver for migrations", err)
		os.Exit(1)
	}

	m, err := migrate.NewWithDatabaseInstance(migrationsPath, "postgres", driver)
	if err != nil {
		fmt.Println("Error while creating migrate instance", err)
		os.Exit(1)
	}

	err = m.Up()
	if err != nil {
		fmt.Println("Error while running migrations", err)
	}

	fmt.Println("Migrations ran")
}
