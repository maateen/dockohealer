# Go parameters
GOCMD=$(shell which go)
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
WORKDIR=cmd/dockohealer
buildTime=$(shell date +'%Y-%m-%d_%T')
gitSHA=$(shell git rev-parse HEAD)
versionString=$(shell git tag --sort=committerdate | tail -1)

all: prepare build
prepare:
		$(GOCMD) mod tidy
		$(GOCMD) mod verify
build:
		$(GOBUILD) -ldflags "-X main.buildTime=$(buildTime) -X main.gitSHA=$(gitSHA) -X main.versionString=$(versionString) $(WORKDIR)/dockohealer.go"
run:
		$(GORUN) -ldflags "-X main.buildTime=$(buildTime) -X main.gitSHA=$(gitSHA) -X main.versionString=$(versionString)" $(WORKDIR)/dockohealer.go
clean:
		cd cmd/dockohealer
		$(GOCLEAN)