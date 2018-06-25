package main

import (
	"net/http"
	"os/exec"
	"os"
	"github.com/fsouza/go-dockerclient"
	"strings"
	"fmt"
	"log"
)

func eventFired(w http.ResponseWriter, r *http.Request) {
	_, err := exec.Command("/bin/bash", "-c", *command).CombinedOutput()

	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
}

func polEventFired(event *docker.APIEvents) {
	eventType := fmt.Sprintf("%s:%s", event.Type, event.Action)

	if getEventCommand(eventType) != "" {
		log.Printf("%s event fired", eventType)

		_, err := exec.Command("/bin/bash", "-c", *command, event.ID).CombinedOutput()

		if err != nil {
			os.Stderr.WriteString(err.Error())
		}
	}
}

func getEventCommand(eventType string) string {
	if *commandFile != "" {
		commands := GetCommands()

		if command, ok := commands.Events[eventType]; ok {
			return command
		}
	} else {
		eventsList := strings.Split(*events, ",")

		if stringInSlice(eventType, eventsList) {
			return *command
		}
	}

	return ""
}
