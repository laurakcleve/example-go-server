package main

import (
	"context"

	"log"
	"os"

	"example/webserver/src/controllers"
	"example/webserver/src/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, expecting PORT and DATABASE_URL to be set otherwise.")
	}

	db.InitDB()
	defer db.DB.Close(context.Background())

	router := gin.Default()

	router.GET("/albums", controllers.GetAlbums)
	router.GET("/items", controllers.GetItems)
	router.POST("/albums", controllers.PostAlbums)

	port := os.Getenv("PORT")
	router.Run(":" + port)
}
