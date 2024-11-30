package main

import (
	"github.com/MentalMentos/techTaskMetr.git/internal/config"
	repo "github.com/MentalMentos/techTaskMetr.git/internal/repository"
	"github.com/MentalMentos/techTaskMetr.git/pkg/helpers"
	zaplogger "github.com/MentalMentos/techTaskMetr.git/pkg/logger/zap"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	_ "github.com/go-playground/validator/v10"
	"net/http"
)

func main() {

	router := gin.Default()
	//fc
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})

	// logger init
	myLogger := zaplogger.New()

	// init env config
	if err := config.DatabaseConnection(); err != nil {
		myLogger.Fatal("[ ENV ]", "failed to connect to database")
	}

	// init config
	config := config.New(myLogger)

	db := config.DatabaseConnection()

	// init database
	postgres, err := repo.NewRepository()
	if err != nil {
		myLogger.Fatal(helpers.PgPrefix, helpers.PgConnectFailed)
	}

	// Keep Alive Postgres
	go db.KeepAlivePostgres(postgres, myLogger)

	// init new app
	myApp := app.New(postgres, myLogger, config)

	// migrations
	if err = db.MigrationUp(config, myLogger); err != nil {
		myLogger.Fatal(helpers.PgPrefix, err.Error())
	}

	// start server
	err = myApp.Run(myLogger, postgres)
	if err != nil {
		myLogger.Fatal(helpers.AppPrefix, err.Error())
	}

}
