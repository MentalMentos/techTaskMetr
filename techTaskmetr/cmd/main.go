package main

import (
	"github.com/MentalMentos/techTaskMetr.git/internal/config"
	"github.com/MentalMentos/techTaskMetr.git/internal/controller"
	"github.com/MentalMentos/techTaskMetr.git/internal/models"
	"github.com/MentalMentos/techTaskMetr.git/internal/repository"
	"github.com/MentalMentos/techTaskMetr.git/internal/routes"
	"github.com/MentalMentos/techTaskMetr.git/internal/service"
	zaplogger "github.com/MentalMentos/techTaskMetr.git/pkg/logger/zap"
	_ "github.com/gin-gonic/gin"
	_ "github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	_ "net/http"
)

func main() {
	// Загрузка переменных окружения
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Логгер
	myLogger := zaplogger.New()

	// Инициализация базы данных
	db := config.DatabaseConnection(myLogger)
	db.Table("tasks").AutoMigrate(&models.Task{})

	// Инициализация зависимостей
	taskRepository := repository.New(db, myLogger)
	taskService := service.NewService(taskRepository, myLogger)
	taskController := controller.NewController(taskService, myLogger)

	// Маршруты
	router := routes.SetupRouter(taskController)

	// Запуск приложения
	if err := router.Run(":8882"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
