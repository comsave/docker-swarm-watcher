package main

import (
	"log"
	"net/http"
	"flag"
)

var (
	command = flag.String("command", "docker-gen", "Command to execute when an event is fired")
)

func main(){
	flag.Parse()

	log.Fatal(http.ListenAndServe(":8888", NewRouter(routes)))
}
