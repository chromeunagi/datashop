package main

import ()

type User struct {
	username string
	passHash []byte
	files    *[]Files
}
