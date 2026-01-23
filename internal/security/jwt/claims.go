// Package jwt provides helpers for issuing and validating JWT tokens.
package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

// Claims describes JWT claims used by the auth service.
//
// It contains standard registered claims along with custom fields
// required to identify the user and the session context.
type Claims struct {
	// UserID identifies the authenticated user.
	UserID string `json:"uid"`

	// App identifies the client application requesting the token.
	App string `json:"app"`

	// Device identifies the device/session for the token.
	Device string `json:"device"`

	jwt.RegisteredClaims
}
