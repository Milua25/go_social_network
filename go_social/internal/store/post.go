package store

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
)

// PostStore struct
type PostStore struct {
	db *sql.DB
}

const QueryDurationTime = time.Minute * 2

// Create Method Set for PostStore
func (ps *PostStore) Create(ctx context.Context, posts *Posts) error {
	query := `
	INSERT INTO posts (content, title, user_id, tags)
	VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at
	`

	ctx, cancel := context.WithTimeout(ctx, QueryDurationTime)

	defer cancel()

	err := ps.db.QueryRowContext(
		ctx,
		query,
		posts.Content,
		posts.Title,
		posts.UserID,
		pq.Array(posts.Tags),
	).Scan(&posts.ID, &posts.CreatedAt, &posts.UpdatedAt)

	if err != nil {
		return err
	}
	return nil
}

// Get Method Set for PostStore
func (ps *PostStore) GetByID(ctx context.Context, post_id int) (*Posts, error) {

	query := `SELECT id, user_id, title, content, created_at, updated_at, tags, version FROM posts WHERE id = $1 `

	ctx, cancel := context.WithTimeout(ctx, QueryDurationTime)

	defer cancel()

	var posts Posts

	err := ps.db.QueryRowContext(ctx, query, post_id).Scan(
		&posts.ID,
		&posts.UserID,
		&posts.Title,
		&posts.Content,
		&posts.CreatedAt,
		&posts.UpdatedAt,
		pq.Array((&posts.Tags)),
		&posts.Version,
	)
	if err != nil {
		return &Posts{}, err
	}

	return &posts, nil
}

// Update Method Set for PostStore
func (ps *PostStore) UpdateByID(ctx context.Context, post *Posts) (*Posts, error) {

	query := `UPDATE posts
	SET title = $1,
    content = $2,
    tags = $3,
    updated_at = NOW(),
	version = version +1 
WHERE id = $4 AND version = $5
RETURNING id, user_id, title, content, tags, created_at, updated_at, version;
`
	ctx, cancel := context.WithTimeout(ctx, QueryDurationTime)

	defer cancel()

	var updated Posts

	err := ps.db.QueryRowContext(ctx, query, post.Title, post.Content, pq.Array(post.Tags), post.ID, post.Version).Scan(
		&updated.ID,
		&updated.UserID,
		&updated.Title,
		&updated.Content,
		pq.Array((&updated.Tags)),
		&updated.CreatedAt,
		&updated.UpdatedAt,
		&updated.Version,
	)
	if err != nil {
		return &Posts{}, err
	}

	return &updated, nil
}

// Delete Method Set for PostStore
func (ps *PostStore) DeleteByID(ctx context.Context, post_id int) error {

	query := `
		DELETE FROM posts
		WHERE id = $1; `

	ctx, cancel := context.WithTimeout(ctx, QueryDurationTime)

	defer cancel()

	result, err := ps.db.ExecContext(ctx, query, post_id)
	if err != nil {
		return err
	}
	// Optional: ensure a row was actually deleted.
	if rows, _ := result.RowsAffected(); rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}
