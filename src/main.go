package main

import (
	"fmt"
	"go-crud/src/config"
	"go-crud/src/server"

	"go-crud/src/database/migrations"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "switch"}
	var migrateUpCmd = &cobra.Command{
		Use:   "migrateUp",
		Short: "Apply the migration up SQL queries",
		Run: func(cmd *cobra.Command, args []string) {
			db, _ := config.ConnectToPostgres()
			migrations.RunUpMigrations(db)
			fmt.Println("Up Migrations applied")
		},
	}
	var migrateDownCmd = &cobra.Command{
		Use:   "migrateDown",
		Short: "Apply the migration down SQL queries",
		Run: func(cmd *cobra.Command, args []string) {
			db, _ := config.ConnectToPostgres()
			migrations.RunDownMigrations(db)
			fmt.Println("Down Migrations applied")
		},
	}
	var serverCmd = &cobra.Command{
		Use:   "run",
		Short: "Run the server",
		Run: func(cmd *cobra.Command, args []string) {
			server.RunServer()
		},
	}
	rootCmd.AddCommand(migrateUpCmd, migrateDownCmd, serverCmd)
	rootCmd.Execute()
	// todo implement up and down migration logic
}
