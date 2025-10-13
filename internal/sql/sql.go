package sql

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB(path string) *sql.DB {
	db, _ := sql.Open("sqlite3", path)
	db.Exec(`CREATE TABLE IF NOT EXISTS example_table (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			example TEXT
		)`)
	return db
}

func InsertExample(db *sql.DB, value string) {
	db.Exec(`INSERT INTO example_table (example) VALUES (?)`, value)
}
