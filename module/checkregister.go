package module

import (
	"database/sql"
	"fmt"
	"net/http"
)

// Manage register form, check if username and email are already existing in db
func CheckRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		fmt.Println(r.Form["register[]"])
		for i := 0; i < len(r.Form["register[]"]); i++ {
			if r.Form["register[]"][i] == "" {
				Error2.ErrorMessage = "Please complete the form"
				Error = append(Error, Error2)
				RegisterPage(w, r)
			}
		}
		if r.Form["register[]"][4] != r.Form["register[]"][5] {
			Error2.ErrorMessage = "Password confirmation isn't correct"
			Error = append(Error, Error2)
			RegisterPage(w, r)
			return
		}
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
		usernameCheck := r.Form["register[]"][0]
		mailCheck := r.Form["register[]"][3]

		rowUsername := db.QueryRow("SELECT username FROM Users WHERE username=?", usernameCheck)
		rowMail := db.QueryRow("SELECT email FROM Users WHERE email=?", mailCheck)

		var username, email string

		err1 := rowUsername.Scan(&username)
		if err1 == nil {
			Error2.ErrorMessage = "Username already used"
			Error = append(Error, Error2)
			RegisterPage(w, r)
			return
		} else {
			if err1 == sql.ErrNoRows {
				goto next
			} else {
				handle500(w, err)
				return
			}
		}
	next:
		err2 := rowMail.Scan(&email)
		if err2 == nil {
			Error2.ErrorMessage = "Mail already used"
			Error = append(Error, Error2)
			RegisterPage(w, r)
			return
		} else {
			if err2 == sql.ErrNoRows {
				AddtoDB(w, r.Form["register[]"])
			} else {
				handle500(w, err)
				return
			}
		}
		http.Redirect(w, r, "#.Mail_Verif", http.StatusFound)
	}
}
