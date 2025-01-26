package router

import (
	"github.com/MentalMentos/techTaskMetr/api_gateway/alice-router/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter(router *gin.Engine) *gin.Engine {
	router := gin.Default()
	// Приветственное сообщение
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Привет, я Alice! Добро пожаловать в задачник!",
		})
	})

	// Группа маршрутов с middleware авторизации
	authRoutes := router.Group("/alice")
	authRoutes.Use(handlers.AuthMiddleware) // Подключаем middleware
	// Ручка для регистрации
	router.POST("/register", handlers.RegisterUserHandler)
	router.POST("/login", handlers.LoginUserHandler)
	// Ручки задачника
	{
		authRoutes.POST("/create", handlers.CreateTaskHandler)
		authRoutes.POST("/done", handlers.DoneTaskHandler)
		authRoutes.POST("/update", handlers.UpdateTaskHandler)
		authRoutes.GET("/list", handlers.ListTasksHandler)
	}
	return router
}
