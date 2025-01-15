package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Здесь идет логика взаимодействия с вашим микросервисом задачника
	baseURL := "http://your-service-url.com/tasks/create"

	jsonValue, _ := json.Marshal(req)
	resp, err := http.Post(baseURL, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil || resp.StatusCode != http.StatusCreated {
		log.Fatalf("Ошибка при создании задачи: %v", err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	var data map[string]interface{}
	json.Unmarshal(body, &data)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Задача успешно создана!",
		"data":    data,
	})
}
