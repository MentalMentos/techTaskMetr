package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DoneTaskRequest структура запроса для завершения задачи
type DoneTaskRequest struct {
	ID          string `json:"id" binding:"required"`          // Уникальный идентификатор задачи
	Title       string `json:"title" binding:"required"`       // Название задачи
	Description string `json:"description" binding:"required"` // Описание задачи
	Status      string `json:"status" binding:"required"`      // Новый статус задачи
	AliceUserID string `json:"alice_user_id,omitempty"`        // Идентификатор пользователя Алисы (опционально)
}

func DoneTaskHandler(c *gin.Context) {
	var req DoneTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Ошибка привязки данных в DoneTaskHandler: %v", err)
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
		"status":      req.Status, // Обычно это "1" или "done" для завершенных задач
		"user_id":     userID,
	}

	// Отправляем запрос в микросервис `todo`
	url := "http://localhost:8882/tasks/done"
	responseData, err := sendAuthorizedRequest("POST", url, token.(string), reqWithUser)
	if err != nil {
		log.Printf("Ошибка выполнения задачи: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Задача успешно выполнена!",
		"data":    responseData,
	})
}
