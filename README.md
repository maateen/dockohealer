# dockohealer

[![Build Status](https://travis-ci.com/maateen/dockohealer.svg?branch=master)](https://travis-ci.com/maateen/dockohealer)
[![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/maateen/dockohealer)](https://hub.docker.com/r/maateen/dockohealer)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/maateen/dockohealer)](https://github.com/maateen/dockohealer)
[![Go Report Card](https://goreportcard.com/badge/github.com/maateen/dockohealer)](https://goreportcard.com/report/github.com/maateen/dockohealer)
![Maintenance](https://img.shields.io/maintenance/yes/2020)
[![GitHub](https://img.shields.io/github/license/maateen/dockohealer?color=green)](https://github.com/maateen/dockohealer/blob/master/LICENSE)

This daemon monitors and restarts unhealthy docker containers instantly. This project is inspired from [docker-autoheal](https://github.com/willfarrell/docker-autoheal) and written on GoLang with love.

## Installation

#### Build from source

```shell script
$ git clone https://github.com/maateen/dockohealer.git
$ cd dockohealer/cmd/dockohealer
$ buildTime=$(date +'%Y-%m-%d_%T')
$ gitSHA=$(git rev-parse HEAD)
$ versionString=$(git tag --sort=committerdate | tail -1)
$ go build -ldflags "-X main.buildTime=$buildTime -X main.gitSHA=$gitSHA -X main.versionString=$versionString"
```

#### Use binary

```shell script
$ export VERSION=v0.3
$ export OS=linux
$ export ARCH=amd64
$ wget https://github.com/maateen/dockohealer/releases/download/$VERSION/dockohealer-$OS-$ARCH-$VERSION
$ mv dockohealer-$OS-$ARCH-$version /usr/local/bin/dockohealer
$ chmod +x /usr/local/bin/dockohealer

```

## Usage

#### Standalone

```shell script
$ dockohealer -version

{"buildTime":"","gitSHA":"","level":"info","msg":"","time":"2020-06-20T19:38:43+06:00","version":""}

$ dockohealer

{"level":"info","msg":"Connected to Docker.","time":"2020-04-04T15:26:05+06:00"}
{"level":"info","msg":"Listening from Docker.","time":"2020-04-04T15:26:05+06:00"}
{"containerID":"5e37d4624fbaa128d1fbdd21e3a4cf0aa78eeff48e8902ef60eca95496d3155c","level":"info","msg":"Container is unhealthy.","time":"2020-04-04T15:26:15+06:00"}
{"containerID":"5e37d4624fbaa128d1fbdd21e3a4cf0aa78eeff48e8902ef60eca95496d3155c","level":"info","msg":"Restarting container.","time":"2020-04-04T15:26:15+06:00"}
{"containerID":"5e37d4624fbaa128d1fbdd21e3a4cf0aa78eeff48e8902ef60eca95496d3155c","level":"info","msg":"Successfully restarted container.","time":"2020-04-04T15:26:15+06:00"}
```

#### Docker Container

```shell script
$ docker run -d \
      --name dockohealer \
      --restart=always \
      -v /var/run/docker.sock:/var/run/docker.sock \
      maateen/dockohealer
```

## Process Manager

In case of running as standalone binary, you should use [systemd](https://www.linode.com/docs/quick-answers/linux/start-service-at-boot/) or [supervisor](http://supervisord.org/) to keep this daemon in running state always.

## Road map

- [x] Restart already unhealthy containers
- [x] Add flags to the daemon
- [x] Launch a dockerized version
