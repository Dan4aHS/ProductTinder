package http_v1

import (
	"ShoesShop/internal/models/mapper"
	"ShoesShop/internal/models/restmodels"
	"ShoesShop/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Controller struct {
	s service.IService
}

func NewController(s service.IService) *Controller {
	return &Controller{s: s}
}

func (c *Controller) AddProductHandler(ctx *gin.Context) {
	var product restmodels.AddProductRequest
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	id, err := c.s.AddProduct(ctx, mapper.ProductAddRequestToDB(product))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

func (c *Controller) DeleteProductHandler(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Params.ByName("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	new_id := int32(id)
	err = c.s.DeleteProduct(ctx, new_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

func (c *Controller) ListProductsHandler(ctx *gin.Context) {
	products, err := c.s.ListProducts(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"products": products})
}
