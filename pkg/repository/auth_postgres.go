package repository

import (
	"fmt"
	"user_segmentation"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user user_segmentation.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s DEFAULT VALUES RETURNING id", usersTable)
	row := r.db.QueryRow(query)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
