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

func (r *UserPostgres) AddToSegment(user_id int, slug string) (int, error) {
	var id int

	query := fmt.Sprintf(`INSERT INTO %s (user_id, segment_id)
						SELECT $1, id AS segment_id
						FROM %s
						WHERE slug=$2 RETURNING id`, usersSegmentsRelationTable, segmentsTable)
	row := r.db.QueryRow(query, user_id, slug)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserPostgres) SegmentRelationExists(user_id int, slug string) (bool, error) {
	var exists bool

	query := fmt.Sprintf(`SELECT EXISTS (
		SELECT * FROM %s
		INNER JOIN %s 
		ON user_id=$1 AND slug=$2 AND segments.id = segment_id)`, usersSegmentsRelationTable, segmentsTable)

	row := r.db.QueryRow(query, user_id, slug)

	if err := row.Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}

func (r *UserPostgres) DeleteSegmentRelation(user_id int, slug string) error {
	query := fmt.Sprintf(`DELETE FROM %s
						USING %s
						WHERE user_id=$1 AND slug=$2 AND segment_id=segments.id`, usersSegmentsRelationTable, segmentsTable)
	_, err := r.db.Exec(query, user_id, slug)
	return err
}

func (r *UserPostgres) GetSegmentRelationHistory(month, year int) ([]user_segmentation.UserSegmentHistory, error) {
	var user_segments_history []user_segmentation.UserSegmentHistory

	query := fmt.Sprintf(`SELECT user_id, slug, operation_type, timestamp
						FROM %s
						INNER JOIN %s ON segments.id = segment_id
						WHERE date_part('year', timestamp)=$1 AND date_part('month', timestamp)=$2`, usersSegmentsHistory, segmentsTable)
	if err := r.db.Select(&user_segments_history, query, year, month); err != nil {
		return nil, err
	}

	return user_segments_history, nil
}
