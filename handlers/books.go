package handlers

import (
	"Go_Lang_Api/db"
	"Go_Lang_Api/models"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BookHandler struct {
	DB *pgxpool.Pool
}

func (h *BookHandler) CreateBook(c *gin.Context,) {

	var newBook models.Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	err := db.InsertBook(
		context.Background(),
		h.DB,
		newBook,
	)

	if err != nil {
     fmt.Println("DB Error:", err)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusCreated,
		newBook,
	)
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
    id := c.Param("id") // gets "1" from /books/1

    // convert string to int
    bookID, err := strconv.Atoi(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid ID",
        })
        return
    }

    // call DB function
    err = db.DeleteBook(context.Background(), h.DB, bookID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    // success response
    c.JSON(http.StatusOK, gin.H{
        "message": "Book deleted successfully",
    })
}


func (h *BookHandler) GetBooks(
	c *gin.Context,
) {

	books, err := db.GetAllBooks(
		context.Background(),
		h.DB,
	)

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		books,
	)
}


