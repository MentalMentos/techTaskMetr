package main

import (
	"github.com/MentalMentos/techTaskMetr.git/internal/config"
	"github.com/MentalMentos/techTaskMetr.git/internal/controller"
	"github.com/MentalMentos/techTaskMetr.git/internal/repository"
	"github.com/MentalMentos/techTaskMetr.git/internal/service"
	"github.com/MentalMentos/techTaskMetr.git/migrations"
	zaplogger "github.com/MentalMentos/techTaskMetr.git/pkg/logger/zap"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	_ "github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
	_ "net/http"
)

func main() {
	myLogger := zaplogger.New()

	cfg := config.New(myLogger)

	db := config.DatabaseConnection(*cfg, myLogger)

	if err := migrations.MigrationUp(cfg, myLogger); err != nil {
		log.Fatal("Main", "Migration failed: ", err)
	}

	repo := repository.New(db, myLogger)
	svc := service.NewService(repo, myLogger)
	ctrl := controller.NewController(svc, myLogger)

	r := gin.Default()
	r.POST("/tasks", func(c *gin.Context) { ctrl.Create(*c, myLogger) })

	log.Info("Main", "Starting server on port 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Main", "Failed to start server")
	}

}
