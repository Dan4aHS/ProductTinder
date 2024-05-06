package service

import (
	"ShoesShop/internal/models/dbmodels"
	"ShoesShop/internal/models/mapper"
	"ShoesShop/internal/repository"
	"context"
)

type Service struct {
	Repo repository.IRepository
}

func NewService(repo repository.IRepository) *Service {
	return &Service{Repo: repo}
}

func (s *Service) AddProduct(ctx context.Context, product dbmodels.Product) (id int32, err error) {
	id, err = s.Repo.AddProduct(ctx, mapper.ProductDBToEntity(product))
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Service) DeleteProduct(ctx context.Context, id int32) (err error) {
	err = s.Repo.DeleteProduct(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) ListProducts(ctx context.Context) ([]dbmodels.Product, error) {
	products, err := s.Repo.ListProducts(ctx, nil)
	if err != nil {
		return nil, err
	}
	return mapper.ProductsEntityToDB(products), nil
}
