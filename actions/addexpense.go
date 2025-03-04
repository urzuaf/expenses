package actions

import (
	"database/sql"
	database "expenses/db"
	"expenses/ui/optionlist"
	"expenses/ui/textinput"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"
)

func CleanScreen() {
	fmt.Print("\033[H\033[2J")
}

func assertNotEmpty(str string) {
	if str == "" {
		log.Fatal("Please, enter a value")
	}
}

func assertNumber(str string) {
	if _, err := strconv.Atoi(str); err != nil {
		log.Fatal("Please, enter a valid number")
	}
}

func assertDate(str string) {
	re := regexp.MustCompile(`^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01])$`)
	if !re.MatchString(str) {
		log.Fatal("Please, enter a valid date")
	}

}

func AddExpense(db *sql.DB) {
	newExpense := database.Expense{}

	newExpense.Amount = textinput.GetInput("Amount of expense", "10000")
	assertNumber(newExpense.Amount)

	newExpense.Description = textinput.GetInput("Description of expense", "bus ticket")
	assertNotEmpty(newExpense.Description)

	newExpense.Category = optionlist.GetOption([]string{"Food", "Entertainment", "Transport", "Health", "Savings", "Gifts", "Travel", "Education", "Other"})
	assertNotEmpty(newExpense.Category)

	newExpense.Date = textinput.GetInput("Date of expense (yyyy-mm-dd)", time.Now().Format("2006-01-02"))
	assertDate(newExpense.Date)

	database.InsertData(db, newExpense)
	fmt.Println("Expense ", newExpense.Description, " added successfully")
}

func ModifyExpense(db *sql.DB) {

	//id to edit
	id := ListExpenses(db)

	newExpense := database.Expense{}

	newExpense.Amount = textinput.GetInput("Amount of expense", "10000")
	assertNumber(newExpense.Amount)

	newExpense.Description = textinput.GetInput("Description of expense", "bus ticket")
	assertNotEmpty(newExpense.Description)

	newExpense.Category = optionlist.GetOption([]string{"Food", "Entertainment", "Transport", "Health", "Savings", "Gifts", "Travel", "Education", "Other"})
	assertNotEmpty(newExpense.Category)

	newExpense.Date = textinput.GetInput("Date of expense (yyyy-mm-dd)", time.Now().Format("2006-01-02"))
	assertDate(newExpense.Date)

	database.EditExpense(db, newExpense, id)
	fmt.Println("Expense ", newExpense.Description, " modified successfully")
}
