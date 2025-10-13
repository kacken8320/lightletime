package sql

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB(path string) *sql.DB {
	db, _ := sql.Open("sqlite3", path)
	db.Exec(`CREATE TABLE IF NOT EXISTS category (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			value TEXT
		)`)
	return db
}

func InsertCategory(db *sql.DB, value string) {
	db.Exec(`INSERT INTO category (value) VALUES (?)`, value)
}
