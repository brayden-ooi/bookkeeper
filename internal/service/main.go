package service

import (
	"database/sql"
	"log"
	"os"

	"github.com/brayden-ooi/bookkeeper/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

var DB *database.Queries

func init() {
	godotenv.Load()

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("dbURL is not found in the env file")
	}

	conn, err := sql.Open("sqlite3", dbURL)
	if err != nil {
		log.Fatal("Cant connect to database:", err)
	}

	DB = database.New(conn)

	conn.Exec("PRAGMA foreign_keys = ON")
}
