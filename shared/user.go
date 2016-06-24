package shared

import (
	"net/http"
)

type User struct {
	username string
	passHash []byte
	files    *[]File
}

func isUserLoggedIn(r *http.Request) bool {
	/*
	stored, ok := getSessionToken(r)
	if !ok {
		return false
	}

	return stored == r.Header.Get("Session-Token")
	*/

	return false
}
