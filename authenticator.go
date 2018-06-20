package main

import (
	"encoding/base64"
	"net/http"
	"strings"
)

func authenticate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !isAuthenticated(w, r) {
			setNotAuthenticated(w)
			return
		}
	})
}


func basicAuth(w http.ResponseWriter, r *http.Request, user, pass []byte) bool {

	s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(s) != 2 {
		return false
	}

	b, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		return false
	}

	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 {
		return false
	}

	return pair[0] == string(user) && pair[1] == string(pass)
}

func isAuthenticated(w http.ResponseWriter, r *http.Request) bool {
	return basicAuth(w, r, []byte(*username), []byte(*password))
}

func setNotAuthenticated(w http.ResponseWriter) {
	w.Header().Set("WWW-Authenticate", `Basic realm="Beware! Protected REALM! "`)
	w.WriteHeader(401)
	w.Write([]byte("401 Unauthorized\n"))
}
