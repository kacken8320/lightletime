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
)

func CreateCategoryTable() {
