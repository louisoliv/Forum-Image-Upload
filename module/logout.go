package module

import (
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	DestroyCookie(w, r)
	database.User = User_info{}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
