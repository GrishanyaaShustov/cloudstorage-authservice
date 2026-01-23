// Package crypto provides helpers for password hashing and secure random generation.
package crypto

import (
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/rand"
)

const (
	// DefaultBcryptCost defines the default bcrypt cost.
	DefaultBcryptCost = bcrypt.DefaultCost
)

// HashPassword hashes the provided password using bcrypt.
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), DefaultBcryptCost)
	if err != nil {
		return "", fmt.Errorf("hash password: %w", err)
	}
	return string(hash), nil
}

// ComparePassword verifies that the provided password matches the bcrypt hash.
func ComparePassword(hash, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return fmt.Errorf("compare password: %w", err)
	}
	return nil
}

// RandomURLSafeString returns a cryptographically secure random string encoded
// using base64 URL encoding without padding.
func RandomURLSafeString(nBytes int) (string, error) {
	if nBytes <= 0 {
		return "", fmt.Errorf("random string: nBytes must be positive")
	}

	b := make([]byte, nBytes)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("random string: %w", err)
	}

	return base64.RawURLEncoding.EncodeToString(b), nil
}
