package main

import (
	"fmt"
)

var version = "v1"
var prefix = fmt.Sprintf("/%s/event", version)

var routes = Routes{
	Route{
		"Create",
		"GET",
		fmt.Sprintf("%s/create", prefix),
		eventFired,
	},
}

