package main

import (
	"net/http"
	"os/exec"
	"log"
)

func eventFired(w http.ResponseWriter, r *http.Request)  {
	var serviceName = r.URL.Query().Get("serviceName")

	cmd := exec.Command(*command, serviceName)
	err := cmd.Run()

	if err != nil {
		log.Printf("%v", err)
	}
}
