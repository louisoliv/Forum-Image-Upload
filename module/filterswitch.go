package module

import (
	"database/sql"
	"fmt"
	"net/http"
)

func FilterSwitch(w http.ResponseWriter, r *http.Request, username, category string) {
	database.User.Filtered = true
	if username == "Visiteur" {
		cookies := r.Cookies()
		database.User.Category = cookies[0].Value
		if cookies[0].Value != "" {
			goto next
		}
		return
	}
	database.User.Category = category
next:
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
	stmt, err := db.Prepare(fmt.Sprintf("SELECT * FROM posts WHERE %s = 1", category))
	if err != nil {
		return
	}
	defer stmt.Close()

	// Execute the query
	rows, err := stmt.Query()
	if err != nil {
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
}
