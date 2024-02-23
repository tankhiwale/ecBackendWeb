package repositories

import "github.com/jackc/pgx/v5/pgxpool"

type IUserRepo interface {
	Save()
	Get()
	GetBy()
	Update()
}
type userRepo struct {
	db *pgxpool.Pool
}

func initializeUserRepo(db *pgxpool.Pool) IUserRepo {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) Save() {
}

func (u *userRepo) Get() {
}

func (u *userRepo) GetBy() {
}

func (u *userRepo) Update() {
}
