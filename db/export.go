package db

import (
	"database/sql"
	"fmt"
	"log"
)

// TODO: Verify that data doesn´t have ","
func GetExpensesCSV(db *sql.DB) []string {
	//Read Data
	//We ignore the id for exporting
	lines := []string{}
	lines = append(lines, "description,category,amount,date")
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

		line := fmt.Sprintf("%s,%s,%d,%s", description, category, amount, date)
		lines = append(lines, line)
	}
	return lines
}
