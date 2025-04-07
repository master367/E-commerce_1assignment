package domain

type Order struct {
	ID         int         `json:"id" bson:"id"`
	UserID     int         `json:"user_id" bson:"user_id"`
	Status     string      `json:"status" bson:"status"`
	Items      []OrderItem `json:"items" bson:"items"`
	TotalPrice float64     `json:"total_price" bson:"total_price"`
}

type OrderItem struct {
	ProductID int `json:"product_id" bson:"product_id"`
	Quantity  int `json:"quantity" bson:"quantity"`
}
