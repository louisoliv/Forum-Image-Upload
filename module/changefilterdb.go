package module

import (
	"database/sql"
	"fmt"
	"net/http"
)

func ChangeFilterDB(w http.ResponseWriter, username, category string, filtered bool) {
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		handle500(w, err)
		return
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		handle500(w, err)
		return
	}
	fmt.Println("Connected to database")
	_, err = db.Exec("UPDATE Users SET filtered = ?, category = ? WHERE username = ?", filtered, category, username)
	if err != nil {
		handle500(w, err)
		return
	}
}
