package shared

import (
	"crypto/sha256"
	"fmt" // TODO remove
	"net/http"
	"strconv"
	"sync"
	"time"
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

func generateSessionToken(r *http.Request) string {
	t := strconv.Itoa(int(time.Now().Unix()))
	s := r.Header.Get("USER") + t
	b := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", string(b[:]))
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
