// Package config contains configuration structures and loaders for the auth service.
package config

import "time"

// RedisConfig describes Redis connection and pool configuration
// used by the auth service.
//
// It defines connection parameters, authentication credentials,
// timeouts, and pool-related settings.
type RedisConfig struct {
	// Host specifies the Redis server host.
	Host string `env:"REDIS_HOST"`

	// Port specifies the Redis server port.
	Port int `env:"REDIS_PORT"`

	// Password defines the Redis authentication password.
	// This value must be kept secret and never committed to VCS.
	Password string `env:"REDIS_PASSWORD"`

	// DB specifies the Redis logical database number.
	DB int `env:"REDIS_DB"`

	// DialTimeout defines the maximum time allowed to establish
	// a connection to the Redis server.
	DialTimeout time.Duration `env:"REDIS_DIAL_TIMEOUT" env-default:"5s"`

	// ReadTimeout defines the maximum duration for read operations.
	ReadTimeout time.Duration `env:"REDIS_READ_TIMEOUT" env-default:"3s"`

	// WriteTimeout defines the maximum duration for write operations.
	WriteTimeout time.Duration `env:"REDIS_WRITE_TIMEOUT" env-default:"3s"`

	// PoolSize defines the maximum number of connections
	// in the Redis connection pool.
	PoolSize int `env:"REDIS_POOL_SIZE" env-default:"10"`

	// MinIdleConns defines the minimum number of idle connections
	// maintained in the Redis connection pool.
	MinIdleConns int `env:"REDIS_MIN_IDLE_CONNS" env-default:"2"`
}
