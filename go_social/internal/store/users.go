package store

import (
	"context"
	"database/sql"
)

type UserStore struct {
	db *sql.DB
}

func (us *UserStore) Create(ctx context.Context, users *Users) error {
	query := `
	INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING id, created_at
	`

	err := us.db.QueryRowContext(ctx, query, users.Username, users.Password, users.Email).Scan(&users.ID, &users.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
