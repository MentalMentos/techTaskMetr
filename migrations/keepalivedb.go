package migrations

import (
	"fmt"
	"github.com/MentalMentos/techTaskMetr.git/internal/config"
	"github.com/MentalMentos/techTaskMetr.git/pkg/helpers"
	"github.com/MentalMentos/techTaskMetr.git/pkg/logger"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MigrationUp(config *config.Config, myLogger logger.Logger) error {
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.Username, config.Password, config.Host, config.Port, config.DBName,
	)
	m, err := migrate.New("file://migrations", dbURL)
	if err != nil {
		myLogger.Fatal(helpers.PgPrefix, fmt.Sprintf("Failed to initialize migration: %v", err))
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		myLogger.Fatal(helpers.PgPrefix, fmt.Sprintf("Migration failed: %v", err))
		return err
	}

	myLogger.Info(helpers.PgPrefix, "Migration applied successfully")
	return nil
}
