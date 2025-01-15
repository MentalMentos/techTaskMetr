package alice_router

import (
	"github.com/gin-gonic/gin"
	"github.com/your-username/your-repo/handler"
	"net/http"
)

// SetupRouter настройка маршрутов
func SetupRouter(router *gin.Engine, controller *handler.Controller) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Привет! Добро пожаловать в задачник!",
		})
	})

	authRoutes := router.Group("/tasks")
	{
		authRoutes.POST("/create", controller.CreateTaskHandler)
		authRoutes.POST("/done", controller.DoneTaskHandler)
		authRoutes.POST("/update", controller.UpdateTaskHandler)
		authRoutes.GET("/list", controller.ListTasksHandler)
	}
}
