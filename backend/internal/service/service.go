package service

import (
	"RestaurantStorage/internal/models"
	"RestaurantStorage/internal/repository"
	"context"
)

type Service struct {
	repo *repository.DB
	ProductService
}

type ProductService interface {
	AddProduct(ctx context.Context, product models.ProductModel) error
	GetProductsList(ctx context.Context) ([]*models.ProductModel, error)
	ChangeProduct(ctx context.Context, name string, amount int) error
	DeleteProduct(ctx context.Context, name string) error
	GetProduct(ctx context.Context, name string) (*models.ProductModel, error)
}

func NewService(repo *repository.DB) *Service {
	return &Service{
		repo: repo,
	}
}
