package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"rest-workshop/src/models"
)

type DB struct {
	conn *sqlx.DB
}

// Create inserts a new item into the database.
func (db *DB) Create(data models.Item) error {
	query := `INSERT INTO items (name, price) VALUES (:name, :price)`
	_, err := db.conn.NamedExec(query, data)
	return err
}

// CreateDB initializes a connection to the PostgreSQL database and creates the necessary tables.
func CreateDB(dsn string) (*DB, error) {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	// Create the items table if it doesn't exist
	createTableQuery := `CREATE TABLE IF NOT EXISTS items (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		price NUMERIC(10, 2) NOT NULL
	)`
	if _, err := db.Exec(createTableQuery); err != nil {
		return nil, fmt.Errorf("failed to create table: %w", err)
	}

	log.Println("Database connection established and table verified/created")
	return &DB{conn: db}, nil
}
