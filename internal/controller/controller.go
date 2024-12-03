package controller

import (
	"github.com/MentalMentos/techTaskMetr.git/internal/data/request"
	"github.com/MentalMentos/techTaskMetr.git/internal/service"
	"github.com/MentalMentos/techTaskMetr.git/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

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

func (controller *Controller) Create(c *gin.Context) {
	var taskRequest request.CreateTaskRequest
	if err := c.ShouldBindJSON(&taskRequest); err != nil {
		HandleError(c, &ApiError{Code: http.StatusBadRequest, Message: "Invalid request payload"})
		return
	}

	taskResp, err := controller.Service.Create(c, taskRequest, controller.logger)
	if err != nil {
		HandleError(c, err)
		return
	}

	JsonResponse(c, http.StatusOK, "Registration successful", taskResp)
}
