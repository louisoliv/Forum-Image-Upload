package module

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
)

var ErrorMessage2 string

func containsOnlySpecialChars(s string) bool {
	for _, char := range s {
		if char != ' ' && char != '\t' && char != '\\' {
			return false
		}
	}
	return true
}

func Post(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		cookieUser, cookieToken, check := Checklog(w, r)
		fmt.Println(cookieToken)
		if !check {
			ErrorMessage2 = "Merci de vous connecter svp"
			database.User.ErrorMessage = ErrorMessage2
			http.Redirect(w, r, "index", http.StatusFound)
			return
		}

		// fmt.Println(r.Form.Get("message"))

		if containsOnlySpecialChars(r.FormValue("message")) {
			fmt.Println(r.FormValue("message"))
			ErrorMessage2 = "Merci de rédiger un post conventionnel"
			database.User.ErrorMessage = ErrorMessage2
			http.Redirect(w, r, "index", http.StatusFound)
			return
		}

		if len(r.Form["category[]"]) == 0 || r.FormValue("message") == "" {
			ErrorMessage2 = "Merci de remplir tous les critères"
			database.User.ErrorMessage = ErrorMessage2
			http.Redirect(w, r, "index", http.StatusFound)
			return
		}

		r.ParseForm()

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
		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS posts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT,
			date_time TEXT,
			token TEXT,
			message TEXT,
			golang BOOLEAN,
			javascript BOOLEAN,
			python BOOLEAN,
			rust BOOLEAN,
			html_css BOOLEAN,
			angular BOOLEAN,
			autre BOOLEAN,
			like INT,
			dislike INT,
			image TEXT
		)`)
		if err != nil {
			handle500(w, err)
			return
		}

		var img string
		filePath := "./styles/img_Upload"
		if _, _, err := r.FormFile("image"); err == nil {
			img, err = UploadImages(w, r, filePath)
			if err != nil {
				// Handle the error returned by UploadImages
				http.Redirect(w, r, "index", http.StatusFound)
				return
			}
		}

		// Generate unique token
		token := GenerateToken()
		if previousMessage != r.FormValue("message") {
			_, err = db.Exec(`INSERT INTO posts (username, date_time, token, message, golang, javascript, python, rust, html_css, angular, autre, like, dislike, image)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
				cookieUser,
				time.Now().Format("2006-01-02 15:04:05"),
				token,
				r.FormValue("message"),
				isCategoryPresent("Golang", r.Form["category[]"]),
				isCategoryPresent("Javascript", r.Form["category[]"]),
				isCategoryPresent("Python", r.Form["category[]"]),
				isCategoryPresent("Rust", r.Form["category[]"]),
				isCategoryPresent("HTML/CSS", r.Form["category[]"]),
				isCategoryPresent("Angular", r.Form["category[]"]),
				isCategoryPresent("Autre", r.Form["category[]"]),
				0,
				0,
				img,
			)
			if err != nil {
				handle500(w, err)
				return
			}

			row1 := db.QueryRow("SELECT Messages FROM Users WHERE username = ?", cookieUser)
			var messages int
			err = row1.Scan(&messages)
			if err != nil {
				handle500(w, err)
				return
			}
			messages += 1
			_, err := db.Exec("UPDATE Users SET Messages = ? WHERE username = ?", messages, cookieUser)
			if err != nil {
				fmt.Println(err)
				handle500(w, err)
				return
			}
			var lastInsertID int
			err = db.QueryRow("SELECT last_insert_rowid()").Scan(&lastInsertID)
			if err != nil {
				fmt.Println("Error retrieving last insert ID:", err)
				handle500(w, err)
				return
			}
			for i := 0; i < len(database.User_public); i++ {
				if database.User_public[i].Username == cookieUser {
					database.User_public[i].Messages = messages
					break
				}
			}
			// Construct the Posts struct with the new entry
			newPost := Posts{
				Id:         lastInsertID,
				Username:   cookieUser,
				Date:       time.Now().Format("2006-01-02 15:04:05"),
				Token:      token,
				Message:    r.FormValue("message"),
				Golang:     isCategoryPresent("Golang", r.Form["category[]"]),
				JavaScript: isCategoryPresent("Javascript", r.Form["category[]"]),
				Python:     isCategoryPresent("Python", r.Form["category[]"]),
				Rust:       isCategoryPresent("Rust", r.Form["category[]"]),
				HTML_CSS:   isCategoryPresent("HTML/CSS", r.Form["category[]"]),
				Angular:    isCategoryPresent("Angular", r.Form["category[]"]),
				Autre:      isCategoryPresent("Autre", r.Form["category[]"]),
				Like:       0,
				Dislike:    0,
				Image:      img,
			}
			database.Posts = append([]Posts{newPost}, database.Posts...)
			database.User.ErrorMessage = ""
			previousMessage = r.FormValue("message")
			// previousUsername = r.FormValue("username")
		} else {
			ErrorMessage2 = "Message already sent"
			database.User.ErrorMessage = ErrorMessage2
			http.Redirect(w, r, "index", http.StatusFound)
			return
		}
		// Call Index function or handle response as needed
		http.Redirect(w, r, "index", http.StatusFound)
	}
}
