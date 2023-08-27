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
	Create(segment user_segmentation.Segment) (int, error)
	Exists(slug string) (bool, error)
	Delete(slug string) error
}

type User interface {
	GetActiveSegment(user_id int) ([]user_segmentation.UserSegment, error)
}

type Repository struct {
	Authorization
	Segment
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Segment:       NewSegmentPostgres(db),
		User:          NewUserPostgres(db),
	}
}
