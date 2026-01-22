package userrepo

import "github.com/jackc/pgx/v5/pgxpool"

type UserRepository struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}
