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

// Create inserts a new post and populates its ID/timestamps.
func (ps *PostStore) Create(ctx context.Context, posts *Post) error {
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

// GetByID fetches a post by id and includes its comments plus commenter info.
func (ps *PostStore) GetByID(ctx context.Context, post_id int) (*Post, error) {

	query := `
	SELECT p.id, p.user_id, p.title, p.content, p.created_at, p.updated_at, p.tags, p.version,
       c.id, c.content, c.created_at, u.id, u.username, u.email
	FROM posts p
	JOIN comments AS c ON c.post_id = p.id
	LEFT JOIN users u ON u.id = c.user_id
	WHERE p.id = $1 AND c.id IS NOT NULL;	`

	ctx, cancel := context.WithTimeout(ctx, QueryDurationTime)

	defer cancel()

	var post Post

	rows, err := ps.db.QueryContext(ctx, query, post_id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var comment Comment
		if err := rows.Scan(
			&post.ID,
			&post.UserID,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
			&post.UpdatedAt,
			pq.Array((&post.Tags)),
			&post.Version,
			&comment.ID,
			&comment.Content,
			&comment.CreatedAt,
			&comment.User.ID,
			&comment.User.Username,
			&comment.User.Email); err != nil {
			return nil, err
		}
		post.Comments = append(post.Comments, comment)

		if err := rows.Err(); err != nil {
			return nil, err
		}
	}

	if post.ID == 0 {
		return nil, sql.ErrNoRows
	}

	return &post, nil
}

// UpdateByID updates a post with optimistic locking on version and returns the updated row.
func (ps *PostStore) UpdateByID(ctx context.Context, post *Post) (*Post, error) {

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

	var updated Post

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
		return &Post{}, err
	}

	return &updated, nil
}

// DeleteByID removes a post by id, returning sql.ErrNoRows when nothing was deleted.
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

// GetUserFeed returns posts for a user feed with comment counts, tags, and search/filtering applied.
func (ps *PostStore) GetUserFeed(ctx context.Context, userID int64, fq PaginatedFeedQuery) ([]*PostWithMetadata, error) {

	query := `
	-- direction must be either ASC/DESC injected safely, not a bind param
	SELECT
	p.id,
	p.user_id,
	p.title,
	p.content,
	p.created_at,
	p.version,
	p.tags,
	u.username,
	COUNT(c.id) AS comments_count
	FROM posts AS p
	LEFT JOIN comments AS c ON c.post_id = p.id
	LEFT JOIN users AS u ON u.id = p.user_id
	JOIN followers AS f
	ON f.follower_id = p.user_id OR p.user_id = $1 -- posts from people they follow
	WHERE 
	p.user_id = $1 AND (p.title ILIKE '%' || $4 || '%' OR p.content ILIKE '%' || $4 || '%' ) AND (p.tags @> $5 OR $5 = '{}') -- include own posts
	GROUP BY p.id, p.user_id, p.title, p.content, p.created_at, p.version, p.tags, u.username
	ORDER BY p.created_at ` + fq.Sort + `
	LIMIT $2 OFFSET $3;`

	rows, err := ps.db.QueryContext(ctx, query, userID, fq.Limit, fq.Offset, fq.Search, pq.Array(fq.Tags))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	postsMetadata := make([]*PostWithMetadata, 0)

	for rows.Next() {

		var p PostWithMetadata
		err := rows.Scan(&p.ID, &p.UserID, &p.Title, &p.Content, &p.CreatedAt, &p.Version, pq.Array(&p.Tags), &p.User.Username, &p.CommentCount)
		if err != nil {
			return nil, err
		}

		postsMetadata = append(postsMetadata, &p)
		if rows.Err() != nil {
			return nil, err
		}

	}

	return postsMetadata, nil
}
