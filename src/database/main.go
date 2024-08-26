package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
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

// GetItem retrieves an item by its ID from the database.
func (db *DB) GetItem(id int, item *models.Item) error {
	query := `SELECT id, name, price FROM items WHERE id = $1`
	return db.conn.Get(item, query, id)
}

// CreateDB initializes a connection to the database and creates the necessary tables.
func CreateDB(dsn string) (*DB, error) {
	var db *sqlx.DB
	var err error

	if dsn == "file::memory:?cache=shared" || dsn == ":memory:" || dsn == "" {
		// Use SQLite for in-memory database
		db, err = sqlx.Connect("sqlite3", dsn)
	} else {
		// Assume PostgreSQL for other DSN formats
		db, err = sqlx.Connect("postgres", dsn)
	}

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
