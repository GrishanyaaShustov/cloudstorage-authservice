// Package authservice provides business logic for authentication
// and user session management.
package authservice

import (
	"context"

	"github.com/GrishanyaaShustov/cloudstorage-authservice/internal/repository/refresh"
	userrepo "github.com/GrishanyaaShustov/cloudstorage-authservice/internal/repository/user"
)

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
	user    userrepo.UserRepository
	refresh refreshrepo.RefreshRepository
}

// New creates a new authentication service instance.
func New(user userrepo.UserRepository, refresh refreshrepo.RefreshRepository) *Service {
	return &Service{
		user:    user,
		refresh: refresh,
	}
}
