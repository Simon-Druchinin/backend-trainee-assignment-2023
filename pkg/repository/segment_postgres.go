package repository

import (
	"fmt"
	"user_segmentation"

	"github.com/jmoiron/sqlx"
)

type SegmentPostgres struct {
	db *sqlx.DB
}

func NewSegmentPostgres(db *sqlx.DB) *SegmentPostgres {
	return &SegmentPostgres{db: db}
}

func (r *SegmentPostgres) CreateSegment(segment user_segmentation.Segment) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (slug) VALUES ($1) RETURNING id", segmentsTable)
	row := r.db.QueryRow(query, segment.Slug)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
