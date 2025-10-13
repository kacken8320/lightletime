package handler

import (
	"database/sql"
	"net/http"
	dbpkg "lightletime/internal/sql"
)

func InsertHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		v := r.URL.Query().Get("example")
		dbpkg.InsertExample(db, v)
		w.Write([]byte("ok"))
	}
}
