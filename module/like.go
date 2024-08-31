package module

import (
	"database/sql"
	"fmt"
	"net/http"
)

func Like(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	liked := true
	disliked := true
	cookieUser, cookieToken, check := Checklog(w, r)
	fmt.Println(cookieUser, cookieToken, check)
	if !check {
		ErrorMessage3 = "Please log in to like or dislike"
		database.User.ErrorMessage3 = ErrorMessage3
		database.User.TokenError = r.FormValue("like")
		http.Redirect(w, r, r.FormValue("from"), http.StatusFound)
		return
	}
	if r.Method == "POST" {
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
		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS like (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT,
			token TEXT
		)`)
		if err != nil {
			fmt.Println(err)
			handle500(w, err)
			return
		}
		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS dislike (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT,
			token TEXT
		)`)
		if err != nil {
			fmt.Println(err)
			handle500(w, err)
			return
		}
		row1 := db.QueryRow("SELECT username, token FROM like WHERE username=? AND token=?", cookieUser, r.FormValue("like"))
		row2 := db.QueryRow("SELECT username, token FROM dislike WHERE username=? AND token=?", cookieUser, r.FormValue("like"))
		row3 := db.QueryRow("SELECT like, dislike FROM posts WHERE token=?", r.FormValue("like"))
		var username, token string
		var like, dislike int
		err = row1.Scan(&username, &token)
		if err != nil {
			if err == sql.ErrNoRows {
				liked = false
			} else {
				fmt.Println(err, "here")
				handle500(w, err)
				return
			}
		}
		err = row2.Scan(&username, &token)
		if err != nil {
			if err == sql.ErrNoRows {
				disliked = false
			} else {
				fmt.Println(err, "here1")
				handle500(w, err)
				return
			}
		}
		err = row3.Scan(&like, &dislike)
		if err != nil {
			fmt.Println(err, "here2")
			handle500(w, err)
			return
		}
		fmt.Println(liked, disliked, like, dislike)
		if !liked && !disliked {
			if r.RequestURI == "/like" {
				like += 1
				_, err = db.Exec(`INSERT INTO like (username, token)
			VALUES (?, ?)`,
					cookieUser,
					r.FormValue("like"),
				)
				if err != nil {
					handle500(w, err)
					return
				}
				_, err := db.Exec("UPDATE posts SET like = ? WHERE token = ?", like, r.FormValue("like"))
				if err != nil {
					fmt.Println(err)
					handle500(w, err)
					return
				}
				for i := 0; i < len(database.Posts); i++ {
					if database.Posts[i].Token == r.FormValue("like") {
						database.Posts[i].Like = like
					}
				}
				newLike := LikE{
					Username: cookieUser,
					Token:    r.FormValue("like"),
				}
				database.LikE = append(database.LikE, newLike)
				fmt.Println(database.LikE)
			}
			if r.RequestURI == "/dislike" {
				dislike += 1
				_, err = db.Exec(`INSERT INTO dislike (username, token)
			VALUES (?, ?)`,
					cookieUser,
					r.FormValue("like"),
				)
				if err != nil {
					handle500(w, err)
					return
				}
				_, err := db.Exec("UPDATE posts SET dislike = ? WHERE token = ?", dislike, r.FormValue("like"))
				if err != nil {
					fmt.Println(err)
					handle500(w, err)
					return
				}
				for i := 0; i < len(database.Posts); i++ {
					if database.Posts[i].Token == r.FormValue("like") {
						database.Posts[i].Dislike = dislike
					}
				}
			}
		}
		if liked && !disliked {
			like -= 1
			_, err = db.Exec("DELETE FROM like WHERE username=? AND token=?", cookieUser, r.FormValue("like"))
			if err != nil {
				handle500(w, err)
				return
			}
			_, err = db.Exec("UPDATE posts SET like = ? WHERE token = ?", like, r.FormValue("like"))
			if err != nil {
				fmt.Println(err)
				handle500(w, err)
				return
			}
			for i := 0; i < len(database.LikE); i++ {
				if database.LikE[i].Username == username && database.LikE[i].Token == r.FormValue("like") {
					database.LikE = append(database.LikE[:i], database.LikE[i+1:]...)
				}
			}
			fmt.Println(database.LikE)
			if r.RequestURI == "/dislike" {
				dislike += 1
				_, err = db.Exec(`INSERT INTO dislike (username, token)
			VALUES (?, ?)`,
					cookieUser,
					r.FormValue("like"),
				)
				if err != nil {
					handle500(w, err)
					return
				}
				_, err := db.Exec("UPDATE posts SET dislike = ? WHERE token = ?", dislike, r.FormValue("like"))
				if err != nil {
					fmt.Println(err)
					handle500(w, err)
					return
				}
			}
			for i := 0; i < len(database.Posts); i++ {
				if database.Posts[i].Token == r.FormValue("like") {
					database.Posts[i].Like = like
					database.Posts[i].Dislike = dislike
				}
			}
		}
		if !liked && disliked {
			dislike -= 1
			_, err = db.Exec("DELETE FROM dislike WHERE username=? AND token=?", cookieUser, r.FormValue("like"))
			if err != nil {
				handle500(w, err)
				return
			}
			_, err = db.Exec("UPDATE posts SET dislike = ? WHERE token = ?", dislike, r.FormValue("like"))
			if err != nil {
				fmt.Println(err)
				handle500(w, err)
				return
			}
			if r.RequestURI == "/like" {
				like += 1
				_, err = db.Exec(`INSERT INTO like (username, token)
			VALUES (?, ?)`,
					database.User.Username,
					r.FormValue("like"),
				)
				if err != nil {
					panic(err)
				}
				_, err := db.Exec("UPDATE posts SET like = ? WHERE token = ?", like, r.FormValue("like"))
				if err != nil {
					fmt.Println(err)
					handle500(w, err)
					return
				}
				newLike := LikE{
					Username: cookieUser,
					Token:    r.FormValue("like"),
				}
				database.LikE = append(database.LikE, newLike)
				fmt.Println(database.LikE)
			}
			for i := 0; i < len(database.Posts); i++ {
				if database.Posts[i].Token == r.FormValue("like") {
					database.Posts[i].Like = like
					database.Posts[i].Dislike = dislike
				}
			}
		}
		http.Redirect(w, r, r.FormValue("from"), http.StatusFound)
	}
}
