package main

import (
	"fmt"
	"net/http"
)

const (
	serverPort = ":41080"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if userLoggedIn() {
		accountHandler(w, r)
	} else {
		loginHandler(w, r)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("login page hit\n")
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("logout page hit\n")
}

func accountHandler(w http.ResponseWriter, r *http.Request) {
	if !userLoggedIn() {
		// TODO send 401 (unauthorized)
		return
	}
	fmt.Printf("account page hit\n")
}

func userLoggedIn() bool {
	return false
}

func generateToken() string {
	return "token"
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/account", accountHandler)
	http.ListenAndServe(serverPort, nil)
}
