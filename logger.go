package main

import (
	"log"
	"net/http"
	"time"
)

func writeLogEntry(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		if !IsAuthenticated(w, r) {
			SetNotAuthenticated(w)
			return
		}

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
