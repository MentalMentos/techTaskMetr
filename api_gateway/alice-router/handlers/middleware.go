package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	UserID       int64  `json:"user_id"`
}

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

// AuthMiddleware промежуточная функция для авторизации/регистрации
func AuthMiddleware(c *gin.Context) {
	var loginReq LoginRequest

	// Попытка авторизации
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		log.Fatal("ошибка аторизации в authmiddleware", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные для авторизации"})
		c.Abort()
		return
	}

	authResp, err := TryLogin(loginReq)
	if err != nil {
		log.Printf("cannot login, will try registr authmiddleware", err)
		// Если авторизация не удалась, пробуем регистрировать
		var registerReq RegisterRequest
		registerReq.Name = "пользователь"
		registerReq.Email = loginReq.Email
		registerReq.Password = loginReq.Password
		log.Printf(registerReq.Email, loginReq.Email, "добавились данные для регистрации authmiddleware")
		authResp, err = TryRegister(registerReq)
		if err != nil {
			log.Fatalf("cannot registr authmiddleware", err)
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
	log.Printf("токены сохранились")
}
