package module

import (
	"database/sql"
	"fmt"
	"os"
)

func LoadComment() {
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
	rows, err := db.Query("SELECT * FROM comments")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var comment Comments
		err := rows.Scan(&comment.Id, &comment.Username, &comment.Date, &comment.TokenComment, &comment.Message_comment, &comment.TokenID, &comment.Like, &comment.Dislike, &comment.Image_comment)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		DBComments = append(DBComments, comment)
		database.Comments = append(database.Comments, comment)
	}
	if err := rows.Err(); err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
}
