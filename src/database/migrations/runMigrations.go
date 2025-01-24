package migrations

import (
	"database/sql"
	"go-crud/src/database"
)

func RunUpMigrations(db *sql.DB) {
	database.CreateDatabaseIfNotExist(db)
	database.CreateUsersTable(db)
}

// migrations down not working fix needed
func RunDownMigrations(db *sql.DB) {
	database.DropDatabaseIfExists(db)
	database.DropUsersTable(db)
}
