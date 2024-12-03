package main

import (
	"github.com/MentalMentos/techTaskMetr.git/internal/config"
	"github.com/MentalMentos/techTaskMetr.git/internal/controller"
	"github.com/MentalMentos/techTaskMetr.git/internal/models"
	"github.com/MentalMentos/techTaskMetr.git/internal/repository"
	"github.com/MentalMentos/techTaskMetr.git/internal/service"
	zaplogger "github.com/MentalMentos/techTaskMetr.git/pkg/logger/zap"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	_ "github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"net/http"
	_ "net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	myLogger := zaplogger.New()

	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome Home!")
	})

	db := config.DatabaseConnection(myLogger)
	db.Table("users").AutoMigrate(&models.Task{})

	taskRepository := repository.New(db, myLogger)
	taskService := service.NewService(taskRepository, myLogger)
	taskController := controller.NewController(taskService, myLogger)

	authRoutes := router.Group("/tasks")
	{
		authRoutes.POST("/create", taskController.Create)
	}

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Main", "Failed to start server")
	}
}
