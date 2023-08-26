package repository

import (
	"user_segmentation"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user user_segmentation.User) (int, error)
	UserExists(user_id int) (bool, error)
}

type Segment interface {
	CreateSegment(segment user_segmentation.Segment) (int, error)
}

type Repository struct {
	Authorization
	Segment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Segment:       NewSegmentPostgres(db),
	}
}
