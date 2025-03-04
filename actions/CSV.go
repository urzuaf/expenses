package actions

import (
	"database/sql"
	"encoding/csv"
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
	fpath, path := filepicker.GetFile([]string{".csv"})
	fmt.Println("Importing data from file: " + fpath)
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("could open file")
		log.Fatal(err)
	}
	defer file.Close()

	// Read the CSV data
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	for index, row := range data {
		//First line are headers
		if index == 0 {
			continue
		}

		//Insert data
		ex := database.Expense{Description: row[0], Category: row[1], Amount: row[2], Date: row[3]}
		database.InsertData(db, ex)
	}
	fmt.Println("Data imported successfully")

}
