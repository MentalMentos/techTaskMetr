package routes

import (
	"github.com/MentalMentos/techTaskMetr.git/internal/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter(controller *controller.Controller) *gin.Engine {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome Home!")
	})

	authRoutes := router.Group("/tasks")
	{
		authRoutes.POST("/create", controller.Create)
		authRoutes.POST("/done", controller.Done)
		authRoutes.POST("/update", controller.Update)
		authRoutes.GET("/list", controller.List)
		authRoutes.POST("/create-with-metrics", controller.TaskWithMetrics)
		authRoutes.POST("/update-with-metrics", controller.TaskWithMetrics)
		authRoutes.POST("/delete-with-metrics", controller.TaskWithMetrics)
		authRoutes.GET("/list-with-metrics", controller.TaskWithMetrics)
	}

	router.GET("/metrics", controller.MetricsHandler)

	return router
}
