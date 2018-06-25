package main

import (
	"log"
	"flag"
	"net/http"
	"github.com/fsouza/go-dockerclient"
	"fmt"
	//"time"
)

var (
	command     = flag.String("c", "/bin/echo", "Command to execute when an event is fired")
	username    = flag.String("u", "username", "Basic authentication username")
	password    = flag.String("p", "password", "Basic authentication password")
	port        = flag.String("port", "8888", "Port to expose -- defaults to 8888")
	socket      = flag.String("s", "", "Docker socket to poll -- e.g. unix:///var/run/docker.sock")
	events      = flag.String("e", "", "Comma separated list of Docker events to listen to")
	commandFile = flag.String("f", "", "Commands yml file")
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
		log.Println(err)
	}

	listener := make(chan *docker.APIEvents)
	err = client.AddEventListener(listener)

	if err != nil {
		log.Println(err)
	}

	defer func() {
		err = client.RemoveEventListener(listener)

		if err != nil {
			log.Println(err)
		}
	}()

	for {
		select {
		case event := <-listener:
			polEventFired(event)
		}
	}
}
