FROM golang:1.22-alpine

WORKDIR /whoislive

COPY go.mod go.sum .

RUN ["go", "mod", "tidy"]

COPY . .

RUN ["go", "install", "github.com/cosmtrek/air@latest"]

ENTRYPOINT ["/go/bin/air"]

