package handler

import (
	"database/sql"
	"net/http"
	dbpkg "lightletime/internal/sql"
)

func DepositHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		v := r.URL.Query().Get("value")
		dbpkg.InsertCategory(db, v)
		w.Write([]byte("ok"))
	}
}
