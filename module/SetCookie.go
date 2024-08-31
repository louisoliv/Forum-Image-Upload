package module

import (
	"net/http"
	"time"
)

func SetCookie(w http.ResponseWriter, username, token string) {
	expire := time.Now().Add(10 * time.Minute)
	cookie := http.Cookie{Name: username, Value: token, Expires: expire}
	http.SetCookie(w, &cookie)
}
