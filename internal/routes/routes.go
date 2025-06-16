package routes

import (
	"github.com/cappit/internal/apikey"
	"github.com/cappit/internal/auth"
	"github.com/cappit/internal/config"
	"github.com/cappit/internal/proxy"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

func Register(e *echo.Echo, db *pgxpool.Pool, rdb *redis.Client, cfg *config.Config) {
	// Auth routes
	e.POST("/auth/register", auth.Register(db))
	e.POST("/auth/login", auth.Login(db))

	// JWT protected routes
	authGroup := e.Group("")
	authGroup.Use(auth.JWTMiddleware(cfg.JWTSecret))

	// API key routes
	authGroup.GET("/apikeys", apikey.ListKeys(db))
	authGroup.POST("/apikeys", apikey.CreateKey(db))
	authGroup.PUT("/apikeys/:id/limits", apikey.SetLimits(db))

	// Proxy route (no auth, key-based only)
	e.Any("/proxy/:apikey/*", proxy.Handler(db, rdb))
}
