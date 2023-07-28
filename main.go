package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"main.go/BubbleFood/controllers"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "dragon"
	dbname   = "ForFoodiesDb"
)

var db *sql.DB

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	fmt.Println("db open")

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	router := gin.Default()
	router.GET("/albums", controllers.GetAlbums)
	// router.POST("/albums", createAlbum)

	router.Run("localhost:8080")

}
