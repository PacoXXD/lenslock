package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type PostgresConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	SSLMode  string
}

func (c PostgresConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", c.Host, c.Port, c.Username, c.Password, c.Database, c.SSLMode)
}

func main() {
	cfg := PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		Username: "baloo",
		Password: "junglebook",
		Database: "lenslocked",
		SSLMode:  "disable",
	}
	conn, err := pgx.Connect(context.Background(), cfg.DSN())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	us := &UserService{
		DB: conn,
	}
	user, err := us.Create("john@example.com", "wan890")
	if err != nil {
		panic(err)
	}
	fmt.Println(user)

}

// 	_, err = conn.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, first_name TEXT, age INTEGER)")
// 	if err != nil {
// 		panic(err)
// 	}

// 	_, err = conn.Exec(context.Background(), `INSERT INTO users (first_name, age)
// 	VALUES ('John', 30),
// 	       ('Jane', 25),
// 	       ('Bob', 40);`)
// 	if err != nil {
// 		panic(err)
// 	}

// 	var first_name string
// 	var age int64
// 	err = conn.QueryRow(context.Background(), "select first_name, age from users where age=$1", 40).Scan(&first_name, &age)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
// 		os.Exit(1)
// 	}

// 	fmt.Println(first_name, age)

// 	_, err = conn.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS orders (id SERIAL PRIMARY KEY, user_id INTEGER, amount INTEGER, description TEXT)")
// 	if err != nil {
// 		panic(err)
// 	}

// 	userID := 2
// 	for i := 0; i < 6; i++ {
// 		amount := i * 10
// 		desc := fmt.Sprintf("Order #%d", i)
// 		_, err = conn.Exec(context.Background(), "INSERT INTO orders (amount, description, user_id) VALUES ($1, $2, $3)", amount, desc, userID)
// 		if err != nil {
// 			panic(err)
// 		}

// 	} }
