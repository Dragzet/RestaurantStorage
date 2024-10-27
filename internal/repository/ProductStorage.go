package repository

import (
	"RestaurantStorage/internal/models"
	"context"
)

type ProductStorage interface {
	ChangeProduct(ctx context.Context, name string, amount int) error
	AddProduct(ctx context.Context, product models.ProductModel) error
	DeleteProduct(ctx context.Context, name string) error
	GetProduct(ctx context.Context, name string) (*models.ProductModel, error)
	GetProductList(ctx context.Context) ([]*models.ProductModel, error)
}
