package migrations

import (
	"fmt"
	"github.com/MentalMentos/techTaskMetr.git/internal/config"
	"github.com/MentalMentos/techTaskMetr.git/pkg/helpers"
	"github.com/MentalMentos/techTaskMetr.git/pkg/logger"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"
	"time"
)

const (
	SongsTable          = "tasks"
	KeepAlivePollPeriod = 3 * time.Second
	MaxTries            = 20
)

func KeepAlivePostgres(database *gorm.DB, myLogger logger.Logger) {
	count := 0
	for {
		time.Sleep(KeepAlivePollPeriod)
		err := database.Ping()
		if err != nil {
			count++
			if count == MaxTries {
				myLogger.Fatal(helpers.PgPrefix, helpers.DisconnectDB)
			}
			myLogger.Info(helpers.PgPrefix, helpers.ReconnectDB)
		}
	}
}

func MigrationUp(config *config.Config, myLogger logger.Logger) error {
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=dissable",
		config.Username, config.Password, config.Host, config.Port, config.DBName)
	m, err := migrate.New(
		"file://internal/migrations/migrations",
		dbURL)
	if m == nil || err != nil {
		myLogger.Fatal(helpers.PgPrefix, helpers.PgMigrateFailed)
		return err
	}
	err = m.Up()
	if err != nil {
		myLogger.Fatal(helpers.PgPrefix, helpers.PgMigrateFailed)
		return fmt.Errorf(" %v", err)
	}

	return nil
}
