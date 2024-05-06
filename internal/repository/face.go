package repository

import (
	"ShoesShop/internal/models/entity"
	"context"
)

type IRepository interface {
	AddProduct(ctx context.Context, prod entity.Product) (id int32, err error)
	DeleteProduct(ctx context.Context, id int32) (err error)
	ListProducts(ctx context.Context, filters map[string]any) ([]entity.Product, error)
}
