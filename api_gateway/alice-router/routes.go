package alice_router

import (
	"github.com/MentalMentos/api_gateway/alice-router/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SetupRouter настройка маршрутов
func SetupRouter(router *gin.Engine) {
	// Приветственное сообщение
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Привет, я Alice! Добро пожаловать в задачник!",
		})
	})

	// Группа маршрутов с middleware авторизации
	authRoutes := router.Group("/alice")
	authRoutes.Use(handlers.AuthMiddleware) // Подключаем middleware

	// Ручки задачника
	{
		authRoutes.POST("/create", handlers.CreateTaskHandler)
		authRoutes.POST("/done", handlers.DoneTaskHandler)
		authRoutes.POST("/update", handlers.UpdateTaskHandler)
		authRoutes.GET("/list", handlers.ListTasksHandler)
	}
}
