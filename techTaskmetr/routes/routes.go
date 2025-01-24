package routes

import (
	"github.com/MentalMentos/techTaskMetr/techTaskmetr/internal/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter(controller *controller.Controller) *gin.Engine {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome Home!")
	})

	taskRoutes := router.Group("/tasks")
	{
		taskRoutes.POST("/create", controller.Create)
		taskRoutes.POST("/done", controller.Done)
		taskRoutes.POST("/update", controller.Update)
		taskRoutes.GET("/list", controller.List)
		taskRoutes.POST("/create-with-metrics", controller.TaskWithMetrics)
		taskRoutes.POST("/update-with-metrics", controller.TaskWithMetrics)
		taskRoutes.POST("/done-with-metrics", controller.TaskWithMetrics)
		taskRoutes.GET("/list-with-metrics", controller.TaskWithMetrics)
	}

	router.GET("/metrics", controller.MetricsHandler)

	return router
}
