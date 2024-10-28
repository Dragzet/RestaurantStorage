package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Обработка запроса
		c.Next()

		// Логируем информацию после обработки запроса
		duration := time.Since(start)
		status := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path
		server, _ := os.Hostname()

		// Пример логирования
		fmt.Printf("Server: %s, Method: %s, Path: %s, Status: %d, Duration: %v\n", server, method, path, status, duration)
	}
}
