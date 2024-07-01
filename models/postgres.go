package models

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type PostgresConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	SSLMode  string
}

// NewPostgresStore creates a new Postgres store
func NewPostgresStore(config PostgresConfig) (*pgxpool.Pool, error) {
	pool, err := pgxpool.Connect(context.Background(), config.DSN())
	if err != nil {
		return nil, err
	}

	// Optionally, you can verify the connection here by pinging the database
	// if err := pool.Ping(context.Background()); err != nil {
	// 	return nil, err
	// }

	return pool, nil
}

func DefaultPostgresConfig() PostgresConfig {
	return PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		Username: "baloo",
		Password: "junglebook",
		Database: "lenslocked",
		SSLMode:  "disable",
	}
}

func (c PostgresConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.Username, c.Password, c.Database, c.SSLMode,
	)
}
