package store

import (
	"context"
	"database/sql"
)

type CommentStore struct {
	db *sql.DB
}

func (com *CommentStore) GetByPostID(ctx context.Context, post_id int64) ([]*Comment, error) {
	query := `SELECT
  c.id,
  c.post_id,
  c.user_id,
  c.content,
  c.created_at,
  u.username,
  u.id AS user_id
FROM comments AS c
JOIN users AS u ON u.id = c.user_id
WHERE c.post_id = $1
ORDER BY c.created_at DESC;
`

	rows, err := com.db.QueryContext(ctx, query, post_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := []*Comment{}

	for rows.Next() {
		var c Comment
		c.User = User{}
		err := rows.Scan(&c.ID, &c.PostID, &c.UserID, &c.Content, &c.CreatedAt, &c.User.Username, &c.User.ID)
		if err != nil {
			return nil, err
		}
		comments = append(comments, &c)
	}
	return comments, nil
}

// Create Method Set for PostStore
func (ps *CommentStore) Create(ctx context.Context, comment *Comment) error {
	query := `
	INSERT INTO comments (post_id, user_id, content)
	VALUES ($1, $2, $3) RETURNING id, created_at
	`

	ctx, cancel := context.WithTimeout(ctx, QueryDurationTime)

	defer cancel()

	err := ps.db.QueryRowContext(
		ctx,
		query,
		comment.PostID,
		comment.UserID,
		comment.Content,
	).Scan(&comment.ID, &comment.CreatedAt)

	if err != nil {
		return err
	}
	return nil
}
