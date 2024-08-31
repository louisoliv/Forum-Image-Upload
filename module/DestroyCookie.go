package module

import "net/http"

func DestroyCookie(w http.ResponseWriter, r *http.Request) {
	cookies := r.Cookies()
	for i := 0; i < len(cookies); i++ {
		cookie := http.Cookie{
			Name:   cookies[i].Name,
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		}
		http.SetCookie(w, &cookie)
	}
}
