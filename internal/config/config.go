package config

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type Config struct {
	PostgresURL string
	RedisAddr   string
	RedisPass   string
	RedisDB     int
	JWTSecret   string
}

func Load() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath("./cfg")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("No config file found in ./cfg. Using environment variables...")
	}

	return &Config{
		PostgresURL: viper.GetString("POSTGRES_URL"),
		RedisAddr:   viper.GetString("REDIS_ADDR"),
		RedisPass:   viper.GetString("REDIS_PASS"),
		RedisDB:     viper.GetInt("REDIS_DB"),
		JWTSecret:   viper.GetString("JWT_SECRET"),
	}
}

func InitPostgres(cfg *Config) *pgxpool.Pool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, cfg.PostgresURL)
	if err != nil {
		log.Fatalf("Unable to connect to Postgres: %v", err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("Unable to ping Postgres: %v", err)
	}

	return pool
}

func InitRedis(cfg *Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPass,
		DB:       cfg.RedisDB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("Unable to connect to Redis: %v", err)
	}

	return rdb
}
