package main

import (
	"fmt"
)

var (
	version = "v1"
	prefix  = fmt.Sprintf("/%s/event", version)
	routes  = []Route{
		{
			"New Event",
			"GET",
			fmt.Sprintf("%s/new", prefix),
			eventFired,
		},
	}
)
