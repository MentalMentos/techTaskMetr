package routes

import (
	authcontroller "github.com/MentalMentos/techTaskMetr/auth/internal/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter(authController *authcontroller.AuthController) *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies(nil) // Доверять всем прокси

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome Home!")
	})

	router.GET("/ip", func(c *gin.Context) {
		// Получаем IP клиента
		clientIP := c.ClientIP() // Автоматически извлекает IP с учётом заголовков X-Forwarded-For, X-Real-IP
		c.JSON(200, gin.H{"ip": clientIP})
	})

	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/register", authController.Register)             // Регистрация
		authRoutes.POST("/login", authController.Login)                   // Вход
		authRoutes.POST("/refresh", authController.RefreshToken)          // Обновление токена
		authRoutes.PUT("/update-password", authController.UpdatePassword) // Обновление пароля
	}

	return router
}
