package module

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
)

func Comment(w http.ResponseWriter, r *http.Request) {
	cookieUser, cookieToken, check := Checklog(w, r)
	fmt.Println(cookieToken)
	if !check {
		ErrorMessage2 = "Merci de vous connecter"
		database.User.ErrorMessage = ErrorMessage2
		http.Redirect(w, r, "index", http.StatusFound)
		return
	}

	// Verification and handle error if the format is not respected
	if r.Method == "POST" {
		fmt.Println(r.FormValue("comment"), r.FormValue("tokenpost"))
		if r.FormValue("comment") == "" {
			ErrorMessage4 = "Merci d'écrire un commentaire"
			database.User.ErrorMessage4 = ErrorMessage4
			http.Redirect(w, r, "index#"+r.FormValue("tokenpost"), http.StatusFound)
			return
		}

		if containsOnlySpecialChars(r.FormValue("comment")) {
			ErrorMessage4 = "Veuillez écrire un commentaire conventionnel svp..."
			database.User.ErrorMessage4 = ErrorMessage4
			http.Redirect(w, r, "index#"+r.FormValue("tokenpost"), http.StatusFound)
			return
		}

		r.ParseForm()

		// fmt.Println("OK")
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
		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS comments (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT,
			date_time TEXT,
			token TEXT,
			message TEXT,
			token_comment TEXT,
			like INTEGER,
			dislike INTEGER,
			image TEXT
		)`)
		if err != nil {
			handle500(w, err)
			return
		}

		// Call of the UploadImages to import and check the file format
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

		token := GenerateToken()
		// Insert data into table
		_, err = db.Exec(`INSERT INTO comments (username, date_time, token, message, token_comment, like, dislike, image)
			VALUES (?, ?, ?, ?, ?, ?, ?,?)`,
			cookieUser,
			time.Now().Format("2006-01-02 15:04:05"),
			r.FormValue("tokenpost"),
			r.FormValue("comment"),
			token,
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
		_, err = db.Exec("UPDATE Users SET Messages = ? WHERE username = ?", messages, cookieUser)
		if err != nil {
			handle500(w, err)
			return
		}
		var lastInsertID int
		err = db.QueryRow("SELECT last_insert_rowid()").Scan(&lastInsertID)
		if err != nil {
			fmt.Println("Error retrieving last insert ID:", err)
			return
		}
		for i := 0; i < len(database.User_public); i++ {
			if database.User_public[i].Username == cookieUser {
				database.User_public[i].Messages = messages
				break
			}
		}
		// Construct the Posts struct with the new entry
		newComment := Comments{
			Id:              lastInsertID,
			Username:        cookieUser,
			Date:            time.Now().Format("2006-01-02 15:04:05"),
			TokenComment:    r.FormValue("tokenpost"),
			Message_comment: r.FormValue("comment"),
			TokenID:         token,
			Like:            0,
			Dislike:         0,
			Image_comment:   img,
		}
		database.Comments = append(database.Comments, newComment)
		database.User.ErrorMessage = ""
		// Call Index function or handle response as needed
		http.Redirect(w, r, "index#"+r.FormValue("tokenpost"), http.StatusFound)
	}
}
