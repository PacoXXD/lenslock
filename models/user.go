package models

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int    `json:"id"`
	PasswordHash string `json:"password_hash"` // Hashed password string `json:"password"`
	Email        string `json:"email"`
}

type UserService struct {
	DB *pgxpool.Pool
}

func (us *UserService) Create(email, password string) (*User, error) {

	email = strings.ToLower(email)
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error while hashing password %s", err)
	}
	passwordHash := string(hashedBytes)

	user := User{
		Email:        email,
		PasswordHash: passwordHash,
	}

	rows := us.DB.QueryRow(context.Background(), "INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING id", email, passwordHash)

	_ = rows.Scan(&user.ID)

	return &user, nil

}

func (us *UserService) Authenticate(email, password string) (*User, error) {
	email = strings.ToLower(email)
	user := User{
		Email: email,
	}
	rows := us.DB.QueryRow(context.Background(), "SELECT id, password_hash FROM users WHERE email = $1", email)
	err := rows.Scan(&user.ID, &user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("Authenticate failed")
	}
	// Check that an encrypted password match one that's stored in the database
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	// If the credentials are valid, return the user object
	if err != nil {
		return nil, fmt.Errorf("invalid email address or password")
	}
	return &user, nil
}
