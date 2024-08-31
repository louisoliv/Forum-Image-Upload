package module

import (
	"fmt"
	"net/http"
	"text/template"
)

// Load forum index
func About(w http.ResponseWriter, r *http.Request) {
	cookieUser, cookieToken, check := Checklog(w, r)
	fmt.Println(cookieUser, cookieToken, check)
	database.User.Username = cookieUser
	database.User.Token = cookieToken
	if r.Method == "GET" {
		t, err := template.ParseFiles("./templates/about.html")
		if err != nil {
			handle500(w, err)
			return
		}

		err = t.Execute(w, nil)
		if err != nil {
			handle500(w, err)
			return
		}
	}
}
