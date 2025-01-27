package migrations

import (
	"database/sql"
	"go-crud/src/database"
)

func RunUpMigrations(db *sql.DB) {
	database.CreateDatabaseIfNotExist(db)
	database.CreateUsersTable(db)
}

func RunDownMigrations(db *sql.DB) {
	// database.DropDatabaseIfExists(db)
	database.DropUsersTable(db)
}
