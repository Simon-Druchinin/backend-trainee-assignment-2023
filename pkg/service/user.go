package service

import (
	"user_segmentation"
	"user_segmentation/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetActiveSegment(user_id int) ([]user_segmentation.UserSegment, error) {
	return s.repo.GetActiveSegment(user_id)
}
