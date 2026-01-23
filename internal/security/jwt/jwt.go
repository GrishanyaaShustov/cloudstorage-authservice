// Package jwt provides helpers for issuing and validating JWT tokens.
package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	// ErrInvalidToken indicates that the token is invalid or cannot be verified.
	ErrInvalidToken = errors.New("jwt: invalid token")

	// ErrExpiredToken indicates that the token has expired.
	ErrExpiredToken = errors.New("jwt: token expired")
)

// Manager issues and validates JWT tokens.
type Manager struct {
	secret []byte
}

// New creates a new JWT Manager.
func New(secret string) *Manager {
	return &Manager{secret: []byte(secret)}
}

// IssueAccessToken creates a signed access token for the specified user context.
func (m *Manager) IssueAccessToken(userID, app, device string, ttl time.Duration) (string, error) {
	now := time.Now()

	claims := Claims{
		UserID: userID,
		App:    app,
		Device: device,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
		},
	}

	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := tok.SignedString(m.secret)
	if err != nil {
		return "", fmt.Errorf("issue access token: %w", err)
	}

	return signed, nil
}

// ParseAccessToken validates the token signature and returns parsed claims.
func (m *Manager) ParseAccessToken(tokenString string) (*Claims, error) {
	parser := jwt.NewParser(
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
	)

	var claims Claims
	_, err := parser.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (any, error) {
		return m.secret, nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}
		return nil, fmt.Errorf("%w: %v", ErrInvalidToken, err)
	}

	return &claims, nil
}
