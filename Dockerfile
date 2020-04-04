# build stage
FROM golang:1.13-alpine3.11 AS builder

ADD . /src

RUN cd /src && \
    GOOS=linux GOARCH=amd64 go build cmd/dockohealer/dockohealer.go

# final stage
FROM alpine:3.11

COPY --from=builder /src/dockohealer /usr/local/bin/

CMD /usr/local/bin/dockohealer