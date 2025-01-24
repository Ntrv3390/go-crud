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
	var migrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "Apply the migration SQL queries",
		Run: func(cmd *cobra.Command, args []string) {
			db, _ := config.ConnectToPostgres()
			migrations.RunMigrations(db)
			fmt.Println("Migrations applied")
		},
	}
	var serverCmd = &cobra.Command{
		Use:   "run",
		Short: "Run the server",
		Run: func(cmd *cobra.Command, args []string) {
			server.Server()
		},
	}
	rootCmd.AddCommand(migrateCmd, serverCmd)
	rootCmd.Execute()
	// todo implement up and down migration logic
}
