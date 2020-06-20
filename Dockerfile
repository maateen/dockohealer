# build stage
FROM golang:1.13-alpine3.11 AS builder

RUN apk update && \
    apk add git upx

ADD . /src

RUN cd /src && \
    buildTime=$(date +'%Y-%m-%d_%T') && \
    gitSHA=$(git rev-parse HEAD) && \
    version=$(git tag --sort=committerdate | tail -1) && \
    cd cmd/dockohealer && \
    go build -ldflags "-X main.buildTime=$buildTime -X main.gitSHA=$gitSHA -X main.version=$version" && \
    upx --brute dockohealer

# final stage
FROM alpine:3.11

COPY --from=builder /src/cmd/dockohealer/dockohealer /usr/local/bin/

CMD /usr/local/bin/dockohealer