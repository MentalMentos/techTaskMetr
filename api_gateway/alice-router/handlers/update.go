package handlers


import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
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
		c.JSON(http.StatusBadRequest, gin.H{"error":= err.Error()})
		return
	}

	// Здесь идет логика взаимодействия с вашим микросервисом задачника
	baseURL := "http://your-service-url.com/tasks/update"

	jsonValue, _ := json.Marshal(req)
	resp, err := http.Post(baseURL, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Fatalf("Ошибка при обновлении задачи: %v", err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	var data map[string]interface{}
	json.Unmarshal(body, &data)

	c.JSON(http.StatusOK, gin.H{
		"message": "Задача успешно обновлена!",
		"data":    data,
	})
}
