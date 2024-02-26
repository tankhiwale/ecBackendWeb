package serivces

import "github.com/tankhiwale/ecBackendWeb/internal/repositories"

type IUserService interface {
}

type userService struct {
	repo repositories.IUserRepo
}

func initializeUserService(repo repositories.IUserRepo) IUserService {
	return &userService{
		repo: repo,
	}
}
