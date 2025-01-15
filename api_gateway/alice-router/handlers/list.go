package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ListTasksRequest структура запроса для получения списка задач
type ListTasksRequest struct {
	AliceUserID string `json:"alice_user_id,omitempty"` // Идентификатор пользователя Алисы (для контекста)
}

// ListTasksHandler обработчик для получения списка задач
func ListTasksHandler(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Отсутствует токен авторизации"})
		return
	}

	url := "http://localhost:8882/tasks/list"
	responseData, err := sendAuthorizedRequest("GET", url, token, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Список задач получен!",
		"data":    responseData,
	})
}
