package repository

import (
	"context"
	"order-service/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	collection *mongo.Collection
}

func NewOrderRepository(client *mongo.Client) *OrderRepository {
	collection := client.Database("order_db").Collection("orders")
	return &OrderRepository{collection: collection}
}

func (r *OrderRepository) CreateOrder(o domain.Order) (int, error) {
	// Генерация ID (имитация автоинкремента)
	count, err := r.collection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		return 0, err
	}
	o.ID = int(count) + 1

	_, err = r.collection.InsertOne(context.Background(), o)
	return o.ID, err
}

func (r *OrderRepository) GetOrder(id int) (domain.Order, error) {
	var order domain.Order
	err := r.collection.FindOne(context.Background(), bson.M{"id": id}).Decode(&order)
	return order, err
}

func (r *OrderRepository) UpdateOrder(id int, o domain.Order) error {
	o.ID = id
	_, err := r.collection.ReplaceOne(context.Background(), bson.M{"id": id}, o)
	return err
}

func (r *OrderRepository) ListOrders(userID int) ([]domain.Order, error) {
	var orders []domain.Order
	cursor, err := r.collection.Find(context.Background(), bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	if err = cursor.All(context.Background(), &orders); err != nil {
		return nil, err
	}
	return orders, nil
}
