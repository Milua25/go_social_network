package store

import (
	"context"
	"database/sql"
)

// PostStore struct
type FollowerStore struct {
	db *sql.DB
}

// UnFollow removes a follower relationship between followerId and userId.
func (f *FollowerStore) UnFollow(ctx context.Context, followerId, userId int64) error {
	query := `
	DELETE FROM followers WHERE user_id = $1 AND follower_id = $2
	`

	ctx, cancel := context.WithTimeout(ctx, QueryDurationTime)
	defer cancel()

	_, err := f.db.ExecContext(ctx, query, userId, followerId)
	return err
}

// Follow creates a follower relationship between followerId and userId.
func (f *FollowerStore) Follow(ctx context.Context, followerId, userId int64) error {
	query := `
	INSERT INTO followers (user_id, follower_id) VALUES ($1, $2);
	`

	ctx, cancel := context.WithTimeout(ctx, QueryDurationTime)
	defer cancel()

	_, err := f.db.ExecContext(ctx, query, userId, followerId)
	return err
}
