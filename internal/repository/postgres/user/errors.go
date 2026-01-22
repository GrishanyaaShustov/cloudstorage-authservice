// Package userrepo provides PostgreSQL-backed user repository implementations.
package userrepo

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

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

func isUnavailable(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code == "57P01" || // admin shutdown
			pgErr.Code == "08006" || // connection failure
			pgErr.Code == "08001"
	}
	return false
}

func isUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code == "23505"
	}
	return false
}

func mapUniqueViolation(err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.ConstraintName {
		case "users_email_uq":
			return ErrEmailConflict
		case "users_login_uq":
			return ErrLoginConflict
		}
	}
	return ErrInternal
}
