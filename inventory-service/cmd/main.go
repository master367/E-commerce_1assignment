package main

import (
	"context"
	"inventory-service/internal/delivery"
	"inventory-service/internal/repository"
	"inventory-service/internal/usecase"
	"inventory-service/pkg/logger"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	logger.Init()

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	defer client.Disconnect(context.Background())

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}
	logger.Logger.Println("Connected to MongoDB")

	repo := repository.NewInventoryRepository(client)
	uc := usecase.NewInventoryUsecase(repo)
	r := gin.Default()

	delivery.NewInventoryHandler(r, uc)

	log.Fatal(r.Run(":8081"))
}
