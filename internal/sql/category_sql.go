package sql

import (
)

type Category struct {
	ID			int	`db:"id"`
	UserID			int	`db:"user_id"`
	ParentCategoryID	int	`db:"parent_category_id"`	//=0 if no parent exists
	Name			string	`db:"name"`
	Multiplier		float64	`db:"multiplier"`
	TimeSpentInMins		int	`db:"time_spent_in_mins"`
	MinimalTimeInMins	int	`db:"minimal_time_in_mins"`	//=0 if not an activity
	SkipCounter		int	`db:"skip_counter"`		//=0 if not an activity
	IsActivity		bool	`db:"isi_activity"`
}

func CreateCategoryTable() *sql.DB {
	db.Exec(`CREATE TABLE IF NOT EXISTS category (
			category_id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER,
			parent_category_id INTEGER,
			name TEXT,
			multiplier float,
			time_spent_in_minutes INTEGER,
			minimal_time_in_minutes INTEGER,
			is_activity BOOLEAN)`)
}

func InsertCategory(db *sql.DB, userID int, parentCategoryID int, name string, multiplier float64, timeSpentInMinutes int, minimalTimeInMinutes int, isActivity bool) {
	db.Exec(`INSERT INTO category (
			user_id,
			parent_category_id,
			name,
			multiplier,
			time_spent_in_minutes,
			minimal_time_in_minutes,
			skip_counter,
			is_activity) VALUES (?)`)
}
