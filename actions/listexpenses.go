package actions

import (
	"database/sql"
	database "expenses/db"
	"expenses/ui/tablelist"
	"strings"

	"github.com/charmbracelet/bubbles/table"
)

func ListExpenses(db *sql.DB) string {
	var tableHeader []table.Column
	var tableRows []table.Row

	//Create table header
	tableHeader = []table.Column{
		{Title: "ID", Width: 4},
		{Title: "Description", Width: 14},
		{Title: "Category", Width: 10},
		{Title: "Amount", Width: 10},
		{Title: "Date", Width: 10},
	}

	//Got every expense as a string
	expenses := database.GetExpenses(db, false)

	//Parse it into table format
	for _, expense := range expenses {
		items := strings.Split(expense, "\\")
		var row table.Row
		for _, item := range items {
			row = append(row, item)
		}
		tableRows = append(tableRows, row)
	}

	//Display table
	choice := tablelist.DisplayTable(tableHeader, tableRows)
	return choice
}
