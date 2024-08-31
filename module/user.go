package module

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/mattn/go-sqlite3"
)

// Load forum index
func User(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	t, err := template.ParseFiles("./templates/user.html")
	if err != nil {
		fmt.Println(err)
		handle500(w, err)
		return
	}
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		fmt.Println(err)
		handle500(w, err)
		return
	}
	defer db.Close()
	// Get the user ID from the session or other authentication mechanism
	// Retrieve user data from the database
	cookieUser, cookieToken, check := Checklog(w, r)
	if !check {
		http.Redirect(w, r, "index", http.StatusFound)
		return
	}
	row := db.QueryRow("SELECT id, username, first_name, last_name, email FROM Users WHERE username = ? AND token = ?", cookieUser, cookieToken)
	// Scan the row data into the user_data struct
	var id int
	var username, first_name, last_name, email string

	if err := row.Scan(&id, &username, &first_name, &last_name, &email); err != nil {
		if err != sql.ErrNoRows {
			http.Redirect(w, r, "index", http.StatusFound)
			return
		}
	}

	fmt.Println(cookieUser, cookieToken, check)
	database.User.Username = cookieUser
	database.User.Firstname = first_name
	database.User.Lastname = last_name
	database.User.Email = email
	database.User.Token = cookieToken
	err = t.Execute(w, database)
	if err != nil {
		fmt.Println(err)
		handle500(w, err)
		return
	}
}

// }
