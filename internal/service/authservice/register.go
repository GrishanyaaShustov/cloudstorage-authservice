package authservice

import (
	"context"
	"strings"
)

func (s *Service) Register(ctx context.Context, req RegisterRequest) (RegisterResponse, error) {
	email := strings.ToLower(req.Email)
	login := req.Login
	password := req.Password
	check := req.CheckPassword

	if !isCheckPasswordCorrect(password, check) {
		return RegisterResponse{}, ErrWrongCheckPassword
	}

	if !isValidEmail(email) {
		return RegisterResponse{}, ErrNotValidEmail
	}

	emailExists, err := s.user.ExistsByEmail(ctx, email)
	if err != nil {
	}
}
