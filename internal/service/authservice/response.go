// Package authservice provides business logic for authentication
// and user session management.
package authservice

// LoginResponse describes the result of a successful login operation.
//
// It contains newly issued access credentials along with refresh
// session identifiers used to maintain user authentication state.
type LoginResponse struct {
	// AccessToken contains a JWT access token used for authenticating requests.
	AccessToken string

	// RefreshTokenID identifies the refresh session stored on the server side.
	RefreshTokenID string

	// RefreshTokenSecret is a secret value associated with the refresh session
	// and is used to validate refresh requests.
	RefreshTokenSecret string
}

// RegisterResponse describes the result of a successful user registration.
//
// It contains access credentials and refresh session identifiers
// allowing the user to be authenticated immediately after registration.
type RegisterResponse struct {
	// AccessToken contains a JWT access token used for authenticating requests.
	AccessToken string

	// RefreshTokenID identifies the refresh session stored on the server side.
	RefreshTokenID string

	// RefreshTokenSecret is a secret value associated with the refresh session
	// and is used to validate refresh requests.
	RefreshTokenSecret string
}
