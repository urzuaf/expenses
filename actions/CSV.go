package actions

import (
	"database/sql"
	database "expenses/db"
	"expenses/ui/filepicker"
	"expenses/ui/textinput"
	"fmt"
	"log"
	"os"
	"strings"
)

func ExportToCSV(db *sql.DB) {
	lines := database.GetExpensesCSV(db)

	path := textinput.GetInput("Enter the output path", "./expenses.csv")
	csvContent := strings.Join(lines, "\n")
	err := os.WriteFile(path, []byte(csvContent), 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func ImportFromCSV(db *sql.DB) {
	path := filepicker.GetFile([]string{".csv"})
	fmt.Println(path)

}
