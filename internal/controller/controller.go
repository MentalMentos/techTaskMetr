package controller

import (
	"github.com/MentalMentos/techTaskMetr.git/internal/data/request"
	"github.com/MentalMentos/techTaskMetr.git/internal/data/response"
	"github.com/MentalMentos/techTaskMetr.git/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	Service service.Service
}

func NewAuthController(Service *service.Service) *Controller {
	return &Controller{
		Service: *Service,
	}
}

// Register контроллер для регистрации пользователей
func (controller *Controller) Create(c *gin.Context) {
	var Request request.CreateTaskRequest
	if err := c.ShouldBindJSON(&Request); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Invalid request payload",
			Data:   nil,
		})
		return
	}

	authResp, err := controller.Service.Create(c, Request)
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
		Status: "Registration successful",
		Data:   authResp,
	})
}
