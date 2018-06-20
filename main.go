package main

import (
	"log"
	"net/http"
	"flag"
)

var (
	command  = flag.String("c", "docker-gen", "Command to execute when an event is fired")
	username = flag.String("u", "username", "Basic authentication username")
	password = flag.String("p", "password", "Basic authentication password")
)

func main() {
	flag.Parse()

	log.Fatal(http.ListenAndServe(":8888", NewRouter(routes)))
}
