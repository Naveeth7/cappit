package routes

import (
	"github.com/cappit/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

func Register(_ *echo.Echo, _ *pgxpool.Pool, _ *redis.Client, _ *config.Config) {}
