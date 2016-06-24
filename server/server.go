package server

import (
	"fmt"
	"net/http"
	//"time"

	"github.com/sg95/datashop/shared"
)

const (
	serverPort = ":41080"
	requestChannelSize = 1024
)

var (
	TokenStore *Tokens
)

type (
	Server struct {
		// Server state
		currentUsers []*shared.User

		// Requests
		logins chan *loginRequest
		uploads chan *uploadRequest
		reads chan *readRequest
	}

	loginRequest struct {

	}

	uploadRequest struct {

	}

	readRequest struct {

	}
)

// Creates a Server and instantiates its resources.
func NewServer() *Server {
	server := new(Server)

	server.currentUsers = make([]*shared.User, 0)
	server.logins = make(chan *loginRequest, requestChannelSize)
	server.uploads = make(chan *uploadRequest, requestChannelSize)
	server.reads = make(chan *readRequest, requestChannelSize)

	return server
}

// rootHandler redirects to an appropriate handler.
func rootHandler(w http.ResponseWriter, r *http.Request) {
	if isUserLoggedIn(r) {
		accountHandler(w, r)
	} else {
		loginHandler(w, r)
	}
}

// loginHandler handles login request.
func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("login page hit\n")
	t := generateSessionToken(r)
	if t == "" {
		return
	}
	fmt.Println("token:", t)
}

// logoutHandler handles requests to log out.
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	if !isUserLoggedIn(r) {
		// TODO send 401 (unauthorized)
		return
	}
	fmt.Printf("logout page hit\n")
}

// accountHandler handles requests to view "my account" page.
func accountHandler(w http.ResponseWriter, r *http.Request) {
	if !isUserLoggedIn(r) {
		// TODO send 401 (unauthorized)
		return
	}
	fmt.Printf("account page hit\n")
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/account", accountHandler)
	http.ListenAndServe(serverPort, nil)
}
