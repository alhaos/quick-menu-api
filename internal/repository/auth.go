package repository

import (
	"context"
	"errors"
	"github.com/alhaos/quick-menu-api/internal/model"
	"github.com/alhaos/quick-menu-api/internal/utils"
	"github.com/jackc/pgx"
	"time"
)

// Login authenticates a user by verifying the provided credentials against stored data.
// Returns:
//   - userID if authentication is successful
//   - error with specific messages for different failure cases:
//   - "user not found" when username doesn't exist
//   - "invalid password" when password doesn't match
//   - wrapped database error for any DB issues
//
// Note: Includes timing delay on password mismatch to mitigate brute-force attacks.
func (r *Repository) Login(user model.User) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), utils.Timeout())
	defer cancel()

	query := "SELECT password_hash, id FROM users WHERE name = $1"

	var passwordHash, id string
	row := r.db.QueryRowEx(ctx, query, nil, user.Name)
	err := row.Scan(&passwordHash, &id)
	if errors.Is(err, pgx.ErrNoRows) {
		return "", errors.New("login failed")
	}
	if err != nil {
		return "", err
	}

	if !utils.VerifyPassword(user.Password, passwordHash) {
		time.Sleep(100 * time.Millisecond)
		return "", errors.New("login failed")
	}

	return id, nil

}
