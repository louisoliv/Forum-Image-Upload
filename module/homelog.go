package module

import (
	"fmt"
	"net/http"
	"text/template"
)

// Load the html login
func HomeLog(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/login.html")
	if err != nil {
		fmt.Println(err)
		handle500(w, err)
		return
	}
	err = t.Execute(w, Error)
	if err != nil {
		fmt.Println(err)
		handle500(w, err)
		return
	}
	Error2.ErrorMessage = ""
	Error = nil
}
