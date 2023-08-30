package service

import (
	"user_segmentation"
	"user_segmentation/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type SegmentService struct {
	repo repository.Segment
}

func NewSegmentService(repo repository.Segment) *SegmentService {
	return &SegmentService{repo: repo}
}

func (s *SegmentService) Create(segment user_segmentation.Segment) (int, error) {
	return s.repo.Create(segment)
}

func (s *SegmentService) Exists(slug string) (bool, error) {
	return s.repo.Exists(slug)
}

func (s *SegmentService) Delete(slug string) error {
	return s.repo.Delete(slug)
}
