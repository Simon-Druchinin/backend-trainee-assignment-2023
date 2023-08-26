package service

import (
	"user_segmentation"
	"user_segmentation/pkg/repository"
)

type Authorization interface {
	CreateUser(user user_segmentation.User) (int, error)
}

type Segment interface {
}

type Service struct {
	Authorization
	Segment
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
