# docker-swarm-watcher
Listen for docker swarm events and exposes and event endpoint.


# Requirements

Only swarm managers can retrieve swarm events.

```yml
version: "3"
services:
  web:
    deploy:
      placement:
        constraints: [node.role == manager]
```

# build / install 
- `go get -d -v ./...`
- `go build ./...`
- `go install -v ./...`

# run 

To expose an event endpoint and listen to docker service:create events

```bash
docker-swarm-watcher -c="/bin/echo hello" -u username -p password -s="unix:///var/run/docker.sock" -e="service:create"
```

To specify specific commands for each event you can use a command file. See example-commands.yml

```bash
docker-swarm-watcher -u username -p password -s="unix:///var/run/docker.sock" -f="/home/user/commands.yml"
```
