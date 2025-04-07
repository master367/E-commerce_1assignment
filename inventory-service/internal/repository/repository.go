package repository

import (
	"context"
	"inventory-service/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type InventoryRepository struct {
	collection *mongo.Collection
}

func NewInventoryRepository(client *mongo.Client) *InventoryRepository {
	collection := client.Database("inventory_db").Collection("products")
	return &InventoryRepository{collection: collection}
}

func (r *InventoryRepository) CreateProduct(p domain.Product) (int, error) {
	// Генерация ID (имитация автоинкремента)
	count, err := r.collection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		return 0, err
	}
	p.ID = int(count) + 1

	_, err = r.collection.InsertOne(context.Background(), p)
	return p.ID, err
}

func (r *InventoryRepository) GetProduct(id int) (domain.Product, error) {
	var product domain.Product
	err := r.collection.FindOne(context.Background(), bson.M{"id": id}).Decode(&product)
	return product, err
}

func (r *InventoryRepository) UpdateProduct(id int, p domain.Product) error {
	p.ID = id
	_, err := r.collection.ReplaceOne(context.Background(), bson.M{"id": id}, p)
	return err
}

func (r *InventoryRepository) DeleteProduct(id int) error {
	_, err := r.collection.DeleteOne(context.Background(), bson.M{"id": id})
	return err
}

func (r *InventoryRepository) ListProducts() ([]domain.Product, error) {
	var products []domain.Product
	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	if err = cursor.All(context.Background(), &products); err != nil {
		return nil, err
	}
	return products, nil
}
