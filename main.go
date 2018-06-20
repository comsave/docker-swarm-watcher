package main

import (
	"log"
	"net/http"
	"flag"
	"fmt"
)

var (
	command  = flag.String("c", "/bin/echo", "Command to execute when an event is fired")
	username = flag.String("u", "username", "Basic authentication username")
	password = flag.String("p", "password", "Basic authentication password")
	port     = flag.String("port", "8888", "Port to expose")
)

func main() {
	flag.Parse()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s",*port), NewRouter(routes)))
}
