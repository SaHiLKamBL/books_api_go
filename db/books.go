package db

import (
	"context"

	"Go_Lang_Api/models"

	"github.com/jackc/pgx/v5/pgxpool"
)


func DeleteBook(ctx context.Context,db *pgxpool.Pool,id int) error{
	 query :=`
	     DELETE FROM books
		 WHERE id=$1
	 `
	 _,err :=db.Exec(
		ctx,
		query,
		id,
	 )

	 return err
}


func InsertBook(ctx context.Context,db *pgxpool.Pool,book models.Book,) error {
	query := `
	INSERT INTO books(
		id,
		title,
		author,
		quantity
	)
	VALUES($1,$2,$3,$4)
	`

	_, err := db.Exec(
		ctx,
		query,
		book.Id,
		book.Title,
		book.Author,
		book.Quantity,
	)

	return err
}


func GetAllBooks(ctx context.Context,db *pgxpool.Pool,) ([]models.Book,error) {

	query := `
	SELECT
		id,
		title,
		author,
		quantity
	FROM books
	`

	rows, err := db.Query(
		ctx,
		query,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []models.Book

	for rows.Next() {

		var book models.Book

		err := rows.Scan(
			&book.Id,
			&book.Title,
			&book.Author,
			&book.Quantity,
		)

		if err != nil {
			return nil, err
		}

		books = append(
			books,
			book,
		)
	}

	return books,nil
}