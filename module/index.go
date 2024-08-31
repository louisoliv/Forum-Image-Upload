package module

import (
	"fmt"
	"html/template"
	"net/http"
)

// Load user page
func Index(w http.ResponseWriter, r *http.Request) {
	cookieUser, cookieToken, _ := Checklog(w, r)
	database.User.Username = cookieUser
	database.User.Token = cookieToken
	if r.Method == "GET" {
		t, err := template.ParseFiles("./templates/blog.html")
		if err != nil {
			fmt.Println(err)
			handle500(w, err)
			return
		}
		err = t.Execute(w, database)
		database.User.ErrorMessage = ""
		database.User.ErrorMessage2 = ""
		database.User.ErrorMessage3 = ""
		database.User.ErrorMessage4 = ""
		database.User.ErrorMessage5 = ""
		if err != nil {
			fmt.Println(err)
			handle500(w, err)
			return
		}
	}
}
