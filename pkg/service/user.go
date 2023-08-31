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

func (s *UserService) AddToSegment(user_id int, slug string) (int, error) {
	return s.repo.AddToSegment(user_id, slug)
}

func (s *UserService) SegmentRelationExists(user_id int, slug string) (bool, error) {
	return s.repo.SegmentRelationExists(user_id, slug)
}

func (s *UserService) DeleteSegmentRelation(user_id int, slug string) error {
	return s.repo.DeleteSegmentRelation(user_id, slug)
}

func (s *UserService) GetSegmentRelationHistory(month, year int) ([]user_segmentation.UserSegmentHistory, error) {
	return s.repo.GetSegmentRelationHistory(month, year)
}
