package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func TryLogin(req LoginRequest) (*AuthResponse, error) {
	authURL := "http://localhost:8881/auth/login"
	jsonValue, _ := json.Marshal(req)

	resp, err := http.Post(authURL, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к auth: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("авторизация не удалась")
	}

	var authResp AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		return nil, fmt.Errorf("ошибка обработки ответа от auth: %v", err)
	}

	return &authResp, nil
}

func TryRegister(req RegisterRequest) (*AuthResponse, error) {
	registerURL := "http://localhost:8881/auth/register"
	jsonValue, _ := json.Marshal(req)

	resp, err := http.Post(registerURL, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к auth: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("регистрация не удалась")
	}

	var authResp AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		return nil, fmt.Errorf("ошибка обработки ответа от auth: %v", err)
	}

	return &authResp, nil
}
