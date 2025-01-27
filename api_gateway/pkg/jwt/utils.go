package pkg

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const secretKey = "your_secret_key" // Ваш секретный ключ для подписи JWT

type AuthClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// CheckJWT проверяет наличие и валидность JWT-токена в заголовке Authorization
func CheckJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := extractTokenFromHeader(c.GetHeader("Authorization"))
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid authorization header"})
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*AuthClaims)
		if !ok || claims.ExpiresAt < time.Now().Unix() {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Expired token"})
			c.Abort()
			return
		}

		// Если токен валиден, сохраняем данные пользователя в контексте запроса
		c.Set("username", claims.Username)
		c.Next()
	}
}

// extractTokenFromHeader извлекает JWT-токен из заголовка Authorization
func extractTokenFromHeader(header string) string {
	if len(header) > 7 && strings.ToLower(header[:7]) == "bearer " {
		return header[7:]
	}
	return ""
}
