package main
import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "./database.db")
	defer db.Close()

	createQuery := `
		CREATE TABLE IF NOT EXISTS category (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			value TEXT
		);`
	db.Exec(createQuery)

	insertQuery := `INSERT INTO category (value) VALUES (?)`
	db.Exec(insertQuery, "example_value");
}
