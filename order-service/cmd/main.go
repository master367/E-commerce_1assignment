package main

import (
	"context"
	"log"
	"order-service/internal/delivery"
	"order-service/internal/repository"
	"order-service/internal/usecase"
	"order-service/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Инициализация логгера
	logger.Init()

	// Подключение к MongoDB
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	defer client.Disconnect(context.Background())

	// Проверка подключения
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}
	logger.Logger.Println("Connected to MongoDB")

	// Инициализация слоев
	repo := repository.NewOrderRepository(client)
	uc := usecase.NewOrderUsecase(repo)
	r := gin.Default()

	// Настройка маршрутов
	delivery.NewOrderHandler(r, uc)

	// Запуск сервера на порту 8082
	log.Fatal(r.Run(":8082"))
}
