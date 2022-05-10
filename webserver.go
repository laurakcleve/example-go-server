package main

import (
	"context"
	"fmt"

	"log"
	"os"

	"example/webserver/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

type Server struct {
	db     *pgx.Conn
	router *gin.Engine
}

func (s *Server) postGoTest(ginConn *gin.Context) {
	if _, err := s.db.Exec(context.Background(), "INSERT INTO gotest(id) VALUES($1)", 3); err != nil {
		// Handling error, if occur
		fmt.Println("Unable to insert due to: ", err)
		return
	}
}

func NewServer(db *pgx.Conn) *Server {
	router := gin.Default()
	server := &Server{db, router}
	return server
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, expecting PORT and DATABASE_URL to be set otherwise.")
	}

	databaseURL := os.Getenv("DATABASE_URL")
	port := os.Getenv("PORT")

	conn, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	server := NewServer(conn)
	defer conn.Close(context.Background())

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	var id int64
	// 	err = conn.QueryRow(context.Background(), "select * from gotest").Scan(&id)
	// 	fmt.Fprintf(w, "The id is: %v", id)
	// })

	// fs := http.FileServer(http.Dir("static/"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Wrap in log.Fatal to output errors, will not output anything otherwise
	// https://github.com/golang/go/issues/11693
	// log.Fatal(http.ListenAndServe(":"+port, nil))

	server.router.GET("/albums", controllers.GetAlbums)
	server.router.POST("/albums", controllers.PostAlbums)
	server.router.POST("/gotest", server.postGoTest)

	server.router.Run("localhost:" + port)
}
