package main

import (
	"net/http"
)

type User struct {
	username string
	passHash []byte
	files    *[]File
}

func isUserLoggedIn(r *http.Request) bool {
	stored, ok := TokenStore.getSessionToken(r)
	if !ok {
		return false
	}

	return stored == r.Header.Get("Session-Token")
}
