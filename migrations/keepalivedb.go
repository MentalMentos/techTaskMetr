package migrations

import (
	"github.com/MentalMentos/techTaskMetr.git/pkg/helpers"
	"github.com/MentalMentos/techTaskMetr.git/pkg/logger"
	"gorm.io/gorm"
	"time"
)

const (
	// SongsTable - table songs name
	SongsTable = "tasks"

	// KeepAlivePollPeriod - period for pinging db
	KeepAlivePollPeriod = 3 * time.Second

	// MaxTries - max tries to connect to db
	MaxTries = 20
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
