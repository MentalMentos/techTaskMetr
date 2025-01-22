package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// CreateTaskRequest структура запроса для создания задачи
type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required"`       // Название задачи (обязательно)
	Description string `json:"description" binding:"required"` // Описание задачи (обязательно)
	Status      string `json:"status" binding:"required"`      // Статус задачи (обязательно)
	AliceUserID string `json:"alice_user_id,omitempty"`        // Идентификатор пользователя Алисы (для контекста)
}

// CreateTaskHandler обработчик для создания новой задачи
func CreateTaskHandler(c *gin.Context) {
	var req CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Fatalf("cannot bind json in create alice", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, exists := c.Get("accesstoken")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Токен не найден"})
		return
	}

	url := "http://localhost:8882/tasks/create"
	responseData, err := sendAuthorizedRequest("POST", url, token.(string), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Задача успешно создана!",
		"data":    responseData,
	})
}
