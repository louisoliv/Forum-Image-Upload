package module

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// Formate given information by user to uniformize data before storage into the db
func AddtoDB(w http.ResponseWriter, register []string) {
	fmt.Println(register)
	addRole := "Member"
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		handle500(w, err)
		return
	}
	defer db.Close()
	S := strings.ToUpper(string(register[1][0]))
	R := []rune(register[1])
	for i := 0; i < len(register[1]); i++ {
		if i != 0 {
			S += strings.ToLower(string(register[1][i]))
		}
	}
	fmt.Println(S)
	register1 := string(R)
	fmt.Println(register[1], register1)
	register[2] = strings.ToUpper(register[2])
	hashpass, err := HashPass(register[4])
	if err != nil {
		handle500(w, err)
		return
	}
	mailtoken := GenerateToken()
	verified := false
	MailVerification(register[3], mailtoken)
	_, err = db.Exec("INSERT INTO Users (username, first_name, last_name, email, role, password, Messages, Date, Localisation, Statut, Loisirs, Date_naissance, Sexe, Verification_token, Verified) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", register[0], S, register[2], register[3], addRole, hashpass, 0, time.Now().Format("2006-01-02 15:04:05"), "", "", "", "", "", mailtoken, verified)
	if err != nil {
		handle500(w, err)
		return
	}
	user := User_public{
		Username:       register[0],
		Email:          register[3],
		Firstname:      S,
		Lastname:       register[2],
		Role:           "Member",
		Date:           time.Now().Format("2006-01-02 15:04:05"),
		Messages:       0,
		Localisation:   "",
		Statut:         "",
		Loisirs:        "",
		Date_Naissance: "",
		Sexe:           "",
	}
	database.User_public = append(database.User_public, user)
}
