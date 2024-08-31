package module

import (
	"database/sql"
	"fmt"
	"os"
)

func LoadLike() {
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
	rows, err := db.Query("SELECT username, token FROM like")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var like LikE
		err := rows.Scan(&like.Username, &like.Token)
		if err != nil {
			os.Exit(1)
		}
		database.LikE = append([]LikE{like}, database.LikE...)
	}
	if err := rows.Err(); err != nil {
		os.Exit(3)
	}
}
