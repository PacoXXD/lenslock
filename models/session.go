package models

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/PacoXXD/lenslock/rand"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	MinBytesPerToken = 32
)

type Session struct {
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	Token     string
	TokenHash string `json:"token_hash"`
}

type SessionService struct {
	DB            *pgxpool.Pool
	BytesPerToken int
}

func (ss *SessionService) Create(userID int) (*Session, error) {
	bytesPerToken := ss.BytesPerToken
	if bytesPerToken < MinBytesPerToken {
		bytesPerToken = MinBytesPerToken
	}
	// type error interface {
	// 	Error() string
	// }

	// var errNoRow error = errors.New("no rows in result set")

	token, err := rand.String(bytesPerToken)
	if err != nil {
		return nil, fmt.Errorf("create token %w", err)
	}
	session := Session{
		UserID:    userID,
		Token:     token,
		TokenHash: ss.hash(token),
	}
	// Try to update the existing session
	commandTag, err := ss.DB.Exec(context.Background(), "UPDATE session SET token_hash=$1 WHERE user_id=$2", session.TokenHash, session.UserID)
	if err != nil {
		return nil, fmt.Errorf("update tokenHash %w", err)
	}

	// If no rows were updated, insert a new session
	if commandTag.RowsAffected() == 0 {
		rows := ss.DB.QueryRow(context.Background(), "INSERT INTO session (user_id, token_hash) VALUES ($1, $2) RETURNING id", session.UserID, session.TokenHash)
		err = rows.Scan(&session.ID)
		if err != nil {
			return nil, fmt.Errorf("create tokenHash %w", err)
		}
	} else {
		// Get the existing session ID
		rows := ss.DB.QueryRow(context.Background(), "SELECT id FROM session WHERE user_id=$1", session.UserID)
		err = rows.Scan(&session.ID)
		if err != nil {
			return nil, fmt.Errorf("get session ID %w", err)
		}
	}

	return &session, nil
}

func (ss *SessionService) Get(token string) (*Session, error) {
	tokenHash := ss.hash(token)
	rows := ss.DB.QueryRow(context.Background(), "SELECT id, user_id, token_hash FROM session WHERE token_hash=$1", tokenHash)
	var session Session
	err := rows.Scan(&session.ID, &session.UserID, &session.TokenHash)
	if err == pgx.ErrNoRows {
		// No existing session found, insert a new session
		rows = ss.DB.QueryRow(context.Background(), "INSERT INTO session (user_id, token_hash) VALUES ($1, $2) RETURNING id", session.UserID, session.TokenHash)
		err = rows.Scan(&session.ID)
		if err != nil {
			return nil, fmt.Errorf("create tokenHash %w", err)
		}
	} else if err != nil {
		return nil, fmt.Errorf("update tokenHash %w", err)
	}

	return &session, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	tokenHash := ss.hash(token)
	var user User

	rows := ss.DB.QueryRow(context.Background(), "SELECT user_id FROM session WHERE token_hash=$1", tokenHash)
	err := rows.Scan(&user.ID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("no session found: %w", err)
		}
		return nil, fmt.Errorf("get user ID %w", err)
	}

	// Get the user details from the users table
	rows = ss.DB.QueryRow(context.Background(), "SELECT email, password_hash FROM users WHERE id=$1", user.ID)
	err = rows.Scan(&user.Email, &user.PasswordHash)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("no user found: %w", err)
		}
		return nil, fmt.Errorf("get user %w", err)
	}

	return &user, nil
}

// func (ss *SessionService) User(token string) (*User, error) {
// 	var user User
// 	tokenHash := ss.hash(token)
// 	rows := ss.DB.QueryRow(context.Background(), "SELECT user_id FROM session WHERE token_hash=$1", tokenHash)
// 	err := rows.Scan(&user.ID)
// 	if err != nil {
// 		return nil, fmt.Errorf("get user %w", err)
// 	}
// 	fmt.Println(user.ID)
// 	// Get the user ID from the session
// 	rows = ss.DB.QueryRow(context.Background(), "SELECT email,password_hash FROM users WHERE id=$1", user.ID)
// 	err = rows.Scan(&user.Email, &user.PasswordHash)
// 	if err != nil {
// 		return nil, fmt.Errorf("user %w", err)
// 	}
// 	fmt.Println(user.Email, user.PasswordHash)
// 	return &user, nil

// }

func (ss *SessionService) hash(token string) string {
	tokenHash := sha256.Sum256([]byte(token))
	return base64.URLEncoding.EncodeToString(tokenHash[:])
}

// type TokenManager struct {
// }

// func (tm *TokenManager) New() {

// }
