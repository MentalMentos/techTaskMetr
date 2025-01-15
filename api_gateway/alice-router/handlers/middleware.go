package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthMiddleware промежуточная функция для авторизации/регистрации
func AuthMiddleware(c *gin.Context) {
	var loginReq LoginRequest

	// Попытка авторизации
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные для авторизации"})
		c.Abort()
		return
	}

	authResp, err := TryLogin(loginReq)
	if err != nil {
		// Если авторизация не удалась, пробуем регистрировать
		var registerReq RegisterRequest
		registerReq.Email = loginReq.Email
		registerReq.Password = loginReq.Password

		authResp, err = TryRegister(registerReq)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Авторизация или регистрация не удались"})
			c.Abort()
			return
		}
	}

	// Сохраняем токены и UserID в контексте
	c.Set("access_token", authResp.AccessToken)
	c.Set("refresh_token", authResp.RefreshToken)
	c.Set("user_id", authResp.UserID)
	c.Next()
}
