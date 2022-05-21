package main

import (
	"context"
	"example/webserver/graph"
	"example/webserver/graph/generated"
	"example/webserver/src/db"
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
		log.Printf("No PORT env found, setting to default port: %s", port)
	}
	
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, expecting PORT and DATABASE_URL to be set otherwise.")
	}

	db.InitDB()
	defer db.DB.Close(context.Background())

	r := gin.Default()

	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())

	r.Run(":" + port)
}
