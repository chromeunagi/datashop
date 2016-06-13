package main

import (
	"fmt"
	"net/http"
)

const (
	serverPort = ":41080"
)

var (
	TokenStore *Tokens
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if isUserLoggedIn(r) {
		accountHandler(w, r)
	} else {
		loginHandler(w, r)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("login page hit\n")
	t := generateSessionToken(r)
	if t == "" {
		return
	}
	fmt.Println("token:", t)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	if !isUserLoggedIn(r) {
		// TODO send 401 (unauthorized)
		return
	}
	fmt.Printf("logout page hit\n")
}

func accountHandler(w http.ResponseWriter, r *http.Request) {
	if !isUserLoggedIn(r) {
		// TODO send 401 (unauthorized)
		return
	}
	fmt.Printf("account page hit\n")
}

func init() {
	TokenStore = NewTokens()
}

func main() {

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/account", accountHandler)
	http.ListenAndServe(serverPort, nil)
}
