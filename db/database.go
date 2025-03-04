package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3" // Importación del driver
)

func main() {
	// Abre o crea un archivo SQLite local
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Crear una tabla si no existe
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL
    );`
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	// Insertar datos de prueba
	_, err = db.Exec("INSERT INTO users (name) VALUES (?)", "Fernando")
	if err != nil {
		log.Fatal(err)
	}

	// Leer los datos
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		fmt.Printf("User: %d - %s\n", id, name)
	}
}
