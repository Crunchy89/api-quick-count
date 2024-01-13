package service

import (
	"github.com/crunchy89/api-quick-count/app/repository"
)

type AuthService interface {
}

type baseAuthService struct {
	userRepo repository.UserRepository
}

func NewAuthService(
	userRepo repository.UserRepository,
) AuthService {
	return &baseAuthService{
		userRepo: userRepo,
	}
}
