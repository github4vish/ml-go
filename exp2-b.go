package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

func main() {
	// Open a connection to the SQLite database
	db, err := sql.Open("sqlite3", "./example.db")
	if err != nil {
		log.Fatalf("failed to open database: %s", err)
	}
	defer db.Close()

	// Create a table (if not exists)
	createTableSQL := `CREATE TABLE IF NOT EXISTS employees (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        age INTEGER,
        email TEXT,
        is_manager BOOLEAN
    );`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("failed to create table: %s", err)
	}

	// Insert sample data into the employees table
	insertSQL := `INSERT INTO employees (name, age, email, is_manager) VALUES (?, ?, ?, ?)`
	_, err = db.Exec(insertSQL, "John Doe", 30, "john.doe@example.com", false)
	if err != nil {
		log.Fatalf("failed to insert data: %s", err)
	}
	_, err = db.Exec(insertSQL, "Jane Smith", 25, "jane.smith@example.com", true)
	if err != nil {
		log.Fatalf("failed to insert data: %s", err)
	}

	// Query the employees table
	rows, err := db.Query("SELECT name, age, email, is_manager FROM employees")
	if err != nil {
		log.Fatalf("failed to query data: %s", err)
	}
	defer rows.Close()

	// Iterate through the result set
	fmt.Println("Employees:")
	for rows.Next() {
		var name string
		var age int
		var email string
		var isManager bool

		// Scan the row into variables
		err := rows.Scan(&name, &age, &email, &isManager)
		if err != nil {
			log.Fatalf("failed to scan row: %s", err)
		}

		fmt.Printf("Name: %s, Age: %d, Email: %s, Is Manager: %t\n", name, age, email, isManager)
	}

	// Check for errors during iteration
	if err = rows.Err(); err != nil {
		log.Fatalf("error during rows iteration: %s", err)
	}
}
