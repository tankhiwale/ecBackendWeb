package repositories

import "github.com/jackc/pgx/v5/pgxpool"

type Repositories struct {
}

func InitializeRepositories(db *pgxpool.Pool) *Repositores {
	return &Repositores{}
}
