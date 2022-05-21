package main

import (
	"context"
	"example/webserver/graph"
	"example/webserver/graph/generated"
	"example/webserver/src/db"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, expecting PORT and DATABASE_URL to be set otherwise.")
	}

	db.InitDB()
	defer db.DB.Close(context.Background())

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
