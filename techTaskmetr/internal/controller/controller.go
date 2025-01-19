package controller

import (
	"github.com/MentalMentos/techTaskMetr.git/internal/data/request"
	"github.com/MentalMentos/techTaskMetr.git/internal/service"
	"github.com/MentalMentos/techTaskMetr.git/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

var (
	tasksCreatedCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "tasks_created_total",
		Help: "Total number of tasks created",
	})

	taskCreationDurationHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "task_creation_duration_seconds",
		Help:    "Histogram of task creation duration in seconds",
		Buckets: prometheus.DefBuckets,
	})
)

func init() {
	// Регистрируем метрики
	prometheus.MustRegister(tasksCreatedCounter)
	prometheus.MustRegister(taskCreationDurationHistogram)
}

type Controller struct {
	Service service.Service
	logger  logger.Logger
}

func NewController(Service *service.Service, logger logger.Logger) *Controller {
	return &Controller{
		Service: *Service,
		logger:  logger,
	}
}

// MetricsHandler возвращает метрики в формате Prometheus
func (controller *Controller) MetricsHandler(c *gin.Context) {
	promHandler := promhttp.Handler()
	promHandler.ServeHTTP(c.Writer, c.Request)
}

// TaskWithMetrics обрабатывает запросы на создание задач с метриками
func (controller *Controller) TaskWithMetrics(c *gin.Context) {
	start := time.Now()
	var taskRequest request.CreateTaskRequest
	if err := c.ShouldBindJSON(&taskRequest); err != nil {
		HandleError(c, &ApiError{Code: http.StatusBadRequest, Message: "Invalid request payload"})
		return
	}

	taskResp, err := controller.Service.Create(c, taskRequest)
	if err != nil {
		HandleError(c, &ApiError{Code: http.StatusBadRequest, Message: "Invalid request payload"})
		return
	}

	// Обновляем метрики
	tasksCreatedCounter.Inc()
	taskCreationDurationHistogram.Observe(time.Since(start).Seconds())

	JsonResponse(c, http.StatusCreated, "Task created with metrics", taskResp)
}

func (controller *Controller) Create(c *gin.Context) {
	var taskRequest request.CreateTaskRequest
	if err := c.ShouldBindJSON(&taskRequest); err != nil {
		HandleError(c, &ApiError{Code: http.StatusBadRequest, Message: "Invalid request payload"})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return
	}

	taskRequest.User_id = int64(userID.(int))
	taskResp, err := controller.Service.Create(c, taskRequest)
	if err != nil {
		HandleError(c, err)
		return
	}

	JsonResponse(c, http.StatusOK, "Tasks created successful", taskResp)
}

func (controller *Controller) Update(c *gin.Context) {
	var taskRequest request.UpdateTaskRequest
	if err := c.ShouldBindJSON(&taskRequest); err != nil {
		HandleError(c, &ApiError{Code: http.StatusBadRequest, Message: "Invalid request payload"})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return
	}
	taskRequest.User_id = int64(userID.(int))
	taskResp, err := controller.Service.Update(c, taskRequest)
	if err != nil {
		HandleError(c, err)
		return
	}

	JsonResponse(c, http.StatusOK, "Task updated successfully", taskResp)
}

func (controller *Controller) Done(c *gin.Context) {
	var taskRequest request.DeleteTaskRequest
	if err := c.ShouldBindJSON(&taskRequest); err != nil {
		HandleError(c, &ApiError{Code: http.StatusBadRequest, Message: "Invalid request payload"})
		return
	}

	taskResp, err := controller.Service.Done(c, taskRequest)
	if err != nil {
		HandleError(c, err)
		return
	}

	JsonResponse(c, http.StatusOK, "Tasks created successful", taskResp)
}

func (controller *Controller) List(c *gin.Context) {
	var taskRequest request.GetTaskRequest
	if err := c.ShouldBindQuery(&taskRequest); err != nil {
		HandleError(c, &ApiError{Code: http.StatusBadRequest, Message: "Invalid request payload"})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return
	}
	taskRequest.User_id = int64(userID.(int))

	taskResp, err := controller.Service.List(c, taskRequest.User_id)
	if err != nil {
		HandleError(c, err)
		return
	}

	JsonResponse(c, http.StatusOK, "Tasks list successful", taskResp)
}
