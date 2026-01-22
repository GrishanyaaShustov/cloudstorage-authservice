// Package config contains configuration structures and loaders for the auth service.
package config

import "time"

// PostgresConfig describes PostgreSQL connection and pool configuration
// used by the auth service.
//
// It defines connection credentials, database name, and parameters
// controlling the behavior of the connection pool.
type PostgresConfig struct {
	// Host specifies the PostgreSQL server host.
	Host string `env:"POSTGRES_HOST"`

	// Port specifies the PostgreSQL server port.
	Port int `env:"POSTGRES_PORT"`

	// User defines the database user name.
	User string `env:"POSTGRES_USER"`

	// Password defines the database user password.
	// This value must be kept secret and never committed to VCS.
	Password string `env:"POSTGRES_PASSWORD"`

	// Name specifies the name of the PostgreSQL database.
	Name string `env:"POSTGRES_NAME"`

	// MaxConns defines the maximum number of connections
	// in the PostgreSQL connection pool.
	MaxConns int32 `env:"POSTGRES_MAX_CONNS" env-default:"10" yaml:"max_conns"`

	// MinConns defines the minimum number of connections
	// maintained in the connection pool.
	MinConns int32 `env:"POSTGRES_MIN_CONNS" env-default:"2" yaml:"min_conns"`

	// MaxConnLifetime defines the maximum amount of time
	// a connection may be reused.
	MaxConnLifetime time.Duration `env:"POSTGRES_MAX_CONN_LIFETIME" env-default:"30m" yaml:"max_conn_lifetime"`

	// MaxConnIdleTime defines how long a connection may remain
	// idle before being closed.
	MaxConnIdleTime time.Duration `env:"POSTGRES_MAX_CONN_IDLE_TIME" env-default:"5m" yaml:"max_conn_idle_time"`

	// HealthCheckPeriod defines how often the connection pool
	// performs health checks on idle connections.
	HealthCheckPeriod time.Duration `env:"POSTGRES_HEALTH_CHECK_PERIOD" env-default:"1m" yaml:"health_check_period"`
}
