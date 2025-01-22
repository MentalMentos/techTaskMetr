package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func RegisterUserHandler(c *gin.Context) {
	var registerReq RegisterRequest

	// Проверяем входящие данные
	if err := c.ShouldBindJSON(&registerReq); err != nil {
		log.Printf("Ошибка регистрации: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные для регистрации"})
		return
	}

	// Попытка регистрации
	authResp, err := TryRegister(registerReq)
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
