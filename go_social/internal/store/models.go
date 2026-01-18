package store

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Post struct {
	ID        int64     `json:"id"`
	Content   string    `json:"content"`
	Title     string    `json:"title"`
	UserID    int64     `json:"user_id"`
	Tags      []string  `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Version   int       `json:"version"`
	Comments  []Comment `json:"comments"`
	User      User      `json:"user"`
}

type PostWithMetadata struct {
	Post
	CommentCount int `json:"comment_count"`
}

type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Password  password  `json:"-"` // not returning the password
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	IsActive  bool      `json:"is_active"`
	RoleID    int64     `json:"role_id"`
	Role      Role      `json:"role"`
}

type password struct {
	text *string
	hash []byte
}

// Set hashes and stores the provided plaintext password.
func (p *password) Set(text string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	p.text = &text
	p.hash = hash

	return nil
}

type Comment struct {
	ID        int64  `json:"id"`
	PostID    int64  `json:"post_id"`
	UserID    int64  `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	User      User   `json:"user"`
}

type Follower struct {
	UserID     int64  `json:"user_id"`
	FollowerID int64  `json:"follower_id"`
	CreatedAt  string `json:"created_at"`
}

type Role struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:description`
	Level       int    `json:"level"`
}
