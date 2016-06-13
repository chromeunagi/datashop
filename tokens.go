package main

import (
	"fmt" // TODO remove
	"net/http"
	"sync"
)

type Tokens struct {
	store map[string]string
	*sync.Mutex
}

func NewTokens() *Tokens {
	return &Tokens{
		make(map[string]string),
		new(sync.Mutex),
	}
}

func (t *Tokens) getSessionToken(r *http.Request) (string, bool) {
	user := r.Header.Get("User")
	passHash := r.Header.Get("Password-Hash")
	key := user + passHash

	fmt.Println("checking for token:", user, passHash, key)

	val, ok := t.store[key]

	fmt.Println("res:", val, ok)
	return val, ok
}
