package handler

import (
	"database/sql"
	"net/http"
	dbpkg "lightletime/internal/sql"
)

func InsertCategory(w http.ResponseWriter, r *http.Request) {
	var c sql.Category;

	// decode json
	err := json.NewDecoder(r.Body).Decode(&u);
	if err != nil {
		log.Fatal(err)
	}

	sql.InsertCategory(c.userID, c.parentCategoryID, c.name, c.multiplier, c.timeSpentInMinutes, c.minimalTimeInMinutes, c.isActivity);

	w.WriteHeader(http.StatusCreated);
	w.Write([]byte("category added (in: category_handler.goi"))
}
