package main

import (
	"context"
	router_alice "github.com/MentalMentos/techTaskMetr/api_gateway/alice-router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	c := context.Background()
	router := router_alice.SetupRouter(&c)

	log.Printf("роутер создался")
	// Настройка маршрутов

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
