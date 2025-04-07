package delivery

import (
	"api-gateway/internal/middleware"
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Middleware для логирования и аутентификации
	r.Use(gin.Logger())
	r.Use(middleware.AuthMiddleware())

	// Группа маршрутов для Inventory Service
	inventory := r.Group("/products")
	{
		inventory.POST("", proxyRequest("http://localhost:8081/products", "POST"))
		inventory.GET("/:id", proxyRequest("http://localhost:8081/products/:id", "GET"))
		inventory.PATCH("/:id", proxyRequest("http://localhost:8081/products/:id", "PATCH"))
		inventory.DELETE("/:id", proxyRequest("http://localhost:8081/products/:id", "DELETE"))
		inventory.GET("", proxyRequest("http://localhost:8081/products", "GET"))
	}

	// Группа маршрутов для Order Service
	orders := r.Group("/orders")
	{
		orders.POST("", proxyRequest("http://localhost:8082/orders", "POST"))
		orders.GET("/:id", proxyRequest("http://localhost:8082/orders/:id", "GET"))
		orders.PATCH("/:id", proxyRequest("http://localhost:8082/orders/:id", "PATCH"))
		orders.GET("", proxyRequest("http://localhost:8082/orders", "GET"))
	}
}

// proxyRequest перенаправляет запросы к соответствующему сервису
func proxyRequest(target string, method string) gin.HandlerFunc {
	return func(c *gin.Context) {
		url := target
		if c.Param("id") != "" {
			url = target[:len(target)-3] + c.Param("id") // Заменяем ":id" на реальный ID
		}

		// Чтение тела запроса
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
			return
		}

		// Создание нового запроса
		req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
			return
		}

		// Копирование заголовков
		req.Header = c.Request.Header

		// Отправка запроса
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Service unavailable"})
			return
		}
		defer resp.Body.Close()

		// Чтение ответа
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
			return
		}

		// Передача ответа клиенту
		c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), respBody)
	}
}
