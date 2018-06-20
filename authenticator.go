package main

import (
	"encoding/base64"
	"net/http"
	"strings"
)

func BasicAuth(w http.ResponseWriter, r *http.Request, user, pass []byte) bool {

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

func IsAuthenticated(w http.ResponseWriter, r *http.Request) bool {
	return BasicAuth(w, r, []byte(*username), []byte(*password))
}

func SetNotAuthenticated(w http.ResponseWriter) {
	w.Header().Set("WWW-Authenticate", `Basic realm="Beware! Protected REALM! "`)
	w.WriteHeader(401)
	w.Write([]byte("401 Unauthorized\n"))
}
