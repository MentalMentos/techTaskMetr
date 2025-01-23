package app

import (
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/internal/config"
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/internal/controller"
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/internal/model"
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/internal/repository"
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/internal/routes"
	"github.com/MentalMentos/ginWeb-Tonik/ginWeb/internal/service"
	zaplogger "github.com/MentalMentos/ginWeb-Tonik/ginWeb/pkg/logger/zap"
	_ "github.com/gin-gonic/gin"
	_ "github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func runApp() {
	if err := godotenv.Load("../../.env"); err != nil {
		panic(".env file not found")
	}
	log := zaplogger.New()
	db := config.DatabaseConnection(log)
	//validate := validator.New()
	db.Table("users").AutoMigrate(&model.User{})

	authRepository := repository.NewRepository(db)
	authService := service.New(authRepository, log)
	authController := controller.NewAuthController(authService, log)

	router := routes.SetupRouter(authController)
	if err := router.Run(":8881"); err != nil {
		log.Fatal("Main", "Failed to start server")
	}
}
