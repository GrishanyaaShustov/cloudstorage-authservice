// Package redis provides helpers for creating and initializing
// Redis clients.
package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/GrishanyaaShustov/cloudstorage-authservice/internal/config"
)

// New creates and initializes a Redis client using the provided configuration.
func New(ctx context.Context, cfg config.RedisConfig) (*redis.Client, error) {
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	rdb := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     cfg.Password,
		DB:           cfg.DB,
		DialTimeout:  cfg.DialTimeout,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})

	pingCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	if err := rdb.Ping(pingCtx).Err(); err != nil {
		_ = rdb.Close()
		return nil, fmt.Errorf("redis ping failed: %w", err)
	}

	return rdb, nil
}

// MustNew creates a Redis client and panics on error.
func MustNew(ctx context.Context, cfg config.RedisConfig) *redis.Client {
	rdb, err := New(ctx, cfg)
	if err != nil {
		panic(err)
	}

	return rdb
}
