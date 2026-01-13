package store

import (
	"context"
	"database/sql"
)

type UserStore struct {
	db *sql.DB
}

// Create inserts a new user and sets the generated ID/CreatedAt fields.
func (us *UserStore) Create(ctx context.Context, users *User) error {
	query := `
	INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING id, created_at
	`

	err := us.db.QueryRowContext(ctx, query, users.Username, users.Password, users.Email).Scan(&users.ID, &users.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

// GetUserByID fetches a single user by its ID.
func (us *UserStore) GetUserByID(ctx context.Context, user_id int) (*User, error) {
	query := `
		SELECT id, username,email, created_at FROM users WHERE id = $1
	`
	ctx, cancel := context.WithTimeout(ctx, QueryDurationTime)

	defer cancel()
	var user User

	err := us.db.QueryRowContext(ctx, query, user_id).Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)
	if err != nil {
		return &User{}, err
	}
	return &user, nil
}

// GetUsers returns all users in the system.
func (us *UserStore) GetUsers(ctx context.Context) ([]*User, error) {
	query := `
		SELECT id, username,email, created_at FROM users;
	`
	ctx, cancel := context.WithTimeout(ctx, QueryDurationTime)

	defer cancel()

	rows, err := us.db.QueryContext(ctx, query)
	if err != nil {
		return []*User{}, err
	}

	defer rows.Close()

	users := make([]*User, 0)

	for rows.Next() {
		var user User

		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt); err != nil {
			return []*User{}, err
		}

		if err := rows.Err(); err != nil {
			return []*User{}, nil
		}

		users = append(users, &user)
	}

	return users, nil
}
