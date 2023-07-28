package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"main.go/BubbleFood/models"
)

var db *sql.DB

func GetAlbums(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	rows, err := db.Query("SELECT id, title, artist, price FROM albums")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var albums []models.Album
	for rows.Next() {
		var a models.Album
		err := rows.Scan(&a.ID, &a.Title, &a.Artist, &a.Price)
		if err != nil {
			panic(err)
		}
		albums = append(albums, a)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, albums)
}
