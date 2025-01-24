package main

import (
	"github.com/MentalMentos/techTaskMetr/auth/config"
	authcontroller "github.com/MentalMentos/techTaskMetr/auth/internal/controller"
	"github.com/MentalMentos/techTaskMetr/auth/internal/model"
	"github.com/MentalMentos/techTaskMetr/auth/internal/repository"
	"github.com/MentalMentos/techTaskMetr/auth/internal/service"
	"github.com/MentalMentos/techTaskMetr/auth/routes"
	zaplogger "github.com/MentalMentos/techTaskMetr/techTaskmetr/pkg/logger/zap"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	myLogger := zaplogger.New()
	if err := godotenv.Load("../.env"); err != nil {
		myLogger.Fatal("Error loading .env file", "")
	}
	// Инициализация базы данных
	db := config.DatabaseConnection(myLogger)
	//todo auth
	db.Table("users").AutoMigrate(&model.User{})

	authRepository := repository.NewRepository(db)
	authService := service.New(authRepository, myLogger)
	authController := authcontroller.NewAuthController(authService, myLogger)
	// Маршруты
	router := routes.SetupRouter(authController)

	if err := router.Run(":8881"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
