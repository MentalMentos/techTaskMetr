package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListTasksHandler(c *gin.Context) {
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

	// Формируем параметры запроса
	queryParams := map[string]interface{}{
		"user_id": userID,
	}

	// Отправляем запрос в микросервис `todo`
	url := "http://localhost:8882/tasks/list"
	responseData, err := sendAuthorizedRequest("GET", url, token.(string), queryParams)
	if err != nil {
		log.Printf("Ошибка получения списка задач: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Список задач успешно получен!",
		"tasks":   responseData,
	})
}
