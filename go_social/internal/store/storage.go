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
		Create(context.Context, *Posts) error
		GetByID(context.Context, int) (*Posts, error)
		UpdateByID(ctx context.Context, post *Posts) (*Posts, error)
		DeleteByID(ctx context.Context, post_id int) error
	}
	Users interface {
		Create(context.Context, *Users) error
	}
	Comments interface {
		Create(context.Context, *Comment) error
		GetByPostID(ctx context.Context, post_id int64) ([]*Comment, error)
	}
}

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
	}
}
