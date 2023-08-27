package repository

import (
	"fmt"
	"user_segmentation"

	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) GetActiveSegment(user_id int) ([]user_segmentation.UserSegment, error) {
	var activeSegments []user_segmentation.UserSegment

	query := fmt.Sprintf(`SELECT user_id, slug
						FROM %s INNER JOIN %s
						ON segments.id = segment_id AND user_id=$1`, usersSegmentsRelationTable, segmentsTable)
	if err := r.db.Select(&activeSegments, query, user_id); err != nil {
		return nil, err
	}

	return activeSegments, nil
}
