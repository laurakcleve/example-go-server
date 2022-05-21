package controllers

import (
	"context"
	"example/webserver/src/db"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Item struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

func GetItems(c *gin.Context) {
	rows, err := db.DB.Query(context.Background(), "select id, name from item")
	if err != nil {
		fmt.Println(err)
		return 
	}

	defer rows.Close()

	var items []Item

	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.ID, &item.Name); err != nil {
			fmt.Println(err)
			return 
		}
		items = append(items, item)
	}

	if rows.Err() != nil {
		fmt.Println(err)
		return 
	}

	c.IndentedJSON(http.StatusOK, items)
}