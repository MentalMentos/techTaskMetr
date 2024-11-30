package main

import (
	"github.com/MentalMentos/techTaskMetr.git/internal/config"
	"github.com/MentalMentos/techTaskMetr.git/internal/controller"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-playground/validator/v10"
	"net/http"
)

func main() {

	router := gin.Default()
	//fc
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})

	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.User{})

	// Repository
	Repository := repository.NewRepo(db)

	// Service
	service := service.New(Repository, validate)

	// Controller
	Controller := controller.NewAuthController(service)

	// TODO: routes

}
