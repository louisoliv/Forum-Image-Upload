package module

import (
	"fmt"
	"net/http"
	"text/template"
)

func Historic_created_post(w http.ResponseWriter, r *http.Request) {
	cookieUser, cookieToken, check := Checklog(w, r)
	if !check {
		http.Redirect(w, r, "index", http.StatusFound)
		return
	}
	database.User.Username = cookieUser
	database.User.Token = cookieToken
	t, err := template.ParseFiles("./templates/history_created_posts.html")
	if err != nil {
		fmt.Println(err)
		handle500(w, err)
		return
	}
	err = t.Execute(w, database)
	if err != nil {
		fmt.Println(err)
		handle500(w, err)
		return
	}
}
