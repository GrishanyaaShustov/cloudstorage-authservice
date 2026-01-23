// Package authservice provides business logic for authentication
// and user session management.
package authservice

import (
	"context"
	"time"
)

// UserRepository defines an interface for user persistence operations
// required by the authentication service.
//
// Implementations are responsible for creating users, checking
// uniqueness constraints, and updating email verification state.
type UserRepository interface {
	// Create inserts a new user record into storage.
	Create(ctx context.Context, email, login, passwordHash string) error

	// ExistsByEmail checks whether a user with the specified email exists.
	ExistsByEmail(ctx context.Context, email string) (bool, error)

	// ExistsByLogin checks whether a user with the specified login exists.
	ExistsByLogin(ctx context.Context, login string) (bool, error)

	// ConfirmEmail marks the user's email as verified.
	ConfirmEmail(ctx context.Context, userID string) error
}

// RefreshRepository defines an interface for managing refresh sessions.
//
// Implementations are responsible for storing, validating, and removing
// refresh session records associated with a user, application, and device.
type RefreshRepository interface {
	// Save stores a refresh session with the provided TTL.
	Save(ctx context.Context, userID, app, device, value string, ttl time.Duration) error

	// Get returns the stored value for a refresh session.
	Get(ctx context.Context, userID, app, device string) (string, error)

	// Exists checks whether a refresh session exists.
	Exists(ctx context.Context, userID, app, device string) (bool, error)

	// Delete removes a refresh session.
	Delete(ctx context.Context, userID, app, device string) error
}

// AuthService defines the contract for authentication use cases.
//
// Implementations are responsible for user registration, login,
// and issuing authentication tokens.
type AuthService interface {
	// Login validates user credentials and returns authentication tokens.
	Login(ctx context.Context, req LoginRequest) (LoginResponse, error)

	// Register creates a new user account and returns the registration result.
	Register(ctx context.Context, req RegisterRequest) (RegisterResponse, error)
}

// Service implements authentication use cases.
type Service struct {
	user    UserRepository
	refresh RefreshRepository
}

// New creates a new authentication service instance.
func New(user UserRepository, refresh RefreshRepository) *Service {
	return &Service{
		user:    user,
		refresh: refresh,
	}
}
