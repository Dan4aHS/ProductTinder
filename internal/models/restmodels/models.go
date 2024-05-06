package restmodels

import "ShoesShop/internal/models/dbmodels"

type AddProductRequest struct {
	Category int32    `json:"category"`
	Name     string   `json:"name"`
	Color    string   `json:"color"`
	Price    int32    `json:"price"`
	Size     int32    `json:"size"`
	Quantity int32    `json:"quantity"`
	Images   []string `json:"image"`
}

type AddProductResponse struct {
	ProductID int32 `json:"product_id"`
}

type DeleteProductRequest struct {
	ProductID int32 `json:"product_id"`
}

type GetProductsResponse struct {
	Products []dbmodels.Product `json:"products"`
}
