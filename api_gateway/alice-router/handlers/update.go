package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdateTaskRequest структура запроса для обновления задачи
type UpdateTaskRequest struct {
	ID          string `json:"id" binding:"required"`          // Уникальный идентификатор задачи
	Title       string `json:"title" binding:"required"`       // Название задачи
	Description string `json:"description" binding:"required"` // Описание задачи
	Status      string `json:"status" binding:"required"`      // Новый статус задачи
}

func UpdateTaskHandler(c *gin.Context) {
	var req UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Ошибка привязки данных в UpdateTaskHandler: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные"})
		return
	}

	// Получаем токен и user_id из контекста
	token, exists := c.Get("access_token")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Токен не найден"})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не найден"})
		return
	}

	// Формируем тело запроса с учетом user_id
	reqWithUser := map[string]interface{}{
		"id":          req.ID,
		"title":       req.Title,
		"description": req.Description,
		"status":      req.Status,
		"user_id":     userID,
	}

	// Отправляем запрос в микросервис `todo`
	url := "http://localhost:8882/tasks/update"
	responseData, err := sendAuthorizedRequest("POST", url, token.(string), reqWithUser)
	if err != nil {
		log.Printf("Ошибка обновления задачи: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Задача успешно обновлена!",
		"data":    responseData,
	})
}
