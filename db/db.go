package db

import (
	"context"
	//"encoding/csv"
	//"fmt"
	"log"
	"os"
//	"strconv"
	//"time"

	"github.com/jackc/pgx/v5/pgxpool"
)




func ConnectDB() *pgxpool.Pool {
	dbURL := os.Getenv("DATABASE_URL")

	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	return pool
}