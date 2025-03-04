package db

import (
	"database/sql"
	"fmt"
	"log"
)

func GetDateStats(db *sql.DB) []string {
	//Read Data
	var lines []string
	rows, err := db.Query("SELECT date, SUM(amount) FROM expenses GROUP BY date")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var date string
		var amount int
		rows.Scan(&date, &amount)

		line := fmt.Sprintf("%s %d", date, amount)
		lines = append(lines, line)
	}

	return lines
}
func GetCategoryStats(db *sql.DB) map[string]int {
	//Read Data
	stats := make(map[string]int)
	rows, err := db.Query("SELECT category, SUM(amount) FROM expenses GROUP BY category")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var category string
		var amount int
		rows.Scan(&category, &amount)

		stats[category] = amount
	}

	return stats
}
