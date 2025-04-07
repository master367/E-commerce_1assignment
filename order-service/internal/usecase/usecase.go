package usecase

import "order-service/internal/domain"

type OrderUsecase struct {
	repo OrderRepository
}

type OrderRepository interface {
	CreateOrder(o domain.Order) (int, error)
	GetOrder(id int) (domain.Order, error)
	UpdateOrder(id int, o domain.Order) error
	ListOrders(userID int) ([]domain.Order, error)
}

func NewOrderUsecase(repo OrderRepository) *OrderUsecase {
	return &OrderUsecase{repo: repo}
}

func (u *OrderUsecase) CreateOrder(o domain.Order) (int, error) {
	return u.repo.CreateOrder(o)
}

func (u *OrderUsecase) GetOrder(id int) (domain.Order, error) {
	return u.repo.GetOrder(id)
}

func (u *OrderUsecase) UpdateOrder(id int, o domain.Order) error {
	return u.repo.UpdateOrder(id, o)
}

func (u *OrderUsecase) ListOrders(userID int) ([]domain.Order, error) {
	return u.repo.ListOrders(userID)
}
