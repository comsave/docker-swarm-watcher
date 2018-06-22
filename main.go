package main

import (
	"log"
	"flag"
	"net/http"
	"github.com/fsouza/go-dockerclient"
	"fmt"
)

var (
	command  = flag.String("c", "/bin/echo", "Command to execute when an event is fired")
	username = flag.String("u", "username", "Basic authentication username")
	password = flag.String("p", "password", "Basic authentication password")
	port     = flag.String("port", "8888", "Port to expose")
	socket   = flag.String("s", "", "Docker socket to poll")
)

func main() {
	flag.Parse()

	if *socket != "" {
		go listenForEvents()
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *port), NewRouter(routes)))
}

func listenForEvents() {
	client, err := docker.NewClient(*socket)

	if err != nil {

	}

	listener := make(chan *docker.APIEvents)
	err = client.AddEventListener(listener)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {

		err = client.RemoveEventListener(listener)
		if err != nil {
			log.Fatal(err)
		}

	}()

	for {
		select {
		case msg := <-listener:
			polEventFired(msg)
		}
	}
}
