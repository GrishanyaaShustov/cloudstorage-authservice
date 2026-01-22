// Package config contains configuration structures and loaders for the auth service.
package config

// HTTPConfig describes HTTP server configuration used by the auth service.
//
// It defines network-related settings required to start and expose
// the HTTP API, including the address the server listens on.
type HTTPConfig struct {
	// Addr specifies the network address the HTTP server binds to.
	// The value is typically provided in the form "host:port".
	Addr string `yaml:"addr"`
}
