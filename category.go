package main

type Category struct {
	ID			int		`json:"id"`
	Name			string		`json:"name"`
	Multiplier		float64		`json:"multiplier"`
	TimeSpentInMinutes	int		`json:"time_spent_in_minutes"`
	MinimalTimeInSeconds	int		`json:"minimal_time_in_seconds"`
	ParentID		int		`json:"parent_id"`
	IsActivity		bool		`json:"is_activity"`
}
