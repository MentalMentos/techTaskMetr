package main

import (
	"context"
	"github.com/MentalMentos/techTaskMetr/api_gateway/alice-router"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()
	c := context.Background()

	router_alice.SetupRouter(&c)

	log.Println("API Gateway running on port :8080")
	router.Run(":8080") // Запускаем сервер на порту 8080
}
