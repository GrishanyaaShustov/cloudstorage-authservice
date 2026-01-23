package authservice

import "errors"

var (
	ErrWrongCheckPassword = errors.New("password and check-password do not match")
	ErrNotValidEmail      = errors.New("not valid email")
	ErrInternal           = errors.New("internal error")
)
