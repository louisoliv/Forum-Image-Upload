package module

import (
	"fmt"
	"net/http"
	"text/template"
)

// Load register form page
func RegisterPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		t, err := template.ParseFiles("./templates/register.html")
		if err != nil {

			fmt.Println(err)
			handle500(w, err)
			return
		}

		err = t.Execute(w, Error)
		if err != nil {
			fmt.Println("here3")
			fmt.Println(err)
			handle500(w, err)
			return
		}
		Error2.ErrorMessage = ""
		Error = nil
	}
}
