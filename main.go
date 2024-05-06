package main

import (
	"ShoesShop/internal/app"
	"ShoesShop/internal/config"
	"ShoesShop/internal/repository/postgres"
	"ShoesShop/internal/service"
	http_v1 "ShoesShop/internal/transport/http-v1"
	"log"
)

func main() {
	cfg := config.GetConfig()
	db, err := postgres.ConnectPostgres(cfg)
	if err != nil {
		log.Fatal(err)
	}
	repo := postgres.NewRepository(db)
	serv := service.NewService(repo)
	controller := http_v1.NewController(serv)
	application := app.InitApp(cfg, controller)
	application.Run()
}
