package cache

import (
	"context"

	"github.com/Milua25/go_social/internal/store"
	"github.com/go-redis/redis/v8"
)

type Storage struct {
	Users interface {
		Get(context.Context, int64) (*store.User, error)
		Set(context.Context, *store.User) error
	}
}

// NewRedisDBStorage returns a cache Storage backed by Redis client.
func NewRedisDBStorage(rdb *redis.Client) Storage {
	return Storage{
		Users: &UserStore{rdb: rdb},
	}
}
