package controller

import (
	"github.com/MentalMentos/techTaskMetr.git/internal/data/request"
	"github.com/MentalMentos/techTaskMetr.git/internal/data/response"
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

func (controller *Controller) Create(c gin.Context, logger logger.Logger) {
	var Request request.CreateTaskRequest
	if err := c.ShouldBindJSON(&Request); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Invalid request payload",
			Data:   nil,
		})
		return
	}

	authResp, err := controller.Service.Create(c, Request, logger)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		Code:   http.StatusOK,
		Status: "successful",
		Data:   authResp,
	})
}
