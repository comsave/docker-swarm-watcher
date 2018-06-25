# docker-swarm-watcher
Listens docker service create &amp; remove events and executes custom command

# run 
`docker-swarm-watcher -command="oh something changed at service $1"`

# build / install 
- `go get -d -v ./...`
- `go build ./...`
- `go install -v ./...`
