package module

import (
	"database/sql"
	"fmt"
	"os"
)

func LoadPost() {
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Println("Connected to database")
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var post Posts
		err := rows.Scan(&post.Id, &post.Username, &post.Date, &post.Token, &post.Message, &post.Golang, &post.JavaScript, &post.Python, &post.Rust, &post.HTML_CSS, &post.Angular, &post.Autre, &post.Like, &post.Dislike, &post.Image)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		database.Posts = append([]Posts{post}, database.Posts...)
	}
	if err := rows.Err(); err != nil {
		os.Exit(3)
	}
}
