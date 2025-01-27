package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required"`       // Название задачи (обязательно)
	Description string `json:"description" binding:"required"` // Описание задачи (обязательно)
}

func CreateTaskHandler(c *gin.Context) {
	var req CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Ошибка привязки данных в CreateTaskHandler: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные"})
	}

	// Получаем токен и user_id из контекста
	token, exists := c.Get("access_token")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Токен не найден"})
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не найден"})
	}

	// Добавляем user_id к запросу
	reqWithUser := map[string]interface{}{
		"user_id":     userID,
		"title":       req.Title,
		"description": req.Description,
	}

	// Отправляем запрос в микросервис `todo`
	url := "http://localhost:8882/tasks/create"
	responseData, err := sendAuthorizedRequest("POST", url, token.(string), reqWithUser)
	if err != nil {
		log.Printf("Ошибка выполнения задачи: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Задача успешно создана!",
		"data":    responseData,
	})
}
