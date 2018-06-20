package main

import (
	"fmt"
)

var version = "v1"
var prefix = fmt.Sprintf("/%s/event", version)

var routes = Routes{
	Route{
		"New Event",
		"GET",
		fmt.Sprintf("%s/new", prefix),
		eventFired,
	},
}

