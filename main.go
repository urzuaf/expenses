package main

import (
	database "expenses/db"
	"fmt"
	"log"

	"expenses/ui/optionlist"
	"expenses/ui/textinput"
)

func cleanScreen() {
	fmt.Print("\033[H\033[2J")
}

func assertNotEmpty(str string) {
	if str == "" {
		log.Fatal("Please, enter a value")
	}
}

func main() {
	//Open and close db
	db := database.Connect()
	defer db.Close()

	//Get user option
	selection := optionlist.GetOption([]string{"Add expense", "List expenses", "Delete expense"})
	cleanScreen()
	switch selection {
	case "Add expense":
		newExpense := database.Expense{}
		newExpense.Description = textinput.GetInput("Description of expense", "bus ticket")
		assertNotEmpty(newExpense.Description)
		newExpense.Category = textinput.GetInput("Category of expense", "transport")
		assertNotEmpty(newExpense.Category)
		newExpense.Amount = textinput.GetInput("Amount of expense", "10000")
		assertNotEmpty(newExpense.Amount)
		newExpense.Date = textinput.GetInput("Date of expense", "2023-03-03")
		assertNotEmpty(newExpense.Date)
		database.InsertData(db, newExpense)
		cleanScreen()
		fmt.Println("Expense added successfully")
	case "List expenses":
		fmt.Println("Listing expenses...")
		database.GetExpenses(db)
	case "Delete expense":
		fmt.Println("Listing expenses...")
		database.GetExpenses(db)
		id := textinput.GetInput("Select the ID you want to delete", "")
		database.DeleteData(db, id)
		fmt.Println("Expense deleted successfully")
	default:
		fmt.Println("No option selected!")
		return
	}

}
