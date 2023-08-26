package service

import (
	"user_segmentation"
	"user_segmentation/pkg/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user user_segmentation.User) (int, error) {
	return s.repo.CreateUser(user)
}

func (s *AuthService) UserExists(user_id int) (bool, error) {
	return s.repo.UserExists(user_id)
}
