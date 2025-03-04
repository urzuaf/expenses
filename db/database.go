package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3" // Importación del driver
)

type Expense struct {
	Description string
	Category    string
	Amount      string
	Date        string
}

func Connect() *sql.DB {
	// Abre o crea un archivo SQLite local
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		log.Fatal(err)
	}
	createMainTable(db)
	return db
}

func createMainTable(db *sql.DB) {
	// create program main table if it doesn't exist
	query := `
    CREATE TABLE IF NOT EXISTS expenses(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        description TEXT NOT NULL,
		category TEXT NOT NULL,
		amount INTEGER NOT NULL,
		date TEXT NOT NULL
		);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func InsertData(db *sql.DB, ex Expense) {
	// Insert data
	_, err := db.Exec("INSERT INTO expenses(description,category,amount,date) VALUES (?,?,?,?)", ex.Description, ex.Category, ex.Amount, ex.Date)
	if err != nil {
		log.Fatal(err)
	}
}
func DeleteData(db *sql.DB, id string) {
	// Insert data
	_, err := db.Exec("DELETE FROM expenses WHERE id=?", id)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateData(db *sql.DB, id string, ex Expense) {
	// Insert data
	_, err := db.Exec("UPDATE expenses SET description=?,category=?,amount=?,date=? WHERE id=?", ex.Description, ex.Category, ex.Amount, ex.Date, id)
	if err != nil {
		log.Fatal(err)
	}
}

func GetExpenses(db *sql.DB) {
	//Read Data
	rows, err := db.Query("SELECT * FROM expenses")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var description string
		var category string
		var amount int
		var date string
		rows.Scan(&id, &description, &category, &amount, &date)
		fmt.Printf("Expense ID: %d - Description: %s - Category: %s - Amount: %d - %s\n", id, description, category, amount, date)
	}
}
