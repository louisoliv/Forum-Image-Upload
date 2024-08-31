package module

import (
	"database/sql"
	"fmt"
	"net/http"
)

func CommentLike(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.FormValue("like"), "here")
	liked := true
	disliked := true
	cookieUser, cookieToken, check := Checklog(w, r)
	fmt.Println(cookieUser, cookieToken, check)
	if !check {
		ErrorMessage2 = "Please log in to like or dislike"
		database.User.ErrorMessage = ErrorMessage2
		http.Redirect(w, r, "index", http.StatusFound)
		return
	}
	if r.Method == "POST" {
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
		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Comments_like (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT,
			token TEXT
		)`)
		if err != nil {
			handle500(w, err)
			return
		}
		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Comments_dislike (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT,
			token TEXT
		)`)
		if err != nil {
			handle500(w, err)
			return
		}
		fmt.Println("Form:", r.FormValue("like"))
		row1 := db.QueryRow("SELECT username, token FROM Comments_like WHERE username=? AND token=?", cookieUser, r.FormValue("like"))
		row2 := db.QueryRow("SELECT username, token FROM Comments_dislike WHERE username=? AND token=?", cookieUser, r.FormValue("like"))
		row3 := db.QueryRow("SELECT like, dislike, token FROM comments WHERE token_comment=?", r.FormValue("like"))
		var username, token, tokenpost string
		var like, dislike int
		err = row1.Scan(&username, &token)
		if err != nil {
			if err == sql.ErrNoRows {
				liked = false
			} else {
				fmt.Println("row1", err)
				handle500(w, err)
				return
			}
		}
		err = row2.Scan(&username, &token)
		if err != nil {
			if err == sql.ErrNoRows {
				disliked = false
			} else {
				fmt.Println("row2", err)
				handle500(w, err)
				return
			}
		}
		err = row3.Scan(&like, &dislike, &tokenpost)
		if err != nil {
			fmt.Println("row3", err)
			handle500(w, err)
			return
		}
		fmt.Println(liked, disliked, like, dislike)
		if !liked && !disliked {
			if r.RequestURI == "/commentlike" {
				like += 1
				_, err = db.Exec(`INSERT INTO Comments_like (username, token)
			VALUES (?, ?)`,
					cookieUser,
					r.FormValue("like"),
				)
				if err != nil {
					handle500(w, err)
					return
				}
				_, err := db.Exec("UPDATE comments SET like = ? WHERE token_comment = ?", like, r.FormValue("like"))
				if err != nil {
					handle500(w, err)
					return
				}
				for i := 0; i < len(database.Comments); i++ {
					if database.Comments[i].TokenID == r.FormValue("like") {
						database.Comments[i].Like = like
					}
				}
			}
			if r.RequestURI == "/commentdislike" {
				dislike += 1
				_, err = db.Exec(`INSERT INTO Comments_dislike (username, token)
			VALUES (?, ?)`,
					cookieUser,
					r.FormValue("like"),
				)
				if err != nil {
					handle500(w, err)
					return
				}
				_, err := db.Exec("UPDATE comments SET dislike = ? WHERE token_comment = ?", dislike, r.FormValue("like"))
				if err != nil {
					handle500(w, err)
					return
				}
				for i := 0; i < len(database.Comments); i++ {
					if database.Comments[i].TokenID == r.FormValue("like") {
						database.Comments[i].Dislike = dislike
					}
				}
			}
		}
		if liked && !disliked {
			like -= 1
			_, err = db.Exec("DELETE FROM Comments_like WHERE username=? AND token=?", cookieUser, r.FormValue("like"))
			if err != nil {
				handle500(w, err)
				return
			}
			_, err = db.Exec("UPDATE comments SET like = ? WHERE token_comment = ?", like, r.FormValue("like"))
			if err != nil {
				fmt.Println(err)
				handle500(w, err)
				return
			}
			if r.RequestURI == "/commentdislike" {
				dislike += 1
				_, err = db.Exec(`INSERT INTO Comments_dislike (username, token)
			VALUES (?, ?)`,
					cookieUser,
					r.FormValue("like"),
				)
				if err != nil {
					handle500(w, err)
					return
				}
				_, err := db.Exec("UPDATE comments SET dislike = ? WHERE token_comment = ?", dislike, r.FormValue("like"))
				if err != nil {
					fmt.Println(err)
					handle500(w, err)
					return
				}
			}
			for i := 0; i < len(database.Comments); i++ {
				if database.Comments[i].TokenID == r.FormValue("like") {
					database.Comments[i].Like = like
					database.Comments[i].Dislike = dislike
				}
			}
		}
		if !liked && disliked {
			dislike -= 1
			_, err = db.Exec("DELETE FROM Comments_dislike WHERE username=? AND token=?", cookieUser, r.FormValue("like"))
			if err != nil {
				handle500(w, err)
				return
			}
			_, err = db.Exec("UPDATE comments SET dislike = ? WHERE token_comment = ?", dislike, r.FormValue("like"))
			if err != nil {
				fmt.Println(err)
				handle500(w, err)
				return
			}
			if r.RequestURI == "/commentlike" {
				like += 1
				_, err = db.Exec(`INSERT INTO Comments_like (username, token)
			VALUES (?, ?)`,
					database.User.Username,
					r.FormValue("like"),
				)
				if err != nil {
					handle500(w, err)
					return
				}
				_, err := db.Exec("UPDATE comments SET like = ? WHERE token_comment = ?", like, r.FormValue("like"))
				if err != nil {
					fmt.Println(err)
					handle500(w, err)
					return
				}
			}
			for i := 0; i < len(database.Comments); i++ {
				if database.Comments[i].TokenID == r.FormValue("like") {
					database.Comments[i].Like = like
					database.Comments[i].Dislike = dislike
				}
			}
		}
		http.Redirect(w, r, "index#"+tokenpost, http.StatusFound)
	}
}
