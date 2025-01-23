package authcontroller

import (
	"github.com/MentalMentos/techTaskMetr/techTaskmetr/auth/internal/data/request"
	"github.com/MentalMentos/techTaskMetr/techTaskmetr/auth/internal/service"
	"github.com/MentalMentos/techTaskMetr/techTaskmetr/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	authService service.Service
	logger      logger.Logger
}

func NewAuthController(authService *service.Service, logger logger.Logger) *AuthController {
	return &AuthController{
		authService: *authService,
		logger:      logger,
	}
}

func (controller *AuthController) Register(c *gin.Context) {
	var userRequest request.RegisterUserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		HandleError(c, &ApiError{Code: http.StatusBadRequest, Message: "Invalid request payload"})
		controller.logger.Fatal("[ CONTROLLER_REGISTER ]", "cannot bind json")
		return
	}

	userRequest.IP = GetClientIP(c)
	authResp, err := controller.authService.Register(c, userRequest)
	if err != nil {
		HandleError(c, err)
		return
	}
	controller.logger.Info("[ CONTROLLER_REGISTER ]", "register successful")
	JsonResponse(c, http.StatusOK, "Registration successful", authResp)
}

func (controller *AuthController) Login(c *gin.Context) {
	var userRequest request.LoginRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		HandleError(c, &ApiError{Code: http.StatusBadRequest, Message: "Invalid request payload"})
		controller.logger.Fatal("[ CONTROLLER_LOGIN ]", "cannot bind json")
		return
	}

	userRequest.IP = GetClientIP(c)
	authResp, err := controller.authService.Login(c, userRequest)
	if err != nil {
		HandleError(c, err)
		return
	}
	controller.logger.Info("[ CONTROLLER_LOGIN ]", "login successful")
	JsonResponse(c, http.StatusOK, "Auth successful", authResp)
}

func (controller *AuthController) UpdatePassword(c *gin.Context) {
	var userRequest request.UpdateUserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		HandleError(c, &ApiError{Code: http.StatusBadRequest, Message: "Invalid request payload"})
		controller.logger.Fatal("[ CONTROLLER_UPDATE ]", "cannot bind json")
		return
	}

	userRequest.IP = GetClientIP(c)
	authResp, err := controller.authService.UpdatePassword(c, userRequest)
	if err != nil {
		HandleError(c, err)
		return
	}
	controller.logger.Info("[ CONTROLLER_UPDATE ]", "update successful")
	JsonResponse(c, http.StatusOK, "Password updated successful", authResp)
}

func (controller *AuthController) RefreshToken(c *gin.Context) {
	var userRequest request.UpdateTokenRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		HandleError(c, &ApiError{Code: http.StatusBadRequest, Message: "Invalid request payload"})
		controller.logger.Fatal("[ CONTROLLER_REFRESH_TOKEN ]", "cannot bind json")
		return
	}

	authResp, err := controller.authService.RefreshToken(c, userRequest.RefreshToken)
	if err != nil {
		HandleError(c, err)
		return
	}
	controller.logger.Info("[ CONTROLLER_REFRESH_TOKEN ]", "token refresh successful")
	JsonResponse(c, http.StatusOK, "Token refreshed successful", authResp)
}
