package main

import (
	"api-gateway/internal/delivery"
	"api-gateway/pkg/logger"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация логгера
	logger.Init()

	// Создание Gin роутера
	r := gin.Default()

	r.Use(corsMiddleware())

	// Настройка маршрутов и middleware
	delivery.SetupRoutes(r)

	// Запуск сервера на порту 8080
	log.Fatal(r.Run(":8080"))
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // Разрешаем все источники
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204) // Обработка preflight-запросов
			return
		}
		c.Next()
	}
}
