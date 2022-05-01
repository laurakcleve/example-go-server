package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	databaseURL := os.Getenv("DATABASE_URL")
	port := os.Getenv("PORT")

	conn, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var id int64
		err = conn.QueryRow(context.Background(), "select * from gotest").Scan(&id)
		fmt.Fprintf(w, "The id is: %v", id)
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Wrap in log.Fatal to output errors, will not output anything otherwise
	// https://github.com/golang/go/issues/11693
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
