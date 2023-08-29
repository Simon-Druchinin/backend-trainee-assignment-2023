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

type User interface {
	GetActiveSegment(user_id int) ([]user_segmentation.UserSegment, error)
	AddToSegment(user_id int, slug string) (int, error)
	SegmentRelationExists(user_id int, slug string) (bool, error)
}

type Service struct {
	Authorization
	Segment
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Segment:       NewSegmentService(repos.Segment),
		User:          NewUserService(repos.User),
	}
}
