// Package authservice provides business logic for authentication
// and user session management.
package authservice

// LoginRequest describes input data required to authenticate a user.
//
// It contains user credentials that are validated by the authentication
// service during the login process.
type LoginRequest struct {
	// Email specifies the user's email address.
	Email string

	// Password contains the user's raw password.
	// It is validated against the stored password hash.
	Password string
}

// RegisterRequest describes input data required to register a new user.
//
// It contains credentials and validation fields necessary to create
// a new user account.
type RegisterRequest struct {
	// Email specifies the user's email address.
	Email string

	// Login specifies the unique login name for the user.
	Login string

	// Password contains the user's raw password.
	Password string

	// CheckPassword contains the password confirmation value
	// used to validate password correctness.
	CheckPassword string
}
