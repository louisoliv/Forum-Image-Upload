package module

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"
)

func PublicUserPageModify(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.FormValue("user"))
	t, err := template.ParseFiles("./templates/user_public_profile_modify.html")
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
	row := db.QueryRow("SELECT id, username, first_name, last_name, email, password FROM Users WHERE id = ?", r.FormValue("user"))
	// Scan the row data into the user_data struct
	var userdata User_info
	if err := row.Scan(&userdata.Id, &userdata.Username, &userdata.Firstname, &userdata.Lastname, &userdata.Email, &userdata.Password); err != nil {
		if err != sql.ErrNoRows {
			fmt.Println(err)
			handle500(w, err)
			return
		}
	}
	database.User = userdata
	err = t.Execute(w, database)
	if err != nil {
		fmt.Println(err)
		handle500(w, err)
		return
	}
}
