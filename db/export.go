package db

import (
	"database/sql"
	"fmt"
	"log"
)

// TODO: Verify that data doesn´t have ","
func GetExpensesCSV(db *sql.DB) []string {
	//Read Data
	lines := []string{}
	lines = append(lines, "id,description,category,amount,date")
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

		line := fmt.Sprintf("%d,%s,%s,%d,%s", id, description, category, amount, date)
		lines = append(lines, line)
	}
	return lines
}
