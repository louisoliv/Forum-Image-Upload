package module

import (
	"fmt"
	"net/http"
	"text/template"
)

func Historic_comments(w http.ResponseWriter, r *http.Request) {
	cookieUser, cookieToken, check := Checklog(w, r)
	fmt.Println(cookieUser, cookieToken, check)
	database.User.Username = cookieUser
	database.User.Token = cookieToken
	t, err := template.ParseFiles("./templates/history_comments.html")
	if err != nil {
		fmt.Println(err)
		handle500(w, err)
		return
	}
	err = t.Execute(w, user_data)
	if err != nil {
		fmt.Println(err)
		handle500(w, err)
		return
	}
}
