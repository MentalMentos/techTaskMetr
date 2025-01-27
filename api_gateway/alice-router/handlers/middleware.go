package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	UserID       int    `json:"user_id"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// AuthMiddleware проверяет авторизацию или регистрирует нового пользователя.
func AuthMiddleware(c *gin.Context) {
	var loginReq LoginRequest

	// Проверяем входящие данные
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		log.Printf("Ошибка авторизации: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные для авторизации"})
		c.Abort()
		return
	}

	// Попытка авторизации
	authResp, err := TryLogin(loginReq)
	if err != nil {
		log.Printf("Ошибка авторизации: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверные данные для входа"})
		c.Abort()
		return
	}
	log.Println(authResp.AccessToken)
	log.Println(authResp.RefreshToken)
	log.Println(authResp.UserID)
	// Сохраняем токены и UserID в контексте
	c.Set("access_token", authResp.AccessToken)
	c.Set("refresh_token", authResp.RefreshToken)
	c.Set("user_id", authResp.UserID)
	c.Next()
}
