package usecase

import "inventory-service/internal/domain"

type InventoryUsecase struct {
	repo InventoryRepository
}

type InventoryRepository interface {
	CreateProduct(p domain.Product) (int, error)
	GetProduct(id int) (domain.Product, error)
	UpdateProduct(id int, p domain.Product) error
	DeleteProduct(id int) error
	ListProducts() ([]domain.Product, error)
}

func NewInventoryUsecase(repo InventoryRepository) *InventoryUsecase {
	return &InventoryUsecase{repo: repo}
}

func (u *InventoryUsecase) CreateProduct(p domain.Product) (int, error) {
	return u.repo.CreateProduct(p)
}

func (u *InventoryUsecase) GetProduct(id int) (domain.Product, error) {
	return u.repo.GetProduct(id)
}

func (u *InventoryUsecase) UpdateProduct(id int, p domain.Product) error {
	return u.repo.UpdateProduct(id, p)
}

func (u *InventoryUsecase) DeleteProduct(id int) error {
	return u.repo.DeleteProduct(id)
}

func (u *InventoryUsecase) ListProducts() ([]domain.Product, error) {
	return u.repo.ListProducts()
}
