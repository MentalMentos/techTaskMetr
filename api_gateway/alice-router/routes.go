package router_alice

import (
	"context"
	"github.com/MentalMentos/techTaskMetr/api_gateway/alice-router/handlers"
	pkg "github.com/MentalMentos/techTaskMetr/api_gateway/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter(c *context.Context) *gin.Engine {
	router := gin.Default()
	router.Use(pkg.CheckJWT())
	// Приветственное сообщение
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Привет, я Alice! Добро пожаловать в задачник!",
		})
	})

	// Группа маршрутов с middleware авторизации
	authRoutes := router.Group("/alice")
	authRoutes.Use(handlers.AuthMiddleware) // Подключаем middleware
	authRoutes.POST("/register", handlers.RegisterUserHandler)
	// Ручки задачника
	{
		authRoutes.POST("/create", handlers.CreateTaskHandler)
		authRoutes.POST("/done", handlers.DoneTaskHandler)
		authRoutes.POST("/update", handlers.UpdateTaskHandler)
		authRoutes.GET("/list", handlers.ListTasksHandler)
	}
	return router
}
