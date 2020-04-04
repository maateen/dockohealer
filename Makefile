# Go parameters
GOCMD=$(shell which go)
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_NAME=dockohealer
GOBIN=bin
GOOS=darwin linux
GOARCH=amd64 386
VERSION=v0.1

all: prepare build
prepare:
		$(GOCMD) mod tidy
		$(GOCMD) mod verify
build:
		mkdir -p $(GOBIN); \
		for goos in $(GOOS); do \
			for goarch in $(GOARCH); do \
				GOOS=$$goos GOARCH=$$goarch $(GOBUILD) cmd/dockohealer/dockohealer.go; \
				upx --brute dockohealer; \
				mv dockohealer $(GOBIN)/$(BINARY_NAME)-$$goos-$$goarch-$(VERSION); \
			done \
		done
clean: 
		$(GOCLEAN) cmd/dockohealer/dockohealer.go
		rm -rf $(GOBIN);