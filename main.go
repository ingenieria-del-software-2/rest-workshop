package main

import (
	"fmt"
	"log"
	"rest-workshop/src/database"
	"rest-workshop/src/router"
)

func main() {
	// Load configuration from config.go
	config := LoadConfig()

	// Build the DSN (Data Source Name) for PostgreSQL
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.DatabaseUser,
		config.DatabasePassword,
		config.DatabaseHost,
		config.DatabasePort,
		config.DatabaseName)

	// Initialize the database connection and create the table
	db, err := database.CreateDB(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Initialize the router and start the server
	r := router.CreateRouter(db)
	address := fmt.Sprintf("%s:%s", config.Host, config.Port)
	if err := r.Run(address); err != nil {
		log.Fatalf("Could not start the server: %v", err)
	}
}
