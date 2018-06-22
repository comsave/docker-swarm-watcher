package main

import (
	"net/http"
	"os/exec"
	"fmt"
	"os"
	"github.com/fsouza/go-dockerclient"
)

func eventFired(w http.ResponseWriter, r *http.Request) {
	var serviceName = r.URL.Query().Get("serviceName")

	_, err := exec.Command("/bin/bash", "-c", fmt.Sprintf("%s %s", *command, serviceName)).CombinedOutput()

	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
}

func polEventFired(msg *docker.APIEvents)  {
	_, err := exec.Command("/bin/bash", "-c", *command).CombinedOutput()

	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
}
