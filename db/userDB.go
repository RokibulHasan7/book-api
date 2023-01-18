package db

import (
	"github.com/go-chi/jwtauth/v5"
)

var TokenAuth *jwtauth.JWTAuth
var Users map[string]string

func InitUser() {
	Users = map[string]string{
		"rakib": "1234",
		"sakib": "12345",
	}
	TokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
}
