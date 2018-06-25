package main

import (
	"path/filepath"
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

type Commands struct {
	Events map[string]string
}

func GetCommands() *Commands {
	c := &Commands{}
	c.parseCommandFile()

	return c
}

func (c *Commands) parseCommandFile() {
	file, _ := filepath.Abs(*commandFile)
	yamlFile, err := ioutil.ReadFile(file)

	if err != nil {
		log.Printf("Config yaml file error:   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Config yaml unmarshal error: %v", err)
	}
}
