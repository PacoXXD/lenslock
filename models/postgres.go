package models

import (
	"context"
	"database/sql"
	"fmt"
	"io/fs"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose/v3"
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

func NewPostgresDB(config PostgresConfig) (*sql.DB, error) {
	db, err := sql.Open("pgx", config.DSN())
	if err != nil {
		return nil, err
	}
	return db, nil
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

func Migrate(db *sql.DB, dir string) error {
	err := goose.SetDialect("postgres")
	if err != nil {
		return fmt.Errorf("migrate: %w", err)
	}
	err = goose.Up(db, dir)
	if err != nil {
		return fmt.Errorf("migrate: %w", err)
	}
	return nil
}

func MigrateFS(db *sql.DB, migrationsFS fs.FS, dir string) error {
	if dir == "" {
		dir = "."
	}
	goose.SetBaseFS(migrationsFS)
	defer func() { goose.SetBaseFS(nil) }()
	return Migrate(db, dir)
}
