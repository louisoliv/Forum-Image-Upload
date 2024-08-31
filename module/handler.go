package module

import (
	"fmt"
	"net/http"
	"text/template"
)

func handleError(w http.ResponseWriter, statusCode int, templateFiles ...string) {
	w.WriteHeader(statusCode)
	t, err := template.ParseFiles(templateFiles...)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Something went wrong\nError %d\n%s", statusCode, err)))
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Something went wrong\nError %d\n%s", statusCode, err)))
		return
	}
}

func handle500(w http.ResponseWriter, err error) {
	handleError(w, 500, "templates/error500.html")
	fmt.Println(err)
}

func handle404(w http.ResponseWriter, _ *http.Request) {
	handleError(w, 404, "templates/error404.html")
}

func handle400(w http.ResponseWriter, _ *http.Request) {
	handleError(w, 400, "templates/error400.html")
}
