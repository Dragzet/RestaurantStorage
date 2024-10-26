package repository

import "context"

type ProductStorage interface {
	ChangeProduct(ctx context.Context, name string, amount int) error
	AddProduct(ctx context.Context, name string, amount int) error
	DeleteProduct(ctx context.Context, name string) error
}
