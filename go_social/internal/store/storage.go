package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNotFound = errors.New("Results not Found")
)

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetByID(context.Context, int) (*Post, error)
		UpdateByID(ctx context.Context, post *Post) (*Post, error)
		DeleteByID(ctx context.Context, postID int) error
		GetUserFeed(ctx context.Context, userID int64, fq PaginatedFeedQuery) ([]*PostWithMetadata, error)
	}
	Users interface {
		Create(context.Context, *sql.Tx, *User) error
		GetUserByID(ctx context.Context, user_id int) (*User, error)
		GetUserByEmail(ctx context.Context, email string) (*User, error)
		GetUsers(ctx context.Context) ([]*User, error)
		CreateAndInvite(ctx context.Context, user *User, invitationExp time.Duration, token string) error
		Activate(ctx context.Context, token string) error
		Delete(ctx context.Context, id int64) error
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

// withTx wraps a callback in a database transaction.
func withTx(db *sql.DB, ctx context.Context, fn func(*sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
