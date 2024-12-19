package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "notification"
)

func InitPostgres() *pgxpool.Pool {
	// Build connection string
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		user,
		password,
		host,
		port,
		dbname,
		"disable",
	)

	fmt.Printf("connecting to postgres: %s:%d\n", host, port)

	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		panic(err)
	}

	// Configure connection pool
	config.MinConns = 25
	config.MaxConns = 40

	// Configure logging
	// Connect to database
	conn, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		panic(err)
	}

	// Test connection
	err = conn.Ping(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to postgres!!")
	return conn
}
