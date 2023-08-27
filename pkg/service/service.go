package service

import (
	"user_segmentation"
	"user_segmentation/pkg/repository"
)

type Authorization interface {
	CreateUser(user user_segmentation.User) (int, error)
	UserExists(user_id int) (bool, error)
}

type Segment interface {
	Create(segment user_segmentation.Segment) (int, error)
	Exists(slug string) (bool, error)
	Delete(slug string) error
}

type Service struct {
	Authorization
	Segment
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Segment:       NewSegmentService(repos.Segment),
	}
}
