package module

import (
	"database/sql"
	"fmt"
	"net/http"
)

func Filter(w http.ResponseWriter, r *http.Request) {
	cookieUser, cookieToken, check := Checklog(w, r)
	fmt.Println(cookieUser, cookieToken, check)
	category := r.URL.Query().Get("category")
	filtered := true
	if cookieUser != "Visiteur" {
		if category == "" {
			filtered = false
			ChangeFilterDB(w, cookieUser, category, filtered)
			database.User.Filtered = false
			http.Redirect(w, r, "index", http.StatusFound)
			return
		}
		ChangeFilterDB(w, cookieUser, category, filtered)
		database.User.Filtered = true
		database.User.Category = category
		database.FilteredPosts = nil
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
		fmt.Println(r.RequestURI, r.FormValue("filter"), category)
		stmt, err := db.Prepare(fmt.Sprintf("SELECT * FROM posts WHERE %s = 1", category))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		// Execute the query
		rows, err := stmt.Query()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		for rows.Next() {
			var post FilteredPosts
			err := rows.Scan(&post.Id, &post.Username, &post.Date, &post.Token, &post.Message, &post.Golang, &post.JavaScript, &post.Python, &post.Rust, &post.HTML_CSS, &post.Angular, &post.Autre, &post.Like, &post.Dislike, &post.Image)
			if err != nil {
				handle500(w, err)
				return
			}
			database.FilteredPosts = append([]FilteredPosts{post}, database.FilteredPosts...)
		}
		if err := rows.Err(); err != nil {
			handle500(w, err)
			return
		}

		defer rows.Close()
		http.Redirect(w, r, "index", http.StatusFound)
	}
	if cookieUser == "Visiteur" {
		DestroyCookie(w, r)
		SetCookie(w, cookieUser, category)
		if category == "" {
			filtered = false
			database.User.Filtered = false
			http.Redirect(w, r, "index", http.StatusFound)
			return
		}
		FilterSwitch(w, r, cookieUser, category)
		http.Redirect(w, r, "index", http.StatusFound)
	}
}
