package http

import (
	"github.com/MentalMentos/techTaskMetr.git/internal/controller"
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

func LoginUserHandler(c *gin.Context) {
	var loginReq LoginRequest

	// Проверяем входящие данные
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		log.Printf("Ошибка регистрации: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные для регистрации"})
		return
	}

	// Попытка регистрации
	authResp, err := controller.Cont
	if err != nil {
		log.Printf("Ошибка регистрации: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось зарегистрировать пользователя"})
		return
	}

	// Возвращаем успешный ответ с токенами
	c.JSON(http.StatusCreated, gin.H{
		"message":       "Пользователь успешно зарегистрирован!",
		"access_token":  authResp.AccessToken,
		"refresh_token": authResp.RefreshToken,
		"user_id":       authResp.UserID,
	})
}

func LoginUserHandler(c *gin.Context) {
	var loginReq LoginRequest

	// Проверяем входящие данные
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		log.Printf("Ошибка регистрации: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные для регистрации"})
		return
	}

	// Попытка регистрации
	authResp, err := TryLogin(*c, loginReq)
	if err != nil {
		log.Printf("Ошибка регистрации: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось зарегистрировать пользователя"})
		return
	}

	// Возвращаем успешный ответ с токенами
	c.JSON(http.StatusCreated, gin.H{
		"message":       "Пользователь успешно зарегистрирован!",
		"access_token":  authResp.AccessToken,
		"refresh_token": authResp.RefreshToken,
		"user_id":       authResp.UserID,
	})
}
