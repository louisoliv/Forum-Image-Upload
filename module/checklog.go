package module

import (
	"database/sql"
	"fmt"
	"net/http"
)

func Checklog(w http.ResponseWriter, r *http.Request) (string, string, bool) {
	cookies := r.Cookies()
	if len(cookies) == 0 {
		fmt.Println("here")
		username := "Visiteur"
		filter := ""
		SetCookie(w, username, filter)
		http.Redirect(w, r, "index", http.StatusFound)
		return "", "", false
	}
	if cookies[0].Name == "Visiteur" {
		database.User.Filtered = false
		if cookies[0].Value != "" {
			FilterSwitch(w, r, cookies[0].Name, cookies[0].Value)
		}
		return cookies[0].Name, cookies[0].Value, false
	}
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		handle500(w, err)
		return "", "", false
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		handle500(w, err)
		return "", "", false
	}
	fmt.Println("Connected to database")
	row1 := db.QueryRow("SELECT username, token, filtered, category FROM Users WHERE username=? AND token=?", cookies[0].Name, cookies[0].Value)
	var username, token, category string
	var filtered bool
	err = row1.Scan(&username, &token, &filtered, &category)
	if err != nil {
		if err == sql.ErrNoRows {
			DestroyCookie(w, r)
			http.Redirect(w, r, "index", http.StatusFound)
		} else {
			handle500(w, err)
			return "", "", false
		}
	}
	fmt.Println(filtered, category)
	if filtered {
		FilterSwitch(w, r, cookies[0].Name, category)
	}
	database.User.Filtered = filtered
	return username, token, true
}
