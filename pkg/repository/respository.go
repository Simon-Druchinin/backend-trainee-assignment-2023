package repository

import (
	"user_segmentation"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user user_segmentation.User) (int, error)
}

type Segment interface {
}

type Repository struct {
	Authorization
	Segment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
