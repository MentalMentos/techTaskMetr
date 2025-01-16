package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// UpdateTaskRequest структура запроса для обновления задачи
type UpdateTaskRequest struct {
	Id          string `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required"`
	AliceUserID string `json:"alice_user_id,omitempty"` // Идентификатор пользователя Алисы (для контекста)
}

// UpdateTaskHandler обработчик для обновления существующей задачи
func UpdateTaskHandler(c *gin.Context) {
	var req UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Fatalf("cannot bind json in update alice", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Отсутствует токен авторизации"})
		return
	}

	url := "http://localhost:8882/tasks/update"
	data, err := sendAuthorizedRequest("POST", url, token, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Задача успешно обновлена!",
		"data":    data,
	})
}
