FROM golang:1.22.2-alpine3.18

RUN apk update && apk add git

ENV GO111MODULE=on

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH/src/github-trending-api

COPY . .

RUN go mod init github-trending-api
RUN go get ./
RUN go build -o app

ENTRYPOINT ["./app"]

EXPOSE 3000