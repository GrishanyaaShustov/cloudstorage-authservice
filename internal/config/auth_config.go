// Package config contains configuration structures and loaders for the auth service.
package config

import "time"

// AuthConfig describes all authentication-related configuration
// options used by the auth service.
//
// The configuration is loaded from environment variables and
// controls token lifetimes, JWT signing, and cookie behavior.
type AuthConfig struct {
	// AccessTokenTTL defines how long an access token remains valid
	// after it has been issued.
	AccessTokenTTL time.Duration `env:"AUTH_ACCESS_TOKEN_TTL"`

	// RefreshTTL defines the lifetime of a refresh token.
	// After this period the user must re-authenticate.
	RefreshTTL time.Duration `env:"AUTH_REFRESH_TTL"`

	// JWTSecret is a secret key used to sign and verify JWT tokens.
	// This value must be kept private and never committed to VCS.
	JWTSecret string `env:"AUTH_JWT_SECRET"`

	// CookieDomain specifies the domain attribute for auth cookies.
	// It allows sharing cookies between subdomains if required.
	CookieDomain string `env:"AUTH_COOKIE_DOMAIN"`

	// CookieSecure defines whether auth cookies should be sent
	// only over HTTPS connections.
	CookieSecure bool `env:"AUTH_COOKIE_SECURE"`

	// CookieSameSite controls the SameSite policy for auth cookies
	// and helps protect against CSRF attacks.
	CookieSameSite string `env:"AUTH_COOKIE_SAMESITE"`
}
