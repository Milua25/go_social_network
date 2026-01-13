package store

import (
	"context"
	"database/sql"
)

// var (
// 	ErrNotFound = "Results not Found"
// )

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetByID(context.Context, int) (*Post, error)
		UpdateByID(ctx context.Context, post *Post) (*Post, error)
		DeleteByID(ctx context.Context, postID int) error
		GetUserFeed(ctx context.Context, userID int64, fq PaginatedFeedQuery) ([]*PostWithMetadata, error)
	}
	Users interface {
		Create(context.Context, *User) error
		GetUserByID(ctx context.Context, user_id int) (*User, error)
		GetUsers(ctx context.Context) ([]*User, error)
	}
	Comments interface {
		Create(context.Context, *Comment) error
		GetByPostID(ctx context.Context, post_id int64) ([]*Comment, error)
	}
	Followers interface {
		Follow(ctx context.Context, followerId, userId int64) error
		UnFollow(ctx context.Context, followerId, userId int64) error
	}
}

// NewPGStorage wires up the concrete Postgres-backed stores into a single Storage facade.
func NewPGStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostStore{
			db: db,
		},
		Users: &UserStore{
			db: db,
		},
		Comments: &CommentStore{
			db: db,
		},
		Followers: &FollowerStore{
			db: db,
		},
	}
}
