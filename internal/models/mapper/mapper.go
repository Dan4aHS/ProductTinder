package mapper

import (
	"ShoesShop/internal/models/dbmodels"
	"ShoesShop/internal/models/entity"
	"ShoesShop/internal/models/restmodels"
)

func ProductEntityToDB(prod entity.Product) dbmodels.Product {
	return dbmodels.Product{
		Id:       prod.Id,
		Category: prod.Category,
		Name:     prod.Name,
		Color:    prod.Color,
		Price:    prod.Price,
		Size:     prod.Size,
		Quantity: prod.Quantity,
		Images:   prod.Images,
	}
}

func ProductDBToEntity(prod dbmodels.Product) entity.Product {
	return entity.Product{
		Id:       prod.Id,
		Category: prod.Category,
		Name:     prod.Name,
		Color:    prod.Color,
		Price:    prod.Price,
		Size:     prod.Size,
		Quantity: prod.Quantity,
		Images:   prod.Images,
	}
}

func ProductsEntityToDB(prods []entity.Product) []dbmodels.Product {
	res := make([]dbmodels.Product, len(prods))
	for i, prod := range prods {
		res[i] = ProductEntityToDB(prod)
	}
	return res
}

func ProductAddRequestToDB(prod restmodels.AddProductRequest) dbmodels.Product {
	return dbmodels.Product{
		Category: prod.Category,
		Name:     prod.Name,
		Color:    prod.Color,
		Price:    prod.Price,
		Size:     prod.Size,
		Quantity: prod.Quantity,
		Images:   prod.Images,
	}
}

func ProductsDBToListResponse(prods []dbmodels.Product) restmodels.GetProductsResponse {
	return restmodels.GetProductsResponse{
		Products: prods,
	}
}
