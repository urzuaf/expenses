package actions

import (
	"database/sql"
	database "expenses/db"
	barchart "expenses/ui/barcharts"
)

func GetStats(db *sql.DB) {
	categoryStats := database.GetCategoryStats(db)
	barchart.DisplayBarchart("Expenses per category", categoryStats)

}
