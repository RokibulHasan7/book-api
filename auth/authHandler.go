package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/RokibulHasan7/book-api/model"

	"github.com/RokibulHasan7/book-api/db"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)

	authHeader := r.Header.Get("Authorization")
	authStr := strings.Split(authHeader, " ")
	fmt.Println(authStr[0])
	fmt.Println(authStr[1])

	err := Checker(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	//claims := token.Claims.(jwt.MapClaims)

	log.Println(db.TokenAuth)
	myToken, ok := GenerateToken(user.UserName)
	expireTime := time.Now().Add(2 * time.Minute)

	fmt.Println("Mytoken: ", myToken)
	if ok != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Cann't Generate Authentication Token."))
		return
	}

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
		//fmt.Println(authHeader)
		//authStr := strings.Split(authHeader, " ")
		//fmt.Println(authStr[0])
		//fmt.Println(authStr[1])
		if len(authHeader) > 7 && strings.ToUpper(authHeader[0:6]) == "BEARER" {
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
			/*claims := jwt.MapClaims{}
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte("secret"), nil
			})
			if err != nil {
				http.Error(w, "Access token is missing or invalid", http.StatusUnauthorized)
				return
			}

			// do something with decoded claims
			for key, val := range claims {
				fmt.Printf("Key: %v, value: %v\n", key, val)
			}
			if err == nil {

				next.ServeHTTP(w, r)
			}
			if !token.Valid {
				http.Error(w, "Access token is missing or invalid", http.StatusUnauthorized)
				return
			}

			// validate the essential claims

			/*token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				return []byte("AllYourBase"), nil
			})
			if err != nil {
				http.Error(w, "Access token is missing or invalid", http.StatusUnauthorized)
				return
			} else if _, ok := err.(*jwt.ValidationError); ok {
				http.Error(w, "Access token is missing or invalid", http.StatusUnauthorized)
				return
			} else if token.Valid {
				next.ServeHTTP(w, r)
			} else {
				http.Error(w, "Access token is missing or invalid", http.StatusUnauthorized)
				return
			}*/
		} else {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	})
}

func BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		fmt.Println(authHeader)
		authStr := strings.Split(authHeader, " ")
		fmt.Println(authStr[0])
		fmt.Println(authStr[1])
		if len(authStr) == 2 && authStr[0] == "Basic" {
			username, password, ok := r.BasicAuth()

			if !ok {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			} else {
				if _, ok := db.Users[username]; ok == false {
					http.Error(w, "User doesn't Exists.", http.StatusBadRequest)
					return
				}

				if pass, _ := db.Users[username]; pass != password {
					http.Error(w, "Wrong Username or Password", http.StatusUnauthorized)
					return
				}
			}

			/*if username != "admin" || password != "1234" {
				w.WriteHeader(400)
				w.Write([]byte("Wrong Username or Password."))
				return
			}*/

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
		/*
			func (r *Request) BasicAuth() (username, password string, ok bool) --->
			It checks the Authorization header and then extracts the username and password from  Base64 encoded value
			and return it. If there is any issue in parsing it will return ok variable as false. So while using this
			function, we first need to check the value of ok variable. If the ok variable is true then we can further
			match the username and password and verify if it is correct.
		*/

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

func GenerateToken(username string) (string, error) {
	expireTime := time.Now().Add(2 * time.Minute)
	_, myToken, err := db.TokenAuth.Encode(map[string]interface{}{
		"aud": username,
		"exp": expireTime.Unix(),
	})
	//log.Println(myToken)

	if err != nil {
		return "", err
	}
	return myToken, nil
}
