# dockohealer

[![Go Report Card](https://goreportcard.com/badge/github.com/maateen/dockohealer)](https://goreportcard.com/report/github.com/maateen/dockohealer)

This daemon monitors and restarts unhealthy docker containers instantly. This project is inspired from [docker-autoheal](https://github.com/willfarrell/docker-autoheal) and written on GoLang with love.

## Installation

```shell script

```

## Usage

```shell script
$ dockohealer
```

```json
{"level":"info","msg":"Connected to Docker.","time":"2020-04-04T15:26:05+06:00"}
{"level":"info","msg":"Listening from Docker.","time":"2020-04-04T15:26:05+06:00"}
{"containerID":"5e37d4624fbaa128d1fbdd21e3a4cf0aa78eeff48e8902ef60eca95496d3155c","level":"info","msg":"Container is unhealthy.","time":"2020-04-04T15:26:15+06:00"}
{"containerID":"5e37d4624fbaa128d1fbdd21e3a4cf0aa78eeff48e8902ef60eca95496d3155c","level":"info","msg":"Restarting container.","time":"2020-04-04T15:26:15+06:00"}
{"containerID":"5e37d4624fbaa128d1fbdd21e3a4cf0aa78eeff48e8902ef60eca95496d3155c","level":"info","msg":"Successfully restarted container.","time":"2020-04-04T15:26:15+06:00"}
```

## Process Manager

You should use [systemd](https://www.linode.com/docs/quick-answers/linux/start-service-at-boot/) or [supervisor](http://supervisord.org/) to keep this daemon in running state always.

## Road map

- [ ] Restart already unhealthy containers
- [ ] Add flags to the daemon
- [ ] Launch a dockerized version