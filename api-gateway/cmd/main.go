package main

import (
	"api-gateway/internal/delivery"
	"api-gateway/pkg/logger"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	logger.Init()

	r := gin.Default()

	r.Use(corsMiddleware())

	delivery.SetupRoutes(r)

	log.Fatal(r.Run(":8080"))
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
