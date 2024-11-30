package main

import (
	"github.com/MentalMentos/techTaskMetr.git/internal/config"
	repo "github.com/MentalMentos/techTaskMetr.git/internal/repository"
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

	db := config.DatabaseConnection()

	postgres := repo.NewRepository(db)

	// Keep Alive Postgres
	//go db.KeepAlivePostgres(postgres, myLogger)

}
