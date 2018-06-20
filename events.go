package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"log"
)

func eventFired(w http.ResponseWriter, r *http.Request)  {
	fmt.Printf("An event was fired %v --- %v", w, r)

	cmd := exec.Command("docker-gen")
	err := cmd.Run()

	if err != nil {
		log.Printf("%v", err)
	}
}
