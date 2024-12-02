package config

import (
	"fmt"
	"github.com/MentalMentos/techTaskMetr.git/pkg/helpers"
	"github.com/MentalMentos/techTaskMetr.git/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"sync"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

var (
	config Config
	once   sync.Once
)

// New returns Config struct with env variables
func New(logger logger.Logger) *Config {
	once.Do(func() {
		config = Config{
			Host:     os.Getenv("PG_HOST"),
			Port:     os.Getenv("PG_PORT"),
			Username: os.Getenv("PG_USER"),
			Password: os.Getenv("PG_PASSWORD"),
			DBName:   os.Getenv("PG_DATABASE_NAME"),
		}
	})
	logger.Info("Config", "Config init")
	return &config
}

func DatabaseConnection(config *Config, logger logger.Logger) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.Username, config.Password, config.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal(helpers.PgPrefix, helpers.PgConnectFailed)
	}
	logger.Info(helpers.PgPrefix, "Database connection done")
	return db
}
