package authservice

import "context"

type UserRepository interface {
	Create(ctx context.Context, email, login, passwordHash string) error
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	ExistsByLogin(ctx context.Context, login string) (bool, error)
	ConfirmEmail(ctx context.Context, userID string) error
}

type AuthService interface {
	Login(ctx context.Context, req LoginRequest) (LoginResponse, error)
	Register(ctx context.Context, req RegisterRequest) (RegisterResponse, error)
}

type Service struct {
	user UserRepository
}

func New(user UserRepository) *Service {
	return &Service{user: user}
}
