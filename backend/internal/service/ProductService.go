package service

import (
	"RestaurantStorage/internal/models"
	"context"
	"strings"
)

func (s *Service) ChangeProduct(ctx context.Context, name string, amount int) error {
	name = strings.ToLower(name)
	return s.repo.ChangeProduct(ctx, name, amount)
}

func (s *Service) DeleteProduct(ctx context.Context, name string) error {
	name = strings.ToLower(name)
	return s.repo.DeleteProduct(ctx, name)
}

func (s *Service) GetProduct(ctx context.Context, name string) (*models.ProductModel, error) {
	name = strings.ToLower(name)
	return s.repo.GetProduct(ctx, name)
}

func (s *Service) GetProductsList(ctx context.Context) ([]*models.ProductModel, error) {
	return s.repo.GetProductList(ctx)
}

func (s *Service) AddProduct(ctx context.Context, product models.ProductModel) error {
	product.Name = strings.ToLower(product.Name)
	return s.repo.AddProduct(ctx, product)
}
