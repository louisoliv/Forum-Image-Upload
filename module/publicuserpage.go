package module

import (
	"fmt"
	"net/http"
	"text/template"
)

func PublicUserPage(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("userpublic")
	t, err := template.ParseFiles("./templates/user_public_profile.html")
	if err != nil {
		fmt.Println(err)
		handle500(w, err)
		return
	}
	database.User.Userpublic = username
	for i := 0; i < len(database.User_public); i++ {
		if database.User_public[i].Username == username {
			break
		}
		if i == len(database.User_public)-1 && database.User_public[i].Username != username {
			http.Redirect(w, r, "index", http.StatusFound)
			return
		}
	}
	err = t.Execute(w, database)
	if err != nil {
		fmt.Println(err)
		handle500(w, err)
		return
	}
}
