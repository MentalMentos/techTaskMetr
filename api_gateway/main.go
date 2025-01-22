package main

import (
	"context"
	alice_router "github.com/MentalMentos/api_gateway/alice-router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	log.Printf("роутер создался")
	// Настройка маршрутов
	alice_router.SetupRouter(router)
	// Запуск сервера
	srv := &http.Server{
		Addr:    ":8880",
		Handler: router,
	}

	go func() {
		log.Printf("сервер запущен, подключён")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Ошибка запуска сервера: %v", err)
		}
	}()

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Выключение сервера...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Ошибка остановки сервера: %v", err)
	}

	log.Println("Сервер остановлен.")
}
