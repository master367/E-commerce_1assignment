package domain

type Product struct {
	ID         int     `json:"id" bson:"id"`
	Name       string  `json:"name" bson:"name"`
	CategoryID int     `json:"category_id" bson:"category_id"`
	Stock      int     `json:"stock" bson:"stock"`
	Price      float64 `json:"price" bson:"price"`
}

type Category struct {
	ID   int    `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}
