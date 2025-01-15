package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

func RegisterHandler(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	registerURL := "http://localhost:8881/register"
	jsonValue, _ := json.Marshal(req)

	// Запрос на регистрацию
	resp, err := http.Post(registerURL, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Ошибка подключения к auth: %v", err)})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := ioutil.ReadAll(resp.Body)
		c.JSON(resp.StatusCode, gin.H{"error": string(body)})
		return
	}

	// Если регистрация успешна, вызываем логин
	loginReq := LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	loginJson, _ := json.Marshal(loginReq)
	loginResp, err := http.Post("http://localhost:8881/login", "application/json", bytes.NewBuffer(loginJson))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Ошибка при логине после регистрации: %v", err)})
		return
	}
	defer loginResp.Body.Close()

	if loginResp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(loginResp.Body)
		c.JSON(loginResp.StatusCode, gin.H{"error": string(body)})
		return
	}

	var loginResponse LoginResponse
	body, _ := ioutil.ReadAll(loginResp.Body)
	if err := json.Unmarshal(body, &loginResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обработки ответа логина"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":      "Регистрация и вход выполнены успешно!",
		"access_token": loginResponse.AccessToken,
	})
}
