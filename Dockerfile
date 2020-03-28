FROM golang:latest

RUN mkdir -p /app

WORKDIR /app

ADD . /app
RUN go get .../.

CMD "go run cmd/main.go"