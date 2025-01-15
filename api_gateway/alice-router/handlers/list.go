package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

// ListTasksRequest структура запроса для получения списка задач
type ListTasksRequest struct {
	AliceUserID string `json:"alice_user_id,omitempty"` // Идентификатор пользователя Алисы (для контекста)
}

// ListTasksHandler обработчик для получения списка задач
func ListTasksHandler(c *gin.Context) {
	var req ListTasksRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Здесь идет логика взаимодействия с вашим микросервисом задачника
	baseURL := "http://your-service-url.com/tasks/list"

	jsonValue, _ := json.Marshal(req)
	resp, err := http.Get(baseURL + "?alice_user_id=" + req.AliceUserID)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Fatalf("Ошибка при получении списка задач: %v", err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	var data map[string]interface{}
	json.Unmarshal(body, &data)

	c.JSON(http.StatusOK, gin.H{
		"message": "Список задач получен!",
		"data":    data,
	})
}
