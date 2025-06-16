package main

import (
	"github.com/cappit/internal/config"
	"github.com/cappit/internal/middleware"
	"github.com/cappit/internal/routes"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	cfg := config.Load()
	db := config.InitPostgres(cfg)
	rdb := config.InitRedis(cfg)

	middleware.Logger(e)

	routes.Register(e, db, rdb, cfg)

	log.Fatal(e.Start(":8080"))
}
