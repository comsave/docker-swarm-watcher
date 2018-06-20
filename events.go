package main

import (
	"net/http"
	"os/exec"
	"os"
)

func eventFired(w http.ResponseWriter, r *http.Request) {

	_, err := exec.Command("/bin/bash", "-c", *command).CombinedOutput()

	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
}
