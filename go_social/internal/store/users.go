package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type UserStore struct {
	db *sql.DB
}

var (
	ErrDuplicateEmail    = errors.New("a user with that email already exists")
	ErrDuplicateUsername = errors.New("a user with that username already exists")
)

// Create inserts a new user within an existing transaction and sets ID/CreatedAt.
func (us *UserStore) Create(ctx context.Context, tx *sql.Tx, users *User) error {
	query := `
	INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING id, created_at
	`

	err := tx.QueryRowContext(ctx, query, users.Username, users.Password.hash, users.Email).Scan(&users.ID, &users.CreatedAt)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		case err.Error() == `pq: duplicate key value violates unique constraint "users_username_key"`:
			return ErrDuplicateUsername
		default:
			return err
		}
	}

	return nil
}

// GetUserByID fetches a single user by its ID.
func (us *UserStore) GetUserByID(ctx context.Context, user_id int) (*User, error) {
	query := `
		SELECT id, username,email, created_at FROM users WHERE id = $1
	`
	ctx, cancel := context.WithTimeout(ctx, QueryDurationTime)

	defer cancel()
	var user User

	err := us.db.QueryRowContext(ctx, query, user_id).Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)
	if err != nil {
		return &User{}, err
	}
	return &user, nil
}

// GetUsers returns all users in the system.
func (us *UserStore) GetUsers(ctx context.Context) ([]*User, error) {
	query := `
		SELECT id, username,email, created_at FROM users;
	`
	ctx, cancel := context.WithTimeout(ctx, QueryDurationTime)

	defer cancel()

	rows, err := us.db.QueryContext(ctx, query)
	if err != nil {
		return []*User{}, err
	}

	defer rows.Close()

	users := make([]*User, 0)

	for rows.Next() {
		var user User

		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt); err != nil {
			return []*User{}, err
		}

		if err := rows.Err(); err != nil {
			return []*User{}, nil
		}

		users = append(users, &user)
	}

	return users, nil
}

// CreateAndInvite creates a user and a related invitation token in a single transaction.
func (us *UserStore) CreateAndInvite(ctx context.Context, user *User, invitationExp time.Duration, token string) error {

	return withTx(us.db, ctx, func(tx *sql.Tx) error {

		//create the user
		if err := us.Create(ctx, tx, user); err != nil {
			return err
		}

		//invite the user
		if err := us.createUserInvitation(ctx, tx, invitationExp, token, user.ID); err != nil {
			return err
		}

		return nil
	})

}

// createUserInvitation stores an invitation token with expiration for a user.
func (us *UserStore) createUserInvitation(ctx context.Context, tx *sql.Tx, invitationExp time.Duration, token string, userId int64) error {
	query := `INSERT INTO user_invitations (token, user_id, expiry) VALUES ($1, $2, $3)`

	ctx, cancel := context.WithTimeout(ctx, QueryDurationTime)
	defer cancel()

	_, err := tx.ExecContext(ctx, query, token, userId, time.Now().Add(invitationExp))

	if err != nil {
		return err
	}
	return nil
}

// Activate marks a user as active based on a valid invitation token.
func (us *UserStore) Activate(ctx context.Context, token string) error {
	return withTx(us.db, ctx, func(tx *sql.Tx) error {
		user, err := us.getUserFromInvitation(ctx, tx, token)
		if err != nil {
			return err
		}

		if err := us.update(ctx, tx, user); err != nil {
			return err
		}

		return us.deleteUserInvitation(ctx, tx, user.ID)
	})
}

// getUserFromInvitation loads the user associated with an invitation token.
func (us *UserStore) getUserFromInvitation(ctx context.Context, tx *sql.Tx, token string) (*User, error) {
	query := `SELECT users.id, users.created_at FROM user_invitations JOIN users ON users.id = user_invitations.user_id WHERE token = $1 AND expirty > now()`
	ctx, cancel := context.WithTimeout(ctx, QueryDurationTime)
	defer cancel()

	var user User
	if err := tx.QueryRowContext(ctx, query, token).Scan(&user.ID, &user.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &user, nil
}

// update marks a user as activated.
func (us *UserStore) update(ctx context.Context, tx *sql.Tx, user *User) error {
	query := `UPDATE users SET activated = true WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryDurationTime)
	defer cancel()

	_, err := tx.ExecContext(ctx, query, user.ID)
	return err
}

// deleteUSerInvitation removes an invitation row once used.
func (us *UserStore) deleteUserInvitation(ctx context.Context, tx *sql.Tx, userId int64) error {
	query := `DELETE FROM user_invitations WHERE user_id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryDurationTime)
	defer cancel()

	_, err := tx.ExecContext(ctx, query, userId)
	return err
}

func (us *UserStore) Delete(ctx context.Context, userId int64) error {
	return withTx(us.db, ctx, func(tx *sql.Tx) error {
		if err := us.delete(ctx, tx, userId); err != nil {
			return err
		}

		if err := us.deleteUserInvitation(ctx, tx, userId); err != nil {
			return err
		}
		return nil
	})
}

// deleteUserInvitation removes an invitation row once used.
func (us *UserStore) delete(ctx context.Context, tx *sql.Tx, userId int64) error {
	query := `DELETE FROM users WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryDurationTime)
	defer cancel()

	_, err := tx.ExecContext(ctx, query, userId)
	return err
}

// func (us *UserStore) Activate(ctx context.Context, token string) error {

// 	return withTx(us.db, ctx, func(tx *sql.Tx) error {

// 		// find the user that this token belongs to
// 		user, err := us.getUserFromInvitation(ctx, tx, token)
// 		if err != nil {
// 			return err
// 		}
// 		// update the user state
// 		user.IsActive = true

// 		if err := us.update(ctx, tx, user); err != nil {
// 			return err
// 		}
// 		// clean up the invitations

// 		if err := us.deleteUSerInvitation(ctx, tx, user.ID); err != nil {
// 			return err
// 		}

// 		return nil
// 	})
// }

// func (us *UserStore) getUserFromInvitation(ctx context.Context, tx *sql.Tx, token string) (*User, error) {
// 	query := `
// 	SELECT u.id, u.username, u.email, u.created_at, u.is_active
// 	FROM users u
// 	JOIN user_invitations ui ON u.id = ui.user_id
// 	WHERE ui.token  = $1 AND ui.expirt > $2
// 	`
// 	ctx, cancel := context.WithTimeout(ctx, QueryDurationTime)
// 	defer cancel()

// 	// store token on DB
// 	hashedTokenByte := sha256.Sum256([]byte(token))

// 	hashedToken := hex.EncodeToString(hashedTokenByte[:])

// 	user := &User{}

// 	err := tx.QueryRowContext(ctx, query, hashedToken).Scan(
// 		&user.ID,
// 		&user.Username,
// 		&user.Email,
// 		&user.CreatedAt,
// 		&user.IsActive,
// 	)

// 	if err != nil {
// 		switch err {
// 		case sql.ErrNoRows:
// 			return nil, ErrNotFound
// 		default:
// 			return nil, err
// 		}
// 	}

// 	return user, nil
// }

// func (us *UserStore) update(ctx context.Context, tx *sql.Tx, user *User) error {
// 	query := `UPDATE users SET username = $1, email= $2, is_active = $3 WHERE id = $4;`

// 	ctx, cancel := context.WithTimeout(ctx, QueryDurationTime)
// 	defer cancel()

// 	_, err := tx.ExecContext(ctx, query, user.Username, user.Email, user.IsActive, user.ID)

// 	if err != nil {

// 		return err
// 	}

// 	return nil
// }

// func (us *UserStore) deleteUSerInvitation(ctx context.Context, tx *sql.Tx, userId int64) error {
// 	query := `DELETE FROM user_invitations WHERE user_id=$1`

// 	ctx, cancel := context.WithTimeout(ctx, QueryDurationTime)
// 	defer cancel()

// 	_, err := tx.ExecContext(ctx, query, userId)

// 	if err != nil {

// 		return err
// 	}

// 	return nil
// }
