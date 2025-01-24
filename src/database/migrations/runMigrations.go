package migrations

import (
	"database/sql"
	"go-crud/src/database"
)

func RunMigrations(db *sql.DB) {
	// run all migrations here
	database.CreateDatabaseIfNotExist(db)
	database.CreateUsersTable(db)
}
