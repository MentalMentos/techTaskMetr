package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// DoneTaskRequest структура запроса для завершения задачи
type DoneTaskRequest struct {
	Id          string `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required"`
	AliceUserID string `json:"alice_user_id,omitempty"` // Идентификатор пользователя Алисы (для контекста)
}

// DoneTaskHandler обработчик для завершения задачи
func DoneTaskHandler(c *gin.Context) {
	var req DoneTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, exists := c.Get("token")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Токен не найден"})
		return
	}

	url := "http://localhost:8882/tasks/done"
	data, err := sendAuthorizedRequest("POST", url, token.(string), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message": "Задача успешно завершена!",
		"data":    data,
	})
}
