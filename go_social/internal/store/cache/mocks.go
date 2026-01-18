package cache

import (
	"context"

	"github.com/Milua25/go_social/internal/store"
	"github.com/stretchr/testify/mock"
)

// NewMockStore returns a mock cache storage for tests.
func NewMockStore() Storage {
	return Storage{
		Users: &MockeUserStore{},
	}
}

type MockeUserStore struct {
	mock.Mock
}

func (m *MockeUserStore) Get(ctx context.Context, user_id int64) (*store.User, error) {
	args := m.Called(user_id)
	user, _ := args.Get(0).(*store.User)
	return user, args.Error(1)
}
func (m *MockeUserStore) Set(ctx context.Context, user *store.User) error {
	args := m.Called(user)
	return args.Error(0)
}
