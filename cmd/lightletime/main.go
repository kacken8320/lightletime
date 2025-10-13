package main
import (
	"net/http"
	"fmt"
	"lightletime/api/handler"
	"lightletime/internal/sql"
)

func main() {
	db := sql.InitDB("./database.db")
	defer db.Close()
	fmt.Println("db initialized")

	fmt.Println("listening on :8080...")
	http.HandleFunc("/deposit", handler.DepositHandler(db))
	http.ListenAndServe(":8080", nil)
}
