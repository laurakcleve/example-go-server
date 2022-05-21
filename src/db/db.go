package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

var DB *pgx.Conn

func InitDB() {
	databaseURL := os.Getenv("DATABASE_URL")

	conn, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	DB = conn
}