package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lib/pq"
)

func CreateDatabaseIfNotExist(db *sql.DB) {
	_, err := db.Exec("CREATE DATABASE crud")
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "42P04" {
			fmt.Println("Database 'crud' already exists.")
		} else {
			log.Fatalf("Error creating database: %v\n", err)
		}
	} else {
		fmt.Println("Database 'crud' created successfully.")
	}
}

func DropDatabaseIfExists(db *sql.DB) {
	_, err := db.Exec("DROP DATABASE IF EXISTS crud")
	if err != nil {
		log.Fatalf("Error dropping database: %v\n", err)
	} else {
		fmt.Println("Database 'crud' dropped successfully.")
	}
}

func CreateUsersTable(db *sql.DB) error {
	checkQuery := `
		SELECT EXISTS (
			SELECT 1 
			FROM information_schema.tables 
			WHERE table_name = 'users'
		);
	`
	var exists bool
	err := db.QueryRow(checkQuery).Scan(&exists)
	if err != nil {
		log.Printf("Error checking if 'users' table exists: %v\n", err)
		return err
	}

	if exists {
		log.Println("Table 'users' already exists.")
		return nil
	}

	createQuery := `
        CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY NOT NULL,
            name VARCHAR(100) NOT NULL UNIQUE,
			age INT
        );
    `
	_, err = db.Exec(createQuery)
	if err != nil {
		log.Printf("Error creating 'users' table: %v\n", err)
		return err
	}

	log.Println("Table 'users' created successfully.")
	return nil
}

func DropUsersTable(db *sql.DB) error {
	dropQuery := `
		DROP TABLE IF EXISTS users;
	`
	_, err := db.Exec(dropQuery)
	if err != nil {
		log.Printf("Error dropping 'users' table: %v\n", err)
		return err
	}
	log.Println("Table 'users' dropped successfully.")
	return nil
}
