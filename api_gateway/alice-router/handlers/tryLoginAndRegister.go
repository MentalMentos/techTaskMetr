package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	UserID       int64  `json:"user_id"`
}

// TryLogin адаптированная попытка авторизации
func TryLogin(req LoginRequest) (*AuthResponse, error) {
	authURL := "http://localhost:8881/auth/login"
	jsonValue, _ := json.Marshal(req)

	resp, err := http.Post(authURL, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, fmt.Errorf("Ошибка подключения к auth: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Авторизация не удалась")
	}

	var authResp AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		return nil, fmt.Errorf("Ошибка обработки ответа от auth: %v", err)
	}

	return &authResp, nil
}

// TryRegister адаптированная попытка регистрации
func TryRegister(req RegisterRequest) (*AuthResponse, error) {
	registerURL := "http://localhost:8881/auth/register"
	jsonValue, _ := json.Marshal(req)

	resp, err := http.Post(registerURL, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, fmt.Errorf("Ошибка подключения к auth: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("Регистрация не удалась")
	}

	// После успешной регистрации выполняем авторизацию
	loginReq := LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}
	return TryLogin(loginReq)
}
