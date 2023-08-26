package service

import (
	"user_segmentation"
	"user_segmentation/pkg/repository"
)

type SegmentService struct {
	repo repository.Segment
}

func NewSegmentService(repo repository.Segment) *SegmentService {
	return &SegmentService{repo: repo}
}

func (s *SegmentService) CreateSegment(segment user_segmentation.Segment) (int, error) {
	return s.repo.CreateSegment(segment)
}
