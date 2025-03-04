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
	selection := optionlist.GetOption([]string{"Add expense", "List expenses", "Delete expense"})

	actions.CleanScreen()
	//Process user option
	switch selection {
	case "Add expense":
		actions.AddExpense(db)
	case "List expenses":
		actions.ListExpenses(db)
	case "Delete expense":
		actions.DeleteExpense(db)
	default:
		fmt.Println("No option selected!")
		return
	}

}
