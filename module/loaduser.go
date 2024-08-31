package module

import (
	"database/sql"
	"fmt"
	"os"
)

func LoadUser() {
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
	rows, err := db.Query("SELECT id, username, first_name, last_name, email, role, token, Messages, Date, Localisation, Statut, Loisirs, Date_naissance, Sexe FROM Users")
	if err != nil {
		fmt.Print("rows:", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var user User_public
		err := rows.Scan(&user.Id, &user.Username, &user.Firstname, &user.Lastname, &user.Email, &user.Role, &user.Token, &user.Messages, &user.Date, &user.Localisation, &user.Statut, &user.Loisirs, &user.Date_Naissance, &user.Sexe)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		database.User_public = append([]User_public{user}, database.User_public...)
	}
	if err := rows.Err(); err != nil {
		os.Exit(3)
	}
}
