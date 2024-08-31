package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"forum/module"

	_ "github.com/mattn/go-sqlite3"
)

const (
	colorYellow = "\033[33m"
	colorGreen  = "\033[32m"
	colorRed    = "\033[31m"
)

// Start the server at the port :5657
func main() {
	fs := http.FileServer(http.Dir("styles"))
	http.Handle("/styles/", http.StripPrefix("/styles/", fs))
	// Load the structs from forum.db
	if !module.LoadedUser {
		module.LoadUser()
		module.LoadedUser = true
	}
	if !module.Loaded {
		module.LoadPost()
		module.Loaded = true
	}
	if !module.Loadcomment {
		module.LoadComment()
		module.Loadcomment = true
	}
	if !module.Loadlike {
		module.LoadLike()
		module.Loadlike = true
	}
	http.HandleFunc("/", module.HomeLog)
	// Actions posts, comments, like, filter, login, logout functions
	http.HandleFunc("/filter", module.Filter)
	http.HandleFunc("/comment", module.Comment)
	http.HandleFunc("/like", module.Like)
	http.HandleFunc("/dislike", module.Like)
	http.HandleFunc("/commentlike", module.CommentLike)
	http.HandleFunc("/commentdislike", module.CommentLike)
	http.HandleFunc("/post", module.Post)
	http.HandleFunc("/login", module.Log)
	http.HandleFunc("/logout", module.Logout)
	// Loading page functions
	http.HandleFunc("/registerform", module.RegisterPage)
	http.HandleFunc("/submitregister", module.CheckRegister)
	http.HandleFunc("/index", module.Index)
	http.HandleFunc("/invite", module.Index)
	http.HandleFunc("/user", module.User)
	http.HandleFunc("/about", module.About)
	http.HandleFunc("/modify_user_profile", module.Modify_data_user)
	http.HandleFunc("/user_public_profile", module.PublicUserPage)
	http.HandleFunc("/historic_created_post", module.Historic_created_post)
	http.HandleFunc("/historic_liked_posts", module.Historic_liked_post)
	http.HandleFunc("/historic_comments", module.Historic_comments)
	http.HandleFunc("/publicuserpageModify", module.PublicUserPageModify)
	http.HandleFunc("/verify", module.VerifyMail)

	// Print main information into terminal when server is starting
	fmt.Println(string(colorYellow), "Starting local Server ...")
	fmt.Println(string(colorGreen), "Server ready on http://localhost:8000")
	fmt.Println(string(colorRed), "To stop  the program: Ctrl + C")
	srv := &http.Server{
		Addr:              "localhost:8000",
		ReadHeaderTimeout: 15 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       30 * time.Second,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
