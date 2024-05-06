package service

import (
	"ShoesShop/internal/models/dbmodels"
	"context"
)

type IService interface {
	AddProduct(ctx context.Context, product dbmodels.Product) (id int32, err error)
	DeleteProduct(ctx context.Context, id int32) (err error)
	ListProducts(ctx context.Context) ([]dbmodels.Product, error)
}
