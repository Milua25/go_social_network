package store

import (
	"context"
	"database/sql"
	"time"
)

// NewMockStore returns a Storage with mocked user store for tests.
func NewMockStore() Storage {
	return Storage{
		Users: &MockUserStore{},
	}
}

type MockUserStore struct {
}

// Create satisfies the User store interface without side effects.
func (m *MockUserStore) Create(ctx context.Context, tx *sql.Tx, u *User) error {
	return nil
}

// GetUserByID returns nil user for tests.
func (m *MockUserStore) GetUserByID(ctx context.Context, user_id int) (*User, error) {

	return nil, nil
}

// GetUserByEmail returns nil user for tests.
func (m *MockUserStore) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	return nil, nil
}

// GetUsers returns nil users slice for tests.
func (m *MockUserStore) GetUsers(ctx context.Context) ([]*User, error) {
	return nil, nil
}
func (m *MockUserStore) CreateAndInvite(ctx context.Context, user *User, invitationExp time.Duration, token string) error {
	return nil
}
func (m *MockUserStore) Activate(ctx context.Context, token string) error {
	return nil
}

func (m *MockUserStore) Delete(ctx context.Context, id int64) error {
	return nil
}
