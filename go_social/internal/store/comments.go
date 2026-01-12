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
		c.User = Users{}
		err := rows.Scan(&c.ID, &c.PostID, &c.UserID, &c.Content, &c.CreatedAt, &c.User.Username, &c.User.ID)
		if err != nil {
			return nil, err
		}
		comments = append(comments, &c)
	}
	return comments, nil
}
