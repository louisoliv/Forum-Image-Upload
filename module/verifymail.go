package module

import (
	"database/sql"
	"fmt"
	"net/http"
)

func VerifyMail(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		fmt.Println(err)
		handle500(w, err)
		return
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		handle500(w, err)
		return
	}
	row := db.QueryRow("SELECT Verified FROM Users WHERE Verification_token = ?", token)
	var verified bool
	err = row.Scan(&verified)
	if err != nil {
		fmt.Println(err)
		return
	}
	verified = true
	_, err = db.Exec("UPDATE Users SET Verified = ? WHERE Verification_token = ?", verified, token)
	if err != nil {
		fmt.Println(err)
		handle500(w, err)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
