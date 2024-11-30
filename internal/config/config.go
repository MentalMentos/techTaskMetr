package config

import (
	"os"
	"sync"
)

type Config struct {
	ServerPort string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

var (
	config Config
	once   sync.Once
)

// New returns Config struct with env variables
func New(logger logger.Logger) *Config {
	once.Do(func() {
		config = Config{
			ServerPort: os.Getenv("server_port"),
			DBHost:     os.Getenv("DB_HOST"),
			DBPort:     os.Getenv("DB_PORT"),
			DBUser:     os.Getenv("DB_USER"),
			DBPassword: os.Getenv("DB_PASS"),
			DBName:     os.Getenv("DB_NAME"),
		}
	})
	logger.Info("Config", "Config init")
	return &config
}
