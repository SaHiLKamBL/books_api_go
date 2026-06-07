package main

import (
	"github.com/gin-gonic/gin"
	"Go_Lang_Api/db"
	"Go_Lang_Api/handlers"
	"fmt"
	"os"
)

func main() {

	dbPool := db.ConnectDB()

	defer dbPool.Close()

	router := gin.Default()

	bookHandler := handlers.BookHandler{
		DB: dbPool,
	}

	router.GET(
		"/books",
		bookHandler.GetBooks,
	)

	router.POST(
		"/books",
		bookHandler.CreateBook,
	)
	fmt.Println(os.Getenv("DATABASE_URL"))
     fmt.Println("Program Started")
	router.Run(":8080")
}