package main

import (
	"encoding/base64"
	"net/http"
	"strings"
)

const realm = "web-dev-auth"

type User struct {
	username string
	password string
}

// 维护系统中的用户列表
var users = []User{
	{username: "root", password: "123456"},
	// ...
}

func validateUser(user *User) bool {
	for _, u := range users {
		if u == *user {
			return true
		}
	}
	return false
}

func validateBasicAuth(authorization string) bool {
	if authorization == "" {
		return false
	}
	rawBase64, hasPrefix := strings.CutPrefix(authorization, "Basic ")
	if !hasPrefix {
		return false
	}
	credentials, err := base64.StdEncoding.DecodeString(rawBase64)
	if err != nil {
		return false
	}
	pair := strings.Split(string(credentials), ":")
	if len(pair) != 2 {
		return false
	}
	username, password := pair[0], pair[1]
	return validateUser(&User{username, password})

}

func main() {
	http.HandleFunc("/protected", func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !validateBasicAuth(authHeader) { // 验证不通过时，
			w.Header().Set("WWW-Authenticate", "Basic realm="+realm)
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			w.Write([]byte("Aha! You got me!"))
		}

	})

	http.ListenAndServe(":3000", nil)
}
