# docker-swarm-watcher
Listens docker service create &amp; remove events and executes custom command

# run 
`docker-swarm-watcher -command="oh something changed at service $1"`

# build 
`go get -d -v ./...`
`go install -v ./...`
`go build ./...`
