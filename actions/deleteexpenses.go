package actions

import (
	"database/sql"
	database "expenses/db"
	"fmt"
)

func DeleteExpense(db *sql.DB) {
	id := ListExpenses(db)
	database.DeleteData(db, id)
	fmt.Printf("Expense %s deleted successfully\n", id)
}
