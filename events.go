package main

import (
	"net/http"
	"os/exec"
	"os"
	"github.com/fsouza/go-dockerclient"
	"strings"
	"fmt"
)

func eventFired(w http.ResponseWriter, r *http.Request) {
	_, err := exec.Command("/bin/bash", "-c", *command).CombinedOutput()

	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
}

func polEventFired(event *docker.APIEvents) {
	eventsList := strings.Split(*events, ",")
	eventType := fmt.Sprintf("%s:%s", event.Type, event.Action)

	if stringInSlice(eventType, eventsList) {
		_, err := exec.Command("/bin/bash", "-c", *command, event.ID).CombinedOutput()

		if err != nil {
			os.Stderr.WriteString(err.Error())
		}
	}
}
