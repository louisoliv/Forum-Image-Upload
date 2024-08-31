package module

import (
	"database/sql"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// var currentUser User_info

// Check if username/password from form is corresponding to db row
func Log(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		if r.Form["log[]"] == nil {
			Error2.ErrorMessage = "Please input your username and password"
			Error = append(Error, Error2)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
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
		fmt.Println("Connected to database")
		usernameCheck := r.Form["log[]"][0]
		passwordCheck := r.Form["log[]"][1]
		row1 := db.QueryRow("SELECT username, password, Verified FROM Users WHERE username=?", usernameCheck)
		var username, password string
		var verified bool
		err = row1.Scan(&username, &password, &verified)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("Username not found")
				Error2.ErrorMessage = "Wrong Username or Password"
				Error = append(Error, Error2)
				HomeLog(w, r)
				return
			} else {
				fmt.Println(err)
				handle500(w, err)
				return
			}
		}
		if !verified {
			Error2.ErrorMessage = "Please valid your mail address by clicking the link sended"
			Error = append(Error, Error2)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		fmt.Println("Username found")
		err = bcrypt.CompareHashAndPassword([]byte(password), []byte(passwordCheck))
		if err != nil {
			fmt.Println("Wrong password")
			Error2.ErrorMessage = "Wrong Username or Password"
			Error = append(Error, Error2)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		} else {
			filtered := false
			category := ""
			DestroyCookie(w, r)
			token := GenerateToken()
			SetCookie(w, username, token)
			_, err := db.Exec("UPDATE Users SET token = ?, filtered = ?, category = ? WHERE username = ?", token, filtered, category, username)
			if err != nil {
				fmt.Println(err)
				handle500(w, err)
				return
			}
			// Get the user ID
			rowID := db.QueryRow("SELECT id token FROM Users WHERE username=?", username)
			var id int
			err = rowID.Scan(&id)
			if err != nil {
				fmt.Println(err)
				handle500(w, err)
				return
			}
			// Store the id and username in the UserInfo struct
			database.User.Id = id
			database.User.Username = username
			database.User.Token = token
			database.User.Filtered = filtered
			database.User.Category = category

			// fmt.Println(currentUser)
			fmt.Println(r.RequestURI, r.Method)
			http.Redirect(w, r, "index", http.StatusFound)
		}
	}
}
