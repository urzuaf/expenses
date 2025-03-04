package main

import (
	"expenses/actions"
	database "expenses/db"
	"fmt"

	"expenses/ui/optionlist"
)

func main() {
	//Open and close db
	db := database.Connect()
	defer db.Close()

	//Get user option
	selection := optionlist.GetOption([]string{"List expenses", "Add expense", "Modify expense", "Delete expense", "Export expenses", "Import expenses"})

	actions.CleanScreen()

	//Process user option
	switch selection {
	case "List expenses":
		actions.ListExpenses(db)
	case "Add expense":
		actions.AddExpense(db)
	case "Modify expense":
		actions.ModifyExpense(db)
	case "Delete expense":
		actions.DeleteExpense(db)
	case "Export expenses":
		actions.ExportToCSV(db)
	case "Import expenses":
		actions.ImportFromCSV(db)
	default:
		fmt.Println("No option selected!")
		return
	}

}
