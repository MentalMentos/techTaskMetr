package main

import (
	"github.com/MentalMentos/techTaskMetr.git/internal/config"
	"github.com/MentalMentos/techTaskMetr.git/internal/repository"
	"github.com/MentalMentos/techTaskMetr.git/internal/service"
	zaplogger "github.com/MentalMentos/techTaskMetr.git/pkg/logger/zap"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	_ "github.com/go-playground/validator/v10"
	"net/http"
)

func main() {
	router := gin.Default()

	myLogger := zaplogger.New()

	Config := config.New(myLogger)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})

	db := config.DatabaseConnection(myLogger)

	repo := repository.New(db, myLogger)

	service := service.NewService(repo, myLogger)

	go db.KeepAlivePostgres(postgres, myLogger)

}
