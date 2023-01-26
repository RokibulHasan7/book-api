package auth

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/RokibulHasan7/book-api/db"
)

func Login(w http.ResponseWriter, r *http.Request) {
	ok := Checker(r)
	if ok != nil {
		http.Error(w, ok.Error(), http.StatusUnauthorized)
		return
	}
	//claims := token.Claims.(jwt.MapClaims)

	log.Println(db.TokenAuth)
	myToken, ok := GenerateToken()
	if ok != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Cann't Generate Authentication Token."))
		return
	}

	fmt.Println("Mytoken: ", myToken)
	expireTime := time.Now().Add(2 * time.Minute)
	// Set Cookies
	http.SetCookie(w, &http.Cookie{
		Name:    "jwt",
		Value:   myToken,
		Expires: expireTime,
	})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("You are Logged In!"))
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// Remove cookies
	http.SetCookie(w, &http.Cookie{
		Name:    "jwt",
		Value:   "",
		Expires: time.Now(),
	})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logged Out!"))
}

func PrimaryAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		authStr := strings.Split(authHeader, " ")
		if len(authHeader) > 7 && authStr[0] == "Bearer" {
			tokenString := authHeader[7:]
			//fmt.Println("TokenString from PrimaryAuth:", tokenString)

			token, err := jwt.Parse(tokenString, func(tkn *jwt.Token) (interface{}, error) {
				if _, ok := tkn.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", tkn.Header["alg"])
				}
				return []byte("secret"), nil
			})

			// Checking token payloads
			_, ok := token.Claims.(jwt.MapClaims)

			if !ok || !token.Valid || err != nil {
				http.Error(w, "Access token is missing or invalid", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		} else if len(authHeader) > 6 && authStr[0] == "Basic" {
			username, password, ok := r.BasicAuth()

			if !ok {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			} else {
				if _, ok := db.Users[username]; ok == false {
					http.Error(w, "Wrong Username or Password.", http.StatusUnauthorized)
					return
				}

				if pass, _ := db.Users[username]; pass != password {
					http.Error(w, "Wrong Username or Password", http.StatusUnauthorized)
					return
				}
			}
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	})
}

func Checker(r *http.Request) error {
	authHeader := r.Header.Get("Authorization")
	fmt.Println(authHeader)
	authStr := strings.Split(authHeader, " ")
	if len(authStr) == 2 && authStr[0] == "Basic" {
		username, password, ok := r.BasicAuth()

		if !ok {
			return errors.New("unauthorized")
		} else {
			if _, ok := db.Users[username]; ok == false {
				return errors.New("user doesn't exits")
			}

			if pass, _ := db.Users[username]; pass != password {
				return errors.New("wrong username or password")
			}
		}

	} else {
		return errors.New("unauthorized")
	}
	return nil
}

func GenerateToken() (string, error) {
	expireTime := time.Now().Add(2 * time.Minute)
	_, myToken, err := db.TokenAuth.Encode(map[string]interface{}{
		"aud": "Rokibul Hasan",
		"exp": expireTime.Unix(),
	})
	//log.Println(myToken)

	if err != nil {
		return "", err
	}
	return myToken, nil
}
