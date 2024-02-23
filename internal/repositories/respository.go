package repositories

import "github.com/jackc/pgx/v5/pgxpool"

type Repositories struct {
	userRepo IUserRepo
}

func InitializeRepositories(db *pgxpool.Pool) *Repositories {
	return &Repositories{
		userRepo: initializeUserRepo(db),
	}
}
