package handlers

import (
	"context"
	"net/http"

	"Go_Lang_Api/db"
	"Go_Lang_Api/models"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BookHandler struct {
	DB *pgxpool.Pool
}

func (h *BookHandler) CreateBook(
	c *gin.Context,
) {

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
