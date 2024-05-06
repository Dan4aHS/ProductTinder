package app

import (
	"ShoesShop/internal/config"
	http_v1 "ShoesShop/internal/transport/http-v1"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type App struct {
	Router *gin.Engine
	Port   string
}

func InitApp(cfg *config.Config, c *http_v1.Controller) *App {
	router := gin.Default()
	router.GET("/products", c.ListProductsHandler)
	router.DELETE("/products/:id", c.DeleteProductHandler)
	router.POST("/products/new", c.AddProductHandler)
	router.Static("/assets", "./assets")
	return &App{
		Router: router,
		Port:   cfg.App.RestPort,
	}
}

func (a *App) Run() {
	connStr := fmt.Sprintf("0.0.0.0:%s", a.Port)
	err := a.Router.Run(connStr)
	if err != nil {
		log.Fatal(err)
	}
}
