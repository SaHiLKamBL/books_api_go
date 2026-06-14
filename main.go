package main

import (
	"github.com/gin-gonic/gin"
	"Go_Lang_Api/db"
	"Go_Lang_Api/handlers"
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"log"
)

func main() {
             err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
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
	router.DELETE(
		"/books/:id",bookHandler.DeleteBook,
	)

	fmt.Println(os.Getenv("DATABASE_URL"))
     fmt.Println("Program Started")
	router.Run(":8080")
}