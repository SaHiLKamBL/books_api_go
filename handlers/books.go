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
    "Go_Lang_Api/util"
)

type BookHandler struct {
	DB *pgxpool.Pool
}

//hash password

func (h *BookHandler) RegisterHandler(c *gin.Context) {
	var u struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	hashedPassword, err := util.HashPassword(u.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	err = db.CreateUser(context.Background(), h.DB, u.Email, hashedPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
}


func (h *BookHandler) LogUser(w )
 

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


