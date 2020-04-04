FROM alpine:3.11

ENV GOOS linux
ENV GOARCH amd64
ENV VERSION v0.1

RUN apk update && \
    apk add wget && \
    wget https://github.com/maateen/dockohealer/releases/download/v0.1/dockohealer-$GOOS-$GOARCH-$VERSION && \
    mv bin/dockohealer-$GOOS-$GOARCH-$VERSION /usr/local/bin/dockohealer && \
    chmod +x /usr/local/bin/dockohealer

CMD /usr/local/bin/dockohealer