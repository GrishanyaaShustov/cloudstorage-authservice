package userrepo

import (
	"context"
	"errors"
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

var (
	// ErrNotFound indicates that the requested entity was not found in storage.
	ErrNotFound = errors.New("repository: not found")

	// ErrEmailConflict indicates a conflict caused by a duplicate email value.
	ErrEmailConflict = errors.New("repository: email conflict")

	// ErrLoginConflict indicates a conflict caused by a duplicate login value.
	ErrLoginConflict = errors.New("repository: login conflict")

	// ErrUnavailable indicates that the repository backend is unavailable.
	ErrUnavailable = errors.New("repository: unavailable")

	// ErrInternal indicates an unexpected internal repository error.
	ErrInternal = errors.New("repository: internal")
)
