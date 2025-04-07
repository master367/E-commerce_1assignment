package delivery

import (
	"api-gateway/internal/middleware"
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(middleware.AuthMiddleware())

	inventory := r.Group("/products")
	{
		inventory.POST("", proxyRequest("http://localhost:8081/products", "POST"))
		inventory.GET("/:id", proxyRequest("http://localhost:8081/products/:id", "GET"))
		inventory.PATCH("/:id", proxyRequest("http://localhost:8081/products/:id", "PATCH"))
		inventory.DELETE("/:id", proxyRequest("http://localhost:8081/products/:id", "DELETE"))
		inventory.GET("", proxyRequest("http://localhost:8081/products", "GET"))
	}

	orders := r.Group("/orders")
	{
		orders.POST("", proxyRequest("http://localhost:8082/orders", "POST"))
		orders.GET("/:id", proxyRequest("http://localhost:8082/orders/:id", "GET"))
		orders.PATCH("/:id", proxyRequest("http://localhost:8082/orders/:id", "PATCH"))
		orders.GET("", proxyRequest("http://localhost:8082/orders", "GET"))
	}
}

func proxyRequest(target string, method string) gin.HandlerFunc {
	return func(c *gin.Context) {
		url := target
		if c.Param("id") != "" {
			url = target[:len(target)-3] + c.Param("id")
		}

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
			return
		}

		req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
			return
		}

		req.Header = c.Request.Header

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Service unavailable"})
			return
		}
		defer resp.Body.Close()

		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
			return
		}

		c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), respBody)
	}
}
