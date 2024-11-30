package config

import (
	"fmt"
	"github.com/MentalMentos/techTaskMetr.git/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
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

const (
	host     = "localhost"
	port     = 5432
	user     = "user"
	password = "1234"
	dbName   = "postgres"
)

func DatabaseConnection() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("connect database error:%v", err)
	}

	return db
}
