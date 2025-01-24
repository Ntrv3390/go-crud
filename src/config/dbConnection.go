package config

import (
	"database/sql"
	"fmt"
	"log"

	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectToPostgres() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	connStr := os.Getenv("POSTGRES_CONNECTION_STRING")
	if connStr == "" {
		log.Fatal("POSTGRES_CONNECTION_STRING is not set in the .env file")
	}

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to the database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Error pinging the database: %v", err)
	}

	fmt.Println("Connected to PostgreSQL successfully!")
	return db, nil
}
