package database

import (
	"database/sql"
	"log"

	"go-crud/src/api/models"
)

func InsertUserQuery(db *sql.DB, name string, age int) error {
	query := `
		INSERT INTO users (name, age) VALUES ($1, $2)
	`
	_, err := db.Exec(query, name, age)
	if err != nil {
		log.Fatal("Error inserting user:", err)
	}
	return err
}

func GetUsersQuery(db *sql.DB) ([]models.User, error) {
	query := `
		SELECT * FROM users ORDER BY id asc
	`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("Error getting users:", err)
		return nil, err
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.Name, &user.Age, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			log.Fatal("Error scanning user:", err)
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func GetUserQuery(db *sql.DB, id string) (*models.User, error) {
	query := `
		SELECT id, name, age FROM users WHERE id = $1
	`
	row := db.QueryRow(query, id)

	var user models.User
	err := row.Scan(&user.Id, &user.Name, &user.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Printf("Error scanning user: %v", err)
		return nil, err
	}

	return &user, nil
}

func PutUserQuery(db *sql.DB, id string, name string, age int) (*models.User, error) {
	selectQuery := `
		SELECT name, age FROM users WHERE id = $1
	`
	var existingUser models.User
	err := db.QueryRow(selectQuery, id).Scan(&existingUser.Name, &existingUser.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("User with ID %v not found.", id)
			return nil, nil
		}
		log.Printf("Error retrieving user: %v", err)
		return nil, err
	}

	updateQuery := `
		UPDATE users
		SET name = $2, age = $3, updatedAt = CURRENT_TIMESTAMP
		WHERE id = $1
		RETURNING id, name, age, updatedAt
	`

	var updatedUser models.User
	err = db.QueryRow(updateQuery, id, name, age).Scan(&updatedUser.Id, &updatedUser.Name, &updatedUser.Age, &updatedUser.UpdatedAt)
	if err != nil {
		log.Printf("Error updating user: %v", err)
		return nil, err
	}

	return &updatedUser, nil
}

func DeleteUserQuery(db *sql.DB, id string) (*models.User, error) {
	var user models.User
	query := `SELECT id, name, age FROM users WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&user.Id, &user.Name, &user.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Printf("Error fetching user: %v", err)
		return nil, err
	}

	deleteQuery := `DELETE FROM users WHERE id = $1`
	_, err = db.Exec(deleteQuery, id)
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		return nil, err
	}
	return &user, nil
}
