# docker-swarm-watcher
Listen for docker swarm events and exposes and event endpoint.


# Requirements

Only swarm managers can retrieve swarm events. If you're not receiving all events selinux might be the cause. 

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

# options

| Name  | Flag | Description
|---|---|---|
| Command | -c | Command to execute when an event is fired |
| username | -u | Basic authentication username |
| password | -p | Basic authentication password |
| port | -port | Port to expose -- defaults to 8888 |
| socket | -s | Docker socket to poll -- e.g. unix:///var/run/docker.sock |
| events | -e | Comma separated list of Docker events to listen to |
| commandFile | -f | Commands yml file |
| maxEventAge | -max-event-age | Replay events if there age is less than x minutes |
